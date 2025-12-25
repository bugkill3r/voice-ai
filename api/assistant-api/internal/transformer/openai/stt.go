// Copyright (c) 2023-2025 RapidaAI
// Licensed under GPL-2.0 with Rapida Additional Terms.
// See LICENSE.md or contact sales@rapida.ai for commercial usage.

package internal_transformer_openai

import (
	"bytes"
	"context"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"sync"
	"time"

	internal_transformer "github.com/rapidaai/api/assistant-api/internal/transformer"
	"github.com/rapidaai/pkg/commons"
	"github.com/rapidaai/protos"
)

const (
	whisperAPIEndpoint = "https://api.openai.com/v1/audio/transcriptions"
	minBufferBytes     = 3200
	maxBufferBytes     = 480000
	maxBufferDuration  = 5 * time.Second
)

type openaiSpeechToText struct {
	*openaiOption
	mu          sync.Mutex
	ctx         context.Context
	cancel      context.CancelFunc
	logger      commons.Logger
	httpClient  *http.Client
	audioBuffer []byte
	bufferStart time.Time
	options     *internal_transformer.SpeechToTextInitializeOptions
}

func (*openaiSpeechToText) Name() string {
	return "openai-whisper-speech-to-text"
}

func NewOpenaiSpeechToText(
	ctx context.Context,
	logger commons.Logger,
	vaultCredential *protos.VaultCredential,
	opts *internal_transformer.SpeechToTextInitializeOptions,
) (internal_transformer.SpeechToTextTransformer, error) {
	openaiOpts, err := NewOpenaiOption(
		logger,
		vaultCredential,
		opts.AudioConfig,
		opts.ModelOptions,
	)
	if err != nil {
		logger.Errorf("openai-stt: credential extraction failed: %+v", err)
		return nil, err
	}

	return &openaiSpeechToText{
		ctx:          ctx,
		options:      opts,
		logger:       logger,
		openaiOption: openaiOpts,
		audioBuffer:  make([]byte, 0, maxBufferBytes),
	}, nil
}

func (stt *openaiSpeechToText) Initialize() error {
	stt.mu.Lock()
	defer stt.mu.Unlock()

	stt.ctx, stt.cancel = context.WithCancel(stt.ctx)
	stt.httpClient = &http.Client{Timeout: 30 * time.Second}
	stt.bufferStart = time.Now()
	stt.logger.Debugf("openai-stt: initialized with model %s", stt.GetModel())
	return nil
}

func (stt *openaiSpeechToText) Transform(
	ctx context.Context,
	audioData []byte,
	opts *internal_transformer.SpeechToTextOption,
) error {
	stt.mu.Lock()
	defer stt.mu.Unlock()

	if stt.httpClient == nil {
		return fmt.Errorf("openai-stt: not initialized")
	}

	stt.audioBuffer = append(stt.audioBuffer, audioData...)

	bufferDuration := time.Since(stt.bufferStart)
	bufferSize := len(stt.audioBuffer)

	if bufferSize < minBufferBytes {
		return nil
	}

	if bufferDuration > maxBufferDuration || bufferSize > maxBufferBytes {
		return stt.transcribeBuffer(ctx)
	}

	return nil
}

func (stt *openaiSpeechToText) transcribeBuffer(ctx context.Context) error {
	if len(stt.audioBuffer) == 0 {
		return nil
	}

	wavData := stt.createWAV(stt.audioBuffer)
	stt.audioBuffer = stt.audioBuffer[:0]
	stt.bufferStart = time.Now()

	var requestBody bytes.Buffer
	writer := multipart.NewWriter(&requestBody)

	part, err := writer.CreateFormFile("file", "audio.wav")
	if err != nil {
		return fmt.Errorf("openai-stt: form file creation failed: %w", err)
	}
	if _, err := io.Copy(part, bytes.NewReader(wavData)); err != nil {
		return fmt.Errorf("openai-stt: audio copy failed: %w", err)
	}

	_ = writer.WriteField("model", stt.GetModel())
	if lang := stt.GetLanguage(); lang != "" {
		_ = writer.WriteField("language", lang)
	}
	_ = writer.WriteField("response_format", "json")
	writer.Close()

	req, err := http.NewRequestWithContext(ctx, "POST", whisperAPIEndpoint, &requestBody)
	if err != nil {
		return fmt.Errorf("openai-stt: request creation failed: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+stt.GetAPIKey())
	req.Header.Set("Content-Type", writer.FormDataContentType())

	resp, err := stt.httpClient.Do(req)
	if err != nil {
		stt.logger.Errorf("openai-stt: API request failed: %v", err)
		return fmt.Errorf("openai-stt: API request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		stt.logger.Errorf("openai-stt: API error %d: %s", resp.StatusCode, string(body))
		return fmt.Errorf("openai-stt: API error %d", resp.StatusCode)
	}

	var result struct {
		Text string `json:"text"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return fmt.Errorf("openai-stt: response parse failed: %w", err)
	}

	if result.Text != "" && stt.options.OnTranscript != nil {
		if err := stt.options.OnTranscript(result.Text, 1.0, "", true); err != nil {
			stt.logger.Errorf("openai-stt: transcript callback failed: %v", err)
		}
	}

	return nil
}

func (stt *openaiSpeechToText) createWAV(pcmData []byte) []byte {
	sampleRate := stt.audioConfig.GetSampleRate()
	if sampleRate == 0 {
		sampleRate = 16000
	}
	channels := uint16(1)
	bitsPerSample := uint16(16)
	byteRate := uint32(sampleRate) * uint32(channels) * uint32(bitsPerSample) / 8
	blockAlign := channels * bitsPerSample / 8
	dataSize := uint32(len(pcmData))

	buf := new(bytes.Buffer)

	buf.WriteString("RIFF")
	binary.Write(buf, binary.LittleEndian, dataSize+36)
	buf.WriteString("WAVE")

	buf.WriteString("fmt ")
	binary.Write(buf, binary.LittleEndian, uint32(16))
	binary.Write(buf, binary.LittleEndian, uint16(1))
	binary.Write(buf, binary.LittleEndian, channels)
	binary.Write(buf, binary.LittleEndian, uint32(sampleRate))
	binary.Write(buf, binary.LittleEndian, byteRate)
	binary.Write(buf, binary.LittleEndian, blockAlign)
	binary.Write(buf, binary.LittleEndian, bitsPerSample)

	buf.WriteString("data")
	binary.Write(buf, binary.LittleEndian, dataSize)
	buf.Write(pcmData)

	return buf.Bytes()
}

func (stt *openaiSpeechToText) Close(ctx context.Context) error {
	stt.mu.Lock()
	defer stt.mu.Unlock()

	if len(stt.audioBuffer) > 0 {
		_ = stt.transcribeBuffer(ctx)
	}

	if stt.cancel != nil {
		stt.cancel()
	}

	stt.logger.Infof("openai-stt: closed")
	return nil
}

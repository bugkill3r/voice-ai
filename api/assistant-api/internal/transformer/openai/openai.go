// Copyright (c) 2023-2025 RapidaAI
// Licensed under GPL-2.0 with Rapida Additional Terms.
// See LICENSE.md or contact sales@rapida.ai for commercial usage.

package internal_transformer_openai

import (
	"fmt"

	internal_audio "github.com/rapidaai/api/assistant-api/internal/audio"
	"github.com/rapidaai/pkg/commons"
	"github.com/rapidaai/pkg/utils"
	"github.com/rapidaai/protos"
)

type openaiOption struct {
	apiKey      string
	logger      commons.Logger
	mdlOpts     utils.Option
	audioConfig *internal_audio.AudioConfig
}

func NewOpenaiOption(
	logger commons.Logger,
	vaultCredential *protos.VaultCredential,
	audioConfig *internal_audio.AudioConfig,
	opts utils.Option,
) (*openaiOption, error) {
	cx, ok := vaultCredential.GetValue().AsMap()["key"]
	if !ok {
		return nil, fmt.Errorf("openai: missing 'key' in vault config")
	}
	return &openaiOption{
		apiKey:      cx.(string),
		logger:      logger,
		mdlOpts:     opts,
		audioConfig: audioConfig,
	}, nil
}

func (opt *openaiOption) GetAPIKey() string {
	return opt.apiKey
}

func (opt *openaiOption) GetModel() string {
	if model, err := opt.mdlOpts.GetString("listen.model"); err == nil {
		return model
	}
	return "whisper-1"
}

func (opt *openaiOption) GetLanguage() string {
	if lang, err := opt.mdlOpts.GetString("listen.language"); err == nil {
		return lang
	}
	return ""
}

func (opt *openaiOption) GetEncoding() string {
	switch opt.audioConfig.Format {
	case internal_audio.Linear16:
		return "pcm16"
	case internal_audio.MuLaw8:
		return "mulaw"
	default:
		return "pcm16"
	}
}

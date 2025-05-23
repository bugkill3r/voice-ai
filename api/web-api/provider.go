package web_api

import (
	"context"

	"github.com/lexatic/web-backend/artifacts"
	config "github.com/lexatic/web-backend/config"
	commons "github.com/lexatic/web-backend/pkg/commons"
	"github.com/lexatic/web-backend/pkg/connectors"
	"github.com/lexatic/web-backend/pkg/utils"
	web_api "github.com/lexatic/web-backend/protos/lexatic-backend"
)

type webProviderApi struct {
	cfg      *config.AppConfig
	logger   commons.Logger
	postgres connectors.PostgresConnector
	redis    connectors.RedisConnector
}

type webProviderRPCApi struct {
	webProviderApi
}

type webProviderGRPCApi struct {
	webProviderApi
}

func NewProviderGRPC(config *config.AppConfig, logger commons.Logger, postgres connectors.PostgresConnector, redis connectors.RedisConnector) web_api.ProviderServiceServer {
	return &webProviderGRPCApi{
		webProviderApi{
			cfg:      config,
			logger:   logger,
			postgres: postgres,
			redis:    redis,
		},
	}
}

func (w *webProviderGRPCApi) GetAllToolProvider(context.Context, *web_api.GetAllToolProviderRequest) (*web_api.GetAllToolProviderResponse, error) {
	providers, err := artifacts.GetToolProviders()
	if err != nil {
		return utils.Error[web_api.GetAllToolProviderResponse](
			err,
			"Unable to get tool providers, please try again in sometime.")
	}
	return utils.PaginatedSuccess[web_api.GetAllToolProviderResponse, []*web_api.ToolProvider](
		uint32(len(providers)), 1,
		providers)
}

// GetAllModel implements lexatic_backend.ProviderServiceServer.
func (w *webProviderGRPCApi) GetAllModel(context.Context, *web_api.GetAllModelRequest) (*web_api.GetAllModelResponse, error) {
	panic("unimplemented")
}

// GetAllProvider implements lexatic_backend.ProviderServiceServer.
func (w *webProviderGRPCApi) GetAllProvider(context.Context, *web_api.GetAllProviderRequest) (*web_api.GetAllProviderResponse, error) {
	panic("unimplemented")
}

// GetModel implements lexatic_backend.ProviderServiceServer.
func (w *webProviderGRPCApi) GetModel(context.Context, *web_api.GetModelRequest) (*web_api.GetModelResponse, error) {
	panic("unimplemented")
}

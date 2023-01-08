package service

import (
	"context"
	"github.com/RizkiMufrizal/gofiber-clean-architecture/configuration"
	"github.com/RizkiMufrizal/gofiber-clean-architecture/model"
	"github.com/RizkiMufrizal/gofiber-clean-architecture/restclient"
)

func NewHttpBinServiceImpl() HttpBinService {
	return &httpBinServiceImpl{}
}

type httpBinServiceImpl struct {
}

func (h *httpBinServiceImpl) PostMethod(ctx context.Context) {
	httpbinRestClient := restclient.NewHttpBinRestClient()
	httpBin := model.HttpBin{
		Name: "rizki",
	}
	var response map[string]interface{}
	httpbinRestClient.PostMethod(ctx, &httpBin, &response)
	configuration.NewLogger().Info("log response service ", response)
}

package api

import (
	"context"
	"exchanger-parser/internal/garantex/models"
	"fmt"
	"github.com/go-resty/resty/v2"
	"net/http"
)

type API interface {
	Parse(ctx context.Context) (models.GreenUsdtRub, error)
}

type api struct {
	client *resty.Client
}

func NewAPI(
	debug bool,
) API {
	return &api{
		client: resty.New().
			EnableTrace().
			SetDebug(debug).
			SetBaseURL(GarantexBaseURL),
	}
}

func (a *api) Parse(
	ctx context.Context,
) (
	models.GreenUsdtRub,
	error,
) {
	var response models.GreenUsdtRub = models.GreenUsdtRub{}

	resp, err := a.client.R().
		SetContext(ctx).
		SetResult(&response).
		Get(GarantexMarketRate)
	if err != nil {
		return response, err
	}

	if resp.StatusCode() != http.StatusOK {
		return response, fmt.Errorf(
			"status code not StatusOK: %d",
			resp.StatusCode(),
		)
	}

	return response, nil
}

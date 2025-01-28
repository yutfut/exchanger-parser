package api

import (
	"context"
	"exchanger-parser/internal/binance/models"
	"fmt"
	"github.com/go-resty/resty/v2"
	"net/http"
)

type API interface {
	Parse(
		ctx context.Context,
		request models.P2PRequest,
	) (
		models.Response,
		error,
	)
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
			SetBaseURL(BinanceBaseURL),
	}
}

func (a *api) Parse(
	ctx context.Context,
	request models.P2PRequest,
) (
	models.Response,
	error,
) {
	response := models.Response{}

	resp, err := a.client.R().
		SetContext(ctx).
		SetResult(&response).
		SetBody(&request).
		Post(BinanceMarketRate)
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

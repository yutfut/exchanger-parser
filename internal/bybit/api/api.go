package api

import (
	"context"
	"exchanger-parser/internal/bybit/models"
	"fmt"
	"github.com/go-resty/resty/v2"
	"net/http"
)

type API interface {
	Parse(ctx context.Context, request models.P2PCourseParams) (models.P2PCourseResponse, error)
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
			SetBaseURL(BybitBaseURL),
	}
}

func (a *api) Parse(
	ctx context.Context,
	request models.P2PCourseParams,
) (
	models.P2PCourseResponse,
	error,
) {
	var response models.P2PCourseResponse = models.P2PCourseResponse{}

	resp, err := a.client.R().
		SetContext(ctx).
		SetBody(request).
		SetResult(&response).
		Post(P2PCourse)
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

package http

import (
	"exchanger-parser/internal/api/models"
	"exchanger-parser/internal/api/useCase"
	"exchanger-parser/pkg/utils"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

func NewRouting(router chi.Router, a Handler) {
	router.Route(
		"/api/v1",
		func(r chi.Router) {
			r.Get("/get", a.Get)
		},
	)
}

type Handler interface {
	Get(w http.ResponseWriter, r *http.Request)
}

type handler struct {
	useCase useCase.UseCase
	logger  *log.Logger
}

func NewHandler(
	useCase useCase.UseCase,
	logger *log.Logger,
) Handler {
	return &handler{
		useCase: useCase,
		logger:  logger,
	}
}

func (h handler) Get(
	w http.ResponseWriter,
	r *http.Request,
) {
	request := models.Request{}

	if err := utils.ParseBody(r, &request); err != nil {
		h.logger.Println(err)
		utils.SendError(
			w,
			http.StatusBadRequest,
			err,
		)
		return
	}

	response, err := h.useCase.Get(
		r.Context(),
		request,
	)
	if err != nil {
		utils.SendError(
			w,
			http.StatusInternalServerError,
			err,
		)
		return
	}

	utils.SendRAWJSON(
		w,
		http.StatusOK,
		response,
	)
}

// grpc
// болше бирж для подстраховки

package http

import (
	"net/http"

	"jamtangan/domain"
)

type httpHandler struct {
	adminUseCase    domain.AdminUseCase
	customerUseCase domain.CustomerUseCase
}

func NewHandler(
	adminUseCase domain.AdminUseCase,
	customerUseCase domain.CustomerUseCase,
) *httpHandler {
	return &httpHandler{
		adminUseCase:    adminUseCase,
		customerUseCase: customerUseCase,
	}
}

func (h *httpHandler) Health(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write([]byte("healthy"))
}

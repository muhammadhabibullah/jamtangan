package http

import (
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

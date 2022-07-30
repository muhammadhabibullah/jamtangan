package http_test

import (
	"testing"

	"github.com/onsi/gomega"

	"jamtangan/domain"
	httpHandler "jamtangan/handler/http"
	mocks "jamtangan/mock"
)

type httpHandlerTest struct {
	httpHandler     domain.HTTPHandler
	adminUseCase    *mocks.AdminUseCase
	customerUseCase *mocks.CustomerUseCase
}

func test(t *testing.T, fn func(*gomega.WithT, *httpHandlerTest)) {
	adminUseCase := new(mocks.AdminUseCase)
	customerUseCase := new(mocks.CustomerUseCase)

	h := httpHandler.NewHandler(adminUseCase, customerUseCase)

	hh := httpHandlerTest{
		httpHandler:     h,
		adminUseCase:    adminUseCase,
		customerUseCase: customerUseCase,
	}

	g := gomega.NewGomegaWithT(t)
	fn(g, &hh)
}

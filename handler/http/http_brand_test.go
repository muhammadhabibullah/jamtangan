package http_test

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/mock"

	"jamtangan/domain"
)

func TestHttpHandler_Brand_Create(t *testing.T) {
	t.Parallel()

	const (
		CreateBrand = "CreateBrand"

		brandName = "CASIO"
	)

	var (
		createdTime, _ = time.Parse(time.RFC3339, "2022-07-29T06:08:59.004834Z")

		brand = domain.Brand{
			ID:        1552898963420483584,
			Name:      brandName,
			CreatedAt: createdTime,
			UpdatedAt: createdTime,
		}
		errDatabase = fmt.Errorf("database error")
	)

	tests := []struct {
		name string
		test func(*GomegaWithT, *httpHandlerTest)
	}{
		{
			name: "invalid method",
			test: func(t *GomegaWithT, hh *httpHandlerTest) {
				req, _ := ioutil.ReadFile("json/brand/create_request.json")
				r := httptest.NewRequest(http.MethodDelete, "/brand", bytes.NewBuffer(req))
				w := httptest.NewRecorder()

				hh.httpHandler.Brand(w, r)
				res := w.Result()
				defer res.Body.Close()

				data, err := ioutil.ReadAll(res.Body)
				t.Expect(err).ShouldNot(HaveOccurred())

				expectedData, _ := ioutil.ReadFile("json/invalid_method_request.json")
				t.Expect(data).Should(MatchJSON(expectedData))
				t.Expect(res.StatusCode).Should(Equal(http.StatusMethodNotAllowed))
			},
		},
		{
			name: "invalid request body",
			test: func(t *GomegaWithT, hh *httpHandlerTest) {
				req, _ := ioutil.ReadFile("json/brand/create_invalid_request.json")
				r := httptest.NewRequest(http.MethodPost, "/brand", bytes.NewBuffer(req))
				w := httptest.NewRecorder()

				hh.httpHandler.Brand(w, r)
				res := w.Result()
				defer res.Body.Close()

				data, err := ioutil.ReadAll(res.Body)
				t.Expect(err).ShouldNot(HaveOccurred())

				expectedData, _ := ioutil.ReadFile("json/brand/create_invalid_response.json")
				t.Expect(data).Should(MatchJSON(expectedData))
				t.Expect(res.StatusCode).Should(Equal(http.StatusBadRequest))
			},
		},
		{
			name: "error missing brand name",
			test: func(t *GomegaWithT, hh *httpHandlerTest) {
				req, _ := ioutil.ReadFile("json/brand/create_missing_name_request.json")
				r := httptest.NewRequest(http.MethodPost, "/brand", bytes.NewBuffer(req))
				w := httptest.NewRecorder()

				hh.httpHandler.Brand(w, r)
				res := w.Result()
				defer res.Body.Close()

				data, err := ioutil.ReadAll(res.Body)
				t.Expect(err).ShouldNot(HaveOccurred())

				expectedData, _ := ioutil.ReadFile("json/brand/create_missing_name_response.json")
				t.Expect(data).Should(MatchJSON(expectedData))
				t.Expect(res.StatusCode).Should(Equal(http.StatusBadRequest))
			},
		},
		{
			name: "error duplicate",
			test: func(t *GomegaWithT, hh *httpHandlerTest) {
				hh.adminUseCase.On(CreateBrand, mock.Anything, brandName).
					Return(domain.Brand{}, domain.ErrDuplicate)

				req, _ := ioutil.ReadFile("json/brand/create_request.json")
				r := httptest.NewRequest(http.MethodPost, "/brand", bytes.NewBuffer(req))
				w := httptest.NewRecorder()

				hh.httpHandler.Brand(w, r)
				res := w.Result()
				defer res.Body.Close()

				data, err := ioutil.ReadAll(res.Body)
				t.Expect(err).ShouldNot(HaveOccurred())

				expectedData, _ := ioutil.ReadFile("json/brand/create_duplicate_response.json")
				t.Expect(data).Should(MatchJSON(expectedData))
				t.Expect(res.StatusCode).Should(Equal(http.StatusConflict))
			},
		},
		{
			name: "error server",
			test: func(t *GomegaWithT, hh *httpHandlerTest) {
				hh.adminUseCase.On(CreateBrand, mock.Anything, brandName).
					Return(domain.Brand{}, errDatabase)

				req, _ := ioutil.ReadFile("json/brand/create_request.json")
				r := httptest.NewRequest(http.MethodPost, "/brand", bytes.NewBuffer(req))
				w := httptest.NewRecorder()

				hh.httpHandler.Brand(w, r)
				res := w.Result()
				defer res.Body.Close()

				data, err := ioutil.ReadAll(res.Body)
				t.Expect(err).ShouldNot(HaveOccurred())

				expectedData, _ := ioutil.ReadFile("json/server_error_response.json")
				t.Expect(data).Should(MatchJSON(expectedData))
				t.Expect(res.StatusCode).Should(Equal(http.StatusInternalServerError))
			},
		},
		{
			name: "success",
			test: func(t *GomegaWithT, hh *httpHandlerTest) {
				hh.adminUseCase.On(CreateBrand, mock.Anything, brandName).
					Return(brand, nil)

				req, _ := ioutil.ReadFile("json/brand/create_request.json")
				r := httptest.NewRequest(http.MethodPost, "/brand", bytes.NewBuffer(req))
				w := httptest.NewRecorder()

				hh.httpHandler.Brand(w, r)
				res := w.Result()
				defer res.Body.Close()

				data, err := ioutil.ReadAll(res.Body)
				t.Expect(err).ShouldNot(HaveOccurred())

				expectedData, _ := ioutil.ReadFile("json/brand/create_response.json")
				t.Expect(data).Should(MatchJSON(expectedData))
				t.Expect(res.StatusCode).Should(Equal(http.StatusOK))
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			test(t, tt.test)
		})
	}
}

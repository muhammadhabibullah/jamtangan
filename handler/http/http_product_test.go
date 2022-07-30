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

func TestHttpHandler_Product_Create(t *testing.T) {
	t.Parallel()

	const (
		CreateProduct = "CreateProduct"

		productID        = 1552703849368653824
		productName      = "Casio G-Shock GX-56BB-1DR King Kong Solar Powered WR 200M Black Resin Band"
		productPrice     = 1450000
		productBrandID   = 1552655170888798208
		productCreatedAt = "2022-07-29T06:08:59.004834Z"
	)

	var (
		createdTime, _ = time.Parse(time.RFC3339, productCreatedAt)

		productRequest = domain.Product{
			Name:    productName,
			Price:   productPrice,
			BrandID: productBrandID,
		}

		product = domain.Product{
			ID:        productID,
			Name:      productName,
			Price:     productPrice,
			BrandID:   productBrandID,
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
				req, _ := ioutil.ReadFile("json/product/create_request.json")
				r := httptest.NewRequest(http.MethodDelete, "/product", bytes.NewBuffer(req))
				w := httptest.NewRecorder()

				hh.httpHandler.Product(w, r)
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
				req, _ := ioutil.ReadFile("json/product/create_invalid_request.json")
				r := httptest.NewRequest(http.MethodPost, "/product", bytes.NewBuffer(req))
				w := httptest.NewRecorder()

				hh.httpHandler.Product(w, r)
				res := w.Result()
				defer res.Body.Close()

				data, err := ioutil.ReadAll(res.Body)
				t.Expect(err).ShouldNot(HaveOccurred())

				expectedData, _ := ioutil.ReadFile("json/product/create_invalid_response.json")
				t.Expect(data).Should(MatchJSON(expectedData))
				t.Expect(res.StatusCode).Should(Equal(http.StatusBadRequest))
			},
		},
		{
			name: "error missing product name",
			test: func(t *GomegaWithT, hh *httpHandlerTest) {
				req, _ := ioutil.ReadFile("json/product/create_missing_name_request.json")
				r := httptest.NewRequest(http.MethodPost, "/product", bytes.NewBuffer(req))
				w := httptest.NewRecorder()

				hh.httpHandler.Product(w, r)
				res := w.Result()
				defer res.Body.Close()

				data, err := ioutil.ReadAll(res.Body)
				t.Expect(err).ShouldNot(HaveOccurred())

				expectedData, _ := ioutil.ReadFile("json/product/create_missing_name_response.json")
				t.Expect(data).Should(MatchJSON(expectedData))
				t.Expect(res.StatusCode).Should(Equal(http.StatusBadRequest))
			},
		},
		{
			name: "error missing product price",
			test: func(t *GomegaWithT, hh *httpHandlerTest) {
				req, _ := ioutil.ReadFile("json/product/create_missing_price_request.json")
				r := httptest.NewRequest(http.MethodPost, "/product", bytes.NewBuffer(req))
				w := httptest.NewRecorder()

				hh.httpHandler.Product(w, r)
				res := w.Result()
				defer res.Body.Close()

				data, err := ioutil.ReadAll(res.Body)
				t.Expect(err).ShouldNot(HaveOccurred())

				expectedData, _ := ioutil.ReadFile("json/product/create_missing_price_response.json")
				t.Expect(data).Should(MatchJSON(expectedData))
				t.Expect(res.StatusCode).Should(Equal(http.StatusBadRequest))
			},
		},
		{
			name: "error brand ID not found",
			test: func(t *GomegaWithT, hh *httpHandlerTest) {
				hh.adminUseCase.On(CreateProduct, mock.Anything, productRequest).
					Return(domain.Product{}, domain.ErrBadRequest)

				req, _ := ioutil.ReadFile("json/product/create_request.json")
				r := httptest.NewRequest(http.MethodPost, "/product", bytes.NewBuffer(req))
				w := httptest.NewRecorder()

				hh.httpHandler.Product(w, r)
				res := w.Result()
				defer res.Body.Close()

				data, err := ioutil.ReadAll(res.Body)
				t.Expect(err).ShouldNot(HaveOccurred())

				expectedData, _ := ioutil.ReadFile("json/product/create_brand_id_not_found_response.json")
				t.Expect(data).Should(MatchJSON(expectedData))
				t.Expect(res.StatusCode).Should(Equal(http.StatusBadRequest))
			},
		},
		{
			name: "error server",
			test: func(t *GomegaWithT, hh *httpHandlerTest) {
				hh.adminUseCase.On(CreateProduct, mock.Anything, productRequest).
					Return(domain.Product{}, errDatabase)

				req, _ := ioutil.ReadFile("json/product/create_request.json")
				r := httptest.NewRequest(http.MethodPost, "/product", bytes.NewBuffer(req))
				w := httptest.NewRecorder()

				hh.httpHandler.Product(w, r)
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
				hh.adminUseCase.On(CreateProduct, mock.Anything, productRequest).
					Return(product, nil)

				req, _ := ioutil.ReadFile("json/product/create_request.json")
				r := httptest.NewRequest(http.MethodPost, "/product", bytes.NewBuffer(req))
				w := httptest.NewRecorder()

				hh.httpHandler.Product(w, r)
				res := w.Result()
				defer res.Body.Close()

				data, err := ioutil.ReadAll(res.Body)
				t.Expect(err).ShouldNot(HaveOccurred())

				expectedData, _ := ioutil.ReadFile("json/product/create_response.json")
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

func TestHttpHandler_Product_GetByID(t *testing.T) {
	t.Parallel()

	const (
		GetProductByID = "GetProductByID"

		productID        int64 = 1552703849368653824
		productName            = "Casio G-Shock GX-56BB-1DR King Kong Solar Powered WR 200M Black Resin Band"
		productPrice           = 1450000
		productBrandID   int64 = 1552655170888798208
		productCreatedAt       = "2022-07-28T17:13:40.190124Z"
	)

	var (
		createdTime, _ = time.Parse(time.RFC3339, productCreatedAt)

		product = domain.Product{
			ID:        productID,
			Name:      productName,
			Price:     productPrice,
			BrandID:   productBrandID,
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
			name: "invalid ID",
			test: func(t *GomegaWithT, hh *httpHandlerTest) {
				r := httptest.NewRequest(http.MethodGet, "/product", nil)
				w := httptest.NewRecorder()

				hh.httpHandler.Product(w, r)
				res := w.Result()
				defer res.Body.Close()

				data, err := ioutil.ReadAll(res.Body)
				t.Expect(err).ShouldNot(HaveOccurred())

				expectedData, _ := ioutil.ReadFile("json/product/get_by_invalid_id_response.json")
				t.Expect(data).Should(MatchJSON(expectedData))
				t.Expect(res.StatusCode).Should(Equal(http.StatusBadRequest))
			},
		},
		{
			name: "error not found",
			test: func(t *GomegaWithT, hh *httpHandlerTest) {
				hh.adminUseCase.On(GetProductByID, mock.Anything, productID).
					Return(domain.Product{}, domain.ErrNotFound)

				target := fmt.Sprintf("/product?id=%d", productID)
				r := httptest.NewRequest(http.MethodGet, target, nil)
				w := httptest.NewRecorder()

				hh.httpHandler.Product(w, r)
				res := w.Result()
				defer res.Body.Close()

				data, err := ioutil.ReadAll(res.Body)
				t.Expect(err).ShouldNot(HaveOccurred())

				expectedData, _ := ioutil.ReadFile("json/product/get_by_id_not_found_response.json")
				t.Expect(data).Should(MatchJSON(expectedData))
				t.Expect(res.StatusCode).Should(Equal(http.StatusNotFound))
			},
		},
		{
			name: "error server",
			test: func(t *GomegaWithT, hh *httpHandlerTest) {
				hh.adminUseCase.On(GetProductByID, mock.Anything, productID).
					Return(domain.Product{}, errDatabase)

				target := fmt.Sprintf("/product?id=%d", productID)
				r := httptest.NewRequest(http.MethodGet, target, nil)
				w := httptest.NewRecorder()

				hh.httpHandler.Product(w, r)
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
				hh.adminUseCase.On(GetProductByID, mock.Anything, productID).
					Return(product, nil)

				target := fmt.Sprintf("/product?id=%d", productID)
				r := httptest.NewRequest(http.MethodGet, target, nil)
				w := httptest.NewRecorder()

				hh.httpHandler.Product(w, r)
				res := w.Result()
				defer res.Body.Close()

				data, err := ioutil.ReadAll(res.Body)
				t.Expect(err).ShouldNot(HaveOccurred())

				expectedData, _ := ioutil.ReadFile("json/product/get_by_id_response.json")
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

func TestHttpHandler_Product_FetchByBrandID(t *testing.T) {
	t.Parallel()

	const (
		FetchProductByBrandID = "FetchProductByBrandID"

		productID        int64 = 1552703849368653824
		productName            = "Casio G-Shock GX-56BB-1DR King Kong Solar Powered WR 200M Black Resin Band"
		productPrice           = 1450000
		productBrandID   int64 = 1552655170888798208
		productCreatedAt       = "2022-07-28T17:13:40.190124Z"
	)

	var (
		createdTime, _ = time.Parse(time.RFC3339, productCreatedAt)

		products = []domain.Product{
			{
				ID:        productID,
				Name:      productName,
				Price:     productPrice,
				BrandID:   productBrandID,
				CreatedAt: createdTime,
				UpdatedAt: createdTime,
			},
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
				target := fmt.Sprintf("/product/brand?id=%d", productBrandID)
				r := httptest.NewRequest(http.MethodDelete, target, nil)
				w := httptest.NewRecorder()

				hh.httpHandler.ProductBrand(w, r)
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
			name: "invalid ID",
			test: func(t *GomegaWithT, hh *httpHandlerTest) {
				r := httptest.NewRequest(http.MethodGet, "/product/brand", nil)
				w := httptest.NewRecorder()

				hh.httpHandler.ProductBrand(w, r)
				res := w.Result()
				defer res.Body.Close()

				data, err := ioutil.ReadAll(res.Body)
				t.Expect(err).ShouldNot(HaveOccurred())

				expectedData, _ := ioutil.ReadFile("json/product/get_by_invalid_id_response.json")
				t.Expect(data).Should(MatchJSON(expectedData))
				t.Expect(res.StatusCode).Should(Equal(http.StatusBadRequest))
			},
		},
		{
			name: "error not found",
			test: func(t *GomegaWithT, hh *httpHandlerTest) {
				hh.customerUseCase.On(FetchProductByBrandID, mock.Anything, productBrandID).
					Return(nil, domain.ErrNotFound)

				target := fmt.Sprintf("/product/brand?id=%d", productBrandID)
				r := httptest.NewRequest(http.MethodGet, target, nil)
				w := httptest.NewRecorder()

				hh.httpHandler.ProductBrand(w, r)
				res := w.Result()
				defer res.Body.Close()

				data, err := ioutil.ReadAll(res.Body)
				t.Expect(err).ShouldNot(HaveOccurred())

				expectedData, _ := ioutil.ReadFile("json/product/get_by_id_not_found_response.json")
				t.Expect(data).Should(MatchJSON(expectedData))
				t.Expect(res.StatusCode).Should(Equal(http.StatusNotFound))
			},
		},
		{
			name: "error server",
			test: func(t *GomegaWithT, hh *httpHandlerTest) {
				hh.customerUseCase.On(FetchProductByBrandID, mock.Anything, productBrandID).
					Return(nil, errDatabase)

				target := fmt.Sprintf("/product/brand?id=%d", productBrandID)
				r := httptest.NewRequest(http.MethodGet, target, nil)
				w := httptest.NewRecorder()

				hh.httpHandler.ProductBrand(w, r)
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
				hh.customerUseCase.On(FetchProductByBrandID, mock.Anything, productBrandID).
					Return(products, nil)

				target := fmt.Sprintf("/product/brand?id=%d", productBrandID)
				r := httptest.NewRequest(http.MethodGet, target, nil)
				w := httptest.NewRecorder()

				hh.httpHandler.ProductBrand(w, r)
				res := w.Result()
				defer res.Body.Close()

				data, err := ioutil.ReadAll(res.Body)
				t.Expect(err).ShouldNot(HaveOccurred())

				expectedData, _ := ioutil.ReadFile("json/product/get_by_brand_id_response.json")
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

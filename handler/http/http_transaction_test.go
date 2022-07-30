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

func TestHttpHandler_Transaction_Create(t *testing.T) {
	t.Parallel()

	const (
		CreateTransaction = "CreateTransaction"

		id                 int64 = 1553196332603150336
		totalPrice         int64 = 13597000
		firstProductPrice  int64 = 4799000
		secondProductPrice int64 = 3999000

		transactionCreatedAt              = "2022-07-30T01:50:37.342581Z"
		firstTransactionProductCreatedAt  = "2022-07-30T01:50:37.342756Z"
		secondTransactionProductCreatedAt = "2022-07-30T01:50:37.342839Z"
	)

	var (
		firstTransactionProduct = domain.TransactionProduct{
			ProductID: 1552676634333548544,
			Quantity:  2,
		}
		secondTransactionProduct = domain.TransactionProduct{
			ProductID: 1552688894279946240,
			Quantity:  1,
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
				req, _ := ioutil.ReadFile("json/transaction/create_request.json")
				r := httptest.NewRequest(http.MethodDelete, "/transaction", bytes.NewBuffer(req))
				w := httptest.NewRecorder()

				hh.httpHandler.Transaction(w, r)
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
				req, _ := ioutil.ReadFile("json/transaction/create_invalid_request.json")
				r := httptest.NewRequest(http.MethodPost, "/transaction", bytes.NewBuffer(req))
				w := httptest.NewRecorder()

				hh.httpHandler.Transaction(w, r)
				res := w.Result()
				defer res.Body.Close()

				data, err := ioutil.ReadAll(res.Body)
				t.Expect(err).ShouldNot(HaveOccurred())

				expectedData, _ := ioutil.ReadFile("json/transaction/create_invalid_response.json")
				t.Expect(data).Should(MatchJSON(expectedData))
				t.Expect(res.StatusCode).Should(Equal(http.StatusBadRequest))
			},
		},
		{
			name: "error missing quantity",
			test: func(t *GomegaWithT, hh *httpHandlerTest) {
				req, _ := ioutil.ReadFile("json/transaction/create_missing_qty_request.json")
				r := httptest.NewRequest(http.MethodPost, "/transaction", bytes.NewBuffer(req))
				w := httptest.NewRecorder()

				hh.httpHandler.Transaction(w, r)
				res := w.Result()
				defer res.Body.Close()

				data, err := ioutil.ReadAll(res.Body)
				t.Expect(err).ShouldNot(HaveOccurred())

				expectedData, _ := ioutil.ReadFile("json/transaction/create_missing_qty_response.json")
				t.Expect(data).Should(MatchJSON(expectedData))
				t.Expect(res.StatusCode).Should(Equal(http.StatusBadRequest))
			},
		},
		{
			name: "error missing product ID",
			test: func(t *GomegaWithT, hh *httpHandlerTest) {
				req, _ := ioutil.ReadFile("json/transaction/create_missing_product_id_request.json")
				r := httptest.NewRequest(http.MethodPost, "/transaction", bytes.NewBuffer(req))
				w := httptest.NewRecorder()

				hh.httpHandler.Transaction(w, r)
				res := w.Result()
				defer res.Body.Close()

				data, err := ioutil.ReadAll(res.Body)
				t.Expect(err).ShouldNot(HaveOccurred())

				expectedData, _ := ioutil.ReadFile("json/transaction/create_missing_product_id_response.json")
				t.Expect(data).Should(MatchJSON(expectedData))
				t.Expect(res.StatusCode).Should(Equal(http.StatusBadRequest))
			},
		},
		{
			name: "error product ID not found",
			test: func(t *GomegaWithT, hh *httpHandlerTest) {
				createTransaction := domain.TransactionDetail{
					TransactionProducts: []*domain.TransactionProduct{
						&firstTransactionProduct,
						&secondTransactionProduct,
					},
				}
				hh.customerUseCase.On(CreateTransaction, mock.Anything, &createTransaction).
					Return(domain.ErrNotFound)

				req, _ := ioutil.ReadFile("json/transaction/create_request.json")
				r := httptest.NewRequest(http.MethodPost, "/transaction", bytes.NewBuffer(req))
				w := httptest.NewRecorder()

				hh.httpHandler.Transaction(w, r)
				res := w.Result()
				defer res.Body.Close()

				data, err := ioutil.ReadAll(res.Body)
				t.Expect(err).ShouldNot(HaveOccurred())

				expectedData, _ := ioutil.ReadFile("json/transaction/create_product_id_not_found_response.json")
				t.Expect(data).Should(MatchJSON(expectedData))
				t.Expect(res.StatusCode).Should(Equal(http.StatusBadRequest))
			},
		},
		{
			name: "error server",
			test: func(t *GomegaWithT, hh *httpHandlerTest) {
				createTransaction := domain.TransactionDetail{
					TransactionProducts: []*domain.TransactionProduct{
						&firstTransactionProduct,
						&secondTransactionProduct,
					},
				}
				hh.customerUseCase.On(CreateTransaction, mock.Anything, &createTransaction).
					Return(errDatabase)

				req, _ := ioutil.ReadFile("json/transaction/create_request.json")
				r := httptest.NewRequest(http.MethodPost, "/transaction", bytes.NewBuffer(req))
				w := httptest.NewRecorder()

				hh.httpHandler.Transaction(w, r)
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
				createTransaction := domain.TransactionDetail{
					TransactionProducts: []*domain.TransactionProduct{
						&firstTransactionProduct,
						&secondTransactionProduct,
					},
				}
				hh.customerUseCase.On(CreateTransaction, mock.Anything, &createTransaction).
					Run(func(args mock.Arguments) {
						arg := args.Get(1).(*domain.TransactionDetail)
						tCreatedAt, _ := time.Parse(time.RFC3339, transactionCreatedAt)
						arg.Transaction = &domain.Transaction{
							ID:         id,
							TotalPrice: totalPrice,
							CreatedAt:  tCreatedAt,
							UpdatedAt:  tCreatedAt,
						}

						ftpCreatedAt, _ := time.Parse(time.RFC3339, firstTransactionProductCreatedAt)
						arg.TransactionProducts[0].TransactionID = id
						arg.TransactionProducts[0].Price = firstProductPrice
						arg.TransactionProducts[0].CreatedAt = ftpCreatedAt
						arg.TransactionProducts[0].UpdatedAt = ftpCreatedAt

						stpCreatedAt, _ := time.Parse(time.RFC3339, secondTransactionProductCreatedAt)
						arg.TransactionProducts[1].TransactionID = id
						arg.TransactionProducts[1].Price = secondProductPrice
						arg.TransactionProducts[1].CreatedAt = stpCreatedAt
						arg.TransactionProducts[1].UpdatedAt = stpCreatedAt
					}).
					Return(nil)

				req, _ := ioutil.ReadFile("json/transaction/create_request.json")
				r := httptest.NewRequest(http.MethodPost, "/transaction", bytes.NewBuffer(req))
				w := httptest.NewRecorder()

				hh.httpHandler.Transaction(w, r)
				res := w.Result()
				defer res.Body.Close()

				data, err := ioutil.ReadAll(res.Body)
				t.Expect(err).ShouldNot(HaveOccurred())

				expectedData, _ := ioutil.ReadFile("json/transaction/create_response.json")
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

func TestHttpHandler_Transaction_GetByID(t *testing.T) {
	t.Parallel()

	const (
		GetTransactionByID = "GetTransactionByID"

		transactionID        int64 = 1553196332603150336
		totalPrice           int64 = 13597000
		transactionCreatedAt       = "2022-07-30T01:50:37.342581Z"

		firstProductID                   int64 = 1552676634333548544
		firstTransactionProductPrice     int64 = 4799000
		firstProductQuantity             int64 = 2
		firstTransactionProductCreatedAt       = "2022-07-30T01:50:37.342756Z"

		secondProductID                   int64 = 1552688894279946240
		secondTransactionProductPrice     int64 = 3999000
		secondProductQuantity             int64 = 1
		secondTransactionProductCreatedAt       = "2022-07-30T01:50:37.342839Z"
	)

	var (
		transactionCreatedTime, _              = time.Parse(time.RFC3339, transactionCreatedAt)
		firstTransactionProductCreatedTime, _  = time.Parse(time.RFC3339, firstTransactionProductCreatedAt)
		secondTransactionProductCreatedTime, _ = time.Parse(time.RFC3339, secondTransactionProductCreatedAt)

		transaction = domain.Transaction{
			ID:         transactionID,
			TotalPrice: totalPrice,
			CreatedAt:  transactionCreatedTime,
			UpdatedAt:  transactionCreatedTime,
		}

		firstTransactionProduct = domain.TransactionProduct{
			TransactionID: transactionID,
			ProductID:     firstProductID,
			Price:         firstTransactionProductPrice,
			Quantity:      firstProductQuantity,
			CreatedAt:     firstTransactionProductCreatedTime,
			UpdatedAt:     firstTransactionProductCreatedTime,
		}

		secondTransactionProduct = domain.TransactionProduct{
			TransactionID: transactionID,
			ProductID:     secondProductID,
			Price:         secondTransactionProductPrice,
			Quantity:      secondProductQuantity,
			CreatedAt:     secondTransactionProductCreatedTime,
			UpdatedAt:     secondTransactionProductCreatedTime,
		}

		transactionDetail = domain.TransactionDetail{
			Transaction: &transaction,
			TransactionProducts: []*domain.TransactionProduct{
				&firstTransactionProduct,
				&secondTransactionProduct,
			},
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
				r := httptest.NewRequest(http.MethodGet, "/transaction", nil)
				w := httptest.NewRecorder()

				hh.httpHandler.Transaction(w, r)
				res := w.Result()
				defer res.Body.Close()

				data, err := ioutil.ReadAll(res.Body)
				t.Expect(err).ShouldNot(HaveOccurred())

				expectedData, _ := ioutil.ReadFile("json/transaction/get_by_invalid_id_response.json")
				t.Expect(data).Should(MatchJSON(expectedData))
				t.Expect(res.StatusCode).Should(Equal(http.StatusBadRequest))
			},
		},
		{
			name: "error not found",
			test: func(t *GomegaWithT, hh *httpHandlerTest) {
				hh.customerUseCase.On(GetTransactionByID, mock.Anything, transactionID).
					Return(domain.TransactionDetail{}, domain.ErrNotFound)

				target := fmt.Sprintf("/transaction?id=%d", transactionID)
				r := httptest.NewRequest(http.MethodGet, target, nil)
				w := httptest.NewRecorder()

				hh.httpHandler.Transaction(w, r)
				res := w.Result()
				defer res.Body.Close()

				data, err := ioutil.ReadAll(res.Body)
				t.Expect(err).ShouldNot(HaveOccurred())

				expectedData, _ := ioutil.ReadFile("json/transaction/get_by_id_not_found_response.json")
				t.Expect(data).Should(MatchJSON(expectedData))
				t.Expect(res.StatusCode).Should(Equal(http.StatusNotFound))
			},
		},
		{
			name: "error server",
			test: func(t *GomegaWithT, hh *httpHandlerTest) {
				hh.customerUseCase.On(GetTransactionByID, mock.Anything, transactionID).
					Return(domain.TransactionDetail{}, errDatabase)

				target := fmt.Sprintf("/transaction?id=%d", transactionID)
				r := httptest.NewRequest(http.MethodGet, target, nil)
				w := httptest.NewRecorder()

				hh.httpHandler.Transaction(w, r)
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
				hh.customerUseCase.On(GetTransactionByID, mock.Anything, transactionID).
					Return(transactionDetail, nil)

				target := fmt.Sprintf("/transaction?id=%d", transactionID)
				r := httptest.NewRequest(http.MethodGet, target, nil)
				w := httptest.NewRecorder()

				hh.httpHandler.Transaction(w, r)
				res := w.Result()
				defer res.Body.Close()

				data, err := ioutil.ReadAll(res.Body)
				t.Expect(err).ShouldNot(HaveOccurred())

				expectedData, _ := ioutil.ReadFile("json/transaction/get_by_id_response.json")
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

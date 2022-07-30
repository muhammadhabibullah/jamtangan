package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"jamtangan/domain"
)

func (h *httpHandler) Transaction(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.getTransactionByID(w, r)
	case http.MethodPost:
		h.createTransaction(w, r)
	default:
		http.Error(w, domain.NewHTTPError(domain.ErrInvalidRequestMethod), http.StatusMethodNotAllowed)
	}
}

// createTransaction
// @Summary      Create new transaction
// @Description  Create new transaction
// @Tags         Transaction
// @Accept       json
// @Produce      json
// @Param 		 request body 	domain.CreateTransactionRequest true "Create transaction request body"
// @Success      200  {object}  domain.TransactionDetail
// @Failure      400  {object}  domain.HTTPError
// @Failure      409  {object}  domain.HTTPError
// @Failure      500  {object}  domain.HTTPError
// @Router       /transaction [post]
func (h *httpHandler) createTransaction(w http.ResponseWriter, r *http.Request) {
	var req domain.CreateTransactionRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, domain.NewHTTPError(err), http.StatusBadRequest)
		return
	}

	if err = req.Validate(); err != nil {
		http.Error(w, domain.NewHTTPError(err), http.StatusBadRequest)
		return
	}

	transactionDetail := req.ToTransactionDetail()
	err = h.customerUseCase.CreateTransaction(r.Context(), &transactionDetail)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			http.Error(w, domain.NewHTTPError(err), http.StatusBadRequest)
			return
		}

		http.Error(w, domain.NewHTTPError(err), http.StatusInternalServerError)
		return
	}

	response, _ := json.MarshalIndent(&transactionDetail, "", "    ")
	_, _ = fmt.Fprintf(w, string(response))
}

// getTransactionByID
// @Summary      Get transaction by ID
// @Description  Get transaction by ID
// @Tags         Transaction
// @Accept       json
// @Produce      json
// @Param 		 id query string true "Transaction ID"
// @Success      200  {object}  domain.TransactionDetail
// @Failure      400  {object}  domain.HTTPError
// @Failure      404  {object}  domain.HTTPError
// @Failure      500  {object}  domain.HTTPError
// @Router       /transaction [get]
func (h *httpHandler) getTransactionByID(w http.ResponseWriter, r *http.Request) {
	transactionIDQuery := r.URL.Query().Get("id")
	transactionID, _ := strconv.ParseInt(transactionIDQuery, 10, 64)
	if transactionID == 0 {
		http.Error(w, domain.NewHTTPError(domain.ErrInvalidID), http.StatusBadRequest)
		return
	}

	transaction, err := h.customerUseCase.GetTransactionByID(r.Context(), transactionID)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			http.Error(w, domain.NewHTTPError(err), http.StatusNotFound)
			return
		}

		http.Error(w, domain.NewHTTPError(err), http.StatusInternalServerError)
		return
	}

	response, _ := json.MarshalIndent(&transaction, "", "    ")
	_, _ = fmt.Fprintf(w, string(response))
}

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

func (h *httpHandler) createTransaction(w http.ResponseWriter, r *http.Request) {
	var createTransaction domain.TransactionDetail
	err := json.NewDecoder(r.Body).Decode(&createTransaction)
	if err != nil {
		http.Error(w, domain.NewHTTPError(err), http.StatusBadRequest)
		return
	}

	if err = createTransaction.Validate(); err != nil {
		http.Error(w, domain.NewHTTPError(err), http.StatusBadRequest)
		return
	}

	err = h.customerUseCase.CreateTransaction(r.Context(), &createTransaction)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			http.Error(w, domain.NewHTTPError(err), http.StatusBadRequest)
			return
		}

		http.Error(w, domain.NewHTTPError(err), http.StatusInternalServerError)
		return
	}

	response, _ := json.Marshal(&createTransaction)
	_, _ = fmt.Fprintf(w, string(response))
}

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

	response, _ := json.Marshal(&transaction)
	_, _ = fmt.Fprintf(w, string(response))
}

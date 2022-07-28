package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"jamtangan/domain"
)

func (h *httpHandler) Brand(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		h.createBrand(w, r)
	default:
		http.Error(w, domain.NewHTTPError(domain.ErrInvalidRequestMethod), http.StatusMethodNotAllowed)
	}
}

func (h *httpHandler) createBrand(w http.ResponseWriter, r *http.Request) {
	var createBrand domain.Brand
	err := json.NewDecoder(r.Body).Decode(&createBrand)
	if err != nil {
		http.Error(w, domain.NewHTTPError(err), http.StatusBadRequest)
		return
	}

	brand, err := h.adminUseCase.CreateBrand(r.Context(), createBrand.Name)
	if err != nil {
		if errors.Is(err, domain.ErrDuplicate) {
			http.Error(w, domain.NewHTTPError(err), http.StatusConflict)
			return
		}

		http.Error(w, domain.NewHTTPError(err), http.StatusInternalServerError)
		return
	}

	response, _ := json.Marshal(&brand)
	_, _ = fmt.Fprintf(w, string(response))
}

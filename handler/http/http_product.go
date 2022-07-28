package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"jamtangan/domain"
)

func (h *httpHandler) Product(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.getProductByID(w, r)
	case http.MethodPost:
		h.createProduct(w, r)
	default:
		http.Error(w, domain.NewHTTPError(domain.ErrInvalidRequestMethod), http.StatusMethodNotAllowed)
	}
}

func (h *httpHandler) ProductBrand(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.fetchProductByBrandID(w, r)
	default:
		http.Error(w, domain.NewHTTPError(domain.ErrInvalidRequestMethod), http.StatusMethodNotAllowed)
	}
}

func (h *httpHandler) createProduct(w http.ResponseWriter, r *http.Request) {
	var createProduct domain.Product
	err := json.NewDecoder(r.Body).Decode(&createProduct)
	if err != nil {
		http.Error(w, domain.NewHTTPError(err), http.StatusBadRequest)
		return
	}

	product, err := h.adminUseCase.CreateProduct(r.Context(), createProduct)
	if err != nil {
		if errors.Is(err, domain.ErrBadRequest) {
			http.Error(w, domain.NewHTTPError(err), http.StatusBadRequest)
			return
		}

		http.Error(w, domain.NewHTTPError(err), http.StatusInternalServerError)
		return
	}

	response, _ := json.Marshal(&product)
	_, _ = fmt.Fprintf(w, string(response))
}

func (h *httpHandler) getProductByID(w http.ResponseWriter, r *http.Request) {
	productIDQuery, ok := r.URL.Query()["id"]
	if !ok || len(productIDQuery) == 0 {
		http.Error(w, domain.NewHTTPError(domain.ErrInvalidID), http.StatusBadRequest)
		return
	}

	productID, _ := strconv.ParseInt(productIDQuery[0], 10, 64)
	if productID == 0 {
		http.Error(w, domain.NewHTTPError(domain.ErrInvalidID), http.StatusBadRequest)
		return
	}

	product, err := h.adminUseCase.GetProductByID(r.Context(), productID)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			http.Error(w, domain.NewHTTPError(err), http.StatusNotFound)
			return
		}

		http.Error(w, domain.NewHTTPError(err), http.StatusInternalServerError)
		return
	}

	response, _ := json.Marshal(&product)
	_, _ = fmt.Fprintf(w, string(response))
}

func (h *httpHandler) fetchProductByBrandID(w http.ResponseWriter, r *http.Request) {
	brandIDQuery, ok := r.URL.Query()["id"]
	if !ok || len(brandIDQuery) == 0 {
		http.Error(w, domain.NewHTTPError(domain.ErrInvalidID), http.StatusBadRequest)
		return
	}

	brandID, _ := strconv.ParseInt(brandIDQuery[0], 10, 64)
	if brandID == 0 {
		http.Error(w, domain.NewHTTPError(domain.ErrInvalidID), http.StatusBadRequest)
		return
	}

	product, err := h.customerUseCase.FetchProductByBrandID(r.Context(), brandID)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			http.Error(w, domain.NewHTTPError(err), http.StatusNotFound)
			return
		}

		http.Error(w, domain.NewHTTPError(err), http.StatusInternalServerError)
		return
	}

	response, _ := json.Marshal(&product)
	_, _ = fmt.Fprintf(w, string(response))
}

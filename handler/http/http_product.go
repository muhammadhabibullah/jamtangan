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

// createProduct
// @Summary      Create new product
// @Description  Create new product
// @Tags         Product
// @Accept       json
// @Produce      json
// @Param 		 request body 	domain.CreateProductRequest true "Create product request body"
// @Success      200  {object}  domain.Product
// @Failure      400  {object}  domain.HTTPError
// @Failure      409  {object}  domain.HTTPError
// @Failure      500  {object}  domain.HTTPError
// @Router       /product [post]
func (h *httpHandler) createProduct(w http.ResponseWriter, r *http.Request) {
	var req domain.CreateProductRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, domain.NewHTTPError(err), http.StatusBadRequest)
		return
	}

	if err = req.Validate(); err != nil {
		http.Error(w, domain.NewHTTPError(err), http.StatusBadRequest)
		return
	}

	product, err := h.adminUseCase.CreateProduct(r.Context(), req.ToProduct())
	if err != nil {
		if errors.Is(err, domain.ErrBadRequest) {
			http.Error(w, domain.NewHTTPError(err), http.StatusBadRequest)
			return
		}

		http.Error(w, domain.NewHTTPError(err), http.StatusInternalServerError)
		return
	}

	response, _ := json.MarshalIndent(&product, "", "    ")
	_, _ = fmt.Fprintf(w, string(response))
}

// getProductByID
// @Summary      Get product by ID
// @Description  Get product by ID
// @Tags         Product
// @Accept       json
// @Produce      json
// @Param 		 id query string true "Product ID"
// @Success      200  {object}  domain.Product
// @Failure      400  {object}  domain.HTTPError
// @Failure      404  {object}  domain.HTTPError
// @Failure      500  {object}  domain.HTTPError
// @Router       /product [get]
func (h *httpHandler) getProductByID(w http.ResponseWriter, r *http.Request) {
	productIDQuery := r.URL.Query().Get("id")
	productID, _ := strconv.ParseInt(productIDQuery, 10, 64)
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

	response, _ := json.MarshalIndent(&product, "", "    ")
	_, _ = fmt.Fprintf(w, string(response))
}

// fetchProductByBrandID
// @Summary      Fetch product by brand ID
// @Description  Fetch product by brand ID
// @Tags         Product
// @Accept       json
// @Produce      json
// @Param 		 id query string true "Brand ID"
// @Success      200  {object}  domain.Product
// @Failure      400  {object}  domain.HTTPError
// @Failure      404  {object}  domain.HTTPError
// @Failure      500  {object}  domain.HTTPError
// @Router       /product/brand [get]
func (h *httpHandler) fetchProductByBrandID(w http.ResponseWriter, r *http.Request) {
	brandIDQuery := r.URL.Query().Get("id")
	brandID, _ := strconv.ParseInt(brandIDQuery, 10, 64)
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

	response, _ := json.MarshalIndent(&product, "", "    ")
	_, _ = fmt.Fprintf(w, string(response))
}

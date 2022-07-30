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

// createBrand
// @Summary      Create new brand
// @Description  Create new brand
// @Tags         Brand
// @Accept       json
// @Produce      json
// @Param 		 request body 	domain.CreateBrandRequest true "Create brand request body"
// @Success      200  {object}  domain.Brand
// @Failure      400  {object}  domain.HTTPError
// @Failure      409  {object}  domain.HTTPError
// @Failure      500  {object}  domain.HTTPError
// @Router       /brand [post]
func (h *httpHandler) createBrand(w http.ResponseWriter, r *http.Request) {
	var req domain.CreateBrandRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, domain.NewHTTPError(err), http.StatusBadRequest)
		return
	}

	if err = req.Validate(); err != nil {
		http.Error(w, domain.NewHTTPError(err), http.StatusBadRequest)
		return
	}

	brand, err := h.adminUseCase.CreateBrand(r.Context(), req.Name)
	if err != nil {
		if errors.Is(err, domain.ErrDuplicate) {
			http.Error(w, domain.NewHTTPError(err), http.StatusConflict)
			return
		}

		http.Error(w, domain.NewHTTPError(err), http.StatusInternalServerError)
		return
	}

	response, _ := json.MarshalIndent(&brand, "", "    ")
	_, _ = fmt.Fprintf(w, string(response))
}

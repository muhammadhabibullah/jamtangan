package domain

import "net/http"

type HTTPHandler interface {
	Health(w http.ResponseWriter, r *http.Request)
	Brand(w http.ResponseWriter, r *http.Request)
	Product(w http.ResponseWriter, r *http.Request)
	ProductBrand(w http.ResponseWriter, r *http.Request)
	Transaction(w http.ResponseWriter, r *http.Request)
}

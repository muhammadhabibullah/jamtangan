package http

import "net/http"

func (h *httpHandler) Health(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write([]byte(`{ "healthy": true }`))
}

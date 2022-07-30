package http_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/onsi/gomega"
)

func TestHttpHandler_Health(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		test func(*GomegaWithT, *httpHandlerTest)
	}{
		{
			name: "healthy",
			test: func(t *GomegaWithT, hh *httpHandlerTest) {
				r := httptest.NewRequest(http.MethodGet, "/health", nil)
				w := httptest.NewRecorder()

				hh.httpHandler.Health(w, r)
				res := w.Result()
				defer res.Body.Close()

				data, err := ioutil.ReadAll(res.Body)
				t.Expect(err).ShouldNot(HaveOccurred())

				expectedData, _ := ioutil.ReadFile("json/health_response.json")
				t.Expect(data).Should(MatchJSON(expectedData))
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

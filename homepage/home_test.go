package homepage

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandlers_Handler(t *testing.T) {
	tests := []struct {
		name           string
		in             *http.Request
		out            *httptest.ResponseRecorder
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "good",
			in:             httptest.NewRequest("GET", "/", nil),
			out:            httptest.NewRecorder(),
			expectedStatus: http.StatusOK,
			expectedBody:   message,
		},
	}

	for _, test := range tests {
		test := test

	}
}

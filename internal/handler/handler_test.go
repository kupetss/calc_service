package handler


import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCalculateHandler(t *testing.T) {
	tests := []struct {
		requestBody string
		expected    string
		statusCode  int
	}{
		{`{"expression": "2+2"}`, `{"result":"4.00"}`, http.StatusOK},
		{`{"expression": "2+"}`, `{"error":"Expression is not valid"}`, http.StatusUnprocessableEntity},
		{`{"expr": "2+2"}`, `{"error":"Internal server error"}`, http.StatusInternalServerError},
	}

	for _, tt := range tests {
		req := httptest.NewRequest(http.MethodPost, "/api/v1/calculate", bytes.NewBufferString(tt.requestBody))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()

		CalculateHandler(rec, req)

		if rec.Code != tt.statusCode {
			t.Errorf("unexpected status code: got %d, expected %d", rec.Code, tt.statusCode)
		}

		if rec.Body.String() != tt.expected {
			t.Errorf("unexpected body: got %s, expected %s", rec.Body.String(), tt.expected)
		}
	}
}

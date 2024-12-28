package handler_test

import (
    "bytes"
    "encoding/json"
    "io/ioutil"
    "net/http"
    "net/http/httptest"
    "reflect"
    "testing"

    "github.com/kupetss/calc_service/internal/handler"
)

func TestCalculateHandler(t *testing.T) {
    tests := []struct {
        name           string
        input          string
        expectedStatus int
        expectedBody   interface{}
    }{
        {
            name:           "valid expression",
            input:          `{"expression": "2+2"}`,
            expectedStatus: http.StatusOK,
            expectedBody:   map[string]interface{}{"result": "4.00"},
        },
        {
            name:           "invalid expression",
            input:          `{"expression": "2//2"}`,
            expectedStatus: http.StatusUnprocessableEntity,
            expectedBody:   map[string]interface{}{"error": "Expression is not valid"},
        },
        {
            name:           "internal server error",
            input:          `{"expression": ""}`,
            expectedStatus: http.StatusInternalServerError,
            expectedBody:   map[string]interface{}{"error": "Internal server error"},
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            req := httptest.NewRequest(http.MethodPost, "/calculate", bytes.NewBufferString(tt.input))
            req.Header.Set("Content-Type", "application/json")

            rec := httptest.NewRecorder()
            handler.CalculateHandler(rec, req)

            if rec.Code != tt.expectedStatus {
                t.Errorf("unexpected status code: got %d, expected %d", rec.Code, tt.expectedStatus)
            }

            var actualBody map[string]interface{}
            body, _ := ioutil.ReadAll(rec.Body)
            if err := json.Unmarshal(body, &actualBody); err != nil {
                t.Fatalf("failed to unmarshal response body: %v", err)
            }

            if !reflect.DeepEqual(actualBody, tt.expectedBody) {
                t.Errorf("unexpected body: got %v, expected %v", actualBody, tt.expectedBody)
            }
        })
    }
}

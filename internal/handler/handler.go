package handler

import (
    "encoding/json"
    "fmt"
    "net/http"
    "github.com/kupetss/calc_service/internal/calculator"
)

type Request struct {
    Expression string `json:"expression"`
}

type Response struct {
    Result string `json:"result,omitempty"`
    Error  string `json:"error,omitempty"`
}

func CalculateHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    var req Request
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        response := Response{Error: "Internal server error"}
        w.WriteHeader(http.StatusInternalServerError)
        json.NewEncoder(w).Encode(response)
        return
    }

    if req.Expression == "" {
        response := Response{Error: "Internal server error"}
        w.WriteHeader(http.StatusInternalServerError)
        json.NewEncoder(w).Encode(response)
        return
    }

    result, err := calculator.Calculate(req.Expression)
    if err != nil {
        response := Response{Error: "Expression is not valid"}
        w.WriteHeader(http.StatusUnprocessableEntity)
        json.NewEncoder(w).Encode(response)
        return
    }

    response := Response{Result: formatResult(result)}
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(response)
}

func formatResult(result float64) string {
    return fmt.Sprintf("%.2f", result)
}

package main

import (
	"fmt"
	"net/http"
	"github.com/kupetss/calc_service/internal/handler"
)

func main() {
	http.HandleFunc("/api/v1/calculate", handler.CalculateHandler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Print(err)
	}
}

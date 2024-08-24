package main

import (
	"fmt"
	"math/rand"
	"net/http"
)

func main() {
	http.HandleFunc("/data", func(w http.ResponseWriter, r *http.Request) {
		if rand.Intn(2) == 0 {
			http.Error(w, "service unavailable", http.StatusServiceUnavailable)
			return
		}
		fmt.Fprintf(w, "success")
	})

	fmt.Println("server running on port 8080")
	http.ListenAndServe(":8080", nil)
}

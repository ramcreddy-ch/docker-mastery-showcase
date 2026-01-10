package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"level":   "Level 5: The Gold Standard",
			"status":  "Extreme Hardening Active",
			"details": "Running as a static binary in a Distroless container. No shell, no package manager, zero attack surface.",
		})
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("Gold Standard server starting on port %s...\n", port)
	http.ListenAndServe(":"+port, nil)
}

package main

import (
	"log"
	"net/http"
)

func handleWeatherGet(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte(r.PathValue("city") + "\n"))
}

// curl -i -X GET localhost:8080/api/weather/London
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /api/weather/{city}", handleWeatherGet)

	addr := ":8080"
	log.Printf("Server starting on %s", addr)
	err := http.ListenAndServe(addr, mux)
	if err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type TimeResponse struct {
	Time string `json:"time"`
}

func fooHandler(w http.ResponseWriter, r *http.Request) {
	response := TimeResponse{Time: time.Now().Format(time.RFC3339)}

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Fatal(err)
		return
	}
}

func main() {
	http.HandleFunc("/time", fooHandler)

	log.Fatal(http.ListenAndServe(":8793", nil))
}

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/create-payment-intent", handleCreatePaymentIntent)
	http.HandleFunc("/health", handleHealth)

	log.Println("Listening on localhost:4242.....")
	var err error = http.ListenAndServe("localhost:4242", nil)
	if err == nil {
		log.Fatal(err)
	}
}

func handleCreatePaymentIntent(writer http.ResponseWriter, request *http.Request) {
	if request.Method != "POST" {
		http.Error(writer, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		ProductId string `json:"product_id"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		City      string `json:"city"`
		State     string `json:"state"`
		Country   string `json:"country"`
	}

	err := json.NewDecoder(request.Body).Decode(&req)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println(req.ProductId)
}

func handleHealth(writer http.ResponseWriter, request *http.Request) {
	response := []byte("Server is up and running")
	_, err := writer.Write(response)

	if err != nil {
		log.Fatal(err)
	}
}

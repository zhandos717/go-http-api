package controllers

import (
	"encoding/json"
	"net/http"
	"strings"
)

func PostPayment(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// Your code to handle POST request for payment
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"message": "PostPayment"})
	} else {
		http.NotFound(w, r)
	}
}

func GetPayments(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		// Your code to handle GET request for payments
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"message": "GetPayments"})
	} else {
		http.NotFound(w, r)
	}
}

func GetPayment(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		// Extract the payment ID from the URL
		id := strings.TrimPrefix(r.URL.Path, "/api/payment/")
		// Your code to handle GET request for a single payment
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"message": "GetPayment", "id": id})
	} else {
		http.NotFound(w, r)
	}
}

func UpdatePayment(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPut {
		// Extract the payment ID from the URL
		id := strings.TrimPrefix(r.URL.Path, "/api/payment/")
		// Your code to handle PUT request for payment
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"message": "UpdatePayment", "id": id})
	} else {
		http.NotFound(w, r)
	}
}

func DeletePayment(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodDelete {
		// Extract the payment ID from the URL
		id := strings.TrimPrefix(r.URL.Path, "/api/payment/")
		// Your code to handle DELETE request for payment
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"message": "DeletePayment", "id": id})
	} else {
		http.NotFound(w, r)
	}
}

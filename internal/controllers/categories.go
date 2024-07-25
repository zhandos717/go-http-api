package controllers

import (
	"encoding/json"
	"net/http"
	"strings"
)

func PostCategory(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// Your code to handle POST request for category
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"message": "PostCategory"})
	} else {
		http.NotFound(w, r)
	}
}

func GetCategories(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		// Your code to handle GET request for categories
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"message": "GetCategories"})
	} else {
		http.NotFound(w, r)
	}
}

func GetCategory(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		// Extract the category ID from the URL
		id := strings.TrimPrefix(r.URL.Path, "/api/category/")
		// Your code to handle GET request for a single category
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"message": "GetCategory", "id": id})
	} else {
		http.NotFound(w, r)
	}
}

func UpdateCategory(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPut {
		// Extract the category ID from the URL
		id := strings.TrimPrefix(r.URL.Path, "/api/category/")
		// Your code to handle PUT request for category
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"message": "UpdateCategory", "id": id})
	} else {
		http.NotFound(w, r)
	}
}

func DeleteCategory(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodDelete {
		// Extract the category ID from the URL
		id := strings.TrimPrefix(r.URL.Path, "/api/category/")
		// Your code to handle DELETE request for category
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"message": "DeleteCategory", "id": id})
	} else {
		http.NotFound(w, r)
	}
}

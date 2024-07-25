package controllers

import (
	"encoding/json"
	"net/http"
	"strings"
)

func PostCategory(w http.ResponseWriter, r *http.Request) {
	// Your code to handle POST request for category
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "PostCategory"})
}

func GetCategories(w http.ResponseWriter, r *http.Request) {
	// Your code to handle GET request for categories
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "GetCategories"})
}

func GetCategory(w http.ResponseWriter, r *http.Request) {
	// Extract the category ID from the URL
	id := strings.TrimPrefix(r.URL.Path, "/api/category/")
	// Your code to handle GET request for a single category
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "GetCategory", "id": id})
}

func UpdateCategory(w http.ResponseWriter, r *http.Request) {
	// Extract the category ID from the URL
	id := strings.TrimPrefix(r.URL.Path, "/api/update/category/")
	// Your code to handle PUT request for category
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "UpdateCategory", "id": id})
}

func DeleteCategory(w http.ResponseWriter, r *http.Request) {
	// Extract the category ID from the URL
	id := strings.TrimPrefix(r.URL.Path, "/api/delete/category/")
	// Your code to handle DELETE request for category
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "DeleteCategory", "id": id})
}

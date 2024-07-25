package controllers

import (
	"encoding/json"
	"net/http"
	"strings"
)

func PostCategory(w http.ResponseWriter, r *http.Request) {
	// Intentional error to test ErrorHandler middleware
	if r.URL.Query().Get("error") == "true" {
		panic("Something went wrong!")
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "PostCategory"})
}

func GetCategories(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "GetCategories"})
}

func GetCategory(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/api/category/")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "GetCategory", "id": id})
}

func UpdateCategory(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/api/update/category/")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "UpdateCategory", "id": id})
}

func DeleteCategory(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/api/delete/category/")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "DeleteCategory", "id": id})
}

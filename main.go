package main

import (
	"net/http"
)

func HomePage(w http.ResponseWriter, r *http.Request) {

}

func main() {
	http.HandleFunc("/", HomePage)
	http.ListenAndServe(":8089", nil)
}

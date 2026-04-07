package handlers

import (
	"fmt"
	"net/http"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {

	// We dont have to handle the httpMethod now because go-chi will handle it
	// if r.Method != http.MethodGet {
	// 	http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	// 	return
	// }

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)


	fmt.Fprint(w, `{"message": "hello world"}`)

}
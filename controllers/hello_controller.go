package controllers

import (
	"net/http"
)

func HelloController(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	w.Write([]byte("E2EE server is running."))
}

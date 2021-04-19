package controllers

import (
	"fmt"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to Raja Kado & Boutique\n"))
	fmt.Fprintf(w, "coba Fprintf")
}
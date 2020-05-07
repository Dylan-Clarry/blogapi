package controllers

import (
	"fmt"
	"net/http"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("controller")
	w.WriteHeader(200)
}
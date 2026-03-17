package http

import (
	"net/http"
)

func Handler(){
	http.HandleFunc("/auth", RegisterHandler) 
	http.ListenAndServe("8080:", nil)
}
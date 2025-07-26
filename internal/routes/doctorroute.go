package routes

import (
	"fmt"
	"net/http"
)

func Registerdocroute(route *http.ServeMux) {
	route.HandleFunc("GET api/v1/doc", Getdocdetails)
}

func Getdocdetails(w http.ResponseWriter, r *http.Request) {
	fmt.Println("############ working doc ##################")
}

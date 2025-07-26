package routes

import (
	"fmt"
	"net/http"
)

func Registerreceproute(route *http.ServeMux) {
	route.HandleFunc("GET api/v1/recep", Getrecepdetails)
}

func Getrecepdetails(w http.ResponseWriter, r *http.Request) {
	fmt.Println("############ working recep ##################")
}

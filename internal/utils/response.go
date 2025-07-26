package utils

import (
	"encoding/json"
	"net/http"
)

func Response(w http.ResponseWriter, status int, data interface{}) error {

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)

}

func Makelikejson(key string, value string) map[string]string {
	m := map[string]string{
		key: value,
	}

	return m

}

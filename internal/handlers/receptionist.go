package handlers

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"Github.com/Gaintlord/hospital_management/internal/models"
	"Github.com/Gaintlord/hospital_management/internal/utils"
)

func Receptionistmux() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var receplogin models.RecepLogin

		err := json.NewDecoder(r.Body).Decode(&receplogin)
		if errors.Is(err, io.EOF) {
			utils.Response(w, http.StatusBadRequest, map[string]string{"message": "empty body recieved"})
		} else {
			utils.Response(w, http.StatusAccepted, &receplogin)

		}
	}
}

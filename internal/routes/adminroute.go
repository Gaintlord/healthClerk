package routes

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"Github.com/Gaintlord/hospital_management/internal/middlerware"
	"Github.com/Gaintlord/hospital_management/internal/models"
	"Github.com/Gaintlord/hospital_management/internal/utils"
	"github.com/golang-jwt/jwt"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Registeradminroute(route *http.ServeMux, db *gorm.DB) {

	route.HandleFunc("POST /api/v1/admin/login", adminLogin(db))
	/*all the routes for doctor*/
	route.HandleFunc("POST /api/v1/admin/doctors", middlerware.Auth(CreateDoc(db)))
	route.HandleFunc("GET /api/v1/admin/doctors/", Getdoc(db))
	route.HandleFunc("GET /api/v1/admin/doctors/{id}", Getdocbyid(db))
	route.HandleFunc("PUT /api/v1/admin/doctors/{id}", Updatedoc(db))
	route.HandleFunc("DELETE /api/v1/admin/doctors/{id}", Deletedoc(db))

	/*for the recptionist*/
	route.HandleFunc("POST /api/v1/admin/recep", Createrecep(db))
	route.HandleFunc("GET /api/v1/admin/recep", Getrecep(db))
	route.HandleFunc("GET /api/v1/admin/recep/{id}", Getrecepbyid(db))
	route.HandleFunc("PUT /api/v1/admin/recep/{id}", Updaterecep(db))
	route.HandleFunc("DELETE /api/v1/admin/recep/{id}", Deleterecep(db))
}

func adminLogin(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var request models.AdninLoginRequest

		err := json.NewDecoder(r.Body).Decode(&request)
		//check the errors
		if errors.Is(err, io.EOF) {
			utils.Response(w, http.StatusBadRequest, map[string]string{"message": "invalid Request"})
		} else if err != nil {
			utils.Response(w, http.StatusBadRequest, map[string]string{"message": "invalid Request"})
		} else {
			var reqadmin models.Admin
			db.First(&reqadmin, "username = ?", request.Username)

			if reqadmin.Username == "" {
				utils.Response(w, http.StatusBadRequest, map[string]string{
					"message": "Invalid username or password",
				})
			}
			err := bcrypt.CompareHashAndPassword([]byte(reqadmin.Password), []byte(request.Password))
			fmt.Println(reqadmin)
			fmt.Println(request.Password)

			if err != nil {
				utils.Response(w, http.StatusBadRequest, map[string]string{
					"message": "Invalid username or password",
				})
			} else {

				token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
					"sub": request.Username,
					"exp": time.Now().Add(time.Hour * 24).Unix(),
				})

				tokenstr, err := token.SignedString([]byte(os.Getenv("SECRET")))
				if err != nil {
					utils.Response(w, http.StatusInternalServerError, utils.Makelikejson("message", "server down"))
				} else {
					cookie := http.Cookie{
						Name:     "Autharization",
						Value:    tokenstr,
						Path:     "/",
						HttpOnly: true,
						Expires:  time.Now().Add(time.Hour * 24),
						Secure:   true,
					}
					http.SetCookie(w, &cookie)
					utils.Response(w, http.StatusAccepted, utils.Makelikejson("message", "welcome back"))
				}
			}
		}
	}

}

// all the doctor logic
func CreateDoc(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("###########im wrapped inside a auth")
	}

}
func Getdoc(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}

}
func Getdocbyid(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}

}
func Updatedoc(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}

}
func Deletedoc(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}

}

// all the receptionist logic
func Createrecep(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}

}
func Getrecep(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}

}
func Getrecepbyid(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}

}
func Updaterecep(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}

}
func Deleterecep(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}

}

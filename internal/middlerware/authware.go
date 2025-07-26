package middlerware

import (
	"fmt"
	"net/http"
	"os"

	"Github.com/Gaintlord/hospital_management/internal/utils"
	"github.com/golang-jwt/jwt/v5"
)

func Auth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("######middleware working#################")
		cookie, err := r.Cookie("Autharization")
		if err != nil {
			utils.Response(w, http.StatusForbidden, utils.Makelikejson("message", "unauthorized"))
			return
		} else {

			token, err := jwt.Parse(cookie.Value, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
				} else {
					return []byte(os.Getenv("SECRET"))
				}
			})
			if err != nil {
				return nil, err
			}
			return token, nil

			fmt.Printf("###########%s###########\n", cookie)
			next.ServeHTTP(w, r)
		}

	}
}

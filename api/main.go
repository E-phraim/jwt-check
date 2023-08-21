package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt"
)

var MySigninKey = []byte(os.Getenv("SECRET_KEY"))

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "TOKEN is Secured")
}

func isAuthorized(endpoint func(http.ResponseWriter, *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header["Token"] != nil {
			token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("Invalid Signing Method")
				}

				aud := "billing.jwtgo.io"
				checkAudience := token.Claims.(jwt.MapClaims).VerifyAudience(aud, false)

				token.Claims.(jwt.MapClaims).VerifyAudience(aud, false)
				if !checkAudience {
					return nil, fmt.Errorf("Invalid seff")
				}

				iss := "jwtgo.io"
				checkIss := token.Claims.(jwt.MapClaims).VerifyIssuer(iss, false)
				if !checkIss {
					return nil, fmt.Errorf(("invalid ISS"))
				}

				return MySigninKey, nil

			})
			if err != nil{
				fmt.Fprintf(w, err.Error())
			}

			if token.Valid{
				endpoint(w, r)
			}
			
		}else{
			fmt.Fprintf(w, "no authorization token provided")
		}
	})
}

func handleRequests() {
	http.Handle("/", isAuthorized(homePage))
	log.Fatal(http.ListenAndServe(":9001", nil))
}

func main() {
	fmt.Printf("jwt check API is running in background")
	handleRequests()
}

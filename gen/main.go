package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

var MySigninKey = []byte(os.Getenv("SECRET_KEY"))

func GetJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["client"] = "kodeforce"
	claims["aud"] = "billing.jwtgo.io"
	claims["iss"] = "jwtgo.io"
	claims["exp"] = time.Now().Add(time.Minute * 1).Unix()

	tokenString, err := token.SignedString(MySigninKey)

	if err != nil {
		fmt.Errorf("Circuit Error, check connection")
		return "", err
	}
	return tokenString, nil

}

func Index(w http.ResponseWriter, r *http.Request) {
	validToken, err := GetJWT()
	fmt.Println(validToken)

	if err != nil {
		fmt.Println("Failed to generate token")
	}
	fmt.Fprint(w, string(validToken))
}

func handleRequests() {
	http.HandleFunc("/", Index)

	log.Fatal(http.ListenAndServe(":8080", nil))
	
}

func main() {
	fmt.Println("jwt Generator is running in background")
	handleRequests()
}

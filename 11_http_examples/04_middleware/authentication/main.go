package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var mySecKey = "MY-SECRET"
var UserInfo = map[string]string{"user1": "password1", "user2": "password2"}

type User struct {
	ID       string `json:"Id"`
	Password string `json:"password"`
}

func hello(w http.ResponseWriter, r *http.Request) {

	var user User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	expectedPwd, ok := UserInfo[user.ID]
	if !ok || expectedPwd != user.Password {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	exp := time.Now().Add(time.Minute * 10)

	claims := jwt.MapClaims{
		"sub": user.ID,
		"exp": exp,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(mySecKey)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Add("token", tokenString)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Token = "+tokenString)
}

func world(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "world")
}

func main() {

	serverMuxA := http.NewServeMux()
	serverMuxA.HandleFunc("/hello", hello)

	serverMuxB := http.NewServeMux()
	serverMuxB.HandleFunc("/world", world)

	go func() {
		http.ListenAndServe("localhost:8081", serverMuxA)
	}()

	http.ListenAndServe("localhost:8082", serverMuxB)
}

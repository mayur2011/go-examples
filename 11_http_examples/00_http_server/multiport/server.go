package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/mux"
)

var mySecKey = []byte("MYSECRET001")
var UserInfo = map[string]string{
	"user1": "password1",
	"user2": "password2",
}

type UserData struct {
	ID       string `json:"id"`
	Password string `json:"password"`
}

type tokenResp struct {
	Token string `json:"token"`
}

// logging middleware code
func loggingMWHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//before executing the handler
		//start := time.Now()
		log.Printf("%s %s", r.Method, r.RequestURI)
		next.ServeHTTP(w, r)
		//after executing the handler
		//log.Printf("Completed %s in %v", r.URL.Path, time.Since(start))
	})
}

func validateJWTToken(w http.ResponseWriter, r *http.Request) (err error) {
	if r.Header["Token"] == nil {
		fmt.Fprintf(w, "can not find token in header")
		return errors.New("TOKEN ERROR")
	}

	// token, err := jwt.ParseWith
	//
	token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("there was an error in parsing")
		}
		return mySecKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			fmt.Fprintf(w, "invalid token")
			return errors.New("TOKEN ERROR")
		}
	}

	if token == nil {
		fmt.Fprintf(w, "invalid token")
		return errors.New("TOKEN ERROR")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		fmt.Fprintf(w, "couldn't parse claims")
		return errors.New("TOKEN ERROR")
	}

	exp := claims["exp"].(float64)
	if int64(exp) < time.Now().Local().Unix() {
		fmt.Fprintf(w, "token expired")
		return errors.New("TOKEN ERROR")
	}

	return nil
}

func ValidateToken(w http.ResponseWriter, r *http.Request) {
	if err := validateJWTToken(w, r); err != nil {
		log.Println("Error while validating JWT Token:", err)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	data := tokenResp{Token: "Valid"}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
	return
}

func generateJWTToken(w http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case "GET":
		fmt.Fprintf(w, "GET method is not supported..!")
		w.WriteHeader(http.StatusBadRequest)
		return

	case "POST":

		var user UserData
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		expectedPwd, ok := UserInfo[user.ID]
		if !ok || expectedPwd != user.Password {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		exp := time.Now().Add(time.Minute * 10).Unix()

		claims := jwt.MapClaims{
			"sub": user.ID,
			"exp": exp,
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, err := token.SignedString(mySecKey)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		data := tokenResp{Token: tokenString}
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(data)
		return
	}
}

func main() {
	fmt.Println("Starting the application...")
	router := mux.NewRouter()
	// regular way to bind middleware for endpoint execution
	/*
		router.Use(loggingMWHandler)
		router.HandleFunc("/GenerateToken", generateJWTToken).Methods("POST")
	*/
	//another way to implement middleware and enpoint function

	//router.Handle("/GenerateToken", loggingMWHandler(http.HandlerFunc(generateJWTToken))).Methods("POST")

	router.Handle("/GenerateToken", loggingMWHandler(http.HandlerFunc(generateJWTToken)))
	router.Handle("/ValidateToken", loggingMWHandler(http.HandlerFunc(ValidateToken))).Methods("POST")

	server := &http.Server{
		Addr:    ":8091",
		Handler: router,
	}
	server.ListenAndServe()
}

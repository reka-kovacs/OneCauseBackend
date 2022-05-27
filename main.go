package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

const (
	PORT = "4200"
)

type JWTData struct {
	// Standard claims are the standard jwt claims from the IETF standard
	// https://tools.ietf.org/html/rfc7519
	jwt.StandardClaims
	CustomClaims map[string]string `json:"custom,omitempty"`
}

type User struct {
	Id       int
	Username string
	Password string
	Token    string
}

var user = User{
	Id:       1,
	Username: "c137@onecause.com",
	Password: "#th@nH@rm#y#r!$100%D0p#",
	Token:    "",
}

var router *mux.Router

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/login", LoginHandler)

	handler := cors.Default().Handler(mux)

	log.Println("Listening for connections on port: ", PORT)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		http.Error(w, "Login failed!", http.StatusUnauthorized)
	}

	var userData map[string]string
	json.Unmarshal(body, &userData)

	hours, minutes, _ := time.Now().Clock()
	if userData["email"] == user.Username && userData["password"] == user.Password && userData["token"] == fmt.Sprintf("%02d%02d", hours, minutes) {
		fmt.Print("valid")
		json.NewEncoder(w).Encode(&user)
	} else {
		http.Error(w, "Login failed!", http.StatusUnauthorized)
	}
}

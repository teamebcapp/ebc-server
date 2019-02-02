package main

import (
	"ebc-server/card"
	postgres "ebc-server/common/db"
	user "ebc-server/user"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	postgres.InitDbConnect()
	defer postgres.PostgresConn.Close()

	// User
	r.HandleFunc("/user", user.GetUser).Methods("GET")
	r.HandleFunc("/users", user.GetUsers).Methods("GET")
	r.HandleFunc("/user", user.PostUser).Methods("POST")
	r.HandleFunc("/user", user.PutUser).Methods("PUT")
	// BC
	r.HandleFunc("/bc", card.GetBusinessCard).Methods("GET")
	r.HandleFunc("/bcs", card.GetBusinessCards).Methods("GET")
	r.HandleFunc("/bc", card.PostBusinessCard).Methods("POST")
	r.HandleFunc("/bc", card.PutBusinessCard).Methods("PUT")

	http.ListenAndServe(":8000", r)
}

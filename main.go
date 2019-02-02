package main

import (
	postgres "ebc-server/common/db"
	user "ebc-server/user"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	postgres.InitDbConnect()
	defer postgres.PostgresConn.Close()

	r.HandleFunc("/user", user.GetUser).Methods("GET")
	r.HandleFunc("/user", user.PostUser).Methods("POST")

	http.ListenAndServe(":8000", r)
}

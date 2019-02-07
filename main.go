package main

import (
	"ebc-server/card"
	postgres "ebc-server/common/db"
	ownerbc "ebc-server/owner-bc"
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
	r.HandleFunc("/user/duplicate", user.GetUserDuplicate).Methods("GET")
	r.HandleFunc("/user/password", user.PutUserPassword).Methods("PUT")
	// BC
	r.HandleFunc("/bc", card.GetBusinessCard).Methods("GET")
	r.HandleFunc("/bcs", card.GetBusinessCards).Methods("GET")
	r.HandleFunc("/bc", card.PostBusinessCard).Methods("POST")
	r.HandleFunc("/bc", card.PutBusinessCard).Methods("PUT")
	// owner_bc
	r.HandleFunc("/owner/bc", ownerbc.GetOwnerBc).Methods("GET")
	r.HandleFunc("/owner/bcs", ownerbc.GetOwnerBcs).Methods("GET")
	r.HandleFunc("/owner/bc", ownerbc.PostOwnerBc).Methods("POST")
	r.HandleFunc("/owner/bc", ownerbc.PutOwnerBc).Methods("PUT")

	http.ListenAndServe(":8000", r)
}

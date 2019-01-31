package user

import (
	"ebc-server/common"
	postgres "ebc-server/common/db"
	"ebc-server/common/utils"
	"encoding/json"
	"log"
	"net/http"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.Header().Add("Access-Control-Allow-Origin", "*")

	var userParam UserParam
	err := json.NewDecoder(r.Body).Decode(&userParam)
	if err != nil {
	}
	resultUser := []User{}
	postgres.PostgresConn.Find(&resultUser)

	result, err := utils.ObjectToJsonByte(resultUser)
	if err != nil {
		log.Fatal(err)
		http.Error(w, err.Error(), 500)
	}
	w.Write(result)

}

func PostUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.Header().Add("Access-Control-Allow-Origin", "*")

	var addUser User
	err := json.NewDecoder(r.Body).Decode(&addUser)
	if err != nil {
		log.Fatal(err)
		http.Error(w, err.Error(), 500)
	}
	postgres.PostgresConn.Create(&addUser)
	//postgres.PostgresConn.Commit()
	result, err := utils.ObjectToJsonByte(common.BaseResult{"success", "200", 1})
	if err != nil {
		log.Fatal(err)
		http.Error(w, err.Error(), 500)
	}
	w.Write(result)
}

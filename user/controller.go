package user

import (
	"ebc-server/common"
	postgres "ebc-server/common/db"
	"ebc-server/common/utils"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/schema"
)

var decoder = schema.NewDecoder()

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.Header().Add("Access-Control-Allow-Origin", "*")
	var userParam UserParam
	// err := json.NewDecoder(r.Form).Decode(&userParam)
	err := decoder.Decode(&userParam, r.URL.Query())
	if err != nil {
		log.Fatal(err)
		http.Error(w, err.Error(), 500)
	}
	// resultUser := []User{}
	var resultUser User
	// postgres.PostgresConn.Find(&resultUser)
	postgres.PostgresConn.Where("user_id = ? AND password = ?", userParam.UserId, userParam.Password).First(&resultUser)
	result, err := utils.ObjectToJsonByte(resultUser)
	if err != nil {
		log.Fatal(err)
		http.Error(w, err.Error(), 500)
	}
	w.Write(result)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.Header().Add("Access-Control-Allow-Origin", "*")
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

func PutUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.Header().Add("Access-Control-Allow-Origin", "*")

	var modifyUser User
	err := json.NewDecoder(r.Body).Decode(&modifyUser)
	if err != nil {
		log.Fatal(err)
		http.Error(w, err.Error(), 500)
	}
	postgres.PostgresConn.Save(&modifyUser)
	//postgres.PostgresConn.Commit()
	result, err := utils.ObjectToJsonByte(common.BaseResult{"success", "200", 1})
	if err != nil {
		log.Fatal(err)
		http.Error(w, err.Error(), 500)
	}
	w.Write(result)
}

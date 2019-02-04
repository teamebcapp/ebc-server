package card

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

func GetBusinessCards(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.Header().Add("Access-Control-Allow-Origin", "*")
	var BCParam BusinessCardParam
	err := decoder.Decode(&BCParam, r.URL.Query())
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), 500)
	}
	var resultUser []BusinessCard
	postgres.PostgresConn.Where("user_id = ?", BCParam.UserId).Find(&resultUser)
	result, err := utils.ObjectToJsonByte(resultUser)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), 500)
	}
	w.Write(result)
}

func GetBusinessCard(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.Header().Add("Access-Control-Allow-Origin", "*")
	var BCParam BusinessCardParam
	err := decoder.Decode(&BCParam, r.URL.Query())
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), 500)
	}
	var resultBC BusinessCard
	postgres.PostgresConn.First(&resultBC, BCParam.BcSeq)
	result, err := utils.ObjectToJsonByte(resultBC)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), 500)
	}
	w.Write(result)
}

func PostBusinessCard(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.Header().Add("Access-Control-Allow-Origin", "*")

	var addBC BusinessCard
	err := json.NewDecoder(r.Body).Decode(&addBC)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), 500)
	}
	postgres.PostgresConn.Create(&addBC)
	result, err := utils.ObjectToJsonByte(common.BaseResult{"success", "200", 1})
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), 500)
	}
	w.Write(result)
}

func PutBusinessCard(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.Header().Add("Access-Control-Allow-Origin", "*")

	var modifyBC BusinessCard
	err := json.NewDecoder(r.Body).Decode(&modifyBC)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), 500)
	}
	postgres.PostgresConn.Save(&modifyBC)
	//postgres.PostgresConn.Commit()
	result, err := utils.ObjectToJsonByte(common.BaseResult{"success", "200", 1})
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), 500)
	}
	w.Write(result)
}

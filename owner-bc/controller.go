package ownerbc

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

func GetOwnerBcs(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.Header().Add("Access-Control-Allow-Origin", "*")
	var ownerBcParam OwnerBcParam
	err := decoder.Decode(&ownerBcParam, r.URL.Query())
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), 500)
	}
	resultOwnerBc := []OwnerBc{}

	if ownerBcParam.OwnerBcSeq != 0 {
		postgres.PostgresConn.Where("owner_bc_seq = ?", ownerBcParam.OwnerBcSeq).Find(&resultOwnerBc)

	} else if ownerBcParam.OwnerUserId != "" {
		postgres.PostgresConn.Where("owner_user_id = ?", ownerBcParam.OwnerUserId).Find(&resultOwnerBc)

	} else {
		http.Error(w, "error parameter", 500)
	}

	result, err := utils.ObjectToJsonByte(resultOwnerBc)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), 500)
	}
	w.Write(result)
}
func GetOwnerBc(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.Header().Add("Access-Control-Allow-Origin", "*")
	var ownerBcParam OwnerBcParam
	err := decoder.Decode(&ownerBcParam, r.URL.Query())
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), 500)
	}
	resultOwnerBc := OwnerBc{}

	if ownerBcParam.OwnerSeq != 0 {
		postgres.PostgresConn.First(&resultOwnerBc, ownerBcParam.OwnerSeq)
	} else {
		http.Error(w, "error parameter", 500)
	}

	result, err := utils.ObjectToJsonByte(resultOwnerBc)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), 500)
	}
	w.Write(result)
}
func PostOwnerBc(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.Header().Add("Access-Control-Allow-Origin", "*")

	var addOwnerBc OwnerBc
	err := json.NewDecoder(r.Body).Decode(&addOwnerBc)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), 500)
	}
	postgres.PostgresConn.Create(&addOwnerBc)
	//postgres.PostgresConn.Commit()
	result, err := utils.ObjectToJsonByte(common.BaseResult{"success", "200", 1})
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), 500)
	}
	w.Write(result)
}

func PutOwnerBc(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.Header().Add("Access-Control-Allow-Origin", "*")

	var modifyOwnerBc OwnerBc
	err := json.NewDecoder(r.Body).Decode(&modifyOwnerBc)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), 500)
	}
	postgres.PostgresConn.Save(&modifyOwnerBc)
	//postgres.PostgresConn.Commit()
	result, err := utils.ObjectToJsonByte(common.BaseResult{"success", "200", 1})
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), 500)
	}
	w.Write(result)
}

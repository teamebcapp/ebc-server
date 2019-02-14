package user

import (
	"ebc-server/common"
	postgres "ebc-server/common/db"
	"ebc-server/common/utils"
	"encoding/json"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"strings"

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
		log.Println(err)
		http.Error(w, err.Error(), 500)
	}
	// resultUser := []User{}
	var resultUser User
	// postgres.PostgresConn.Find(&resultUser)
	postgres.PostgresConn.Where("user_id = ?", userParam.UserId).First(&resultUser)
	resultUser.Password = ""
	result, err := utils.ObjectToJsonByte(resultUser)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), 500)
	}
	w.Write(result)
}

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.Header().Add("Access-Control-Allow-Origin", "*")
	var userParam UserParam
	// err := json.NewDecoder(r.Form).Decode(&userParam)
	err := decoder.Decode(&userParam, r.URL.Query())
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), 500)
	}
	// resultUser := []User{}
	var resultUser User
	// postgres.PostgresConn.Find(&resultUser)
	postgres.PostgresConn.Where("user_id = ? AND password = ?", userParam.UserId, userParam.Password).First(&resultUser)
	resultUser.Password = ""
	result, err := utils.ObjectToJsonByte(resultUser)
	if err != nil {
		log.Println(err)
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
		log.Println(err)
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
		log.Println(err)
		http.Error(w, err.Error(), 500)
	}
	postgres.PostgresConn.Create(&addUser)
	//postgres.PostgresConn.Commit()
	result, err := utils.ObjectToJsonByte(common.BaseResult{"success", "200", 1})
	if err != nil {
		log.Println(err)
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
		log.Println(err)
		http.Error(w, err.Error(), 500)
	}
	postgres.PostgresConn.Save(&modifyUser)
	//postgres.PostgresConn.Commit()
	result, err := utils.ObjectToJsonByte(common.BaseResult{"success", "200", 1})
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), 500)
	}
	w.Write(result)
}

// GetUserDuplicate 아이디 중복체크
func GetUserDuplicate(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.Header().Add("Access-Control-Allow-Origin", "*")
	var userParam UserParam
	// err := json.NewDecoder(r.Form).Decode(&userParam)
	err := decoder.Decode(&userParam, r.URL.Query())
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), 500)
	}
	// resultUser := []User{}
	var resultUser User
	// postgres.PostgresConn.Find(&resultUser)
	postgres.PostgresConn.Where("user_id = ?", userParam.UserId).First(&resultUser)

	duplicate := common.BaseResult{"no duplicate", "200", 0}
	if resultUser.UserSeq != 0 {
		duplicate.ResultMessage = "duplicate"
		duplicate.ResultCount = 1
	}
	result, err := utils.ObjectToJsonByte(duplicate)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), 500)
	}
	w.Write(result)
}

// PutUserPassword 패스워드 변경
func PutUserPassword(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.Header().Add("Access-Control-Allow-Origin", "*")

	var modifyPassword UserPassword
	err := json.NewDecoder(r.Body).Decode(&modifyPassword)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), 500)
	}
	var resultUser User
	postgres.PostgresConn.Where("user_id = ? AND password = ?", modifyPassword.UserId, modifyPassword.PrevPassword).First(&resultUser)

	var result []byte
	if resultUser.UserSeq == 0 {
		log.Println(modifyPassword.UserId, " : 패스워드 변경 실패!!")
		result, err = utils.ObjectToJsonByte(common.BaseResult{"패스워드 확인", "200", 0})
	} else {
		resultUser.Password = modifyPassword.Password
		postgres.PostgresConn.Save(&resultUser)
		result, err = utils.ObjectToJsonByte(common.BaseResult{"success", "200", 1})
		log.Println(resultUser.UserId, " : 패스워드 변경 성공!!")
	}

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), 500)
	}
	w.Write(result)

}

func PostUserFileUpload(w http.ResponseWriter, r *http.Request) {
	var file multipart.File
	var fileHeader *multipart.FileHeader
	var err error
	var uploadedFileName string
	// POST된 파일 데이터를 얻는다

	userId := r.FormValue("userid")
	log.Println("User ID :", userId)
	file, fileHeader, err = r.FormFile("image")
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), 500)
		return
	}
	uploadedFileName = fileHeader.Filename
	fileFormat := strings.Split(uploadedFileName, ".")
	if len(fileFormat) < 1 {
		log.Println("file format err")
		http.Error(w, err.Error(), 500)
		return
	}
	// 서버 측에 저장하기 위해 빈 파일을 만든다
	var saveImage *os.File
	saveImage, err = os.Create("./" + userId + "." + fileFormat[1])
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), 500)
		return
	}
	defer saveImage.Close()
	defer file.Close()
	size, err := io.Copy(saveImage, file)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), 500)
		return
		//os.Exit(1)
	}
	log.Println("쓴 Byte 수=>", size)

	result, _ := utils.ObjectToJsonByte(common.BaseResult{"success", "200", int(size)})
	w.Write(result)
}

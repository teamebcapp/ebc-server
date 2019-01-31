package utils

import (
	"encoding/json"
	"net/http"
)

type Param interface {
}
type Result interface {
}

func JsonToObject(r *http.Request) interface{} {
	// b, err := ioutil.ReadAll(r.Body)
	// if err != nil {
	// 	return nil
	// }

	// // Unmarshal
	// var object interface{}
	// err = json.Unmarshal(b, &object)
	// if err != nil {
	// 	return nil
	// }
	var object interface{}
	json.NewDecoder(r.Body).Decode(&object)
	return object
}

func ObjectToJsonByte(object interface{}) ([]byte, error) {
	json, err := json.Marshal(object)
	if err != nil {
		return nil, err
	}
	return json, err
}

func ObjectToJsonString(object interface{}) (string, error) {
	json, err := json.Marshal(object)
	jsonString := string(json)
	if err != nil {
		return jsonString, err
	}
	return jsonString, err
}

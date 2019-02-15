package interceptor

import (
	"ebc-server/auth"
	"log"
	"net/http"
)

type HttpMethod func(w http.ResponseWriter, r *http.Request)

type Login struct {
	method HttpMethod
}

func SetMethod(method HttpMethod) *Login {
	result := Login{}
	result.method = method
	return &result
}
func (l *Login) Execute(w http.ResponseWriter, r *http.Request) {
	// before
	token := r.Header.Get("access_token")
	if token == "" {
		log.Println("access_token null")
		http.Error(w, "access_token null", 403)
		return
	} else {
		if isValid, err := auth.ValidToken(token); isValid == false {
			log.Println(err)
			http.Error(w, err, 403)
			return
		}
	}

	l.method(w, r)
	//after
}

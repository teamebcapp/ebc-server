package interceptor

import "net/http"

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
	l.method(w, r)
	//after
}

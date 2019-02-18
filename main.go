package main

import (
	"ebc-server/card"
	postgres "ebc-server/common/db"
	"ebc-server/common/interceptor"
	ownerbc "ebc-server/owner-bc"
	user "ebc-server/user"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	postgres.InitDbConnect()
	defer postgres.PostgresConn.Close()

	//setAop()
	// login
	r.HandleFunc("/token/valid", user.InvalidToken).Methods("GET")
	r.HandleFunc("/login", user.Login).Methods("GET")
	// User
	r.HandleFunc("/user", interceptor.SetMethod(user.GetUser).Execute).Methods("GET")
	r.HandleFunc("/users", interceptor.SetMethod(user.GetUsers).Execute).Methods("GET")
	r.HandleFunc("/user", interceptor.SetMethod(user.PostUser).Execute).Methods("POST")
	r.HandleFunc("/user", interceptor.SetMethod(user.PutUser).Execute).Methods("PUT")
	r.HandleFunc("/user/duplicate", interceptor.SetMethod(user.GetUserDuplicate).Execute).Methods("GET")
	r.HandleFunc("/user/password", interceptor.SetMethod(user.PutUserPassword).Execute).Methods("PUT")
	r.HandleFunc("/user/profile", interceptor.SetMethod(user.PostUserFileUpload).Execute).Methods("POST")
	// BC
	r.HandleFunc("/bc", interceptor.SetMethod(card.GetBusinessCard).Execute).Methods("GET")
	r.HandleFunc("/bcs", interceptor.SetMethod(card.GetBusinessCards).Execute).Methods("GET")
	r.HandleFunc("/bc", interceptor.SetMethod(card.PostBusinessCard).Execute).Methods("POST")
	r.HandleFunc("/bc", interceptor.SetMethod(card.PutBusinessCard).Execute).Methods("PUT")
	// owner_bc
	r.HandleFunc("/owner/bc", interceptor.SetMethod(ownerbc.GetOwnerBc).Execute).Methods("GET")
	r.HandleFunc("/owner/bcs", interceptor.SetMethod(ownerbc.GetOwnerBcs).Execute).Methods("GET")
	r.HandleFunc("/owner/bc", interceptor.SetMethod(ownerbc.PostOwnerBc).Execute).Methods("POST")
	r.HandleFunc("/owner/bc", interceptor.SetMethod(ownerbc.PutOwnerBc).Execute).Methods("PUT")
	r.HandleFunc("/owner/bc", interceptor.SetMethod(ownerbc.DeleteOwnerBc).Execute).Methods("DELETE")

	http.ListenAndServe(":8000", r)
}

// func setAop() {

// 	beanFactory := aop.NewClassicBeanFactory()
// 	beanFactory.RegisterBean("auth", new(myAop.LoginAop))
// 	aspect := aop.NewAspect("aspect_1", "auth")
// 	aspect.SetBeanFactory(beanFactory)
// 	pointcut := aop.NewPointcut("pointcut_1").Execution(`Login()`)
// 	aspect.AddPointcut(pointcut)

// 	aspect.AddAdvice(&aop.Advice{Ordering: aop.Before, Method: "Before", PointcutRefID: "pointcut_1"})
// 	aspect.AddAdvice(&aop.Advice{Ordering: aop.After, Method: "After", PointcutRefID: "pointcut_1"})
// 	aspect.AddAdvice(&aop.Advice{Ordering: aop.Around, Method: "Around", PointcutRefID: "pointcut_1"})

// 	gogapAop := aop.NewAOP()
// 	gogapAop.SetBeanFactory(beanFactory)
// 	gogapAop.AddAspect(aspect)
// 	proxy, _ := gogapAop.GetProxy("auth")

// 	proxy.Method(new(myAop.LoginAop))
// }

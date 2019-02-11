package myAop

import (
	"fmt"

	"github.com/gogap/aop"
)

type LoginAop struct {
}

// use join point to get Args from real method
func (p *LoginAop) Before(jp aop.JoinPointer) {
	username := ""
	jp.Args().MapTo(func(u, p string) {
		username = u
	})

	fmt.Printf("Before Login: %s\n", username)
}

// the args is same as Login
func (p *LoginAop) After(username, password string) {
	fmt.Printf("After Login: %s %s\n", username, password)
}

// use join point to around the real func of login
func (p *LoginAop) Around(pjp aop.ProceedingJoinPointer) {
	fmt.Println("@Begin Around")

	ret := pjp.Proceed("fakeName", "fakePassword")
	ret.MapTo(func(loginResult bool) {
		fmt.Println("@Proceed Result is", loginResult)
	})

	fmt.Println("@End Around")
}

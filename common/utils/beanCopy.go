package utils

import (
	"log"
	"reflect"
)

// BeanCopy 동일한 변수 이름 일시 복사
func BeanCopy(src interface{}, dest interface{}) {
	srcElements := reflect.ValueOf(src).Elem()
	destElements := reflect.ValueOf(dest).Elem()

	for i := 0; i < srcElements.NumField(); i++ {
		srcValue := srcElements.Field(i)
		srcType := srcElements.Type().Field(i)
		//srcTag := srcType.Tag  //`gorm:`, `json:` ....
		for j := 0; j < destElements.NumField(); j++ {
			destValue := destElements.Field(j)
			destType := destElements.Type().Field(j)

			if srcType.Name == destType.Name {
				destValue.Set(srcValue)
				break
			}
		}
	}
}

// PrintBeans 모든 멤버 프린트
func PrintBeans(src interface{}) {
	srcElements := reflect.ValueOf(src).Elem()

	for i := 0; i < srcElements.NumField(); i++ {
		srcValue := srcElements.Field(i)
		srcType := srcElements.Type().Field(i)
		log.Println(srcType.Name, " : ", srcValue.Interface())
	}

}

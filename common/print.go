package common

import (
	"fmt"
	"reflect"
)

func PrintStruct(t interface{}) {
	s := reflect.ValueOf(t)
	typeOfT := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("%d: %s %s = %v\n", i,
			typeOfT.Field(i).Name, f.Type(), f.Interface())
	}
}

func Convent(src interface{}, dest interface{}) {
	srcVal := reflect.ValueOf(src).Elem()
	srcType := srcVal.Type()
	destVal := reflect.ValueOf(dest).Elem()
	destType := destVal.Type()
	for i := 0; i < srcVal.NumField(); i++ {
		for j := 0; j < destVal.NumField(); j++ {
			if srcType.Field(i).Name == destType.Field(j).Name {
				destVal.Field(j).Set(srcVal.Field(i))
			}
		}
	}
}

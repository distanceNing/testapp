package types

import (
	"fmt"
	"github.com/pkg/errors"
	"reflect"
	"strings"
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

type Options struct {
	IgnoreCase bool
}

func Convent(src interface{}, dest interface{}, opts *Options) error {
	srcVal := reflect.ValueOf(src).Elem()
	srcType := srcVal.Type()
	destVal := reflect.ValueOf(dest).Elem()
	destType := destVal.Type()
	if srcType.Kind() != reflect.Struct || destType.Kind() != reflect.Struct {
		return errors.New("only support struct type")
	}

	for i := 0; i < srcVal.NumField(); i++ {
		for j := 0; j < destVal.NumField(); j++ {
			//if destVal.Field(j).Type() != srcVal.Field(i).Type() {
			//	continue
			//}
			name1 := srcType.Field(i).Name
			name2 := destType.Field(j).Name
			if opts != nil && opts.IgnoreCase {
				name1 = strings.ToUpper(name1)
				name2 = strings.ToUpper(name2)
			}
			if name1 == name2 {
				v := srcVal.Field(i).Convert(destVal.Field(j).Type())
				destVal.Field(j).Set(v)
			}
		}
	}
	return nil
}

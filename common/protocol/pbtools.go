package protocol

import (
	"fmt"
	"github.com/distanceNing/testapp/pb"
	"github.com/golang/protobuf/jsonpb"
	"github.com/pkg/errors"
	"reflect"
)

func PbToJson(pb interface{}) {
	s := reflect.ValueOf(pb)
	typeOfT := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("%d: %s %s = %v\n", i,
			typeOfT.Field(i).Name, f.Type(), f.Interface())
	}
}

func test() error {
	v := pb.Student{}
	str := `{"name": "xxx","male":true,"scores":[1,2,3]}`

	err := jsonpb.UnmarshalString(str, &v)
	if err != nil {

	}

	if v.Name != "xxx" {
		return errors.New("err")
	}
	return nil
}

package types

import (
	"time"
)

type Rsp struct {
	v map[string]interface{}
}

func NewRsp() *Rsp {
	return &Rsp{make(map[string]interface{})}
}

func (rsp *Rsp) GetV() map[string]interface{} {
	return rsp.v
}

func (rsp *Rsp) Set(k string, v interface{}) {
	rsp.v[k] = v
}

func (rsp *Rsp) SetV(k string, v interface{}) {
	rsp.v[k] = v
}

func StrToTime(str string) (time.Time, error) {
	return time.Parse("2006-01-02 15:04:05", str)
}

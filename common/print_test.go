package common

import (
	"testing"
)

func TestPrintStruct(t *testing.T) {
	type args struct {
		t interface{}
	}
	type Test struct {
		Name string
	}
	tests := []struct {
		name string
		args args
	}{
		{"base", args{Test{"name"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			PrintStruct(tt.args.t)
		})
	}
}

func TestConvent(t *testing.T) {
	type T1 struct {
		Name string
		Id   string
		YY   string
	}
	type T2 struct {
		Name string
		Id   string
		XX   string
	}
	t1 := T1{"1", "2", "yy"}
	t2 := T2{}
	Convent(&t1, &t2)

	if t1.Name != t2.Name || t1.Id != t2.Id {
		t.Error(t1)
		t.Error(t2)
	}
}

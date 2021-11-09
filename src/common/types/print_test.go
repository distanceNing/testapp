package types

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
		Xx   int8
	}
	type T2 struct {
		Name string
		Id   string
		XX   int
	}
	t1 := T1{"1", "2", 11}
	t2 := T2{}
	_ = Convent(&t1, &t2, nil)
	if t1.Name != t2.Name || t1.Id != t2.Id {
		t.Error(t1)
		t.Error(t2)
	}

	_ = Convent(&t1, &t2, &Options{true})
	if t1.Name != t2.Name || t1.Id != t2.Id || int(t1.Xx) != t2.XX {
		t.Error(t1)
		t.Error(t2)
	}
	t3 := 2
	err := Convent(&t1, &t3, &Options{true})
	if err != nil {
		t.Error(err)
	}
}

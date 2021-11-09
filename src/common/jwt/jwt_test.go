package jwt

import (
	"github.com/distanceNing/testapp/src/common/crypto"
	"github.com/distanceNing/testapp/src/common/parser"
	"io"
	"strings"
	"testing"
)

var keyConf = "key.conf.yaml"

type KeyConf struct {
	PrivateKey1 string `yaml:"privite_key1"`
	PrivateKey2 string `yaml:"privite_key2"`
	CompleteKey string `yaml:"complete_key"`
	AppleKeyId  string `yaml:"applekey_id"`
}

func TestJWTGenerator_Get(t *testing.T) {
	type args struct {
		p *JWTParams
	}
	k := KeyConf{}
	err := parser.ParseYamlFile(keyConf, &k)
	if err != nil {
		return
	}
	priKey := crypto.GenerateECKeyFromString(k.PrivateKey2)
	p := JWTParams{"a", k.AppleKeyId, priKey, "com.tencent.xin", ""}

	ag := NewAppleJWTGenerator()
	get, err := ag.Get(&p)
	if err == nil {
		t.Errorf("Get() error = %v, wantErr %v", err, true)
		return
	}

	p.IssuerId = "69a6de6e-7b3e-47e3-e053-5b8c7c11a4d1"
	get, err = ag.Get(&p)
	if err != nil {
		t.Errorf("Get() error = %v, wantErr %v", err, false)
		return
	}
	strs := strings.Split(get, ".")
	if len(strs) != 3 {
		t.Errorf("Get() strs len is %d , want 3 ", len(strs))
		return
	}
	t.Errorf("token : %s", get)

	//tests := []struct {
	//	name    string
	//	args    args
	//	want    string
	//	wantErr bool
	//}{
	//	{"return empty token", args{&p}, "", false},
	//}

	//g := NewJWTGenerator()
	//for _, tt := range tests {
	//	t.Run(tt.name, func(t *testing.T) {
	//		got, err := g.Get(tt.args.p)
	//		if (err != nil) != tt.wantErr {
	//			t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
	//			return
	//		}
	//		if got != tt.want {
	//			t.Errorf("Get() got = %v, want %v", got, tt.want)
	//		}
	//	})
	//}
}

func TestParseECPrivateKeyFromPEM(t *testing.T) {
	type args struct {
		r io.Reader
	}
	k := KeyConf{}
	err := parser.ParseYamlFile(keyConf, &k)
	if err != nil {
		return
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"base", args{strings.NewReader(k.CompleteKey)}, false}, // TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := crypto.ParseECPrivateKeyFromPEM(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseECPrivateKeyFromPEM() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

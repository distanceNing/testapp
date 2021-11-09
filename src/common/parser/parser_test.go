package parser

import "testing"

func TestParseYamlFile(t *testing.T) {
	type St struct {
		Key1 string `yaml:"key1"`
		Key2 string `yaml:"key2"`
		Key3 string `yaml:"key3"`
		Id   string `yaml:"id"`
	}
	v := St{}
	type args struct {
		path string
		v    interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"base", args{"test.yaml", &v}, false},
		{"file not exist", args{"test1.yaml", &v}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ParseYamlFile(tt.args.path, tt.args.v)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseYamlFile() error = %v, wantErr %v", err, tt.wantErr)
			}

			if err == nil {
				if v.Key1 != "1" || v.Key2 != "2" || v.Id != "4" {
					t.Errorf("ParseYamlFile() error = %v", v)
				}
			}

		})
	}
}

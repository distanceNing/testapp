package protocol

import "testing"

func Test_test(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{"a", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := test(); (err != nil) != tt.wantErr {
				t.Errorf("test() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

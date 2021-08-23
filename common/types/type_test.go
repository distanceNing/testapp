package types

import "testing"

func TestStrToTime(t *testing.T) {
	formatTimeStr := "2017-04-21 13:33:37"
	tests := []struct {
		name string
	}{
		{"base"}, // TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			StrToTime(formatTimeStr)
		})
	}
}

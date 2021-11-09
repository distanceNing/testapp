package logic

import (
	"testing"
)

func TestNewServiceController(t *testing.T) {
	ctl := NewServiceController()

	req := Request{}
	rsp := Response{}
	if err := ctl.Proc(&req, &rsp); err != nil {
		t.Errorf("controller proc failed")
	}

}

package consul_test

import (
	"github.com/agitdevcenter/gopkg/consul"
	"testing"
)

func TestGenerate(t *testing.T) {
	_ = consul.NewAgent()
}

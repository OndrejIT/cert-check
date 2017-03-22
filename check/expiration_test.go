package check

import (
	"bytes"
	conf "github.com/spf13/viper"
	"testing"
)

var yamlConf = []byte(`
expires: 0
`)

func init() {
	conf.SetConfigType("yaml")
	conf.ReadConfig(bytes.NewBuffer(yamlConf))
}

func TestExpCheck(t *testing.T) {
	exp_msq := expCheck("github.com")
	if len(exp_msq) != 0 {
		t.Error("Expiration check error.", exp_msq)
	}
}

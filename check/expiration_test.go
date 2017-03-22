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
	if check := expCheck("github.com"); check != nil {
		t.Error("Expiration check error.", check)
	}
}

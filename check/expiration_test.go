package check

import (
	"testing"
	"bytes"
	conf "github.com/spf13/viper"
)

var yamlConf = []byte(`
notofy:
  expires: 10000
`)

func init() {
	conf.SetConfigType("yaml")
	conf.ReadConfig(bytes.NewBuffer(yamlConf))
}

func TestExpCheck(t *testing.T) {
	if check := expCheck("github.com"); check == nil {
		t.Error("Expiration check error.", check)
	}
}

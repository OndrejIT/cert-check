package config

import (
	flag "github.com/spf13/pflag"
	conf "github.com/spf13/viper"
)

func FlagParser() {
	flag.BoolP("debug", "d", false, "Set debug mode.")
	flag.StringP("config", "c", ".", "Set config path.")
	flag.Parse()
	flagToConfig()
}

func flagToConfig() {
	conf.BindPFlag("debug", flag.Lookup("debug"))
	conf.BindPFlag("config", flag.Lookup("config"))
}

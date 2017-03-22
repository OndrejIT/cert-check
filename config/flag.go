package config

import (
	flag "github.com/spf13/pflag"
	conf "github.com/spf13/viper"
)

var hosts []string

func FlagParser() {
	flag.StringP("config", "c", ".", "Set config path.")
	flag.BoolP("debug", "d", false, "Enable debug mode.")
	flag.BoolP("verbose", "v", false, "Enable verbose mode.")
	flag.StringSliceVarP(&hosts, "address", "a", []string{}, "Set check address (multiple).")
	flag.IntP("expires", "e", 14, "Set expires alert (day).")
	flag.Parse()
	flagToConfig()
}

func flagToConfig() {
	conf.BindPFlag("config", flag.Lookup("config"))
	conf.BindPFlag("debug", flag.Lookup("debug"))
	conf.BindPFlag("verbose", flag.Lookup("verbose"))
	conf.Set("fHosts", hosts)
	conf.BindPFlag("expires", flag.Lookup("expires"))
}

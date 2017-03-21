package config

import (
	log "github.com/Sirupsen/logrus"
	conf "github.com/spf13/viper"
)

func Setup() {
	FlagParser()

	if conf.GetBool("debug") {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetFormatter(&log.JSONFormatter{})
	}

	conf.SetConfigName("config")
	conf.AddConfigPath(conf.GetString("config"))

	conf.SetEnvPrefix("")
	conf.AutomaticEnv()

	if err := conf.ReadInConfig(); err != nil {
		log.Fatal(err)
	}

}

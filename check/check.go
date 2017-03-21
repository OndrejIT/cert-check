package check

import (
	log "github.com/Sirupsen/logrus"
	conf "github.com/spf13/viper"
)

func Check() {
	hosts := conf.GetStringSlice("hosts")
	fHosts := conf.GetStringSlice("fHosts")
	for _, host := range fHosts {
		hosts = append(hosts, host)
	}

	for host := range hosts {
		if err := expCheck(hosts[host]); err != nil {
			log.Warn(err)
		}
	}
}

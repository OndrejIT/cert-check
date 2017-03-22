package check

import (
	"fmt"
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
		exp_msq := expCheck(hosts[host])
		if len(exp_msq) != 0 {
			log.Warn(fmt.Sprintf("------ %s ------", hosts[host]))
		}

		for msq := range exp_msq {
			log.Warn(exp_msq[msq])
		}
	}
}

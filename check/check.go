package check

import (
	"crypto/tls"
	"errors"
	"fmt"
	log "github.com/Sirupsen/logrus"
	conf "github.com/spf13/viper"
	"time"
)

func expCheck(host string) error {
	expires := conf.GetInt("notofy.expires")
	conn, err := tls.Dial("tcp", host, nil)
	if err != nil {
		return err
	}
	defer conn.Close()

	time_now := time.Now()
	for _, certs := range conn.ConnectionState().VerifiedChains {
		expire := int(certs[0].NotAfter.Sub(time_now).Hours() / 24)
		if expire < expires {
			return errors.New(fmt.Sprintf("Host: %s, Expires: %d days, Expires on: %s",
				certs[0].Subject.CommonName,
				expire,
				certs[0].NotAfter),
			)

		}

	}
	return nil
}

func Check() {
	hosts := conf.GetStringSlice("hosts")
	for host := range conf.GetStringSlice("hosts") {
		if err := expCheck(hosts[host] + ":443"); err != nil {
			log.Error(err)
		}
	}
}

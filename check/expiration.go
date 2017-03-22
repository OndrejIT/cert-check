package check

import (
	"crypto/tls"
	"fmt"
	log "github.com/Sirupsen/logrus"
	conf "github.com/spf13/viper"
	"time"
)

func expCheck(host string) []string {
	expires := conf.GetInt("expires")
	conn, err := tls.Dial("tcp", host+":443", nil)
	if err != nil {
		return []string{err.Error()}
	}
	defer conn.Close()

	if conf.GetBool("verbose") {
		log.Info(fmt.Sprintf("------ %s ------", host))
	}

	time_now := time.Now()
	var exp_msg []string
	for _, certs := range conn.ConnectionState().VerifiedChains {
		for number, cert := range certs {
			expire := int(cert.NotAfter.Sub(time_now).Hours() / 24)
			if conf.GetBool("verbose") {
				log.Info(fmt.Sprintf("Cert: %d, Name: %s, Expires: %d days, Expires on: %s",
					number,
					cert.Subject.CommonName,
					expire,
					cert.NotAfter),
				)
			}
			if expire < expires {
				exp_msg = append(exp_msg, fmt.Sprintf("Cert: %d, Name: %s, Expires: %d days, Expires on: %s",
					number,
					cert.Subject.CommonName,
					expire,
					cert.NotAfter),
				)

			}
		}

	}
	return exp_msg
}

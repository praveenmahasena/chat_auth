package helpers

import (
	"errors"
	"fmt"
	"net/smtp"
	"sync"

	"github.com/spf13/viper"
)

var (
	c = map[string]string{
		"URL":       "",
		"FROM":      "",
		"EPASSWORD": "",
	}
)

func Mailer(JWT, mailID string, wg *sync.WaitGroup, ch chan<- error) {
	defer wg.Done()
	defer close(ch)

	for e := range c {
		val := viper.GetString(e)
		if val == "" {
			ch <- errors.New("error during reading config file")
			return
		}
		c[e] = val
	}

	to := []string{mailID}
	smtpHost, smtpPort := "smtp.gmail.com", "587"

	auth := smtp.PlainAuth("", c["FROM"], c["EPASSWORD"], smtpHost)

	M := fmt.Sprintf("To: %v\r\n"+

		"Subject: Verify your EmailID\r\n"+

		"\r\n"+

		"%v\r\n", c["FROM"], c["URL"]+JWT)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, c["FROM"], to, []byte(M))

	if err != nil {
		ch <- err
		return
	}

}

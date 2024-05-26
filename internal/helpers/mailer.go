package helpers

import (
	"fmt"
	"net/smtp"
	"sync"

	"github.com/spf13/viper"
)

const urlParam = "URL"
const fromParam = "FROM"
const ePasswordParam = "EPASSWORD"

func Mailer(JWT, mailID string, wg *sync.WaitGroup, ch chan<- error) {
	defer wg.Done()
	defer close(ch)

	authUrl := viper.GetString(urlParam)
	from := viper.GetString(fromParam)
	password := viper.GetString(ePasswordParam)

	to := []string{mailID}
	smtpHost, smtpPort := "smtp.gmail.com", "587" // TODO: change these and get them into .env

	auth := smtp.PlainAuth("", from, password, smtpHost)

	M := fmt.Sprintf("To: %v\r\n"+

		"Subject: Verify your EmailID\r\n"+

		"\r\n"+

		"%v\r\n", from, authUrl+JWT)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, []byte(M))

	if err != nil {
		ch <- err
		return
	}

}

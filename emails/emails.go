package emails

import (
	"net/smtp"
	"log"
	"os"
)

// Email is placeholder for 
type Email struct {
	Recipient 		[]string
	Sender 			string
	Body 			string
}



// SendEmail to reciepient
func (e *Email)	SendEmail() {
	smtpHost,_  := os.LookupEnv("SMTP_SERVER")
	smtpPort,_  := os.LookupEnv("SMTP_PORT")
	username,_  := os.LookupEnv("SMTP_LOGIN")
	passw,_     := os.LookupEnv("SMTP_PASSW")

	auth := smtp.PlainAuth("", username, passw, smtpHost)
	if err := smtp.SendMail(smtpHost+":"+smtpPort, auth, username, e.Recipient, []byte(e.Body)); err != nil {
		log.Println(err)
	}
}
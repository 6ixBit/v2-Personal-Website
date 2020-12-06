package emails

import (
	"fmt"
	"net/smtp"
	"log"
)

// Email is placeholder for 
type Email struct {
	SMTPServerAddr 	string
	Recipient 		string
	Sender 			string
	Body 			string
}

// SendEmail to reciepient
func (e *Email)	SendEmail() {
	conn, err := smtp.Dial(e.SMTPServerAddr)
	if err != nil { log.Println(err) }

	// Set sender and recipient
	if err := conn.Mail(e.Sender); err != nil { log.Println(err) }
	if err := conn.Rcpt(e.Recipient); err != nil { log.Println(err) }

	// Define closer to attach body to
	w, err := conn.Data()
	if err != nil { log.Println(err) }

	// Stream data to email writer
	fmt.Fprintf(w, e.Body)
	fmt.Println("Email succesfuly sent.")

	if err := w.Close(); err != nil { log.Println(err) }
	if	err := conn.Close(); err != nil { log.Println(err)}
}
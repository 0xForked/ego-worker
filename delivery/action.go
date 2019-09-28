package delivery

import (
	"crypto/tls"
	"fmt"
	"github.com/aasumitro/ego-worker/data"
	"github.com/aasumitro/ego-worker/helper"
	"log"
	"net"
	"net/smtp"
)

const (
	ACCOUNT = "account"
)

func ToQueue(msg data.Message) {
	err := data.StoreOutbox(msg)
	helper.CheckError(err, "Failed store queue: ")
}

func ToEmail(msg data.Message, template string, config *helper.EmailDefaultConfig) {
	// Setup headers
	headers := make(map[string]string)
	headers["From"] = config.User
	headers["To"] = msg.TO
	headers["Subject"] = msg.Subject
	// Setup message
	message := ""
	for k, v := range headers {
		message += fmt.Sprintf("%s: %s\n", k, v)
	}
	message += "Content-Type: text/html"
	message += "\r\n" + template

	// Connect to the SMTP Server
	servername := fmt.Sprintf("%s:%s", config.Host, config.Port)
	host, _, _ := net.SplitHostPort(servername)
	auth := smtp.PlainAuth("", config.User, config.Pass, host)

	// TLS config
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         host,
	}

	// Here is the key, you need to call tls.Dial instead of smtp.Dial
	// for smtp servers running on 465 that require an ssl connection
	// from the very beginning (no starttls)
	conn, err := tls.Dial("tcp", servername, tlsConfig)
	if err != nil {
		log.Panic(err)
	}

	c, err := smtp.NewClient(conn, host)
	if err != nil {
		log.Panic(err)
	}

	// Auth
	if err = c.Auth(auth); err != nil {
		log.Panic(err)
	}

	// To && From
	if err = c.Mail(config.User); err != nil {
		log.Panic(err)
	}

	if err = c.Rcpt(msg.TO); err != nil {
		log.Panic(err)
	}

	// Data
	w, err := c.Data()
	if err != nil {
		log.Panic(err)
	}

	_, err = w.Write([]byte(message))
	if err != nil {
		log.Panic(err)
	}

	err = w.Close()
	if err != nil {
		log.Panic(err)
	}

	fmt.Printf("Sent message to %s!\n", msg.TO)

	data.MoveOutboxToSent(msg)

	c.Quit()
}

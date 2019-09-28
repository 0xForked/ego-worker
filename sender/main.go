package main

import (
	"bytes"
	"fmt"
	"github.com/aasumitro/ego-worker/data"
	"github.com/aasumitro/ego-worker/delivery"
	"github.com/aasumitro/ego-worker/helper"
	"html/template"
	"log"
	"strings"
	"time"
)

type TemplateData struct {
	Subject string
	Message string
}

func main() {
	log.Printf(" [*] Waiting for messages to send. To exit press CTRL+C")
	workerSchedule(15*time.Second, workerTask)
}

func workerSchedule(d time.Duration, f func(time.Time)) {
	for x := range time.Tick(d) {
		f(x)
	}
}

func workerTask(t time.Time) {
	config := helper.GetConfig().EmailDefault
	outbox := data.FetchOutbox()
	if outbox != nil {
		for _, msg := range outbox {
			// get email template
			tpl := loadEmailTemplate(msg)
			// do action to send email
			go delivery.ToEmail(msg, tpl, config)
		}
	}
}

func loadEmailTemplate(msg data.Message) string {
	tmplFile := fmt.Sprintf("./sender/template/%s_message.html", strings.ToLower(msg.Template))
	tmpl, err := template.ParseFiles(tmplFile)
	helper.CheckError(err, "Failed load template")
	var tpl bytes.Buffer
	if err := tmpl.Execute(&tpl, TemplateData{
		Subject: msg.Subject,
		Message: msg.Message,
	}); err != nil {
		log.Println(err)
	}
	return tpl.String()
}

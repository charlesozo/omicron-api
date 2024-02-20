package main

import (
	"bytes"
	"gopkg.in/gomail.v2"
	"text/template"
	"time"
)

func sendMail(templatePath, receiverEmail, apiKey string) error {
	var body bytes.Buffer
	t, err := template.ParseFiles(templatePath)
	if err != nil {

		return err
	}
	data := map[string]string{"USER_API_KEY": apiKey}

	if err := t.Execute(&body, data); err != nil {

		return err
	}
	//send with gomail
	m := gomail.NewMessage()
	m.SetHeader("From", "noreply@omicronwabot.netlify.app")
	m.SetHeader("To", receiverEmail)
	m.SetHeader("Subject", "Activation email from omicronwabot.netlify.app")
	m.SetAddressHeader("Cc", receiverEmail, "")
	m.SetDateHeader("X-Date", time.Now())
	m.SetBody("text/html", body.String())
	d := gomail.NewDialer("smtp.gmail.com", 465, "omicronwabot@gmail.com", "yortcnypogaicury")
	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}

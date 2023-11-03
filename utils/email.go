package utils

import (
	"bytes"
	"crypto/tls"
	//"github.com/k3a/html2text"
	"github.com/spf13/cast"
	"gopkg.in/gomail.v2"
	"html/template"
	"os"
	"path/filepath"
)

type EmailData struct {
	URL       string
	FirstName string
	Subject   string
}

func ParseTemplateDir(dir string) (*template.Template, error) {
	var paths []string

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			paths = append(paths, path)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return template.ParseFiles(paths...)
}

func SendMail(email string, data *EmailData) error {
	from := AppSettings.SMTPParams.Username
	to := email
	smtpHost := AppSettings.SMTPParams.Host
	smtpPort := AppSettings.SMTPParams.Port
	smtpUsername := AppSettings.SMTPParams.Username
	smtpPassword := AppSettings.SMTPParams.Password

	var body bytes.Buffer

	temp, err := ParseTemplateDir("templates")
	if err != nil {
		return err
	}

	err = temp.ExecuteTemplate(&body, "verify.html", &data)
	if err != nil {
		return err
	}

	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", to)
	m.SetHeader("Subject", data.Subject)
	m.SetBody("text/html", body.String())
	//m.AddAlternative("text/plain", html2text.HTML2Text(body.String()))

	d := gomail.NewDialer(smtpHost, cast.ToInt(smtpPort), smtpUsername, smtpPassword)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}

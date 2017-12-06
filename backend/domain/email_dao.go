package domain

import (
	"errors"
	"fmt"

	"github.com/1backend/1backend/backend/config"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

const pwResetEmail = `Hi %s,

You can reset your password here: %ssiteUrl/reset/%s

Cheers
1Backend Team`

func SendPasswordReset(secret string, user *User) error {
	from := mail.NewEmail("1Backend", "passwordreset@1backend.com")
	subject := "1Backend Password Reset"
	to := mail.NewEmail(user.Name, user.Email)

	name := user.Name
	if name == "" {
		name = user.Nick
	}

	if config.C.SiteUrl == "" {
		return errors.New("Site url config")
	}
	if config.C.SendGridKey == "" {
		return errors.New("Email api key missing")
	}
	content := mail.NewContent("text/plain", fmt.Sprintf(pwResetEmail, config.C.SiteUrl, name, secret))
	m := mail.NewV3MailInit(from, subject, to, content)

	request := sendgrid.GetRequest(config.C.SendGridKey, "/v3/mail/send", "https://api.sendgrid.com")
	request.Method = "POST"
	request.Body = mail.GetRequestBody(m)
	_, err := sendgrid.API(request)
	return err
}

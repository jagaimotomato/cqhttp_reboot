package pkg

import "gopkg.in/gomail.v2"

/**
*  @Author: tkykc
*  @github: https://github.com/takoyakiccc
*  @date: 2021/12/29
 */

type EMail struct {
	User string
	MailTo []string
	Password string
	Host string
	Port int
	Alias string
	Subject string
	Body string
}

func (e *EMail) Send() error {
	mail := gomail.NewMessage()
	mail.SetHeader("From", mail.FormatAddress(e.User, e.Alias)) // 这种方式可以添加别名，即“XX官方”
	mail.SetHeader("To", e.MailTo...)
	mail.SetHeader("Subject", e.Subject)
	mail.SetBody("text/plain", e.Body)
	dial := gomail.NewDialer(e.Host, e.Port, e.User, e.Password)
	return dial.DialAndSend(mail)
}
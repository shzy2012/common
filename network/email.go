package network

import (
	"crypto/tls"
	"strings"

	"github.com/shzy2012/common/log"
	"gopkg.in/gomail.v2"
)

//Email 邮件
type Email struct {
	ServerSMTP     string
	ServerPort     int
	ServerUser     string
	ServerPassword string
}

//NewEmail 邮件
func NewEmail(smtp string, port int, user, passwd string) *Email {

	return &Email{
		ServerSMTP:     smtp,
		ServerPort:     port,
		ServerUser:     user,
		ServerPassword: passwd,
	}
}

//Send 发送邮件
func (x *Email) Send(from, to, subject, context string) error {

	log.Printf("[SendEmail]: \n from=>%s \n to=>%s \n subject=>%s \n context=>%s\n", from, to, subject, context)

	m := gomail.NewMessage()
	m.SetHeader("From", from)
	addres := strings.Split(to, ",")
	m.SetHeader("To", addres...)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", context)
	//m.SetAddressHeader("Cc", "baohaoze@laiye.com", "baohaoze")
	//添加附件
	// m.Attach("users.csv", gomail.SetCopyFunc(func(w io.Writer) error {
	// 	_, err := w.Write([]byte(file))
	// 	return err
	// }))

	d := gomail.NewDialer(x.ServerSMTP, x.ServerPort, x.ServerUser, x.ServerPassword)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true} //跳过https验证

	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}

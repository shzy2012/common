package network

import (
	"fmt"
	"testing"
)

func Test_SendEmail(t *testing.T) {

	smtp := "smtp.exmail.qq.com"
	port := 465
	user := "xxx"
	passwd := "xxx"

	email := NewEmail(smtp, port, user, passwd)
	err := email.Send("sh-notify@163.com", "demo@163.com", "测试发送", "邮件内容")
	if err != nil {
		fmt.Println("发送失败")
	} else {
		fmt.Println("发送成功")
	}

}

package go_email

import (
	"fmt"
	"testing"
)

func TestEmail(t *testing.T) {
	myEmail := &EmailConf{
		ServerHost: "smtpdm.aliyun.com",
		ServerPort: 465,
		FromEmail:  "donotreply@directmail.coinexpress.cc",
		FromPasswd: "siz8ivQbruGasVh9NtqO",
		FromName:   "CoinExpress",
	}
	myEmail.InitEmail()
	send := &EmailSend{
		Msg:   nil,
		Toers: "xiaka53@vip.qq.com",
		CCers: "",
	}
	subject := "这是主题"
	body := `这是正文<br>
            <h3>这是标题</h3>
             Hello <a href = "http://www.latelee.org">主页</a><br>`
	err := send.Get().SendEmail(subject, body)
	fmt.Println(err)
}

package go_email

import "testing"

func TestEmail(t *testing.T) {
	myEmail := &EmailConf{
		ServerHost: "******",
		ServerPort: 465,
		FromEmail:  "******",
		FromPasswd: "*****",
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
	send.Get().SendEmail(subject, body)
}

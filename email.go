package go_email

import (
	"github.com/go-gomail/gomail"
	"strings"
)

type EmailConf struct {
	// ServerHost 邮箱服务器地址，如腾讯企业邮箱为smtp.exmail.qq.com
	ServerHost string
	// ServerPort 邮箱服务器端口，如腾讯企业邮箱为465
	ServerPort int
	// FromEmail　发件人邮箱地址
	FromEmail string
	// FromPasswd 发件人邮箱密码（注意，这里是明文形式）
	FromPasswd string
	//FromName 发件人别名
	FromName string
}

type EmailSend struct {
	//发送邮件存储信息
	Msg *gomail.Message
	// Toers 接收者邮件，如有多个，则以英文逗号(“,”)隔开，不能为空
	Toers string
	// CCers 抄送者邮件，如有多个，则以英文逗号(“,”)隔开，可以为空
	CCers string
}

// 全局变量，因为发件人账号、密码，需要在发送时才指定
// 注意，由于是小写，外面的包无法使用
var serverHost, fromEmail, fromPasswd, fromName string
var serverPort int

func (s *EmailSend) Get() *EmailSend {
	s.Msg = gomail.NewMessage()
	if len(s.Toers) == 0 {
		return s
	}
	var toers []string
	for _, tmp := range strings.Split(s.Toers, ",") {
		toers = append(toers, strings.TrimSpace(tmp))
	}
	// 收件人可以有多个
	s.Msg.SetHeader("To", toers...)
	//抄送列表
	if len(s.CCers) != 0 {
		for _, tmp := range strings.Split(s.CCers, ",") {
			toers = append(toers, strings.TrimSpace(tmp))
		}
		s.Msg.SetHeader("Cc", toers...)
	}
	// 发件人
	// 第三个参数为发件人别名，如"隔壁老王"，可以为空（此时则为邮箱名称）
	s.Msg.SetAddressHeader("From", fromEmail, fromName)
	return s
}

func (ep *EmailConf) InitEmail() {
	serverHost = ep.ServerHost
	serverPort = ep.ServerPort
	fromEmail = ep.FromEmail
	fromName = ep.FromName
	fromPasswd = ep.FromPasswd
}

// SendEmail body支持html格式字符串
func (s *EmailSend) SendEmail(subject, body string) {
	// 主题
	s.Msg.SetHeader("Subject", subject)

	// 正文
	s.Msg.SetBody("text/html", body)

	d := gomail.NewPlainDialer(serverHost, serverPort, fromEmail, fromPasswd)
	// 发送
	if err := d.DialAndSend(s.Msg); err != nil {
		panic(err)
	}
}

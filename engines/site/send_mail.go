package site

import (
	"github.com/astaxie/beego"
	"github.com/benmanns/goworker"
	"github.com/kapmahc/lotus/engines/base"
	gomail "gopkg.in/gomail.v2"
)

//SMTP smtp config
type SMTP struct {
	From     string
	Username string
	Password string
	Host     string
	Port     int
}

func sendMailWorker(queue string, args ...interface{}) error {
	to := args[0].(string)
	subject := args[1].(string)
	body := args[2].(string)

	if beego.AppConfig.String("runmode") == "prod" {
		var cfg SMTP
		if err := Get("smtp", &cfg); err != nil {
			return err
		}
		msg := gomail.NewMessage()
		msg.SetHeader("From", cfg.From)
		msg.SetHeader("To", to)
		// msg.SetHeader("Cc", cc...)
		// msg.SetHeader("Bcc", bcc...)

		msg.SetHeader("Subject", subject)
		msg.SetBody("text/html", body)
		// for _, f := range files {
		// 	msg.Attach(f)
		// }

		dia := gomail.NewDialer(cfg.Host, cfg.Port, cfg.Username, cfg.Password)
		return dia.DialAndSend(msg)
	}

	beego.Debug("send mail to", to, ":", subject, "\n", body)
	return nil
}

func init() {
	goworker.Register(base.SendMailJob, sendMailWorker)
}

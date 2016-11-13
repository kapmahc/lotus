package auth

import (
	log "github.com/Sirupsen/logrus"
	"github.com/kapmahc/lotus/web"
	gomail "gopkg.in/gomail.v2"
)

//SMTP smtp config
type SMTP struct {
	From     string `json:"form"`
	Username string `json:"username"`
	Password string `json:"-"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
}

//Worker register worker
func (p *Engine) Worker() {
	p.Server.RegisterTask(userEmailQueue, func(to, subject, body string) (interface{}, error) {

		var err error
		if web.IsProduction() {
			var cfg SMTP
			err = p.Dao.Get("smtp", &cfg)
			if err == nil {
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
				err = dia.DialAndSend(msg)
			}
		} else {
			log.Debugf("send mail to %s: %s \n %s", to, subject, body)
		}

		return "done", err
	})
}

const userEmailQueue = "auth.user.email"

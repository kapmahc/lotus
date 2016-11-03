package auth

import (
	"github.com/RichardKnop/machinery/v1/signatures"
	"github.com/kapmahc/lotus/web"
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
			p.Logger.Debug("send mail to %s: %s \n %s", to, subject, body)
		}

		return nil, err
	})
}

func (p *Engine) sendMail(to, subject, body string) {
	task := signatures.TaskSignature{
		Name: userEmailQueue,
		Args: []signatures.TaskArg{
			signatures.TaskArg{
				Type:  "string",
				Value: to,
			},
			signatures.TaskArg{
				Type:  "string",
				Value: subject,
			},
			signatures.TaskArg{
				Type:  "string",
				Value: body,
			},
		},
	}
	if _, err := p.Server.SendTask(&task); err != nil {
		p.Logger.Error("add mail job: %v", err)
	}
}

const userEmailQueue = "auth.user.email"

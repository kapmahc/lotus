package auth

import (
	"bytes"
	"encoding/gob"

	"github.com/kapmahc/lotus/web/jobber"
	logging "github.com/op/go-logging"
)

//MailSender mail sender
type MailSender struct {
	Jobber jobber.Jobber   `inject:""`
	Logger *logging.Logger `inject:""`
}

//Send add send mail to queue
func (p *MailSender) Send(to, subject, body string) error {
	return p.Jobber.Push(
		"emails",
		map[string]string{"to": to, "subject": subject, "body": body},
	)
}

func (p *MailSender) register() {
	p.Jobber.Register("emails", func(args []byte) error {
		buf := bytes.NewBuffer(args)
		arg := make(map[string]string)
		dec := gob.NewDecoder(buf)
		if err := dec.Decode(&arg); err != nil {
			return err
		}
		if IsProduction() {
			// TODO
		} else {
			p.Logger.Debugf("%+v", arg)
		}
		return nil
	})
}

package auth

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/SermoDigital/jose/jws"
	"github.com/astaxie/beego"
	"github.com/kapmahc/lotus/engines/base"
)

const (
	actResetPassword = "ResetPassword"
	actConfirm       = "Confirm"
	actUnlock        = "Unlock"
)

func (p *Controller) sendMail(act, email, uid string) {
	cm := jws.Claims{}
	cm.Set("act", act)
	cm.Set("uid", uid)
	tkn, err := base.SumToken(cm, 1)
	p.Check(err)

	st, err := template.New("").Parse(p.T(fmt.Sprintf("auth-emails.%s-subject", act)))
	p.Check(err)
	bt, err := template.New("").Parse(p.T(fmt.Sprintf("auth-emails.%s-body", act)))
	p.Check(err)

	model := struct {
		Href string
	}{
		Href: fmt.Sprintf(
			"%s%s?token=%s",
			beego.AppConfig.String("homeurl"),
			p.URLFor("auth.Controller.Get"+act),
			string(tkn),
		),
	}

	var subject bytes.Buffer
	var body bytes.Buffer

	err = st.Execute(&subject, model)
	p.Check(err)
	bt.Execute(&body, model)
	p.Check(err)

	base.SendMail(email, subject.String(), body.String())
}

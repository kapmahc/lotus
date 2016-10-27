package base

import (
	"github.com/astaxie/beego"
	"github.com/benmanns/goworker"
)

//SendMail send mail
func SendMail(to, subject, body string) {
	goworker.Enqueue(&goworker.Job{
		Queue: QueueLow,
		Payload: goworker.Payload{
			Class: sendMailJob,
			Args:  []interface{}{to, subject, body},
		},
	})
}

func sendMailWorker(queue string, args ...interface{}) error {
	to := args[0].(string)
	subject := args[1].(string)
	body := args[2].(string)
	if beego.AppConfig.String("runmode") == "prod" {
		// TODO
	} else {
		beego.Debug("send mail to", to, ":", subject, "\n", body)
	}
	return nil
}

const sendMailJob = "SendMail"

func init() {
	goworker.Register(sendMailJob, sendMailWorker)
}

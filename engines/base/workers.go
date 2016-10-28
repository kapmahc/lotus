package base

import "github.com/benmanns/goworker"

//SendMail send mail
//Notice only support string,int,float args
func SendMail(to, subject, body string) {

	goworker.Enqueue(&goworker.Job{
		Queue: QueueLow,
		Payload: goworker.Payload{
			Class: SendMailJob,
			Args:  []interface{}{to, subject, body},
		},
	})
}

//SendMailJob send-mail job class name
const SendMailJob = "SendMail"

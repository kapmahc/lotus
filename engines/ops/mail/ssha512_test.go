package mail_test

import (
	"testing"

	"github.com/kapmahc/lotus/engines/ops/mail"
)

func TestSsha512(t *testing.T) {
	var user mail.User
	password := "123456"
	if err := user.SetPassword(password, 8); err != nil {
		t.Fatal(err)
	}
	t.Logf("doveadm pw -t {SSHA512}%s -p %s", user.Password, password)
	rst, err := user.ChkPassword(password)
	if err != nil {
		t.Fatal(err)
	}
	if !rst {
		t.Fatalf("check password failed")
	}
}

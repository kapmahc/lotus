package web_test

import (
	"crypto/aes"
	"crypto/sha512"
	"testing"

	"github.com/kapmahc/lotus/web"
)

const hello = "Hello, Lotus!"

func TestHmac(t *testing.T) {
	en := web.Hmac{Hash: sha512.New, Key: []byte("123456")}
	code := en.Sum([]byte(hello))
	t.Log(string(code))
	if !en.Chk([]byte(hello), code) {
		t.Fatalf("check password failed")
	}
}

func TestAes(t *testing.T) {
	cip, err := aes.NewCipher([]byte("1234567890123456"))
	if err != nil {
		t.Fatal(err)
	}
	en := web.Aes{Cip: cip}

	code, err := en.Encrypt([]byte(hello))
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(code))
	plain, err := en.Decrypt(code)
	if err != nil {
		t.Fatal(err)
	}
	if string(plain) != hello {
		t.Fatalf("want %s get %s", hello, string(plain))
	}
}

func TestSsha512(t *testing.T) {
	var ssha web.Ssha
	code, err := ssha.Sum(hello, 8)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("doveadm pw -t {SSHA512}%s -p %s", code, hello)
	rst, err := ssha.Chk(hello, code)
	if err != nil {
		t.Fatal(err)
	}
	if !rst {
		t.Fatalf("check ssha512 failed")
	}
}

package site

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
)

var aesCip cipher.Block

//Encrypt encrypt buffer
func Encrypt(buf []byte) ([]byte, error) {
	iv := make([]byte, aes.BlockSize)
	if _, err := rand.Read(iv); err != nil {
		return nil, err
	}
	cfb := cipher.NewCFBEncrypter(aesCip, iv)
	val := make([]byte, len(buf))
	cfb.XORKeyStream(val, buf)

	return append(val, iv...), nil
}

//Decrypt decrypt buffer
func Decrypt(buf []byte) ([]byte, error) {
	bln := len(buf)
	cln := bln - aes.BlockSize
	ct := buf[0:cln]
	iv := buf[cln:bln]

	cfb := cipher.NewCFBDecrypter(aesCip, iv)
	val := make([]byte, cln)
	cfb.XORKeyStream(val, ct)
	return val, nil
}

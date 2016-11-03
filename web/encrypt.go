package web

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha512"
	"encoding/base64"
	"hash"
)

//Aes aes
type Aes struct {
	Cip cipher.Block `inject:"aes.cip"`
}

//Encrypt encrypt buffer
func (p *Aes) Encrypt(buf []byte) ([]byte, error) {
	iv := make([]byte, aes.BlockSize)
	if _, err := rand.Read(iv); err != nil {
		return nil, err
	}
	cfb := cipher.NewCFBEncrypter(p.Cip, iv)
	val := make([]byte, len(buf))
	cfb.XORKeyStream(val, buf)

	return append(val, iv...), nil
}

//Decrypt decrypt buffer
func (p *Aes) Decrypt(buf []byte) ([]byte, error) {
	bln := len(buf)
	cln := bln - aes.BlockSize
	ct := buf[0:cln]
	iv := buf[cln:bln]

	cfb := cipher.NewCFBDecrypter(p.Cip, iv)
	val := make([]byte, cln)
	cfb.XORKeyStream(val, ct)
	return val, nil
}

//Hmac hmac
type Hmac struct {
	Key  []byte           `inject:"hmac.key"`
	Hash func() hash.Hash `inject:"hmac.hash"`
}

//Sum sum a password
func (p *Hmac) Sum(plain []byte) []byte {
	mac := hmac.New(p.Hash, p.Key)
	return mac.Sum(plain)
}

//Chk check password
func (p *Hmac) Chk(plain, code []byte) bool {
	return hmac.Equal(p.Sum(plain), code)
}

//Ssha sha512 with salt
type Ssha struct {
}

//Sum sum
func (p *Ssha) Sum(plain string, num uint) (string, error) {
	salt := make([]byte, num)
	if _, err := rand.Read(salt); err != nil {
		return "", err
	}
	return p.sum(plain, salt)
}

//Chk check
func (p *Ssha) Chk(plain, code string) (bool, error) {
	buf, err := base64.StdEncoding.DecodeString(code)
	if err != nil {
		return false, err
	}

	if len(buf) <= sha512.Size {
		return false, err
	}
	salt := buf[sha512.Size:]
	rst, err := p.sum(plain, salt)
	if err != nil {
		return false, err
	}
	return rst == code, nil
}

func (p *Ssha) sum(plain string, salt []byte) (string, error) {
	buf := append([]byte(plain), salt...)
	code := sha512.Sum512(buf)
	return base64.StdEncoding.EncodeToString(append(code[:], salt...)), nil
	// return base64.StdEncoding.EncodeToString(append(p.Hash.Sum(append(plain, salt...)), salt...)), nil
}

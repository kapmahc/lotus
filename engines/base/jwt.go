package base

import (
	"time"

	"github.com/SermoDigital/jose/crypto"
	"github.com/SermoDigital/jose/jws"
	"github.com/SermoDigital/jose/jwt"
	"github.com/astaxie/beego"
	uuid "github.com/satori/go.uuid"
)

//ParseToken parse jwt token
func ParseToken(buf []byte) (jwt.Claims, error) {
	tk, err := jws.ParseJWT(buf)
	if err != nil {
		return nil, err
	}
	err = tk.Validate(jwtKey, jwtMethod)
	return tk.Claims(), err
}

//SumToken sum jwt token
func SumToken(cm jws.Claims, days int) ([]byte, error) {
	kid := uuid.NewV4()
	now := time.Now()
	cm.SetNotBefore(now)
	cm.SetExpiration(now.AddDate(0, 0, days))
	cm.Set("kid", kid)

	jt := jws.NewJWT(cm, jwtMethod)
	return jt.Serialize(jwtKey)

}

var jwtKey []byte
var jwtMethod crypto.SigningMethod

func init() {
	jwtKey = []byte(beego.AppConfig.String("jwtkey"))
	jwtMethod = crypto.SigningMethodHS512
}

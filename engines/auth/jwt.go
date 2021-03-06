package auth

import (
	"fmt"
	"net/http"
	"time"

	"github.com/SermoDigital/jose/crypto"
	"github.com/SermoDigital/jose/jws"
	"github.com/SermoDigital/jose/jwt"
	"github.com/gin-gonic/gin"
	logging "github.com/op/go-logging"
	uuid "github.com/satori/go.uuid"
)

//Jwt jwt helper
type Jwt struct {
	Key    []byte               `inject:"jwt.key"`
	Method crypto.SigningMethod `inject:"jwt.method"`
	Logger *logging.Logger      `inject:""`
	Dao    *Dao                 `inject:""`
}

//Validate check jwt
func (p *Jwt) Validate(buf []byte) (jwt.Claims, error) {
	tk, err := jws.ParseJWT(buf)
	if err != nil {
		return nil, err
	}
	err = tk.Validate(p.Key, p.Method)
	return tk.Claims(), err
}

//MustAdminHandler check must have admin role
func (p *Jwt) MustAdminHandler() gin.HandlerFunc {
	return p.MustRolesHandler("admin")
}

//MustRolesHandler check must have one roles at least
func (p *Jwt) MustRolesHandler(roles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		u := c.MustGet("user").(*User)
		for _, a := range p.Dao.Authority(u.ID, "-", 0) {
			for _, r := range roles {
				if a == r {
					return
				}
			}
		}
		c.String(http.StatusForbidden, "don't have roles %s", roles)
	}
}

//CurrentUserHandler inject current user
func (p *Jwt) CurrentUserHandler(must bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		tkn, err := jws.ParseFromRequest(c.Request, jws.Compact)
		if err != nil {
			if must {
				c.String(http.StatusInternalServerError, err.Error())
				c.Abort()
			}
			return
		}

		if err := tkn.Verify(p.Key, p.Method); err != nil {
			if must {
				c.String(http.StatusUnauthorized, err.Error())
				c.Abort()
			}
			return
		}
		var user User
		data := tkn.Payload().(map[string]interface{})
		if err := p.Dao.Db.Where("uid = ?", data["uid"]).First(&user).Error; err != nil {
			if must {
				c.String(http.StatusUnauthorized, err.Error())
				c.Abort()
			}
			return
		}
		if !user.IsAvailable() {
			if must {
				c.String(http.StatusForbidden, "bad user status")
				c.Abort()
			}
			return
		}
		c.Set("user", &user)
	}
}

func (p *Jwt) key(kid string) string {
	return fmt.Sprintf("token://%s", kid)
}

//Sum create jwt token
func (p *Jwt) Sum(cm jws.Claims, days int) ([]byte, error) {
	kid := uuid.NewV4()
	now := time.Now()
	cm.SetNotBefore(now)
	cm.SetExpiration(now.AddDate(0, 0, days))
	cm.Set("kid", kid)
	//TODO using kid

	jt := jws.NewJWT(cm, p.Method)
	return jt.Serialize(p.Key)
}

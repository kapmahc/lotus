package site

type fmInstall struct {
	Name                 string `form:"name" valid:"Required; MaxSize(32)"`
	Email                string `form:"email" valid:"Email; MaxSize(255)"`
	Password             string `form:"password" valid:"Required; MaxSize(128)"`
	PasswordConfirmation string `form:"passwordConfirmation"`
}

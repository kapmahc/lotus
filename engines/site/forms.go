package site

type fmInstall struct {
	Name                 string `form:"name"`
	Email                string `form:"email"`
	Password             string `form:"password"`
	PasswordConfirmation string `form:"password_confirmation"`
}

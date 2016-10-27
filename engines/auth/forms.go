package auth

type fmSignUp struct {
	Name                 string `form:"name" valid:"Required; MaxSize(32)"`
	Email                string `form:"email" valid:"Email; MaxSize(255)"`
	Password             string `form:"password" valid:"Required; MaxSize(128)"`
	PasswordConfirmation string `form:"passwordConfirmation"`
}

type fmSignIn struct {
	Email    string `form:"email" valid:"Email"`
	Password string `form:"password" valid:"Required"`
}

type fmEmail struct {
	Email string `form:"email" valid:"Email"`
}

type fmResetPassword struct {
	Token                string `form:"token" valid:"Required"`
	Password             string `form:"password" valid:"Required; MaxSize(128)"`
	PasswordConfirmation string `form:"passwordConfirmation"`
}

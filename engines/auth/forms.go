package auth

type fmSignUp struct {
	Name                 string `form:"name" valid:"Required; MaxSize(32)"`
	Email                string `form:"email" valid:"Email; MaxSize(255)"`
	Password             string `form:"password" valid:"Required; MaxSize(128); MinSize(6)"`
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
	Password             string `form:"password" valid:"Required; MaxSize(128); MinSize(6)"`
	PasswordConfirmation string `form:"passwordConfirmation"`
}

type fmInfo struct {
	Name string `form:"name" valid:"Required; MaxSize(255)"`
	Home string `form:"home" valid:"MaxSize(255)"`
	Logo string `form:"logo" valid:"MaxSize(255)"`
}

type fmChangePassword struct {
	CurrentPassword      string `form:"currentPassword" valid:"Required"`
	NewPassword          string `form:"newPassword" valid:"Required; MaxSize(128); MinSize(6)"`
	PasswordConfirmation string `form:"passwordConfirmation"`
}

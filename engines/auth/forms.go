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

type fmInstall struct {
	Name                 string `form:"name" valid:"Required; MaxSize(32)"`
	Email                string `form:"email" valid:"Email; MaxSize(255)"`
	Password             string `form:"password" valid:"Required; MaxSize(128)"`
	PasswordConfirmation string `form:"passwordConfirmation"`
}

type fmContent struct {
	Content string `form:"content" valid:"Required"`
}

type fmLocale struct {
	Code    string `form:"code" valid:"Required; MaxSize(255)"`
	Message string `form:"message" valid:"Required"`
}

type fmSiteBase struct {
	Title       string `form:"title" valid:"Required"`
	SubTitle    string `form:"subTitle" valid:"Required; MaxSize(32)"`
	Keywords    string `form:"keywords" valid:"Required"`
	Description string `form:"description" valid:"Required"`
	Copyright   string `form:"copyright" valid:"Required"`
}

type fmSiteAuthor struct {
	Name  string `form:"name" valid:"Required; MaxSize(32)"`
	Email string `form:"email" valid:"Email; MaxSize(255)"`
}

type fmSeo struct {
	Google string `form:"google"`
	Baidu  string `form:"baidu"`
}

type fmSMTP struct {
	Host                 string `form:"host" valid:"Required"`
	Port                 int    `form:"port" valid:"Required"`
	Username             string `form:"username" valid:"Required"`
	Password             string `form:"password" valid:"Required"`
	PasswordConfirmation string `form:"passwordConfirmation"`
}

type fmNavBar struct {
	Header string `form:"header"`
	Footer string `form:"footer"`
}

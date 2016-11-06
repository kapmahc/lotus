package auth

type fmSignUp struct {
	Name                 string `form:"name" binding:"max=32,min=2"`
	Email                string `form:"email" binding:"email"`
	Password             string `form:"password" binding:"max=128,min=6"`
	PasswordConfirmation string `form:"passwordConfirmation" binding:"eqfield=Password"`
}

type fmSignIn struct {
	Email    string `form:"email" binding:"email"`
	Password string `form:"password" binding:"required"`
}

type fmEmail struct {
	Email string `form:"email" binding:"email"`
}

type fmResetPassword struct {
	Token                string `form:"token" binding:"required"`
	Password             string `form:"password" binding:"max=128,min=6"`
	PasswordConfirmation string `form:"passwordConfirmation" binding:"eqfield=Password"`
}

type fmInfo struct {
	Name string `form:"name" binding:"max=32,min=2"`
	Home string `form:"home" binding:"max=255"`
	Logo string `form:"logo" binding:"max=255"`
}

type fmChangePassword struct {
	CurrentPassword      string `form:"currentPassword" binding:"required"`
	NewPassword          string `form:"newPassword" binding:"max=128,min=6"`
	PasswordConfirmation string `form:"passwordConfirmation" binding:"eqfield=Password"`
}

type fmContent struct {
	Content string `form:"content" binding:"required"`
}

type fmLocale struct {
	Code    string `form:"code" binding:"required,max=255"`
	Message string `form:"message" binding:"required"`
}

type fmSiteBase struct {
	Title       string `form:"title" binding:"required,max=255"`
	SubTitle    string `form:"subTitle" binding:"required,max=32"`
	Keywords    string `form:"keywords" binding:"required"`
	Description string `form:"description" binding:"required,max=500"`
	Copyright   string `form:"copyright" binding:"required,max=128"`
}

type fmSiteAuthor struct {
	Name  string `form:"name" binding:"min=2,max=32"`
	Email string `form:"email" binding:"email"`
}

type fmSeo struct {
	Google string `form:"google" binding:"required"`
	Baidu  string `form:"baidu" binding:"required"`
}

type fmSMTP struct {
	Host                 string `form:"host" binding:"required"`
	Port                 int    `form:"port" binding:"required"`
	Username             string `form:"username" binding:"required"`
	Password             string `form:"password" binding:"required"`
	PasswordConfirmation string `form:"passwordConfirmation" binding:"eqfield=Password"`
}

type fmNavBar struct {
	Header string `form:"header" binding:"required"`
	Footer string `form:"footer" binding:"required"`
}

package auth

type fmSignUp struct {
	Name                 string `form:"name" validate:"max=32,min=2"`
	Email                string `form:"email" validate:"required,email"`
	Password             string `form:"password" validate:"max=128,min=6"`
	PasswordConfirmation string `form:"passwordConfirmation" validate:"eqfield=Password"`
}

type fmSignIn struct {
	Email    string `form:"email" validate:"required,email"`
	Password string `form:"password" validate:"required"`
}

type fmEmail struct {
	Email string `form:"email" validate:"required,email"`
}

type fmResetPassword struct {
	Token                string `form:"token" validate:"required"`
	Password             string `form:"password" validate:"max=128,min=6"`
	PasswordConfirmation string `form:"passwordConfirmation" validate:"eqfield=Password"`
}

type fmInfo struct {
	Name string `form:"name" validate:"max=32,min=2"`
	Home string `form:"home" validate:"max=255"`
	Logo string `form:"logo" validate:"max=255"`
}

type fmChangePassword struct {
	CurrentPassword      string `form:"currentPassword" validate:"required"`
	NewPassword          string `form:"newPassword" validate:"max=128,min=6"`
	PasswordConfirmation string `form:"passwordConfirmation" validate:"eqfield=Password"`
}

type fmContent struct {
	Content string `form:"content" validate:"required"`
}

type fmLocale struct {
	Code    string `form:"code" validate:"required,max=255"`
	Message string `form:"message" validate:"required"`
}

type fmInstall struct {
	Title                string `form:"title" validate:"required,max=255"`
	SubTitle             string `form:"subTitle" validate:"required,max=32"`
	Username             string `form:"username" validate:"max=32,min=2"`
	Email                string `form:"email" validate:"required,email"`
	Password             string `form:"password" validate:"max=128,min=6"`
	PasswordConfirmation string `form:"passwordConfirmation" validate:"eqfield=Password"`
}

type fmSiteBase struct {
	Title       string `form:"title" validate:"required,max=255"`
	SubTitle    string `form:"subTitle" validate:"required,max=32"`
	Keywords    string `form:"keywords" validate:"required"`
	Description string `form:"description" validate:"required,max=500"`
	Copyright   string `form:"copyright" validate:"required,max=128"`
}

type fmSiteAuthor struct {
	Name  string `form:"name" validate:"min=2,max=32"`
	Email string `form:"email" validate:"required,email"`
}

type fmSeo struct {
	Google string `form:"google" validate:"required"`
	Baidu  string `form:"baidu" validate:"required"`
}

type fmSMTP struct {
	Host                 string `form:"host" validate:"required"`
	Port                 uint   `form:"port" validate:"gte=1,lte=65535"`
	Username             string `form:"username" validate:"required"`
	Password             string `form:"password" validate:"required"`
	PasswordConfirmation string `form:"passwordConfirmation" validate:"eqfield=Password"`
}

type fmNavBar struct {
	Header string `form:"header" validate:"required"`
	Footer string `form:"footer" validate:"required"`
}

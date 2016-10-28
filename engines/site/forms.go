package site

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

type fmBase struct {
	Title       string `form:"title" valid:"Required"`
	SubTitle    string `form:"subTitle" valid:"Required; MaxSize(32)"`
	Keywords    string `form:"keywords" valid:"Required"`
	Description string `form:"description" valid:"Required"`
	Copyright   string `form:"copyright" valid:"Required"`
}

type fmAuthor struct {
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

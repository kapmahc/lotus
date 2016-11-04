package vpn

type fmAddUser struct {
	Email                string `form:"email" valid:"Email"`
	Password             string `form:"password" valid:"Required; MaxSize(128); MinSize(6)"`
	PasswordConfirmation string `form:"passwordConfirmation"`
	StartUp              string `form:"start_up" valid:"Required"`
	ShutDown             string `form:"shut_down" valid:"Required"`
	Details              string `form:"details"`
	Enable               bool   `form:"enable" valid:"Required"`
}

type fmResetUserPassword struct {
	Password             string `form:"password" valid:"Required; MaxSize(128); MinSize(6)"`
	PasswordConfirmation string `form:"passwordConfirmation"`
}

type fmEditUserProfile struct {
	StartUp  string `form:"start_up" valid:"Required"`
	ShutDown string `form:"shut_down" valid:"Required"`
	Details  string `form:"details"`
	Enable   bool   `form:"enable" valid:"Required"`
}

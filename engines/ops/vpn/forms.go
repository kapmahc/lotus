package vpn

type fmAddUser struct {
	Email                string `form:"email" validate:"required,email"`
	Password             string `form:"password" validate:"min=6,max=128"`
	PasswordConfirmation string `form:"passwordConfirmation" validate:"eqfield=Password"`
	StartUp              string `form:"start_up" validate:"required"`
	ShutDown             string `form:"shut_down" validate:"required"`
	Details              string `form:"details"`
	Enable               bool   `form:"enable"`
}

type fmResetUserPassword struct {
	Password             string `form:"password" validate:"min=6,max=128"`
	PasswordConfirmation string `form:"passwordConfirmation" validate:"eqfield=Password"`
}

type fmEditUserProfile struct {
	StartUp  string `form:"start_up" validate:"required"`
	ShutDown string `form:"shut_down" validate:"required"`
	Details  string `form:"details"`
	Enable   bool   `form:"enable"`
}

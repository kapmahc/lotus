package vpn

type fmAddUser struct {
	Email                string `form:"email" binding:"email"`
	Password             string `form:"password" binding:"min=6,max=128"`
	PasswordConfirmation string `form:"passwordConfirmation" binding:"eqfield=Password"`
	StartUp              string `form:"start_up" binding:"required"`
	ShutDown             string `form:"shut_down" binding:"required"`
	Details              string `form:"details"`
	Enable               bool   `form:"enable"`
}

type fmResetUserPassword struct {
	Password             string `form:"password" binding:"min=6,max=128"`
	PasswordConfirmation string `form:"passwordConfirmation" binding:"eqfield=Password"`
}

type fmEditUserProfile struct {
	StartUp  string `form:"start_up" binding:"required"`
	ShutDown string `form:"shut_down" binding:"required"`
	Details  string `form:"details"`
	Enable   bool   `form:"enable"`
}

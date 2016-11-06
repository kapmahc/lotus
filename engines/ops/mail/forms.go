package mail

type fmAddUser struct {
	DomainID             uint   `form:"domain_id" binding:"required"`
	Email                string `form:"email" binding:"email"`
	Name                 string `form:"name" binding:"required,max=128"`
	Password             string `form:"password" binding:"min=6,max=128"`
	PasswordConfirmation string `form:"passwordConfirmation" binding:"eqfield=Password"`
}

type fmEditUserProfile struct {
	Name string `form:"name" binding:"required,max=128"`
}

type fmResetUserPassword struct {
	Password             string `form:"password" binding:"min=6,max=128"`
	PasswordConfirmation string `form:"passwordConfirmation" binding:"eqfield=Password"`
}

type fmDomain struct {
	Name string `form:"name" binding:"required,max=128"`
}

type fmAlias struct {
	DomainID    uint   `form:"domain_id" binding:"required"`
	Source      string `form:"source" binding:"email"`
	Destination string `form:"destination" binding:"email"`
}

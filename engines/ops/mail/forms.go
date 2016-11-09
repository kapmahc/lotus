package mail

type fmAddUser struct {
	DomainID             uint   `form:"domain_id" validate:"required"`
	Email                string `form:"email" validate:"required,email"`
	Name                 string `form:"name" validate:"required,max=128"`
	Password             string `form:"password" validate:"min=6,max=128"`
	PasswordConfirmation string `form:"passwordConfirmation" validate:"eqfield=Password"`
}

type fmEditUserProfile struct {
	Name string `form:"name" validate:"required,max=128"`
}

type fmResetUserPassword struct {
	Password             string `form:"password" validate:"min=6,max=128"`
	PasswordConfirmation string `form:"passwordConfirmation" validate:"eqfield=Password"`
}

type fmDomain struct {
	Name string `form:"name" validate:"required,max=128"`
}

type fmAlias struct {
	DomainID    uint   `form:"domain_id" validate:"required"`
	Source      string `form:"source" validate:"required,email"`
	Destination string `form:"destination" validate:"required,email"`
}

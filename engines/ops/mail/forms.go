package mail

type fmAddUser struct {
	DomainID             uint   `form:"domain_id" valid:"Required"`
	Email                string `form:"email" valid:"Email; MaxSize(255)"`
	Name                 string `form:"name" valid:"Required; MaxSize(128)"`
	Password             string `form:"password" valid:"Required; MaxSize(128); MinSize(6)"`
	PasswordConfirmation string `form:"passwordConfirmation"`
}

type fmEditUserProfile struct {
	Name string `form:"name" valid:"Required; MaxSize(128)"`
}

type fmResetUserPassword struct {
	Password             string `form:"password" valid:"Required; MaxSize(128); MinSize(6)"`
	PasswordConfirmation string `form:"passwordConfirmation"`
}

type fmDomain struct {
	Name string `form:"name" valid:"Required; MaxSize(128)"`
}

type fmAlias struct {
	DomainID    uint   `form:"domain_id" valid:"Required"`
	Source      string `form:"source" valid:"Email; MaxSize(255)"`
	Destination string `form:"destination" valid:"Email; MaxSize(255)"`
}

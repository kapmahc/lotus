package forum

type fmTag struct {
	Name string `form:"name" valid:"Required; MaxSize(32)"`
}

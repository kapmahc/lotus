package reading

type fmNote struct {
	BookID uint   `form:"book_id" valid:"Required"`
	Body   string `form:"body" valid:"Required"`
}

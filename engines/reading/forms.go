package reading

type fmNote struct {
	BookID uint   `form:"book_id" validate:"required"`
	Body   string `form:"body" validate:"required"`
}

package reading

type fmNote struct {
	BookID uint   `form:"book_id" binding:"required"`
	Body   string `form:"body" binding:"required"`
}

package forum

type fmTag struct {
	Name string `form:"name" validate:"required,max=32"`
}

type fmArticle struct {
	Title   string   `form:"title" validate:"required,max=255"`
	Summary string   `form:"summary" validate:"required,max=800"`
	Body    string   `form:"body" validate:"required"`
	Tags    []string `form:"tags"`
}

type fmComment struct {
	ArticleID uint   `form:"article_id" validate:"required"`
	Body      string `form:"body" validate:"required"`
}

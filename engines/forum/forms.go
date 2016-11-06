package forum

type fmTag struct {
	Name string `form:"name" binding:"required,max=32"`
}

type fmArticle struct {
	Title   string   `form:"title" binding:"required,max=255"`
	Summary string   `form:"summary" binding:"required,max=800"`
	Body    string   `form:"body" binding:"required"`
	Tags    []string `form:"tags"`
}

type fmComment struct {
	ArticleID uint   `form:"article_id" binding:"required"`
	Body      string `form:"body" binding:"required"`
}

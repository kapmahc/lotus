package forum

type fmTag struct {
	Name string `form:"name" valid:"Required; MaxSize(32)"`
}

type fmArticle struct {
	Title   string   `form:"name" valid:"Required; MaxSize(255)"`
	Summary string   `form:"summary" valid:"Required; MaxSize(800)"`
	Body    string   `form:"body" valid:"Required"`
	Tags    []string `form:"tags"`
}

type fmComment struct {
	ArticleID uint   `form:"article_id" valid:"Required"`
	Body      string `form:"body" valid:"Required"`
}

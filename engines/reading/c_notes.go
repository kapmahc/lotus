package reading

import (
	"github.com/astaxie/beego/orm"
	"github.com/kapmahc/lotus/engines/auth"
	"github.com/kapmahc/lotus/engines/base"
)

//NewNote new note
// @router /notes/new [get]
func (p *Controller) NewNote() {
	p.MustSignIn()
	title := p.T("reading-pages.new-note")
	p.Data["title"] = title
	p.Data["form"] = p.NewForm(
		"fm-new-note",
		title,
		base.MethodPost,
		p.URLFor("reading.Controller.CreateNote"),
		[]base.Field{
			&base.HiddenField{
				ID:    "book_id",
				Value: p.GetString("book_id"),
			},
			&base.Textarea{
				ID:     "body",
				Label:  p.T("attributes.body"),
				Helper: p.T("site-pages.can-markdown"),
			},
		},
	)
	p.TplName = "auth/form.html"

}

//CreateNote create note
// @router /notes [post]
func (p *Controller) CreateNote() {
	p.MustSignIn()
	var fm fmNote
	fl, er := p.ParseForm(&fm)
	bid := p.GetString("book_id")

	if er == nil {
		var book Book
		o := orm.NewOrm()
		err := o.QueryTable(&book).Filter("id", bid).One(&book)
		p.Check(err)
		_, err = o.Insert(&Note{
			Book: &book,
			Body: fm.Body,
			User: p.CurrentUser(),
		})
		p.Check(err)
		fl.Notice(p.T("site-pages.success"))
		p.Redirect(fl, "reading.Controller.ShowBook", ":id", bid)
	} else {
		fl.Error(er.Error())
		p.Redirect(fl, "reading.Controller.NewNote", "book_id", bid)
	}
}

func (p *Controller) canNote() (Note, bool) {
	var note Note
	err := orm.NewOrm().
		QueryTable(&note).
		Filter("id", p.Ctx.Input.Param(":id")).
		One(&note)
	p.Check(err)

	return note, (p.IsSignIn() && (note.User.ID == p.CurrentUser().ID || p.CurrentUser().Has(auth.AdminRole)))
}

//EditNote edit note
// @router /notes/:id/edit [get]
func (p *Controller) EditNote() {
	note, can := p.canNote()
	if !can {
		p.Abort("403")
	}

	title := p.T("reading-pages.edit-note", note.ID)
	p.Data["title"] = title
	p.Data["form"] = p.NewForm(
		"fm-edit-note",
		title,
		base.MethodPost,
		p.URLFor("reading.Controller.UpdateNote", ":id", note.ID),
		[]base.Field{
			&base.HiddenField{
				ID:    "book_id",
				Value: note.Book.ID,
			},
			&base.Textarea{
				ID:     "body",
				Label:  p.T("attributes.body"),
				Value:  note.Body,
				Helper: p.T("site-pages.can-markdown"),
			},
		},
	)
	p.TplName = "auth/form.html"
}

//UpdateNote update note
// @router /notes/:id [post]
func (p *Controller) UpdateNote() {
	note, can := p.canNote()
	if !can {
		p.Abort("403")
	}

	var fm fmNote
	fl, er := p.ParseForm(&fm)

	if er == nil {
		note.Body = fm.Body
		_, err := orm.NewOrm().Update(&note, "updated_at", "body")
		p.Check(err)
		fl.Notice(p.T("site-pages.success"))
		p.Redirect(fl, "reading.Controller.ShowBook", ":id", note.Book.ID)
	} else {
		fl.Error(er.Error())
		p.Redirect(fl, "reading.Controller.EditNote", ":id", note.ID)
	}
}

//DestroyNote destroy note
// @router /notes/:id [delete]
func (p *Controller) DestroyNote() {
	note, can := p.canNote()
	if !can {
		p.Abort("403")
	}
	_, err := orm.NewOrm().Delete(&note)
	p.Check(err)

	p.Data["json"] = map[string]string{
		"to": p.URLFor("reading.Controller.ShowBook", ":id", note.Book.ID),
	}
	p.ServeJSON()
}

package web

import (
	"net/http"
	"net/url"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/go-playground/form"
	"github.com/gorilla/mux"
	validator "gopkg.in/go-playground/validator.v9"
)

//H hash
type H map[string]interface{}

//ParseForm parse form from http request
func ParseForm(rq *http.Request, vt *validator.Validate, fm interface{}) error {
	if err := rq.ParseForm(); err != nil {
		return err
	}
	dec := form.NewDecoder()
	dec.RegisterCustomTypeFunc(func(vals []string) (interface{}, error) {
		return time.Parse("2006-01-02", vals[0])
	}, time.Time{})
	if err := dec.Decode(fm, rq.PostForm); err != nil {
		return err
	}
	return vt.Struct(fm)
}

//URLFor build url
func URLFor(router *mux.Router, name string, params map[string]string, pairs ...string) *url.URL {
	url, err := router.Get(name).URL(pairs...)
	if err != nil {
		log.Error(err)
	}
	if params != nil {
		qry := url.Query()
		for k, v := range params {
			qry.Set(k, v)
		}
		url.RawQuery = qry.Encode()
	}
	return url
}

//Redirect rediect response
func Redirect(w http.ResponseWriter, r *http.Request, u *url.URL) {
	http.Redirect(w, r, u.String(), http.StatusFound)
}

//RedirectHandler rediect response handler
func RedirectHandler(fn func(http.ResponseWriter, *http.Request) (*url.URL, error)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		u, e := fn(w, r)
		if e == nil {
			Redirect(w, r, u)
		} else {
			http.Error(w, e.Error(), http.StatusInternalServerError)
		}
	}
}

//JSON json response
func JSON(fn func(wrt http.ResponseWriter, req *http.Request) (interface{}, error)) http.HandlerFunc {
	//TODO
	return func(wrt http.ResponseWriter, req *http.Request) {
		// if val, err := fn(c); err == nil {
		// 	c.JSON(http.StatusOK, val)
		// } else {
		// 	if !c.IsAborted() {
		// 		c.AbortWithStatus(http.StatusInternalServerError)
		// 	}
		// 	c.Writer.WriteString(err)
		// }
	}
}

package web

import (
	"net/http"
	"net/url"

	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
)

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
func Redirect(fn func(wrt http.ResponseWriter, req *http.Request) (string, error)) http.HandlerFunc {
	return func(wrt http.ResponseWriter, req *http.Request) {
		//TODO
		// if url, err := fn(c); err == nil {
		// 	c.Redirect(http.StatusTemporaryRedirect, url)
		// } else {
		// 	if !c.IsAborted() {
		// 		c.AbortWithStatus(http.StatusInternalServerError)
		// 	}
		// 	c.Writer.WriteString(err)
		// }
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

package hello

import (
	"appengine"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
)

var templates, _ = template.ParseFiles("templates/index.html")

func init() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/query", query_handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	templates.Execute(w, nil)
	c := appengine.NewContext(r)
	c.Infof("this logger is of opinion that Dima is very cool")
}

func query_handler(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	message, err := ioutil.ReadAll(r.Body)

	if err != nil {
		c.Infof("Error :( %s ", err)
	}

	c.Infof("It says %s ", message)
	fmt.Fprint(w, string(message)+". that is what she said!")
}

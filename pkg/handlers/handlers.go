package handlers

import (
	"net/http"

	"github.com/dwarakanathreddy/goweb/renders"
)

func Home(w http.ResponseWriter, r *http.Request) {
	// data := *r
	// fmt.Println(data.Body)
	// fmt.Println(r)
	// fmt.Fprintf(w, "home")
	renders.RenderTemplate(w, "home.page.tmpl")
}

func About(w http.ResponseWriter, r *http.Request) {
	renders.RenderTemplate(w, "about.page.tmpl")
}

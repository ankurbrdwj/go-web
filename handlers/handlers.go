package handlers

import ("net/http"
 "github.com/ankurbrdwj/go-web/renderers" )


func Home(w http.ResponseWriter, r *http.Request) {
	renderers.RenderTemplate(w, "home.page.tmpl")
}

func About(w http.ResponseWriter, r *http.Request) {
	renderers.RenderTemplate(w, "about.page.tmpl")
}
package handlers

import (
	"net/http"
	"github.com/urdogan0000/test/pkg/render"
)

func Home(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "homePage.tmpl")
}
func About(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "aboutPage.tmpl")
}

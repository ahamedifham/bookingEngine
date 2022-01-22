package handlers

import (
	"net/http"
	"github.com/ahamedifham/bookingEngine/pkg/render"
)

func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w,"home.page.go.tmpl")
}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w,"about.page.go.tmpl")
}


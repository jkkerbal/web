package controllers

import (
	"github.com/jkkerbal/web/models"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

func Login(w http.ResponseWriter, req *http.Request) {

	templatePath := filepath.Join("templates", "login.html")

	t, err := template.ParseFiles(templatePath)
	if err != nil {
		log.Println(err.Error())

	}
	t.ExecuteTemplate(w, "login", nil)

	if req.Method == "POST" {
		username := req.FormValue("username")
		password := req.FormValue("password")
		user := models.AuthenticateUser(username, password)

		log.Println(user)
	}

}

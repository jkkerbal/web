package controllers

import (
	"database/sql"
	"github.com/jkkerbal/web/models"
	_ "github.com/lib/pq"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

func Register(w http.ResponseWriter, r *http.Request) {

	templatePath := filepath.Join("templates", "register.html")

	t, err := template.ParseFiles(templatePath)
	if err != nil {
		log.Println(err.Error())

	}

	t.ExecuteTemplate(w, "register", nil)

	if r.Method == "POST" {
		username := r.FormValue("username")
		password := r.FormValue("password")
		log.Println(username)
		log.Println(password)

		user := models.NewUser(username, password)

		connString := "dbname=postgres sslmode=disable user=postgres"

		db, err := sql.Open("postgres", connString)

		if err != nil {
			log.Println(err.Error())

		}

		_, err = db.Query("INSERT INTO users (username, password, salt) VALUES ($1, $2, $3)", user.UserName, user.Password, user.Salt)
		if err != nil {
			log.Println(err.Error())

		}
	}

}

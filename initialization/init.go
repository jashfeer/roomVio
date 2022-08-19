package initialization

import (
	"fmt"
	"text/template"

	"gorm.io/gorm"
	"github.com/gorilla/sessions"

	"main.go/database"
)

var Tpl *template.Template
var Db *gorm.DB
var Store= sessions.NewCookieStore([]byte("super-secret"))


func Init() {
	var err error
	Db, err = database.OpenDB()
	if err != nil {
		panic(err)
	}
	fmt.Println("you connected to your database.")
	Tpl = template.Must(template.ParseGlob("templates/*.html"))
	
}

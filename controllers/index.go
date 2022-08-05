package controllers

import (
	//"fmt"
	"net/http"

	"main.go/initialization"
)

func Index(res http.ResponseWriter, req *http.Request) {
	// if alreadyLoggedIn(res, req) {
	// 	http.Redirect(res, req, "/welcome", http.StatusSeeOther)
	// 	return
	// }
	u := getUser(res, req)
	//fmt.Println("index u :    ", u)
	initialization.Tpl.ExecuteTemplate(res, "index.html", u)
}


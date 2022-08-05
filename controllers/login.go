package controllers

import (
	"fmt"
	"net/http"

	"golang.org/x/crypto/bcrypt"
	"main.go/initialization"
)

func LoginFrome(res http.ResponseWriter,req *http.Request){
	if alreadyLoggedIn(res,req) {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}
	initialization.Tpl.ExecuteTemplate(res, "login.html", nil)
}

func LoginProcess(res http.ResponseWriter,req *http.Request){
	if req.Method != "POST" {
		initialization.Tpl.ExecuteTemplate(res, "login.html", nil)
		return
	}
	ur:=user{}
	email := req.FormValue("email")
	password := req.FormValue("password")
	result :=initialization.Db.Where("email = ?", email).First(&ur)
	if result.Error != nil {
		initialization.Tpl.ExecuteTemplate(res, "login.html", "check username and password")
		return
	}
	result.Error = bcrypt.CompareHashAndPassword([]byte(ur.Password), []byte(password))
	if result.Error == nil {
		session,_:=initialization.Store.Get(req,"session")
		session.Values["email"]=ur.Email
		session.Save(req,res)



		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}
	fmt.Println("incorrect password")
	initialization.Tpl.ExecuteTemplate(res, "login.html",nil)
}
package controllers

import (
	"net/http"

	"main.go/initialization"
)

func Logout(res http.ResponseWriter,req *http.Request){
	if !alreadyLoggedIn(res,req) {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}
	session,_:=initialization.Store.Get(req,"session")
	delete(session.Values,"email")
	session.Save(req,res)
	http.Redirect(res, req, "/", http.StatusSeeOther)
}
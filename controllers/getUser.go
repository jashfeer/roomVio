package controllers

import (
	//"fmt"
	"net/http"

	"main.go/initialization"
)

func alreadyLoggedIn(res http.ResponseWriter, req *http.Request) bool {
	// fmt.Println("alreadyLoggedIn")
	session, _ := initialization.Store.Get(req, "session")
	_, ok := session.Values["email"]
	return ok
}

func getUser(res http.ResponseWriter, req *http.Request) user {
	// fmt.Println("getUser")
	session, _ := initialization.Store.Get(req, "session")
	email, ok := session.Values["email"]
	var u user
	if ok {
		//db.QueryRow("SELECT id,firstname,lastname,email,password,key FROM users WHERE id=$1", id).Scan(&u.Id, &u.Firstname, &u.Lastname, &u.Email, &u.Password, &u.Key)
		initialization.Db.Where("email = ?", email).First(&u)
		return u
	}
	return u
}

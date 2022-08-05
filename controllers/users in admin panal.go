package controllers

import (
	"fmt"
	"net/http"

	"main.go/initialization"
)



var pageUsersList int64
var tottalPageUsersList int64


func UsersListNextPage(res http.ResponseWriter, req *http.Request) {
	if pageUsersList == tottalPageUsersList {
		http.Redirect(res, req, "/users", http.StatusSeeOther)
		} else {
		pageUsersList = pageUsersList+ 1
		http.Redirect(res, req, "/users", http.StatusSeeOther)

	}
}
func UsersListPreviousPage(res http.ResponseWriter, req *http.Request) {
	if pageUsersList== 1 {
		http.Redirect(res, req, "/users", http.StatusSeeOther)
		} else {
		pageUsersList = pageUsersList - 1
		http.Redirect(res, req, "/users", http.StatusSeeOther)

	}
}







//finding all users from Db

func UsersList(res http.ResponseWriter,req *http.Request){
	if !alreadyLoggedIn(res,req){
		http.Redirect(res, req, "/login", http.StatusSeeOther)
		return
	}
	u:=getUser(res,req)
	if u.Useraccess!="admin"{
		http.Redirect(res, req, "/", http.StatusSeeOther)
	}
	if req.Method != "GET" {
		http.Error(res, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}
	if pageUsersList == 0 {
		pageUsersList = 1
	}
	limit := 3
	offset := (pageUsersList - 1) * int64(limit)
	var count int64

	type All struct{
		Usr   []user
		Page       int64
		TottalPage int64
	}
	all:=All{}

	urs:=make([]user,0)
	result:=initialization.Db.Order("created_at asc").Limit(limit).Offset(int(offset)).Find(&urs).Count(&count)
	if result.Error != nil {
		http.Error(res, http.StatusText(500), 500)
		return
	}
	all.Usr=urs
	if tottalPageUsersList == 0 {
		num := count / int64(limit)
		bal := count % int64(limit)
		if bal != 0 {
			tottalPageUsersList = num + 1
		} else {
			tottalPageUsersList = num
		}
		fmt.Println("num",num)
		fmt.Println("bal",bal)
		fmt.Println("tott",tottalPageUsersList)
	}
	all.Page = pageUsersList
	all.TottalPage = tottalPageUsersList

	initialization.Tpl.ExecuteTemplate(res, "users.html", all)
}

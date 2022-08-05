package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"main.go/initialization"
)



var pageAdminPanal int64
var tottalPageAdminPanal int64


func AdminPanalNextPage(res http.ResponseWriter, req *http.Request) {
	if pageAdminPanal == tottalPageAdminPanal {
		http.Redirect(res, req, "/admin/panal", http.StatusSeeOther)
		} else {
		pageAdminPanal = pageAdminPanal+ 1
		http.Redirect(res, req, "/admin/panal", http.StatusSeeOther)

	}
}
func AdminPanalPreviousPage(res http.ResponseWriter, req *http.Request) {
	if pageAdminPanal== 1 {
		http.Redirect(res, req, "/admin/panal", http.StatusSeeOther)
		} else {
		pageAdminPanal = pageAdminPanal - 1
		http.Redirect(res, req, "/admin/panal", http.StatusSeeOther)

	}
}




func AdminPanal(res http.ResponseWriter,req *http.Request){
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

	if pageAdminPanal == 0 {
		pageAdminPanal = 1
	}
	limit := 2
	offset := (pageAdminPanal - 1) * int64(limit)
	var count int64
	
	type All struct{
		Bks   []booking
		Page       int64
		TottalPage int64
	}
	all:=All{}
	bks:=make([]booking,0)
	result:=initialization.Db.Order("booking_id desc").Limit(limit).Offset(int(offset)).Find(&bks).Count(&count)
	if result.Error != nil {
		http.Error(res, http.StatusText(500), 500)
		return
	}
	all.Bks=bks
	if tottalPageAdminPanal == 0 {
		num := count / int64(limit)
		bal := count % int64(limit)
		if bal != 0 {
			tottalPageAdminPanal = num + 1
		} else {
			tottalPageAdminPanal = num
		}
	}
	all.Page = pageAdminPanal
	all.TottalPage = tottalPageAdminPanal

	initialization.Tpl.ExecuteTemplate(res, "admin.html", all)

}




func Booked(res http.ResponseWriter,req *http.Request){
	if !alreadyLoggedIn(res,req) {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}
	bookingid,err := strconv.Atoi(req.FormValue("bookingid"))
	if err != nil {
		fmt.Println("roomCapacity Err")

		//panic(err)
	}
	bk:=booking{}
	initialization.Db.Model(&bk).Where("booking_id = ?", bookingid).Update("booking_status", "Booked")

	http.Redirect(res, req, "/admin/panal", http.StatusSeeOther)


}





func Cancel(res http.ResponseWriter,req *http.Request){
	if !alreadyLoggedIn(res,req) {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}
	bookingid,err := strconv.Atoi(req.FormValue("bookingid"))
	if err != nil {
		fmt.Println("roomCapacity Err")

		//panic(err)
	}
	bk:=booking{}
	initialization.Db.Model(&bk).Where("booking_id = ?", bookingid).Update("booking_status", "Canceled")

	http.Redirect(res, req, "/admin/panal", http.StatusSeeOther)

}
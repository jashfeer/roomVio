package controllers

import (
	"net/http"

	"main.go/initialization"
)

//finding all rooms from Db
var pageShowRoomsToUser int64
var tottalPageShowRoomsToUser int64


func ShowRoomsToUserNextPage(res http.ResponseWriter, req *http.Request) {
	
	if pageShowRoomsToUser == tottalPageShowRoomsToUser {
		http.Redirect(res, req, "/showRoomsToUser", http.StatusSeeOther)
		} else {
		pageShowRoomsToUser = pageShowRoomsToUser + 1
		http.Redirect(res, req, "/showRoomsToUser", http.StatusSeeOther)

	}
}
func ShowRoomsToUserPreviousPage(res http.ResponseWriter, req *http.Request) {
	if pageShowRoomsToUser== 1 {
		http.Redirect(res, req, "/showRoomsToUser", http.StatusSeeOther)
		} else {
		pageShowRoomsToUser = pageShowRoomsToUser - 1
		http.Redirect(res, req, "/showRoomsToUser", http.StatusSeeOther)

	}
}




func ShowRoomsToUser(res http.ResponseWriter, req *http.Request) {

	if req.Method != "GET" {
		http.Error(res, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}
	if pageShowRoomsToUser == 0 {
		pageShowRoomsToUser = 1
	}
	limit := 6
	offset := (pageShowRoomsToUser - 1) * int64(limit)
	var count int64
	msg :=req.FormValue("msg")
	

	type userAndRomms struct {
		U user
		Msg string
		CheckIn string
		CheckOut string
		Page       int64
		TottalPage int64
		Rms []RoomAndHotel
	}
	all := userAndRomms{}

	if msg == "1"{
		all.Msg ="***Check avilable rooms on your dates"
	}else{
		all.Msg =""
	}

	u := getUser(res, req)
	all.U = u

	rms := make([]RoomAndHotel, 0)


	// result := initialization.Db.Find(&rms)
	result := initialization.Db.Table("rooms").
		Joins("join hotels on rooms.room_hotel_id = hotels.hotel_id").Limit(limit).Offset(int(offset)).Scan(&rms).Count(&count)


	if result.Error != nil {
		http.Error(res, http.StatusText(500), 500)
		return
	}
	all.Rms = rms
	if tottalPageShowRoomsToUser == 0 {
		num := count / int64(limit)
		bal := count % int64(limit)
		if bal != 0 {
			tottalPageShowRoomsToUser = num + 1
		} else {
			tottalPageShowRoomsToUser = num
		}
	}
	all.Page = pageShowRoomsToUser
	all.TottalPage = tottalPageShowRoomsToUser


	initialization.Tpl.ExecuteTemplate(res, "showRoomsToUser.html", all)
}

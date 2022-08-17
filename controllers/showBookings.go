package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"main.go/initialization"
)

var pageBookings int64
var tottalPageBookings int64

func UserBookingsNextPage(res http.ResponseWriter, req *http.Request) {
	if pageBookings == tottalPageBookings {
		http.Redirect(res, req, "/userBookings", http.StatusSeeOther)
	} else {
		pageBookings = pageBookings + 1
		http.Redirect(res, req, "/userBookings", http.StatusSeeOther)

	}
}
func UserBookingsPreviousPage(res http.ResponseWriter, req *http.Request) {
	if pageBookings == 1 {
		http.Redirect(res, req, "/userBookings", http.StatusSeeOther)
	} else {
		pageBookings = pageBookings - 1
		http.Redirect(res, req, "/userBookings", http.StatusSeeOther)

	}
}

func ShowUserBookings(res http.ResponseWriter, req *http.Request) {
	if !alreadyLoggedIn(res, req) {
		http.Redirect(res, req, "/login", http.StatusSeeOther)
		return
	}
	pageBookings=0
	tottalPageBookings=0
	http.Redirect(res, req, "/userBookings", http.StatusSeeOther)

}



func UserBookings(res http.ResponseWriter, req *http.Request) {
	if !alreadyLoggedIn(res, req) {
		http.Redirect(res, req, "/login", http.StatusSeeOther)
		return
	}

	type All struct {
		U          user
		Bookings   []booking
		Page       int64
		TottalPage int64
	}
	all := All{}
	all.U = getUser(res, req)

	var count int64
	if pageBookings == 0 {
		pageBookings = 1
	}
	limit := 5
	offset := (pageBookings - 1) * int64(limit)

	bks := make([]booking, 0)
	result := initialization.Db.Where("booking_user_id = ?", all.U.Id).Order("booking_id desc").Limit(limit).Offset(int(offset)).Find(&bks).Count(&count)
	if result.Error != nil {
		http.Error(res, http.StatusText(500), 500)
		return
	}

	all.Bookings = bks

	if tottalPageBookings == 0 {
		num := count / int64(limit)
		bal := count % int64(limit)
		if bal != 0 {
			tottalPageBookings = num + 1
		} else {
			tottalPageBookings = num
		}
	}

	all.Page = pageBookings
	all.TottalPage = tottalPageBookings

	initialization.Tpl.ExecuteTemplate(res, "userBookings.html", all)
}






func UserCancel(res http.ResponseWriter, req *http.Request) {
	if !alreadyLoggedIn(res, req) {
		http.Redirect(res, req, "/login", http.StatusSeeOther)
		return
	}
	bookingid, err := strconv.Atoi(req.FormValue("bookingid"))
	if err != nil {
		fmt.Println("roomCapacity Err")

		//panic(err)
	}
	bk := booking{}
	initialization.Db.Model(&bk).Where("booking_id = ?", bookingid).Update("booking_status", "Canceled")

	http.Redirect(res, req, "/userBookings", http.StatusSeeOther)


}

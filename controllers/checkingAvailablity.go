package controllers

import (
	"fmt"
	"net/http"
	"time"

	"main.go/initialization"
)


func CheckingAvailablity(res http.ResponseWriter, req *http.Request) {
	if !alreadyLoggedIn(res, req) {
		http.Redirect(res, req, "/login", http.StatusSeeOther)
		return
	}
	if req.Method != "POST" {
		initialization.Tpl.ExecuteTemplate(res, "index.html", 300)
		return
	}

	checkIn =req.FormValue("roomsCheckIn")
	checkOut =req.FormValue("roomsCheckOut")


	BookingCheckIn, err := time.Parse("2006-01-02", checkIn)
	if err != nil {
		fmt.Println(err)
		return
	}
	BookingCheckOut, err := time.Parse("2006-01-02", checkOut)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("iNNN:",BookingCheckIn)
	fmt.Println("out:",BookingCheckOut)


	var roomId [] int64
	var bookingRoomId [] int64
	var availabeRoomId [] int64



	initialization.Db.Model(&booking{}).Distinct("booking_room_id").
	Where(" ($1 <= bookings.booking_check_in OR (bookings.booking_check_in <= $1 AND $1 < bookings.booking_check_out)) AND ($2>= bookings.booking_check_out OR (bookings.booking_check_in < $2 AND $2 <= bookings.booking_check_out)) ",BookingCheckIn,BookingCheckOut).
	Select("booking_room_id").
	Scan(&bookingRoomId)	
	
	initialization.Db.Model(&room{}).Select("room_id").Scan(&roomId)

	var flage int
	for _, allId := range roomId {
		flage = 0
		for _, notavailabeId := range bookingRoomId {
			if(allId==notavailabeId){
				flage = 1
			}
		}
		if flage ==0{
			availabeRoomId = append(availabeRoomId, allId)

		}

    }

	count :=len(availabeRoomId)
	rm:=RoomAndHotel{}
	rms:=[]RoomAndHotel{}


	for _, roomId := range availabeRoomId {
		initialization.Db.Table("rooms").
			Joins("join hotels on rooms.room_hotel_id = hotels.hotel_id ").
			Where("rooms.room_id= ?",roomId).
			Scan(&rm)
			rms = append(rms,rm)
		
	}

	type All struct{
		HotelFilter hotel
		RoomFilter room
		CheckIn string
		CheckOut string
		U user
		MinPrice int
		MaxPrice int
		Rms []RoomAndHotel
	}
	all:=All{}
	all.MinPrice=0
	all.MaxPrice=10000
	all.CheckIn=checkIn
	all.CheckOut=checkOut
	u:=getUser(res,req)
	all.U=u
	all.Rms=rms


	availabeRooms = availabeRoomId

	initialization.Tpl.ExecuteTemplate(res, "avilableRoomList.html", all)






	 fmt.Println("couont:::::::::::::::",count)

	// fmt.Println("Booking room id:::::::::::::::",bookingRoomId)
	// fmt.Println("all room id:::::::::::::::",roomId)
	// fmt.Println("availabe room id:::::::::::::::",availabeRoomId)
	// fmt.Println("alll rooooommsss:::::::::::::::",rms)


}
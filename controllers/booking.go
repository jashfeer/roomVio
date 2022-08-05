package controllers

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	razorpay "github.com/razorpay/razorpay-go"
	"main.go/initialization"
)

var roomId,hotelid,days,roomnumber,tottelAmount,bookingAmount,balanceAmount int
var checkIn,checkOut,hotelname string
var availabeRooms [] int64


func FilteringProcess(res http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		initialization.Tpl.ExecuteTemplate(res, "index.html", 300)
		return
	}

	ht:=hotel{}
	rm:=room{}
	var err error 
	
	HotelParking := req.FormValue("parking")
	if HotelParking == "true" {
		ht.HotelParking = true
	} else {
		ht.HotelParking = false
	}

	ht.HotelStarrating,err =  strconv.Atoi(req.FormValue("starrating"))
	if err != nil {
		fmt.Println("starrating Err")

		//panic(err)
	}

	HotelWifi := req.FormValue("wifi")
	if HotelWifi == "true" {
		ht.HotelWifi = true
	} else {
		ht.HotelWifi = false
	}

	HotelGameroom := req.FormValue("gameroom")
	if HotelGameroom == "true" {
		ht.HotelGameroom = true
	} else {
		ht.HotelGameroom = false
	}

	HotelFitnesroom := req.FormValue("fitnesroom")
	if HotelFitnesroom == "true" {
		ht.HotelFitnesroom = true
	} else {
		ht.HotelFitnesroom = false
	}

	HotelPlayground := req.FormValue("playground")
	if HotelPlayground == "true" {
		ht.HotelPlayground = true
	} else {
		ht.HotelPlayground = false
	}

	HotelSwimmingpool := req.FormValue("swimmingpool")
	if HotelSwimmingpool == "true" {
		ht.HotelSwimmingpool = true
	} else {
		ht.HotelSwimmingpool = false
	}

	HotelCoffeshop := req.FormValue("coffeeshop")
	if HotelCoffeshop == "true" {
		ht.HotelCoffeshop = true
	} else {
		ht.HotelCoffeshop = false
	}

	HotelRestaurant := req.FormValue("restaurant")
	if HotelRestaurant == "true" {
		ht.HotelRestaurant = true
	} else {
		ht.HotelRestaurant = false
	}


	minRate,err :=  strconv.Atoi(req.FormValue("minprice"))
	if err != nil {
		fmt.Println("minprice Err")

		//panic(err)
	}

	maxRAte,err :=  strconv.Atoi(req.FormValue("maxprice"))
	if err != nil {
		fmt.Println("maxprice Err")

		//panic(err)
	}

	rm.RoomType=req.FormValue("roomtype")

	rm.RoomCapacity,err = strconv.Atoi(req.FormValue("roomCapacity"))
	if err != nil {
		fmt.Println("roomCapacity Err")

		//panic(err)
	}

	
	RoomAc := req.FormValue("ac")
	if RoomAc == "true" {
		rm.RoomAc = true
	} else {
		rm.RoomAc = false
	}

	RoomTv :=req.FormValue("tv")
	if RoomTv == "true" {
		rm.RoomTv = true
	} else {
		rm.RoomTv = false
	}

	RoomAttachedBathroom :=req.FormValue("attachedBathroom")
	if RoomAttachedBathroom == "true" {
		rm.RoomAttachedBathroom = true
	} else {
		rm.RoomAttachedBathroom = false
	}

	RoomService :=req.FormValue("roomService")
	if RoomService == "true" {
		rm.RoomService = true
	} else {
		rm.RoomService = false
	}
 
	RoomWindowView :=req.FormValue("windowView")
	if RoomWindowView == "true" {
		rm.RoomWindowView = true
	} else {
		rm.RoomWindowView = false
	}


	RoomBalconyTerruse :=req.FormValue("balconyterruse")
	if RoomBalconyTerruse == "true" {
		rm.RoomBalconyTerruse = true
	} else {
		rm.RoomBalconyTerruse = false
	}


	// checkIn =req.FormValue("checkIn")
	// checkOut =req.FormValue("checkOut")


	fmt.Println("START.................")

	//count :=len(availabeRooms)
	Room := RoomAndHotel{}
	Rooms := []RoomAndHotel{}
	var s []int

	for _, roomId := range availabeRooms {
		initialization.Db.Table("rooms").
			Joins("join hotels on rooms.room_hotel_id = hotels.hotel_id ").
			// Where("rooms.room_id= ?",roomId).

			Where(&room{RoomId:int(roomId),RoomType: rm.RoomType, RoomCapacity: rm.RoomCapacity, RoomAc: rm.RoomAc, RoomTv: rm.RoomTv, RoomAttachedBathroom: rm.RoomAttachedBathroom, RoomService: rm.RoomService, RoomWindowView: rm.RoomWindowView, RoomBalconyTerruse: rm.RoomBalconyTerruse }).
			Where("rooms.room_price BETWEEN ? AND ?",minRate,maxRAte).
			Where("hotels.hotel_parking = ? OR hotels.hotel_parking = ?", true,ht.HotelParking).
			Where("hotels.hotel_starrating >= ?",ht.HotelStarrating).
			Where("hotels.hotel_wifi = ? OR hotels.hotel_wifi = ?", true,ht.HotelWifi).
			Where("hotels.hotel_gameroom = ? OR hotels.hotel_gameroom = ?", true,ht.HotelGameroom).
			Where("hotels.hotel_fitnesroom = ? OR hotels.hotel_fitnesroom = ?", true,ht.HotelFitnesroom).
			Where("hotels.hotel_playground = ? OR hotels.hotel_playground = ?", true,ht.HotelPlayground).
			Where("hotels.hotel_swimmingpool = ? OR hotels.hotel_swimmingpool = ?", true,ht.HotelSwimmingpool).
			Where("hotels.hotel_coffeshop = ? OR hotels.hotel_coffeshop= ?", true,ht.HotelCoffeshop).
			Where("hotels.hotel_restaurant = ? OR hotels.hotel_restaurant= ?", true,ht.HotelRestaurant).

			Scan(&Room)
			
			if Room.RoomNumber != 0{
				flg := 0

				for _, a := range s {
					if a == Room.RoomId {
						flg = 1
					}
				}
				s = append(s,Room.RoomId)

				if flg!=1 {
					Rooms = append(Rooms,Room)
				}
			}

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

	

	all := All{}
	all.CheckIn=checkIn
	all.CheckOut=checkOut

	all.HotelFilter= ht
	all.RoomFilter= rm
	u := getUser(res, req)
	all.MinPrice=minRate
	all.MaxPrice=maxRAte
	all.U  = u
	all.Rms= Rooms
	
	initialization.Tpl.ExecuteTemplate(res, "avilableRoomList.html", all)

}




func Booking(res http.ResponseWriter, req *http.Request){
	if !alreadyLoggedIn(res, req) {
		http.Redirect(res, req, "/login", http.StatusSeeOther)
		return
	}
	if checkIn=="" || checkOut=="" {
		http.Redirect(res, req, "/showRoomsToUser", http.StatusSeeOther)
		return
	}
	
	roomId, _ = strconv.Atoi(req.FormValue("roomId"))

	type PageVariables struct{
		OrderId string
	}

	type All struct{
		U user
		Rm RoomAndHotel
		BookingAmount int
		TottalAmount int
		CheckIn string
		CheckOut string
		Days int
		HomePageVars PageVariables
	}

	all:=All{}
	all.U=getUser(res, req)
	all.CheckIn=checkIn
	all.CheckOut=checkOut

	result := initialization.Db.Table("rooms").
	Joins("join hotels on rooms.room_hotel_id = hotels.hotel_id").Where("rooms.room_id= ?", roomId).Scan(&all.Rm)
	if result.Error != nil {
		http.Error(res, http.StatusText(500), 500)
		return
	}

	hotelid = all.Rm.HotelId
	roomnumber=all.Rm.RoomNumber
	hotelname=all.Rm.HotelName

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

	difference:= BookingCheckOut.Sub(BookingCheckIn)
	all.Days=int(difference.Hours()/24)
	days=all.Days
	all.TottalAmount=all.Rm.RoomPrice *days

	all.BookingAmount = (all.TottalAmount /10 )* 2

	tottelAmount=all.TottalAmount
	bookingAmount=all.BookingAmount
	balanceAmount=tottelAmount-bookingAmount


	// PAYMENT-create order_id from the server




	client:= razorpay.NewClient("rzp_test_X63gYA8NkKt6KM","Kvx3qVDLs2QGNTWejU4phXZx")
	data := map[string]interface{}{
		"amount": all.BookingAmount*100,
		"currency":"INR",
		"receipt": "some_receipt_id",
	}
	body,err:=client.Order.Create(data,nil)

	if err !=nil{
		fmt.Printf("Problem in getting repository information %v\n",err)
		os.Exit(1)
	}

	//Save Order_id from the body
	value:=body["id"]
	str:=value.(string)

	all.HomePageVars =PageVariables{
		OrderId: str,
	}

	initialization.Tpl.ExecuteTemplate(res, "bookingForm.html", all)



}





func ConformBooking(res http.ResponseWriter, req *http.Request){
	if !alreadyLoggedIn(res, req) {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}




	booking:=booking{}
    var err error
	booking.BookingRoomId=roomId
	booking.BookingHotelId=hotelid
	booking.BookingHotelName=hotelname
	booking.BookingRoomNumber=roomnumber
	booking.BookingAmount=bookingAmount
	booking.BalanceAmount=balanceAmount
	booking.TottalAmount=tottelAmount


	booking.BookingCheckIn, err = time.Parse("2006-01-02", checkIn)
	if err != nil {
		fmt.Println(err)
		return
	}
	booking.BookingCheckOut, err = time.Parse("2006-01-02", checkOut)
	if err != nil {
		fmt.Println(err)
		return
	}


	difference:= booking.BookingCheckOut.Sub(booking.BookingCheckIn)
	booking.BookingDays=int(difference.Hours()/24)

	u := getUser(res, req)
	booking.BookingUserId=u.Id
	booking.BookingUserEmail=u.Email
	booking.BookingStatus="Processing"
	booking.CreatedAt=time.Now()


	booking.BookingRazorpayPaymentId=req.FormValue("razorpay_payment_id")
	booking.BookingRazorpayOrderId=req.FormValue("razorpay_order_id")
	booking.BookingRazorpaySignature=req.FormValue("razorpay_signature")

	fmt.Println("razorpay_payment_id :",booking.BookingRazorpayPaymentId)
	fmt.Println("razorpay_order_id",booking.BookingRazorpayOrderId)
	fmt.Println("razorpay_signature",booking.BookingRazorpaySignature)


	fmt.Println(booking)
	initialization.Db.Select("booking_room_id", "booking_days", "booking_check_in", "booking_check_out", "booking_user_id", "booking_status", "booking_razorpay_payment_id", "booking_razorpay_order_id", "booking_razorpay_signature","created_at","booking_hotel_id","booking_hotel_name","booking_room_number","booking_user_email","booking_amount","balance_amount","tottal_amount").Create(&booking)

	http.Redirect(res, req, "/showUserBookings", http.StatusSeeOther)



}
package controllers

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"main.go/database"
	"main.go/initialization"
)

var hotelId int

var pageRoom int64
var tottalPageRoom int64

func RoomNextPage(res http.ResponseWriter, req *http.Request) {
	if pageRoom == tottalPageRoom {
		http.Redirect(res, req, "/room", http.StatusSeeOther)
	} else {
		pageRoom = pageRoom + 1
		http.Redirect(res, req, "/room", http.StatusSeeOther)

	}
}
func RoomPreviousPage(res http.ResponseWriter, req *http.Request) {
	if pageRoom == 1 {
		http.Redirect(res, req, "/room", http.StatusSeeOther)
	} else {
		pageRoom = pageRoom - 1
		http.Redirect(res, req, "/room", http.StatusSeeOther)

	}
}

func Room(res http.ResponseWriter, req *http.Request) {
	if !alreadyLoggedIn(res, req) {
		http.Redirect(res, req, "/login", http.StatusSeeOther)
		return
	}
	u := getUser(res, req)
	if u.Useraccess != "admin" {
		http.Redirect(res, req, "/", http.StatusSeeOther)
	}
	if req.Method != "GET" {
		http.Error(res, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	if pageRoom == 0 {
		pageRoom = 1
	}
	limit := 5
	offset := (pageRoom - 1) * int64(limit)
	var count int64

	type All struct {
		Rms        []room
		Page       int64
		TottalPage int64
	}
	all := All{}

	rms := make([]room, 0)

	id := req.FormValue("id")
	if id != "" {
		hotelId, _ = strconv.Atoi(id)
	}

	result := initialization.Db.Where("room_hotel_id = ?", hotelId).Order("room_id asc").Limit(limit).Offset(int(offset)).Find(&rms).Count(&count)
	if result.Error != nil {
		http.Error(res, http.StatusText(500), 500)
		return
	}

	all.Rms = rms
	if tottalPageRoom == 0 {
		num := count / int64(limit)
		bal := count % int64(limit)
		if bal != 0 {
			tottalPageRoom = num + 1
		} else {
			tottalPageRoom = num
		}
	}
	all.Page = pageRoom
	all.TottalPage = tottalPageRoom

	initialization.Tpl.ExecuteTemplate(res, "rooms.html", all)
}

func AddRoomFrom(res http.ResponseWriter, req *http.Request) {
	if !alreadyLoggedIn(res, req) {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}
	id := req.FormValue("hotelid")
	hotelId, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("HOTEL ID err1111")
		//panic(err)
	}
	fmt.Println("hotelId in AddRoomFrom: ", hotelId)
	initialization.Tpl.ExecuteTemplate(res, "add_room_form.html", hotelId)
}

func AddRoomProcess(res http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		initialization.Tpl.ExecuteTemplate(res, "add_room_form.html", 300)
		return
	}
	rm := room{}
	var err error
	// id := req.FormValue("HotelId")
	// rm.HotelId, err = strconv.Atoi(id)
	// // rm.HotelId=7
	// if err != nil {
	// 	fmt.Println("HOTEL ID err2222")
	// 	//panic(err)
	// }
	rm.RoomHotelId = hotelId
	rm.RoomNumber, err = strconv.Atoi(req.FormValue("roomNumber"))
	if err != nil {
		fmt.Println("roomNumberErr")

		//panic(err)
	}
	rm.RoomSize, err = strconv.Atoi(req.FormValue("roomSize"))
	if err != nil {
		fmt.Println("roomSizeErr")
		//panic(err)
	}
	rm.RoomType = req.FormValue("roomType")

	rm.RoomCapacity, err = strconv.Atoi(req.FormValue("roomCapacity"))
	if err != nil {
		fmt.Println("roomSizeErr")
		//panic(err)
	}
	ac := req.FormValue("ac")
	if ac == "true" {
		rm.RoomAc = true
	} else {
		rm.RoomAc = false
	}
	tv := req.FormValue("tv")
	if tv == "true" {
		rm.RoomTv = true
	} else {
		rm.RoomTv = false
	}
	wifi := req.FormValue("wifi")
	if wifi == "true" {
		rm.RoomWifi = true
	} else {
		rm.RoomWifi = false
	}
	attachedBathroom := req.FormValue("attachedBathroom")
	if attachedBathroom == "true" {
		rm.RoomAttachedBathroom = true
	} else {
		rm.RoomAttachedBathroom = false
	}
	roomService := req.FormValue("roomService")
	if roomService == "true" {
		rm.RoomService = true
	} else {
		rm.RoomService = false
	}
	windowView := req.FormValue("windowView")
	if windowView == "true" {
		rm.RoomWindowView = true
	} else {
		rm.RoomWindowView = false
	}
	balconyterruse := req.FormValue("balconyterruse")
	if balconyterruse == "true" {
		rm.RoomBalconyTerruse = true
	} else {
		rm.RoomBalconyTerruse = false
	}

	//uploading images.............

	req.ParseMultipartForm(10 << 20)
	//uploading image 1.....

	image1, handler1, err := req.FormFile("image1")
	if err != nil {
		fmt.Println("Error retriving image1")
		fmt.Println(err)
		return
	}
	defer image1.Close()
	ImageFile1, err := ioutil.TempFile("assets/roomimages", "upload-*.png")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer ImageFile1.Close()
	fileBytes1, err := ioutil.ReadAll(image1)
	if err != nil {
		fmt.Println(err)
	}
	ImageFile1.Write(fileBytes1)

	bucket1:= "roomvio"
	key1:=ImageFile1.Name()
	body1,_:=  handler1.Open()
	location1 := database.UploadImage(bucket1,key1,body1)

	//uploading image 2....
	image2, handler2, err := req.FormFile("image2")
	if err != nil {
		fmt.Println("Error retriving image2")
		fmt.Println(err)
		return
	}
	defer image2.Close()

	ImageFile2, err := ioutil.TempFile("assets/roomimages", "upload-*.png")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer ImageFile2.Close()
	fileBytes2, err := ioutil.ReadAll(image2)
	if err != nil {
		fmt.Println(err)
	}
	ImageFile2.Write(fileBytes2)

	bucket2:= "roomvio"
	key2:=ImageFile2.Name()
	body2,_:=  handler2.Open()
	location2 := database.UploadImage(bucket2,key2,body2)

	//uploading image 3.....
	image3, handler3, err := req.FormFile("image3")
	if err != nil {
		fmt.Println("Error retriving image3")
		fmt.Println(err)
		return
	}
	defer image3.Close()
	ImageFile3, err := ioutil.TempFile("assets/roomimages", "upload-*.png")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer ImageFile3.Close()
	fileBytes3, err := ioutil.ReadAll(image3)
	if err != nil {
		fmt.Println(err)
	}
	ImageFile3.Write(fileBytes3)

	bucket3:= "roomvio"
	key3:=ImageFile3.Name()
	body3,_:=  handler3.Open()
	location3 := database.UploadImage(bucket3,key3,body3)

	//uploading image 4....
	image4, handler4, err := req.FormFile("image4")
	if err != nil {
		fmt.Println("Error retriving image3")
		fmt.Println(err)
		return
	}
	defer image3.Close()

	ImageFile4, err := ioutil.TempFile("assets/roomimages", "upload-*.png")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer ImageFile4.Close()
	fileBytes4, err := ioutil.ReadAll(image4)
	if err != nil {
		fmt.Println(err)
	}
	ImageFile4.Write(fileBytes4)

    bucket4:= "roomvio"
	key4:=ImageFile4.Name()
	body4,_:=  handler4.Open()
	location4 := database.UploadImage(bucket4,key4,body4)
	

	rm.RoomImage1 = location1
	rm.RoomImage2 = location2
	rm.RoomImage3 = location3
	rm.RoomImage4 = location4

	rm.RoomPrice, err = strconv.Atoi(req.FormValue("price"))
	if err != nil {
		fmt.Println("priceErr")
		//panic(err)
	}
	rm.RoomIsActive = true
	rm.CreatedAt = time.Now()
	rm.RoomDiscretion = req.FormValue("roomDescription")

	fmt.Println("Successful")

	fmt.Println(rm)

	initialization.Db.Select("room_number", "room_size", "room_type", "room_capacity", "room_ac", "room_tv", "room_wifi", "room_attached_bathroom", "room_service", "room_window_view", "room_balcony_terruse", "room_image1", "room_image2", "room_image3", "room_image4", "room_price", "room_is_active", "room_hotel_id", "created_at", "room_discretion").Create(&rm)
	http.Redirect(res, req, "/hotel", http.StatusSeeOther)

}

package controllers

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"main.go/initialization"
)
var hotelId int

var pageRoom int64
var tottalPageRoom int64


func RoomNextPage(res http.ResponseWriter, req *http.Request) {
	if pageRoom == tottalPageRoom {
		http.Redirect(res, req, "/room", http.StatusSeeOther)
		} else {
		pageRoom = pageRoom+ 1
		http.Redirect(res, req, "/room", http.StatusSeeOther)

	}
}
func RoomPreviousPage(res http.ResponseWriter, req *http.Request) {
	if pageRoom== 1 {
		http.Redirect(res, req, "/room", http.StatusSeeOther)
		} else {
		pageRoom = pageRoom - 1
		http.Redirect(res, req, "/room", http.StatusSeeOther)

	}
}





func Room(res http.ResponseWriter, req *http.Request) {
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


	if pageRoom == 0 {
		pageRoom = 1
	}
	limit := 2
	offset := (pageRoom - 1) * int64(limit)
	var count int64
	
	type All struct{
		Rms   []room
		Page       int64
		TottalPage int64
	}
	all:=All{}

	rms:=make([]room,0)

	id := req.FormValue("id")
	if id !=""{
		hotelId, _ = strconv.Atoi(id)
	}


	result:= initialization.Db.Where("room_hotel_id = ?", hotelId).Order("room_id asc").Limit(limit).Offset(int(offset)).Find(&rms).Count(&count)
	if result.Error != nil {
		http.Error(res, http.StatusText(500), 500)
		return
	}

	all.Rms=rms
	if tottalPageRoom == 0 {
	 num := count / int64(limit)
	 bal := count % int64(limit)
	 if bal != 0 {
		 tottalPageRoom = num + 1
	 } else {
		 tottalPageRoom= num
	 }
 }
 all.Page = pageRoom
 all.TottalPage = tottalPageRoom

	initialization.Tpl.ExecuteTemplate(res, "rooms.html", all)
}

func AddRoomFrom(res http.ResponseWriter, req *http.Request) {
	if !alreadyLoggedIn(res,req) {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}
	id := req.FormValue("hotelid")
	hotelId, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("HOTEL ID err1111")
		//panic(err)
	}
	fmt.Println("hotelId in AddRoomFrom: ",hotelId)
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
	rm.RoomHotelId=hotelId
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

	image1, handler, err := req.FormFile("image1")
	if err != nil {
		fmt.Println("Error retriving image1")
		fmt.Println(err)
		return
	}
	defer image1.Close()
	fmt.Printf("uploaded image1: %+v\n", handler.Filename)
	fmt.Printf("Size image1: %+v\n", handler.Size)
	fmt.Printf("MIME image1: %+v\n", handler.Header)

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

	//uploading image 2....
	image2, handler, err := req.FormFile("image2")
	if err != nil {
		fmt.Println("Error retriving image2")
		fmt.Println(err)
		return
	}
	defer image2.Close()
	fmt.Printf("uploaded image2: %+v\n", handler.Filename)
	fmt.Printf("Size image2: %+v\n", handler.Size)
	fmt.Printf("MIME image2: %+v\n", handler.Header)

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

	//uploading image 3.....
	image3, handler, err := req.FormFile("image3")
	if err != nil {
		fmt.Println("Error retriving image3")
		fmt.Println(err)
		return
	}
	defer image3.Close()
	fmt.Printf("uploaded image3: %+v\n", handler.Filename)
	fmt.Printf("Size image3: %+v\n", handler.Size)
	fmt.Printf("MIME image3: %+v\n", handler.Header)

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

	//uploading image 4....
	image4, handler, err := req.FormFile("image4")
	if err != nil {
		fmt.Println("Error retriving image4")
		fmt.Println(err)
		return
	}
	defer image4.Close()
	fmt.Printf("uploaded image4: %+v\n", handler.Filename)
	fmt.Printf("Size image4: %+v\n", handler.Size)
	fmt.Printf("MIME image4: %+v\n", handler.Header)

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

	rm.RoomImage1 = ImageFile1.Name()
	rm.RoomImage2 = ImageFile2.Name()
	rm.RoomImage3 = ImageFile3.Name()
	rm.RoomImage4 = ImageFile4.Name()

	rm.RoomPrice, err = strconv.Atoi(req.FormValue("price"))
	if err != nil {
		fmt.Println("priceErr")
		//panic(err)
	}
	rm.RoomIsActive = true
	rm.CreatedAt= time.Now()
	rm.RoomDiscretion = req.FormValue("roomDescription")


	

	fmt.Println("Successful")

	fmt.Println(rm)

	initialization.Db.Select("room_number", "room_size", "room_type", "room_capacity", "room_ac", "room_tv", "room_wifi", "room_attached_bathroom", "room_room_service", "room_window_view", "room_balcony_terruse", "room_image1", "room_image2", "room_image3", "iroom_mage4", "room_price", "room_is_active", "room_hotel_id","created_at","room_discretion").Create(&rm)
	http.Redirect(res, req, "/hotel", http.StatusSeeOther)

}

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

var pageHotel int64
var tottalPageHotel int64


func HotelNextPage(res http.ResponseWriter, req *http.Request) {
	if pageHotel == tottalPageHotel {
		http.Redirect(res, req, "/hotel", http.StatusSeeOther)
		} else {
		pageHotel = pageHotel+ 1
		http.Redirect(res, req, "/hotel", http.StatusSeeOther)

	}
}
func HotelPreviousPage(res http.ResponseWriter, req *http.Request) {
	if pageHotel== 1 {
		http.Redirect(res, req, "/hotel", http.StatusSeeOther)
		} else {
		pageHotel = pageHotel - 1
		http.Redirect(res, req, "/hotel", http.StatusSeeOther)

	}
}





func Hotel(res http.ResponseWriter, req *http.Request) {
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

	if pageHotel == 0 {
		pageHotel = 1
	}
	limit := 5
	offset := (pageHotel - 1) * int64(limit)
	var count int64
	
	type All struct{
		Hts   []hotel
		Page       int64
		TottalPage int64
	}
	all:=All{}



	hts:=make([]hotel,0)
	result:= initialization.Db.Order("hotel_id asc").Limit(limit).Offset(int(offset)).Find(&hts).Count(&count)
	if result.Error != nil {
		http.Error(res, http.StatusText(500), 500)
		return
	}
  	 all.Hts=hts
	   if tottalPageHotel == 0 {
		num := count / int64(limit)
		bal := count % int64(limit)
		if bal != 0 {
			tottalPageHotel = num + 1
		} else {
			tottalPageHotel= num
		}
	}
	all.Page = pageHotel
	all.TottalPage = tottalPageHotel
   

	initialization.Tpl.ExecuteTemplate(res, "hotel&room.html", all)
}




func AddHotelFrom(res http.ResponseWriter, req *http.Request) {
	if !alreadyLoggedIn(res,req) {
		http.Redirect(res, req, "/login", http.StatusSeeOther)
		return
	}
	initialization.Tpl.ExecuteTemplate(res, "add_hotel_form.html", nil)
	
}

func AddHotelProcess(res http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		initialization.Tpl.ExecuteTemplate(res, "add_hotel_form.html", 300)
		return
	}
	ht := hotel{}
	ht.HotelName = req.FormValue("hotelname")
	ht.HotelCity = req.FormValue("city")
	ht.HotelPhone = req.FormValue("contactnumber")
	ht.HotelStarrating,_ = strconv.Atoi(req.FormValue("star"))

	if ht.HotelName== ""||ht.HotelCity==""||ht.HotelPhone=="" {
		initialization.Tpl.ExecuteTemplate(res, "add_hotel_form.html", 300)
		return
	}

	cctv := req.FormValue("cctv")
	if cctv == "true" {
		ht.HotelCctv = true
	} else {
		ht.HotelCctv = false
	}
	parking := req.FormValue("parking")
	if parking == "true" {
		ht.HotelParking = true
	} else {
		ht.HotelParking = false
	}
	wifi := req.FormValue("wifi")
	if wifi == "true" {
		ht.HotelWifi = true
	} else {
		ht.HotelWifi = false
	}
	gameroom := req.FormValue("gameroom")
	if gameroom == "true" {
		ht.HotelGameroom = true
	} else {
		ht.HotelGameroom = false
	}
	fitnesscenter := req.FormValue("fitnesscenter")
	if fitnesscenter == "true" {
		ht.HotelFitnesroom = true
	} else {
		ht.HotelFitnesroom = false
	}
	playground := req.FormValue("playground")
	if playground == "true" {
		ht.HotelPlayground = true
	} else {
		ht.HotelPlayground = false
	}
	swimming := req.FormValue("swimming")
	if swimming == "true" {
		ht.HotelSwimmingpool = true
	} else {
		ht.HotelSwimmingpool = false
	}
	coffeeshop := req.FormValue("coffeeshop")
	if coffeeshop == "true" {
		ht.HotelCoffeshop = true
	} else {
		ht.HotelCoffeshop = false
	}
	restaurant := req.FormValue("restaurant")
	if restaurant == "true" {
		ht.HotelRestaurant = true
	} else {
		ht.HotelRestaurant = false
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
	ImageFile1, err := ioutil.TempFile("assets/hotelsImages", "upload-*.png")
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

	ImageFile2, err := ioutil.TempFile("assets/hotelsImages", "upload-*.png")
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

	ImageFile3, err := ioutil.TempFile("assets/hotelsImages", "upload-*.png")
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
		fmt.Println("Error retriving image4")
		fmt.Println(err)
		return
	}
	defer image4.Close()

	ImageFile4, err := ioutil.TempFile("assets/hotelsImages", "upload-*.png")
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


	ht.HotelImage1 = location1
	ht.HotelImage2 = location2
	ht.HotelImage3 = location3
	ht.HotelImage4 = location4

	ht.HotelDiscretion = req.FormValue("hotelDescription")

	ht.HotelIsactive=true
	ht.CreatedAt= time.Now()


	fmt.Println("Successful")

	fmt.Println(ht)

	initialization.Db.Select("created_at","hotel_name", "hotel_city", "hotel_phone", "hotel_starrating", "hotel_cctv", "hotel_parking", "hotel_wifi", "hotel_gameroom", "hotel_fitnesroom", "hotel_playground", "hotel_swimmingpool", "hotel_coffeshop", "hotel_restaurant", "hotel_image1", "hotel_image2", "hotel_image3", "hotel_image4","hotel_isactive","hotel_discretion").Create(&ht)
	http.Redirect(res, req, "/hotel", http.StatusSeeOther)

}

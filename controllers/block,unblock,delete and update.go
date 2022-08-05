package controllers

import (
	"fmt"
	"net/http"

	"main.go/initialization"
)


func BlockUnblockUserProcess(res http.ResponseWriter, req *http.Request){
	if req.Method != "GET" {
		http.Error(res, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}
	id := req.FormValue("id")

	fmt.Println("id:  ",id)

	if id == "" {
		http.Error(res, http.StatusText(400), http.StatusBadRequest)
		return
	}

	var u user
	
	initialization.Db.Where("id = ?", id).Find(&u)

	fmt.Println("isactive :  ",u.Isactive)
	if u.Isactive {
		//result:=initialization.Db.Model(&user{}).Where("id = ?", id).Update("isactive", false)
		result:=initialization.Db.Model(&u).Where("id = ?", id).Update("isactive", false)

		if result.Error != nil {
		http.Error(res, http.StatusText(500), http.StatusInternalServerError)
		return
		}
		http.Redirect(res, req, "/users", http.StatusSeeOther)
	}else{
		result:=initialization.Db.Model(&u).Where("id = ?", id).Update("isactive", true)
		if result.Error != nil {
		http.Error(res, http.StatusText(500), http.StatusInternalServerError)
		return
		}
		http.Redirect(res, req, "/users", http.StatusSeeOther)
	}
}




func BlockUnblockHotelProcess(res http.ResponseWriter, req *http.Request){
	if req.Method != "GET" {
		http.Error(res, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}
	id := req.FormValue("id")

	fmt.Println("id:  ",id)

	if id == "" {
		http.Error(res, http.StatusText(400), http.StatusBadRequest)
		return
	}

	var b hotel
	
	initialization.Db.Where("id = ?", id).Find(&b)

	fmt.Println("isactive :  ",b.HotelIsactive)
	if b.HotelIsactive {
		result:=initialization.Db.Model(&hotel{}).Where("id = ?", id).Update("isactive", false)
		if result.Error != nil {
		http.Error(res, http.StatusText(500), http.StatusInternalServerError)
		return
		}
		http.Redirect(res, req, "/hotel", http.StatusSeeOther)
	}else{
		result:=initialization.Db.Model(&hotel{}).Where("id = ?", id).Update("isactive", true)
		if result.Error != nil {
		http.Error(res, http.StatusText(500), http.StatusInternalServerError)
		return
		}
		http.Redirect(res, req, "/hotel", http.StatusSeeOther)
	}
}


func DeleteUserProcess(res http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		http.Error(res, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}
	id := req.FormValue("id")
	if id == "" {
		http.Error(res, http.StatusText(400), http.StatusBadRequest)
		return
	}
	ur:=user{}
	result:=initialization.Db.Where("id = ?", id).Delete(&ur)
	if result.Error != nil {
		http.Error(res, http.StatusText(500), http.StatusInternalServerError)
		return
	}
	http.Redirect(res, req, "/users", http.StatusSeeOther)
}



func UpdateUserProcess(res http.ResponseWriter, req *http.Request) {

}

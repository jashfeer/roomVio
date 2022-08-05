package controllers

import (
	"fmt"
	"net/http"
	"time"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"main.go/initialization"
)

func SignupFrom(res http.ResponseWriter, req *http.Request) {
	if alreadyLoggedIn(res,req) {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}
	initialization.Tpl.ExecuteTemplate(res, "signup.html", nil)
}

func SignupProcess(res http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		initialization.Tpl.ExecuteTemplate(res, "signup.html", 300)
		return
	}
	ur := user{}
	//ur.Id = uuid.NewV4()
	ur.Firstname = req.FormValue("firstname")
	ur.Lastname = req.FormValue("lastname")
	ur.Email = req.FormValue("email")
	ur.Phone = req.FormValue("phone")
	Password := req.FormValue("password")
	confirmPassword := req.FormValue("password2")
	ur.Useraccess = "user"
	ur.Isactive = true
	fmt.Println(ur)

	if Password != confirmPassword {
		http.Error(res, "Confirm Password is incorrect", http.StatusForbidden)
		return
	}
	result:=initialization.Db.Where("email = ?", ur.Email).First(&ur)
	fmt.Println("errrrrr :",result.Error)
	
	
	switch {
	case result.Error == gorm.ErrRecordNotFound:
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(Password), bcrypt.DefaultCost)
		if err != nil {
			http.Error(res, "Server error, unable to create your account.", 500)
			return
		}
		ur.Password = string(hashedPassword)
		ur.Id = uuid.NewV4()

		ur.CreatedAt = time.Now()
		fmt.Println("user :",ur)

		initialization.Db.Select("id", "firstname", "lastname", "email", "phone", "password", "useraccess", "isactive","created_at").Create(&ur)
		if err != nil {
			http.Error(res, "Server error, unable to create your account.", 500)
			return
		}
		initialization.Db.Where("email = ?", ur.Email).First(&ur)

		//create session
		session, _ := initialization.Store.Get(req, "session")
		session.Values["email"] = ur.Email
		session.Save(req, res)


		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	case result.Error != nil:
		http.Error(res, "Server error, unable to create your account.", 500)
		return
	default:
		fmt.Println("sorrry")
		//n:="user already exist"
		initialization.Tpl.ExecuteTemplate(res, "signup.html",nil)

	//	http.Redirect(res, req, "/signup", 301)

	}

}

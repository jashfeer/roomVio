package main

import (
	"fmt"
	"net/http"

	"main.go/controllers"
	"main.go/initialization"
)

func main() {
	fmt.Println("main func")
	initialization.Init()

	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets", fs))

	//user templates
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/login", controllers.LoginFrome)
	http.HandleFunc("/login/process", controllers.LoginProcess)
	http.HandleFunc("/signup", controllers.SignupFrom)
	http.HandleFunc("/signup/process", controllers.SignupProcess)
	http.HandleFunc("/logout", controllers.Logout)

	http.HandleFunc("/showRoomsToUser", controllers.ShowRoomsToUser)
	http.HandleFunc("/showRoomsToUser/next", controllers.ShowRoomsToUserNextPage)
	http.HandleFunc("/showRoomsToUser/previous", controllers.ShowRoomsToUserPreviousPage)

	http.HandleFunc("/showRoomsToUser/checkingAvailablity", controllers.CheckingAvailablity)

	http.HandleFunc("/filtering", controllers.FilteringProcess)

	http.HandleFunc("/showUserBookings", controllers.ShowUserBookings)
	http.HandleFunc("/userBookings", controllers.UserBookings)
	http.HandleFunc("/userBookings/user/cancel", controllers.UserCancel)
	http.HandleFunc("/conformBooking/next", controllers.UserBookingsNextPage)
	http.HandleFunc("/conformBooking/previous", controllers.UserBookingsPreviousPage)

	//admin templates
	http.HandleFunc("/admin/panal", controllers.AdminPanal)
	http.HandleFunc("/admin/panal/next", controllers.AdminPanalNextPage)
	http.HandleFunc("/admin/panal/previous", controllers.AdminPanalPreviousPage)

	http.HandleFunc("/users", controllers.UsersList)
	http.HandleFunc("/BlockUnblockHotel", controllers.BlockUnblockHotelProcess)
	http.HandleFunc("/BlockUnblockUser", controllers.BlockUnblockUserProcess)
	http.HandleFunc("/deleteUser", controllers.DeleteUserProcess)
	http.HandleFunc("/updateUser", controllers.UpdateUserProcess)
	http.HandleFunc("/users/next", controllers.UsersListNextPage)
	http.HandleFunc("/users/previous", controllers.UsersListPreviousPage)

	http.HandleFunc("/hotel", controllers.Hotel)
	http.HandleFunc("/addHotel", controllers.AddHotelFrom)
	http.HandleFunc("/addHotel/process", controllers.AddHotelProcess)
	http.HandleFunc("/hotel/next", controllers.HotelNextPage)
	http.HandleFunc("/hotel/previous", controllers.HotelPreviousPage)

	http.HandleFunc("/room", controllers.Room)
	http.HandleFunc("/room/addRoom", controllers.AddRoomFrom)
	http.HandleFunc("/addRoom/process", controllers.AddRoomProcess)
	http.HandleFunc("/room/next", controllers.RoomNextPage)
	http.HandleFunc("/room/previous", controllers.RoomPreviousPage)

	http.HandleFunc("/booking", controllers.Booking)
	http.HandleFunc("/conformBooking", controllers.ConformBooking)
	http.HandleFunc("/conformBooking/booked", controllers.Booked)
	http.HandleFunc("/conformBooking/cancel", controllers.Cancel)

	http.ListenAndServe(":8080", nil)
}

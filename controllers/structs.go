package controllers

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type user struct {
	gorm.Model

	Id         uuid.UUID
	Firstname  string
	Lastname   string
	Email      string
	Phone      string
	Password   string
	Useraccess string
	Isactive   bool
	CreatedAt  time.Time
}

type hotel struct {
	gorm.Model

	HotelId           int
	HotelName         string
	HotelCity         string
	HotelPhone        string
	HotelStarrating   int
	HotelCctv         bool
	HotelParking      bool
	HotelWifi         bool
	HotelGameroom     bool
	HotelFitnesroom   bool
	HotelPlayground   bool
	HotelSwimmingpool bool
	HotelCoffeshop    bool
	HotelRestaurant   bool
	HotelImage1       string
	HotelImage2       string
	HotelImage3       string
	HotelImage4       string
	HotelIsactive     bool
	CreatedAt    time.Time
	HotelDiscretion   string
}

type room struct {
	gorm.Model

	RoomId               int
	RoomNumber           int
	RoomSize             int
	RoomType             string
	RoomCapacity         int
	RoomAc               bool
	RoomTv               bool
	RoomWifi             bool
	RoomAttachedBathroom bool
	RoomService          bool
	RoomWindowView       bool
	RoomBalconyTerruse   bool
	RoomImage1           string
	RoomImage2           string
	RoomImage3           string
	RoomImage4           string
	RoomPrice            int
	RoomIsActive         bool
	RoomHotelId          int
	RoomDiscretion       string
	CreatedAt        time.Time
}



type booking struct {
	gorm.Model

	BookingId                int
	BookingRoomId            int
	BookingDays              int
	BookingCheckIn           time.Time
	BookingCheckOut          time.Time
	BookingUserId            uuid.UUID
	BookingStatus            string
	BookingRazorpayPaymentId string
	BookingRazorpayOrderId   string
	BookingRazorpaySignature string
	CreatedAt                time.Time
	BookingHotelId           int
	BookingHotelName         string
	BookingRoomNumber        int
	BookingUserEmail         string
	BookingAmount            int
	BalanceAmount            int
	TottalAmount             int
}




type RoomAndHotel struct {
	HotelId              int
	HotelName            string
	HotelCity            string
	HotelPhone           string
	HotelStarrating      int
	HotelCctv            bool
	HotelParking         bool
	HotelWifi            bool
	HotelGameroom        bool
	HotelFitnesroom      bool
	HotelPlayground      bool
	HotelSwimmingpool    bool
	HotelCoffeshop       bool
	HotelRestaurant      bool
	HotelImage1          string
	HotelImage2          string
	HotelImage3          string
	HotelImage4          string
	HotelIsactive        bool
	HotelDiscretion      string
	RoomId               int
	RoomNumber           int
	RoomSize             int
	RoomType             string
	RoomCapacity         string
	RoomAc               bool
	RoomTv               bool
	RoomWifi             bool
	RoomAttachedBathroom bool
	RoomService          bool
	RoomWindowView       bool
	RoomBalconyTerruse   bool
	RoomImage1           string
	RoomImage2           string
	RoomImage3           string
	RoomImage4           string
	RoomPrice            int
	RoomIsActive         bool
	RoomHotelId          int
	RoomDiscretion       string
	BookingId                int
	BookingRoomId            int
	BookingDays              int
	BookingCheckIn           time.Time
	BookingCheckOut          time.Time
	BookingUserId            uuid.UUID
	BookingStatus            string
	BookingRazorpayPaymentId string
	BookingRazorpayOrderId   string
	BookingRazorpaySignature string
	CreatedAt                time.Time
	BookingHotelId           int
	BookingHotelName         string
	BookingRoomNumber        int
	BookingUserEmail         string
	BookingAmount            int
	BalanceAmount            int
	TottalAmount             int
}



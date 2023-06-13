package app

import (
	"my-app/controller"

	log "github.com/sirupsen/logrus"
)

func mapUrls() {

	//LOGIN
	router.POST("/user-login", controller.UserLogin)

	//URL DE FUNCIONES HOTEL
	router.POST("/add-hotel", controller.InsertHotel)
	router.GET("/get-hotel/:id", controller.GetHotelById)
	router.GET("/all-hotels", controller.GetHotels)

	//URL DE FUNCIONES USUARIO
	router.POST("/create-user", controller.InsertUser)
	router.GET("/user/:id", controller.GetUserById)
	router.GET("/all-users", controller.GetUsers)

	//URL DE FUNCIONES RESERVA
	router.POST("/new-reservation", controller.InsertReservation)
	router.GET("/particular-reservation/:id", controller.GetReservationById)
	router.GET("/all-reservation", controller.GetReservations)
	router.GET("/user/reservations-by-user/:id", controller.GetReservationsByUser)
	router.GET("hotel/reservations-by-hotel/:id", controller.GetReservationsByHotel)

	log.Info("Finishing mappings configurations")
}

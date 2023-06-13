package client

import (
	"my-app/model"

	log "github.com/sirupsen/logrus"
)

// Inserta hotel en la db
func InsertReservation(reservation model.Reservation) model.Reservation {

	result := Db.Create(&reservation)

	if result.Error != nil {
		log.Error("Error en la reserva")
	}

	log.Debug("Reserva confirmada", reservation.Id)
	return reservation
}

// GET de reserva particular by id de reserva
func GetReservationById(id int) model.Reservation {
	var reservation model.Reservation

	Db.Where("id = ?", id).First(&reservation)
	log.Debug("Reserva: ", reservation)

	return reservation
}

// GET de todas las reservas //Punto importante del tp
// Puede verlo tanto el user como el admin->( debe ver detalles d users x reserva)<- comprobar !
func GetReservations() model.Reservations {
	var reservations model.Reservations
	Db.Find(&reservations)

	log.Debug("Reservas: ", reservations)

	return reservations
}

// GET todas las reservas de algun hotel particular
func GetReservationsByHotel(hotelId int) model.Reservations {
	var reservations model.Reservations

	Db.Where("hotel_id = ?", hotelId).Find(&reservations)
	log.Debug("Reserva: ", reservations)

	return reservations
}

// GET de todas las reservas by user  --> unicamente el admin!
func GetReservationsByUser(userId int) model.Reservations {
	var reservations model.Reservations

	Db.Where("user_id = ?", userId).Find(&reservations)
	log.Debug("Reserva: ", reservations)

	return reservations
}

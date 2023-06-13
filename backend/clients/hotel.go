package client

import (
	"my-app/model"
	"time"

	log "github.com/sirupsen/logrus"
)

//FUNCION HOTELES ->DISPONIBLES<- VER!

// hotel por un ID propio
func GetHotelById(id int) model.Hotel {
	var hotel model.Hotel

	Db.Where("id = ?", id).First(&hotel)
	log.Debug("Hotel: ", hotel)

	return hotel
}

// insertar nuevos hoteles --> unicamente by admin
func InsertHotel(hotel model.Hotel) model.Hotel {

	result := Db.Create(&hotel)

	if result.Error != nil {
		log.Error("Error de insercion")
	}

	log.Debug("Insercion Completa ", hotel.Id)
	return hotel
}

// encuentra los hoteles -->> hoteles DISPONBLES, ->VER!
func GetHotels() model.Hotels {
	var hotels model.Hotels
	Db.Find(&hotels)

	log.Debug("Hoteles Disponibles: ", hotels)

	return hotels
}

func GetHotelsByIdAndDates(id int, startDate time.Time, endDate time.Time) model.Hotels {
	var hotels model.Hotels
	Db.Where("id = ? AND fecha_inicio <= ? AND fecha_fin >= ?", id, endDate, startDate).Find(&hotels)

	log.Debug("Hoteles Disponibles por ID y Fechas: ", hotels)

	return hotels
}

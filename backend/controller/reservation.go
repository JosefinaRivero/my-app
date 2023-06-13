package controller

import (
	"my-app/dto"
	service "my-app/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func GetReservationById(c *gin.Context) {
	log.Debug("User id to load: " + c.Param("id"))

	id, _ := strconv.Atoi(c.Param("id"))
	var reservationDto dto.ReservationDto

	reservationDto, err := service.Service.GetReservationById(id)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, reservationDto)
}

func GetReservations(c *gin.Context) {

	reservationsDto, err := service.Service.GetReservations()

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, reservationsDto)
}

func GetReservationsByUser(c *gin.Context) {
	log.Debug("User id to load: " + c.Param("id"))

	id, _ := strconv.Atoi(c.Param("id"))
	var userReservations dto.UserReservationsDto

	userReservations, err := service.Service.GetReservationsByUser(id)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, userReservations)
}

func GetReservationsByHotel(c *gin.Context) {
	log.Debug("User id to load: " + c.Param("id"))

	id, _ := strconv.Atoi(c.Param("id"))
	var hotelReservations dto.HotelReservationsDto

	hotelReservations, err := service.Service.GetReservationsByHotel(id)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, hotelReservations)

}

//acá nos falta que haga las piezas disponibles y los errores de reservacion además de las reservaciones disponibles

func InsertReservation(c *gin.Context) {
	var reservationDto dto.ReservationDto
	err := c.BindJSON(&reservationDto)

	// Error Parsing json param
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	reservationDto, er := service.Service.InsertReservation(reservationDto)
	// Error del Insert
	if er != nil {
		c.JSON(http.StatusBadRequest, er.Error())
	}

	c.JSON(http.StatusCreated, reservationDto)
}

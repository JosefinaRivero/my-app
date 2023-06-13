package controller

import (
	"my-app/dto"
	service "my-app/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

//ERRORES VER!

func GetHotelById(c *gin.Context) {
	log.Debug("User id to load: " + c.Param("id"))

	id, _ := strconv.Atoi(c.Param("id"))
	var hotelDto dto.HotelDto

	hotelDto, err := service.Service.GetHotelById(id)

	if err != nil {
		c.JSON(http.StatusNotFound, err.Error()) //CAMBIAR

		return //c.JSON(err.Status(),err) posible solución
	}
	c.JSON(http.StatusOK, hotelDto)
}

func GetHotels(c *gin.Context) {
	var hotelsDto dto.HotelsDto
	hotelsDto, err := service.Service.GetHotels()

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, hotelsDto)
}

func InsertHotel(c *gin.Context) { // esto lo tiene distinto, lo tiene con verificación de admin
	var hotelDto dto.HotelDto
	err := c.BindJSON(&hotelDto)

	// Error Parsing json param
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	userDto, er := service.Service.InsertHotel(hotelDto) // en vez de userDto ver si es hotelDto
	//Error del insert
	if er != nil {
		c.JSON(er.Status(), er)
	}

	c.JSON(http.StatusCreated, userDto) // vuelve a usar
}

package controller

import (
	"my-app/dto"
	service "my-app/service"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// ------->>>>>> VER  --> TOKEN

func UserLogin(c *gin.Context) {
	var loginDto dto.UserDto

	err := c.BindJSON(&loginDto)

	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	loginDto, er := service.Service.UserLogin(loginDto)

	if er != nil {
		c.JSON(http.StatusUnauthorized, er.Error())
	}

	// VER TOKEN

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = loginDto.Id
	claims["role"] = loginDto.Role
	claims["expiration"] = time.Now().Add(time.Hour * 24).Unix()

	c.JSON(http.StatusAccepted, loginDto)
}

func InsertUser(c *gin.Context) { //verifica si el email ya existe , si los campos están vacios.. eso falta
	var userDto dto.UserDto
	err := c.BindJSON(&userDto)

	// Error Parsing json param
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	userDto, er := service.Service.InsertUser(userDto)
	// Error del Insert
	if er != nil {
		c.JSON(http.StatusBadRequest, er.Error())
	}

	c.JSON(http.StatusCreated, userDto)
}

func GetUserById(c *gin.Context) {
	log.Debug("User id to load: " + c.Param("id"))

	id, _ := strconv.Atoi(c.Param("id"))
	var userDto dto.UserDto

	userDto, err := service.Service.GetUserById(id)

	if err != nil {
		c.JSON(http.StatusNotFound, err.Error())
		return
	}
	c.JSON(http.StatusOK, userDto)
}

func GetUsers(c *gin.Context) {

	var usersDto dto.UsersDto

	usersDto, err := service.Service.GetUsers()

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, usersDto)
}

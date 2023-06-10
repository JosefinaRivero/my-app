package client

import (
	"my-app/model"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

var Db *gorm.DB

//LOGIN DE USUARIOS / VER DIFERENCIAR ENTRE USER Y ADMIN

// Crea un usuario
func InsertUser(user model.User) model.User {

	result := Db.Create(&user)

	if result.Error != nil {
		log.Error("Error de Insercion de Usuario")
	}

	log.Debug("Usuario Creado:", user.Id)
	return user
}

// GET de user by id
func GetUserById(id int) model.User {
	var user model.User

	Db.Where("id = ?", id).First(&user)
	log.Debug("Usuario: ", user)

	return user
}

// GET de user by email
func GetUserByEmail(email string) model.User {
	var user model.User

	Db.Where("email = ?", email).First(&user)
	log.Debug("Usuario: ", user)

	return user
}

// GET all users
func GetUsers() model.Users {
	var users model.Users
	Db.Find(&users)

	log.Debug("Usuarios: ", users)

	return users
}

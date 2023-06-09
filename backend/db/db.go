package db

import (
	"my-app/clients"
	"my-app/model"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	log "github.com/sirupsen/logrus"
)

var (
	db  *gorm.DB
	err error
)

func init() {
	// DB Connections Paramters
	DBName := "arqsoft"
	DBUser := "manu"
	DBPass := "hello"
	//DBPass := os.Getenv("MVC_DB_PASS")
	DBHost := "localhost"

	// DB Connections Paramters
	//DBName := "uccarqsoft"
	//DBUser := "rovulcano"
	//DBPass := "rocioucc333"
	//DBHost := "localhost"
	// ------------------------

	db, err = gorm.Open("mysql", DBUser+":"+DBPass+"@tcp("+DBHost+":3306)/"+DBName+"?charset=utf8&parseTime=True")

	if err != nil {
		log.Info("Connection Failed to Open")
		log.Fatal(err)
	} else {
		log.Info("Connection Established")
	}

	// Add all clients here
	client.Db = db

}

func StartDbEngine() {
	// Migrate all model classes
	db.AutoMigrate(&model.Hotel{})
	db.AutoMigrate(&model.Reservation{})
	db.AutoMigrate(&model.User{})

	log.Info("Finishing Migration Database Tables")
}

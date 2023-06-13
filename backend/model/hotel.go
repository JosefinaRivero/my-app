package model

// estructura hotel
type Hotel struct {
	Id               string  `gorm:"primaryKey"`
	Name             string  `gorm:"type:varchar(50); not null"`
	HotelDescription string  `gorm:"type:varchar(500)"`           //descripcion de 500caract
	Price            float64 `gorm:"type:decimal(9,3); not null"` //precio de hotel//indica la cantidad de cifras ant y desp de la coma /price(100.000.000)
	RoomQuant        int     `gorm:"type:int; not null"`          //cambio
	RoomDescription  string  `gorm:"type:varchar(500)"`           //cambio
}
type Hotels []Hotel

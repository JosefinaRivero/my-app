package model

// estructura RESERVA
type Reservation struct {
	Id            int     `gorm:"primaryKey"`
	StartDate     string  `gorm:"type:varchar(10); not null"`  //DD-MM-YYYY
	DepartureDate string  `gorm:"type:varchar(10); not null"`  //DD-MM-YYYY
	RPrice        float64 `gorm:"type:decimal(8,2); not null"` //precio de la reserva
	UserId        int     `gorm:"foreignkey:UserId"`           //usuario por reserva
	HotelId       int     `gorm:"foreignkey:HotelId"`
}

type Reservations []Reservation

package dto

//RESERVAS GENERALES por user+hotel --> todas las reservas

type ReservationDto struct {
	Id            int     `json:"id"` //reserva id
	StartDate     string  `json:"start_date"`
	DepartureDate string  `json:"departure_date"`
	RPrice        float64 `json:"reservation_price"`
	UserId        int     `json:"user_id"`  //user id
	HotelId       int     `json:"hotel_id"` //hotel id
}

type ReservationsDto []ReservationDto

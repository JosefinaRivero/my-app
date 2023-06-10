package dto

//RESERVAS POR HOTEL

type HotelReservationsDto struct {
	HotelId          int     `json:"hotel_id"`
	HotelName        string  `json:"hotel_name"`
	HotelDescription string  `json:"hotel_description"`
	HotelPrice       float64 `json:"hotel_price"`
	HotelRoomQuant   int     `json:"hotel_room_quantity"`
	RoomDescription  string  `json:"room_description"`

	Reservations ReservationsDto `json:"reservations"`
}

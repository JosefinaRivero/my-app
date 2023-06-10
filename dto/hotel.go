package dto

//HOTEL

type HotelDto struct {
	Id               int     `json:"id"`
	Name             string  `json:"name"`
	HotelDescription string  `json:"hotel description"`
	Price            float64 `json:"price"`
	RoomQuant        int     `json:"room quantity"`
	RoomDescription  string  `json:"room description"`
}

type HotelsDto []HotelDto

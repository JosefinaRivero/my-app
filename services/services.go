package service

import (
	"errors"
	"math"
	"my-app/client"
	"my-app/dto"
	"my-app/model"
	"time"
)

type service struct{}

type serviceInterface interface {
	//funciones del controller
	GetHotelById(id int) (dto.HotelDto, error)
	GetHotels() (dto.HotelsDto, error)
	InsertHotel(hotelDto dto.HotelDto) (dto.HotelDto, error) //--> admin

	UserLogin(loginDto dto.UserDto) (dto.UserDto, error) //ver token
	InsertUser(userDto dto.UserDto) (dto.UserDto, error)
	GetUserById(id int) (dto.UserDto, error)
	GetUsers() (dto.UsersDto, error)

	GetReservationById(id int) (dto.ReservationDto, error)
	GetReservations() (dto.ReservationsDto, error)
	GetReservationsByUser(userId int) (dto.UserReservationsDto, error) //-->admin
	GetReservationsByHotel(hotelId int) (dto.HotelReservationsDto, error)
	InsertReservation(reservationDto dto.ReservationDto) (dto.ReservationDto, error)

	ReservationStatus(hotelId int, startDate time.Time, departureDate time.Time) bool //ver
}

var (
	Service serviceInterface
)

func init() {
	Service = &service{}
}

func (s *service) InsertUser(userDto dto.UserDto) (dto.UserDto, error) {
	var user model.User

	user.Name = userDto.Name
	user.LastName = userDto.LastName
	user.Email = userDto.Email
	user.Password = userDto.Password
	//user.Role = "Customer" --> vean

	user = client.InsertUser(user)

	userDto.Id = user.Id //crea id
	//userDto.Role = user.Role //crea el rol

	if user.Id == 0 {
		return userDto, errors.New("error")
	}

	return userDto, nil
}

func (s *service) GetUserById(id int) (dto.UserDto, error) {

	var user model.User = client.GetUserById(id)

	var userDto dto.UserDto

	if user.Id == 0 {
		return userDto, errors.New("Este usuario no existe")
	}

	userDto.Id = user.Id
	userDto.Name = user.Name
	userDto.LastName = user.LastName
	userDto.Email = user.Email
	//userDto.Role = user.Role

	return userDto, nil
}

func (s *service) GetUsers() (dto.UsersDto, error) {
	var users model.Users = client.GetUsers()
	var usersDto dto.UsersDto

	for _, user := range users {
		var userDto dto.UserDto
		userDto.Id = user.Id
		userDto.Name = user.Name
		userDto.LastName = user.LastName
		userDto.Email = user.Email
		//userDto.Role = user.Role

		usersDto = append(usersDto, userDto)
	}

	return usersDto, nil
}

func (s *service) UserLogin(loginDto dto.UserDto) (dto.UserDto, error) {

	var user = client.GetUserByEmail(loginDto.Email)

	if user.Id == 0 {
		return loginDto, errors.New("Este usuario no existe")
	}

	if user.Password != loginDto.Password {
		return loginDto, errors.New("Contraseña incorrecta. *-1 Intento*")
	}
	loginDto.Id = user.Id
	loginDto.Name = user.Name
	loginDto.LastName = user.LastName
	//loginDto.Role = user.Role
	return loginDto, nil
}

func (s *service) InsertHotel(hotelDto dto.HotelDto) (dto.HotelDto, error) {
	var hotel model.Hotel

	hotel.Name = hotelDto.Name
	hotel.Description = hotelDto.Description
	hotel.RoomQuant = hotelDto.RoomQuant
	hotel.RPrice = hotelDto.RPrice
	hotel.RoomDescription = hotelDto.RoomDescription

	hotel = client.InsertHotel(hotel)

	hotelDto.Id = hotel.Id //crea un id al hotel

	if hotel.Id == 0 {
		return hotelDto, errors.New("Error")
	}

	return hotelDto, nil
}

func (s *service) GetHotels() (dto.HotelsDto, error) {

	var hotels model.Hotels = client.GetHotels() //todos los hoteles
	var hotelsDto dto.HotelsDto

	for _, hotel := range hotels {
		var hotelDto dto.HotelDto

		hotelDto.Id = hotel.Id
		hotelDto.Name = hotel.Name
		hotelDto.RoomQuant = hotel.RoomQuant
		hotelDto.Description = hotel.Description
		hotelDto.Price = hotel.Price
		hotelDto.RoomDescription = hotel.RoomDescription

		hotelsDto = append(hotelsDto, hotelDto)
	}

	return hotelsDto, nil
}

func (s *service) GetHotelById(id int) (dto.HotelDto, error) {

	var hotel model.Hotel = client.GetHotelById(id) //h9tel by id
	var hotelDto dto.HotelDto

	if hotel.Id == 0 {
		return hotelDto, errors.New("Este hotel no existe")
	}
	hotelDto.Id = hotel.Id
	hotelDto.Name = hotel.Name
	hotelDto.RoomQuant = hotel.RoomQuant
	hotelDto.Description = hotel.Description
	hotelDto.Price = hotel.Price
	hotelDto.RoomDescription = hotel.RoomDescription

	return hotelDto, nil
}

// VEEEEEEEEEEEER

func (s *service) InsertReservation(reservationDto dto.ReservationDto) (dto.ReservationDto, error) {

	userDto := client.GetUserById(reservationDto.UserId)
	hotelDto := client.GetHotelById(reservationDto.HotelId)

	if hotelDto.Id == 0 {
		return reservationDto, errors.New("Este hotel no existe")
	}

	if userDto.Id == 0 {
		return reservationDto, errors.New("Este usuario no existe")
	}

	timeStart, _ := time.Parse("02-01-2006 15:04", reservationDto.StartDate)
	timeEnd, _ := time.Parse("02-01-2006 15:04", reservationDto.DepartureDate)

	if timeStart.After(timeEnd) {
		return reservationDto, errors.New("a reservation cant end before it starts")
	}

	if s.ReservationStatus(reservationDto.HotelId, timeStart, timeEnd) {
		var reservation model.Reservation

		reservation.StartDate = reservationDto.StartDate
		reservation.DepartureDate = reservationDto.DepartureDate
		reservation.HotelId = reservationDto.HotelId
		reservation.UserId = reservationDto.UserId

		hours := timeEnd.Sub(timeStart).Hours()
		nights := math.Ceil(hours / 24)
		price := hotelDto.Price
		reservation.RPrice = price * nights

		reservation = client.InsertReservation(reservation)

		reservationDto.Id = reservation.Id
		reservationDto.RPrice = reservation.RPrice

		return reservationDto, nil
	}

	return reservationDto, errors.New("Hotel no disponible") //x habitaciones completas
}

func (s *service) GetReservationById(id int) (dto.ReservationDto, error) {
	var reservation model.Reservation
	var reservationDto dto.ReservationDto

	reservation = client.GetReservationById(id)

	if reservation.Id == 0 {
		return reservationDto, errors.New("Reserva Inexistente")
	}

	reservationDto.Id = reservation.Id
	reservationDto.UserId = reservation.UserId
	reservationDto.HotelId = reservation.HotelId
	reservationDto.StartDate = reservation.StartDate
	reservationDto.DepartureDate = reservation.DepartureDate
	reservationDto.RPrice = reservation.RPrice

	return reservationDto, nil
}

func (s *service) GetReservations() (dto.ReservationsDto, error) {

	var reservations model.Reservations = client.GetReservations()
	var reservationsDto dto.ReservationsDto

	for _, reservation := range reservations {
		var reservationDto dto.ReservationDto

		reservationDto.Id = reservation.Id
		reservationDto.UserId = reservation.UserId
		reservationDto.HotelId = reservation.HotelId
		reservationDto.StartDate = reservation.StartDate
		reservationDto.DepartureDate = reservation.DepartureDate
		reservationDto.RPrice = reservation.RPrice

		reservationsDto = append(reservationsDto, reservationDto)
	}

	return reservationsDto, nil
}

func (s *service) GetReservationsByUser(userId int) (dto.UserReservationsDto, error) {
	var user model.User = client.GetUserById(userId)
	var userReservationsDto dto.UserReservationsDto
	var reservationsDto dto.ReservationsDto

	if user.Id == 0 {
		return userReservationsDto, errors.New("Este usuario no existe")
	}
	var reservations model.Reservations = client.GetReservationsByUser(userId)

	userReservationsDto.UserId = user.Id
	userReservationsDto.UserName = user.Name
	userReservationsDto.UserLastName = user.LastName
	userReservationsDto.UserEmail = user.Email
	userReservationsDto.UserPassword = user.Password

	for _, reservation := range reservations {
		var reservationDto dto.ReservationDto

		reservationDto.Id = reservation.Id
		reservationDto.StartDate = reservation.StartDate
		reservationDto.DepartureDate = reservation.DepartureDate
		reservationDto.HotelId = reservation.HotelId
		reservationDto.UserId = reservation.UserId
		reservationDto.RPrice = reservation.RPrice

		reservationsDto = append(reservationsDto, reservationDto)
	}

	userReservationsDto.Reservations = reservationsDto

	return userReservationsDto, nil
}

func (s *service) GetReservationsByHotel(hotelId int) (dto.HotelReservationsDto, error) {
	var hotel model.Hotel = client.GetHotelById(hotelId)
	var hotelReservations dto.HotelReservationsDto
	var reservationsDto dto.ReservationsDto

	if hotel.Id == 0 {
		return hotelReservations, errors.New("Este hotel no existe")
	}

	var reservations model.Reservations = client.GetReservationsByHotel(hotelId)

	hotelReservations.HotelId = hotel.Id
	hotelReservations.HotelName = hotel.Name
	hotelReservations.HotelDescription = hotel.Description
	hotelReservations.HotelRoomQuant = hotel.RoomQuant
	hotelReservations.HotelRoomDescription = hotel.RoomDescription
	hotelReservations.HotelPrice = hotel.Price

	for _, reservation := range reservations {
		var reservationDto dto.ReservationDto
		reservationDto.Id = reservation.Id
		reservationDto.StartDate = reservation.StartDate
		reservationDto.DepartureDate = reservation.DepartureDate
		reservationDto.HotelId = reservation.HotelId
		reservationDto.UserId = reservation.UserId
		reservationDto.RPrice = reservation.RPrice

		reservationsDto = append(reservationsDto, reservationDto)
	}

	hotelReservations.Reservations = reservationsDto

	return hotelReservations, nil
}

//VEEEEEEEEEEER

func (s *service) ReservationStatus(hotelId int, startDate time.Time, departureDate time.Time) bool {

	hotel := client.GetHotelById(hotelId)
	reservations := client.GetReservationsByHotel(hotelId)

	roomsAvailable := hotel.Price

	for _, reservation := range reservations {

		reservationStart, _ := time.Parse("02-01-2006 15:04", reservation.StartDate)
		reservationEnd, _ := time.Parse("02-01-2006 15:04", reservation.DepartureDate)

		if reservationStart.After(startDate) && reservationEnd.Before(departureDate) ||
			reservationStart.Before(startDate) && reservationEnd.After(startDate) ||
			reservationStart.Before(departureDate) && reservationEnd.After(departureDate) ||
			reservationStart.Before(startDate) && reservationEnd.After(departureDate) ||
			reservationStart.Equal(startDate) || reservationEnd.Equal(departureDate) {
			roomsAvailable--
		}
		if roomsAvailable == 0 {
			return false
		}
	}

	return true
}

package model

// estructura user
type User struct {
	Id       int    `gorm:"primaryKey"`
	Name     string `gorm:"type:varchar(50);not null"`
	LastName string `gorm:"type:varchar(50);not null"`
	Account  string `gorm:"type:varchar(10);not null;unique"`
	Email    string `gorm:"type:varchar(30);not null;"`
	Password string `gorm:"type:varchar(10);not null"`
}

type Users []User

package models

import (
	"dating-sim/db"
	"time"

	"gorm.io/gorm"
)

type Preferance struct {
	gorm.Model
	InterestedIn  string `json:"interested_in"`
	AgeRange      uint   `json:"age_range"`
	MaxDistanceKm uint   `json:"max_distance_km"`
}

type Picture struct {
	gorm.Model
	Path         string `json:"path"`
	Mime         string `json:"mime"`
	OriginalName string `json:"orignal_name,omitempty"`
	UserID       uint   `json:"user_id"`
}

type User struct {
	gorm.Model
	Username       string     `gorm:"not null" json:"username"`
	Email          string     `gorm:"not null" json:"email"`
	Password       string     `gorm:"not null" json:"password,omitempty"`
	Gender         string     `gorm:"not null" json:"gender"`
	Age            uint       `json:"age"`
	Bio            string     `json:"bio"`
	DateOfBirth    *time.Time `json:"date_of_birth"`
	Location       string     `json:"location"`
	PreferanceID   uint       `json:"preferance_id"`
	Preferance     Preferance `gorm:"constraint:OnDelete:CASCADE" json:"preferance"`
	ProfilePicture string     `json:"profile_picture"`
	Pictures       []Picture  `gorm:"foreignKey:UserID" json:"pictures"`
	Premium        *time.Time `json:"premium"`
}

func (user *User) Get() error {
	if query := db.DB.Omit("password").First(&user); query.Error != nil {
		return query.Error
	}
	return nil
}

func (user *User) GetAll(users *[]User, size int, page int) error {
	if query := db.DB.
		Model(&user).
		Omit("password, premium").
		Scopes((db.Paginate(size, page))).
		Find(&users); query.Error != nil {
		return query.Error
	}
	return nil
}

func (user *User) GetWithPassword() error {
	if query := db.DB.First(&user); query.Error != nil {
		return query.Error
	}
	return nil
}

func (user *User) Create() error {
	if query := db.DB.Create(&user); query.Error != nil {
		return query.Error
	}
	return nil
}

func (user *User) Update() error {
	if query := db.DB.Updates(&user); query.Error != nil {
		return query.Error
	}
	return nil
}

func (user *User) GetUserByUsername() error {
	if query := db.DB.First(&user, "username = ?", user.Username); query.Error != nil {
		return query.Error
	}
	return nil
}

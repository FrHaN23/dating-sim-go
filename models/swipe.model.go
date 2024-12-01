package models

import (
	"dating-sim/db"

	"gorm.io/gorm"
)

type Swipe struct {
	gorm.Model
	UserID    uint `gorm:"not null"`
	ProfileID uint `gorm:"not null"`
	Action    bool `gorm:"not null"` // The action: 1 ="right" (like) or 0 ="left" (pass)
}

func (s *Swipe) ProcessSwipe(userID, profileID uint, swipeType bool) error {
	// Save the swipe action in the database using GORM
	swipe := Swipe{
		UserID:    uint(userID),
		ProfileID: uint(profileID),
		Action:    swipeType,
	}

	// Use GORM to create the swipe record
	return db.DB.Create(&swipe).Error
}

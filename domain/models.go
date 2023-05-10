package domain

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint           `gorm:"primaryKey"`
	Name      string         `gorm:"size:255" json:"name"`
	Phone     string         `gorm:"size:20" json:"phone"`
	Email     string         `gorm:"size:255" json:"email"`
	Photo     string         `gorm:"size:255" json:"photo"`
	Password  string         `json:"password"`
	Type      string         `json:"type"`
	Units     []Unit         `gorm:"many2many:unit_owners" json:"units"`
	CreatedAt time.Time      `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt time.Time      `json:"updatedAt" gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}

type Unit struct {
	ID        uint           `gorm:"primaryKey"`
	Avenue    string         `gorm:"size:255" json:"name"`
	Number    string         `json:"password"`
	Type      string         `json:"type"`
	Photo     string         `gorm:"size:255" json:"photo"`
	Owners    []User         `gorm:"many2many:unit_owners" json:"owners"`
	CreatedAt time.Time      `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt time.Time      `json:"updatedAt" gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}

type Guest struct {
	ID        uint           `gorm:"primaryKey"`
	Name      string         `gorm:"size:255" json:"name"`
	DocNumber string         `gorm:"size:30" json:"docNumber"`
	DocType   string         `gorm:"size:30" json:"docType"`
	Photo     string         `gorm:"size:255" json:"photo"`
	Unit      Unit           `json:"unit"`
	UnitID    uint           `json:"unitID"`
	Phone     string         `gorm:"size:20" json:"phone"`
	Email     string         `gorm:"size:255" json:"email"`
	Bookings  []Booking      `gorm:"many2many:booking_guests" json:"bookings"`
	CreatedAt time.Time      `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt time.Time      `json:"updatedAt" gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}

type Booking struct {
	ID        uint           `gorm:"primaryKey"`
	Name      string         `gorm:"size:255" json:"name"`
	StartAt   time.Time      `json:"startAt"`
	EndAt     time.Time      `json:"endAt"`
	Guests    []Guest        `gorm:"many2many:booking_guests" json:"guests"`
	CreatedAt time.Time      `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt time.Time      `json:"updatedAt" gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}

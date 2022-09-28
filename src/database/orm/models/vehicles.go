package models

import "time"

type Vehicle struct {
	VehicleId   uint      `gorm:"primaryKey" json:"vehicle_id" form:"vehicle_id"`
	VehicleName string    `json:"vehicle_name" form:"vehicle_name"`
	Location    string    `json:"location,omitempty" form:"location"`
	Description string    `json:"description,omitempty" form:"description"`
	Price       int       `json:"price,omitempty" form:"price"`
	Status      string    `json:"status,omitempty" form:"status"`
	Stock       int       `json:"stock,omitempty" form:"stock"`
	Category    string    `json:"category,omitempty" form:"category"`
	Image       string    `json:"image,omitempty" form:"image"`
	CreatedAt   time.Time `gorm:"default:current_timestamp" json:"created_at" form:"created_at"`
	UpdatedAt   time.Time `gorm:"default:current_timestamp" json:"updated_at" form:"updated_at"`
	Rating      float32   `json:"rating,omitempty" form:"rating"`
	TotalRented int       `json:"total_rented,omitempty" form:"total_rented"`
}

type Vehicles []Vehicle

package models

import (
	"encoding/json"
	"time"
)

type Vehicle struct {
	VehicleId   uint      `gorm:"primaryKey" json:"vehicle_id" form:"vehicle_id"`
	VehicleName string    `json:"vehiclename,omitempty" form:"vehiclename"`
	Location    string    `json:"location,omitempty" form:"location"`
	Description string    `json:"description,omitempty" form:"description"`
	Price       json.Number       `json:"price,omitempty" form:"price"`
	Status      string    `json:"status,omitempty" form:"status"`
	Stock       int       `json:"stock,string,omitempty" form:"stock"`
	Category    string    `json:"category,omitempty" form:"category"`
	Image       string    `json:"image,omitempty" form:"image"`
	CreatedAt   time.Time `gorm:"default:current_timestamp" json:"created_at" form:"created_at"`
	UpdatedAt   time.Time `gorm:"default:current_timestamp" json:"updated_at" form:"updated_at"`
	Rating      json.Number   `json:"rating,omitempty" form:"rating"`
	TotalRented int       `json:"total_rented,string,omitempty" form:"total_rented"`
}

type Vehicles []Vehicle

package models

import (
	"time"
)

type History struct {
	HistoryId     uint      `gorm:"primaryKey" json:"history_id"`
	VehicleId     uint      `json:"vehicle_id"`
	Vehicle       Vehicle   `json:"vehicle"`
	UserId        string    `gorm:"type:uuid" json:"user_id"`
	User          User      `json:"user"`
	StartRental   string    `json:"start_rental"`
	EndRental     string    `json:"end_rental"`
	Prepayment    int       `json:"prepayment"`
	PaymentStatus string    `json:"payment_status"`
	ReturnStatus  string    `json:"return_status"`
	CreatedAt     time.Time `gorm:"default:current_timestamp" json:"created_at"`
	UpdatedAt     time.Time `gorm:"default:current_timestamp" json:"updated_at"`
}

type Histories []History

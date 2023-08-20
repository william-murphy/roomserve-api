package models

import (
	"time"
)

type NewFloor struct {
	Name       string `json:"name"`
	BuildingID uint   `json:"buildingId"`
}

type Floor struct {
	ID         uint      `json:"id" gorm:"primaryKey"`
	CreatedAt  time.Time `json:"-"`
	UpdatedAt  time.Time `json:"-"`
	Name       string    `json:"name" gorm:"size:64;not null;default:concat('Floor ',currval('floors_id_seq'))"`
	BuildingID uint      `json:"-"`
	Building   *Building `json:"building,omitempty"`
	Rooms      []Room    `json:"-"`
}

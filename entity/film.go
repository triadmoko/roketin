package entity

import (
	"github.com/jinzhu/gorm"
)

type Film struct {
	*gorm.Model
	Title       string
	Description string
	Duration    string
	Artist      string
	Genre       string
	Filename    string
}

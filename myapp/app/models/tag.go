package models

import (
	"time"
)

type Tag struct {
	Id int			`gorm:"primary_key" json:"id"`
	Title string		`sql:"size:255" json:"title"`
	Views uint64		`json:"views"`
	Created time.Time	`json:"created"`
}

package models

import (
	"time"
)

type Gif struct {
	Id int			`gorm:"primary_key" json:"id"`
	Title string		`sql:"size:255" json:"title"`
	Description string	`sql:"size:255" json:"description"`
	Views uint64		`json:"views"`
	Featured bool		`json:"featured"`
	Created time.Time	`gorm:"type:timestamp;default:current_timestamp"json:"created"`
	CategoryId int		`json:"category_id"`
}

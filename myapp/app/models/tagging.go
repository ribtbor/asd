package models

import (
	"time"
)

type Tagging struct {
	Id int64		`gorm:"primary_key" json:"id"`
	VideoId	int64		`json:"video_id"`
	TagId int64		`json:"tag_id"`
	Created time.Time	`json:"created"`
	UserId string		`sql:"size:255" json:"user_id"`		
}

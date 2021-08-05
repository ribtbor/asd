package models

import (
	"time"
)

type Pornstar struct {
	Id uint64		`gorm:"primary_key" json:"id"`
	First string		`sql:"size:255" json:"first"`
	Last string		`sql:"size:255" json:"last"`
	Bio string		`sql:"size:255" json:"bio"`
	Views uint64		`json:"views"`
	Featured bool		`json:"featured"`
	Created time.Time	`json:"created"`

}

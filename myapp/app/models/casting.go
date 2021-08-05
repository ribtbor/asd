package models

type Casting struct {
	Id int64		`gorm:"primary_key" json:"id"`
	VideoId	int64		`json:"video_id"`
	PornstarId int64	`json:"pornstar_id"`		
}

package controllers

import (
	"github.com/revel/revel"
	"myapp/app/models"
	"errors"
	"strconv"
)

type Pornstar struct {
	*revel.Controller
}

func (c Pornstar) Index() revel.Result {
	id := c.Params.Get("id")
	i, _ := strconv.Atoi(id)

	pornstar := models.Pornstar{}
	result := DB.Where("id = ?",i).Find(&pornstar)
	err := result.Error
	if err != nil {
		return c.RenderError(errors.New("Record Not Found"))
	}
	castings := []models.Casting{}
	result = DB.Where("pornstar_id = ?",i).Find(&castings)
	err = result.Error
	if err != nil {
		return c.RenderError(errors.New("Record Not Found"))
	}
	videos := []models.Video{}
	for i:=0; i < len(castings) ; i++ {
		video := models.Video{}
		result = DB.Where("id = ?", castings[i].VideoId).Find(&video)
		err = result.Error 
		if err != nil {
			return c.RenderError(errors.New("Record Not Found"))
		}
		videos = append(videos, video)
	}

	return c.Render(pornstar, videos)
}

func (c Pornstar) Pornstars() revel.Result {
	x := c.Params.Route.Get("page")
	page, _ := strconv.Atoi(x)
	offset := 0
	if page > 0 {
		offset = (page - 1) * 24
	}
	pagination:=models.Pagination{}
	var resul int
	DB.Table("pornstars").Count(&resul)

	pagination.Last = (resul/24) + 1
	pagination.Next = page + 1
	pagination.Previous = page - 1
	if page == 0 {
		pagination.Previous = 1
		pagination.Next = 2
	}else if page == 1 {
		pagination.Previous = 1
	}else if page == pagination.Last{
		pagination.Next = 1
	}

	iterator:=0
	for y:=page+1; y<=page+5; y++ {
		pagination.Middle[iterator]=y
		iterator++
	}
	
	pornstars :=[]models.Pornstar{}
	
	result := DB.Order("id asc").Limit(24).Offset(offset).Find(&pornstars);
	err := result.Error
	if err != nil {
		return c.RenderError(errors.New("Record Not Found"))
	}
	return c.Render(pornstars,pagination)

}

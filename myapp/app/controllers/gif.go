
package controllers

import (
	"github.com/revel/revel"
	"myapp/app/models"
	"errors"
	"strconv"
)

type Gif struct {
	*revel.Controller
}

func (c Gif) Index() revel.Result {
	id := c.Params.Route.Get("id")
	i, _ := strconv.Atoi(id)
	video := models.Gif{}
	result := DB.Where("id = ?", i).Find(&video)
	err := result.Error
	if err != nil {
		return c.RenderError(errors.New("Record Not Found"))
	}
	next := i + 1
	previous := i - 1
	return c.Render(video,next,previous)
}

func (c Gif) Gifs() revel.Result {
	x := c.Params.Route.Get("page")
	page, _ := strconv.Atoi(x)
	offset := 0
	if page > 0 {
		offset = (page - 1) * 12
	}
	pagination := models.Pagination{}
	var resul int
	DB.Table("gifss").Count(&resul)

	count := 0
	DB.Model(&Gif{}).Count(&count)

	pagination.Last = (count/12)
	pagination.Next = page + 1
	pagination.Previous = page - 1
	if page == 0 {
		pagination.Previous = 1
		pagination.Next = 2
	} else if page == 1 {
		pagination.Previous = 1
	} else if page == pagination.Last{
		pagination.Last = 1
	}

	iterator:=0
	for y:=page+1; y<=page+5; y++ {
		pagination.Middle[iterator]=y
		iterator++
	}

	videos :=[]models.Gif{}
	result := DB.Order("id desc").Limit(12).Offset(offset).Find(&videos);
	err := result.Error
	if err != nil {
		return c.RenderError(errors.New("Record Not Found"))
	}
	return c.Render(videos, pagination, count)
}

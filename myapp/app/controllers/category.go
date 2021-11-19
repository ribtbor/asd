package controllers

import (
	"github.com/revel/revel"
	"myapp/app/models"
	"errors"
	"strconv"
)

type Category struct {
	*revel.Controller
}

func (c Category) Index() revel.Result {
        x := c.Params.Route.Get("page")
	id := c.Params.Route.Get("id")
        i, _ := strconv.Atoi(id)

	page, _ := strconv.Atoi(x)

        category := models.Category{}
        result := DB.Where("id = ?", i).Find(&category)
        err := result.Error
        if err != nil {
                return c.RenderError(errors.New("Record Not Found"))
        }

	offset := 0
	if page > 0 {
		offset = (page - 1) * 24
	}
	pagination := models.Pagination{}
	var resul int
	DB.Table("gifs").Where("category_id=?",i).Count(&resul)
	
	pagination.Last = (resul/24) + 1
	pagination.Next = page + 1
	pagination.Previous = page - 1
	if page == 0 {
		pagination.Previous = 1
		pagination.Next = 2
	} else if page == 1 {
		pagination.Previous = 1
	} else if page == pagination.Last {
		pagination.Last = 1
	}

	iterator := 0
	for y:=page+1; y<=page+5; y++ {
		pagination.Middle[iterator]=y
		iterator++
	}

	videos := []models.Gif{}
	result = DB.Where("category_id = ?",i).Offset(offset).Find(&videos)
	err = result.Error
	if err != nil {
		return c.RenderError(errors.New("Record Not Found"))
	}
        return c.Render(category, videos, pagination)
}

func (c Category) Categories() revel.Result {
	categories :=[]models.Category{}
	
	result := DB.Order("id asc").Find(&categories);
	err := result.Error
	if err != nil {
		return c.RenderError(errors.New("Record Not Found"))
	}
	return c.Render(categories)

}

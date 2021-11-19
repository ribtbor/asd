package controllers

import (
	"github.com/revel/revel"
	"myapp/app/models"
	"errors"
	"strconv"
)

type Tag struct {
	*revel.Controller
}

func (c Tag) Index() revel.Result {
        id := c.Params.Route.Get("id")
        i, _ := strconv.Atoi(id)
        tag := models.Tag{}
        result := DB.Where("id = ?", i).Find(&tag)
        err := result.Error
        if err != nil {
                return c.RenderError(errors.New("Record Not Found"))
        }

	tagging := []models.Tagging{}
	result = DB.Where("tag_id = ?", i).Find(&tagging)
	err = result.Error
	if err != nil {
		return c.RenderError(errors.New("Record Not Found"))
	}	

	videos := []models.Gif{}
	for i:=0; i < len(tagging) ; i++ {
		video := models.Gif{}
		result = DB.Where("id = ?", tagging[i].VideoId).Find(&video)
		err = result.Error 
		if err != nil {
			return c.RenderError(errors.New("Record Not Found"))
		}
		videos = append(videos, video)
	}
	return c.Render(tag,videos)
}

func (c Tag) Tags() revel.Result {
	tags :=[]models.Tag{}
	
	result := DB.Order("id asc").Find(&tags);
	err := result.Error
	if err != nil {
		return c.RenderError(errors.New("Record Not Found"))
	}
	return c.Render(tags)
}

func (c Tag) Page() revel.Result {
	id := c.Params.Route.Get("id")
        i, _ := strconv.Atoi(id)
        
	page := c.Params.Route.Get("page")
	x, _ := strconv.Atoi(page)

	offset := (x - 1) * 24
	
	tag := models.Tag{}
        result := DB.Where("id = ?", i).Find(&tag)
        err := result.Error
        if err != nil {
                return c.RenderError(errors.New("Record Not Found"))
        }

        tagging := []models.Tagging{}
        result = DB.Where("tag_id = ?", i).Find(&tagging)
        err = result.Error
        if err != nil {
                return c.RenderError(errors.New("Record Not Found"))
        }

        videos := []models.Video{}
        for i:=0; i < len(tagging) ; i++ {
                video := models.Video{}
                result = DB.Where("id = ?", tagging[i].VideoId).Limit(24).Offset(offset).Find(&video)
                err = result.Error
                if err != nil {
                        return c.RenderError(errors.New("Record Not Found"))
                }
                videos = append(videos, video)
        }
	c.ViewArgs["tag"] = tag
	c.ViewArgs["videos"] = videos
        return c.RenderTemplate("Tag/Index.html")
}

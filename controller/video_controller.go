package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jirenmaa/gogin-api/entity"
	"github.com/jirenmaa/gogin-api/service"
	"github.com/jirenmaa/gogin-api/validators"
)

type VideoController interface {
	FindAll() []entity.Video
	Save(ctx *gin.Context) error
	Search(ctx *gin.Context) ([]entity.Video, error)
}

type controller struct {
	service service.VideoService
}

type searchTitle struct {
	Title string `json:"title"`
}

var validate *validator.Validate

func New(service service.VideoService) VideoController {
	validate = validator.New()
	validate.RegisterValidation("test", validators.ValidateVideoTitle)

	return &controller{service: service}
}

func (c *controller) FindAll() []entity.Video {
	return c.service.FindAll()
}

func (c *controller) Save(ctx *gin.Context) error {
	var video entity.Video

	err := ctx.ShouldBindJSON(&video)
	if err != nil {
		return err
	}

	err = validate.Struct(video)
	if err != nil {
		return err
	}

	c.service.Save(video)

	return nil
}

func (c *controller) Search(ctx *gin.Context) ([]entity.Video, error) {
	var title searchTitle

	err := ctx.ShouldBindJSON(&title)
	if err != nil {
		return []entity.Video{}, err
	}

	videos := c.service.Search(title.Title)
	return videos, nil
}

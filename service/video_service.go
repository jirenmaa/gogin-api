package service

import (
	"github.com/jirenmaa/gogin-api/entity"
	"strings"
)

type VideoService interface {
	Save(entity.Video) entity.Video
	FindAll() []entity.Video
	Search(title string) []entity.Video
}

type bucket struct {
	videos []entity.Video
}

func New() VideoService {
	return &bucket{}
}

func (service *bucket) Save(video entity.Video) entity.Video {
	service.videos = append(service.videos, video)

	return video
}

func (service *bucket) FindAll() []entity.Video {
	return service.videos
}

func (service *bucket) Search(title string) []entity.Video {
	var tempVideos []entity.Video
	// split the searched title into each word
	var words []string = strings.Split(title, " ")

	for _, word := range words {
		for _, video := range service.videos {
			if strings.Contains(strings.ToLower(video.Title), strings.ToLower(word)) {
				tempVideos = append(tempVideos, video)
			}
		}
	}

	return tempVideos
}

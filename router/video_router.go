package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jirenmaa/gogin-api/controller"
	"github.com/jirenmaa/gogin-api/service"
	"net/http"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func VideoRouter(server *gin.Engine) {
	videosRouter := server.Group("/api")
	{
		videosRouter.GET("/videos", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, videoController.FindAll())
		})

		videosRouter.POST("/videos", func(ctx *gin.Context) {
			err := videoController.Save(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "Video has been saved"})
			}
		})

		videosRouter.GET("/search_videos", func(ctx *gin.Context) {
			videos, err := videoController.Search(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, videos)
			}
		})
	}
}

package controller

import (
	"xseon-zero/handler"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	FileLinkHandler *handler.FileLinkHandler
}

func NewController(fileLinkHandler *handler.FileLinkHandler) *Controller {
	return &Controller{
		FileLinkHandler: fileLinkHandler,
	}
}

func (controller *Controller) SetupRoutes(r *gin.Engine) {
	r.POST("/file-links", controller.FileLinkHandler.CreateFileLink)
	r.GET("/file-links", controller.FileLinkHandler.GetAllFileLinks)
	r.GET("/file-links/:id", controller.FileLinkHandler.GetFileLinkByID)
	r.PUT("/file-links/:id", controller.FileLinkHandler.UpdateFileLink)
	r.DELETE("/file-links/:id", controller.FileLinkHandler.DeleteFileLink)
}

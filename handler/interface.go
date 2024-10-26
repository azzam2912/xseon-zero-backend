package handler

import (
    "github.com/gin-gonic/gin"
)

type Handler interface {
    CreateFileLink(c *gin.Context)
    GetFileLinkByID(c *gin.Context)
    GetAllFileLinks(c *gin.Context)
    UpdateFileLink(c *gin.Context)  // New method
    DeleteFileLink(c *gin.Context)  // New method
}
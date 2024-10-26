package handler

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"
	"xseon-zero/domain/response"
	"xseon-zero/usecase/filelink"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type FileLinkHandler struct {
	fileLinkUseCase filelink.FileLinkUseCase
}

func NewFileLinkHandler(fileLinkUseCase filelink.FileLinkUseCase) *FileLinkHandler {
	return &FileLinkHandler{
		fileLinkUseCase: fileLinkUseCase,
	}
}

func (h *FileLinkHandler) CreateFileLink(c *gin.Context) {
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Printf("Error reading request body: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var req struct {
		Link     string `json:"link"`
		Caption  string `json:"caption"`
		Category string `json:"category"`
	}

	if err := json.Unmarshal(body, &req); err != nil {
		log.Printf("Error unmarshaling request: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fileLink, err := h.fileLinkUseCase.CreateFileLink(req.Link, req.Caption, req.Category)
	if err != nil {
		log.Printf("Error creating file link: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, fileLink)
}

func (h *FileLinkHandler) GetFileLinkByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		log.Printf("Error parsing UUID: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	fileLink, err := h.fileLinkUseCase.GetFileLinkByID(id)
	if err != nil {
		log.Printf("Error getting file link: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, fileLink)
}

func (h *FileLinkHandler) GetAllFileLinks(c *gin.Context) {
	fileLinks, err := h.fileLinkUseCase.GetAllFileLinks()
	if err != nil {
		log.Printf("Error getting all file links: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, fileLinks)
}

func (h *FileLinkHandler) UpdateFileLink(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		log.Printf("Error parsing UUID: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Printf("Error reading request body: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var req struct {
		Link     string `json:"link"`
		Caption  string `json:"caption"`
		Category string `json:"category"`
	}

	if err := json.Unmarshal(body, &req); err != nil {
		log.Printf("Error unmarshaling request: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fileLink, err := h.fileLinkUseCase.UpdateFileLink(id, req.Link, req.Caption, req.Category)
	if err != nil {
		log.Printf("Error updating file link: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := response.Response{
		Msg:       "",
		Data:      fileLink,
		Timestamp: time.Now(),
	}
	c.JSON(http.StatusOK, response)
}

func (h *FileLinkHandler) DeleteFileLink(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		log.Printf("Error parsing UUID: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	err = h.fileLinkUseCase.DeleteFileLink(id)
	if err != nil {
		log.Printf("Error deleting file link: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "File link deleted successfully"})
}

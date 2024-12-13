package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"test/internal/dto"
)

type ClickService interface {
	Update(dto.UpdateRequest) (*dto.UpdateResponse, error)
	GetStats(dto.GetStatsRequest) (*dto.GetStatsResponse, error)
	Save(dto.SaveRequest) (*dto.SaveResponse, error)
}
type ClickController struct {
	service ClickService
}

func NewClickController(service ClickService) *ClickController {
	return &ClickController{service}
}

func (c *ClickController) Update(ctx *gin.Context) {
	request := dto.UpdateRequest{}
	request.ID = ctx.Param("id")
	_, err := c.service.Update(request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "clicked"})
}

func (c *ClickController) GetStats(ctx *gin.Context) {
	request := dto.GetStatsRequest{}
	request.BannerId = ctx.Param("id")

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	stats, err := c.service.GetStats(request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, stats)
}

func (c *ClickController) Save(ctx *gin.Context) {
	request := dto.SaveRequest{}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	if _, err := c.service.Save(request); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"ok": "banner"})
}

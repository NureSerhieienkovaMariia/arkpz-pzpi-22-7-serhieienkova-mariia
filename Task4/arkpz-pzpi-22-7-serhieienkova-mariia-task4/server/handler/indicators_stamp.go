package handler

import (
	"clinic/server/structures"
	"github.com/gin-gonic/gin"
	"strconv"
)

func (h *Handler) createIndicatorsStamp(c *gin.Context) {
	var input structures.IndicatorsStamp
	if err := c.BindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err := h.services.IndicatorsStampAction.Create(input)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Indicators stamp created"})
}

func (h *Handler) getAllIndicatorsStamps(c *gin.Context) {
	stamps, err := h.services.IndicatorsStampAction.GetAll()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, stamps)
}

func (h *Handler) getIndicatorsStampById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid id"})
		return
	}

	stamp, err := h.services.IndicatorsStampAction.GetById(id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, stamp)
}

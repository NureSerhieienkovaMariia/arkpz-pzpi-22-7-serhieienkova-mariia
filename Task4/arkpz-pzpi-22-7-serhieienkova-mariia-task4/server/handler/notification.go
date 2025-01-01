package handler

import (
	"clinic/server/structures"
	"github.com/gin-gonic/gin"
	"strconv"
)

func (h *Handler) createNotification(c *gin.Context) {
	var input structures.Notification
	if err := c.BindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	id, err := h.services.NotificationAction.Create(input)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"id": id})
}

func (h *Handler) getAllNotifications(c *gin.Context) {
	notifications, err := h.services.NotificationAction.GetAll()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, notifications)
}

func (h *Handler) getNotificationById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid id"})
		return
	}
	notification, err := h.services.NotificationAction.Get(id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, notification)
}

func (h *Handler) getAllNotificationsByPatientID(c *gin.Context) {
	patientID, err := strconv.Atoi(c.Param("patient_id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid patient id"})
		return
	}

	notifications, err := h.services.NotificationAction.GetAllByPatientID(patientID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, notifications)
}

package handler

import (
	"clinic/server/structures"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) createPatient(c *gin.Context) {
	var input structures.CreatePatientInput
	if err := c.BindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	id, err := h.services.PatientAction.Create(input)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// Add entries to userPatientsTable for relatives and doctors
	for _, relativeId := range input.Relatives {
		h.services.UserPatientAction.Create(structures.UserPatient{UserId: relativeId, PatientId: id})
	}
	for _, doctorId := range input.Doctors {
		h.services.UserPatientAction.Create(structures.UserPatient{UserId: doctorId, PatientId: id})
	}

	c.JSON(200, gin.H{"id": id})
}

func (h *Handler) getAllPatients(c *gin.Context) {
	patients, err := h.services.PatientAction.GetAll()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, patients)
}

func (h *Handler) getPatientById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid id"})
		return
	}
	patient, err := h.services.PatientAction.GetById(id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, patient)
}

func (h *Handler) updatePatient(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid id"})
		return
	}

	var input structures.UpdatePatientInput
	if err := c.BindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := h.services.PatientAction.Update(id, input); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"status": "updated"})
}
func (h *Handler) deletePatient(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid id"})
		return
	}
	if err := h.services.PatientAction.Delete(id); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"status": "deleted"})
}

func (h *Handler) getPatientFullInfo(c *gin.Context) {
	patientID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid patient id")
		return
	}

	fullInfo, err := h.services.PatientAction.GetFullInfo(patientID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, fullInfo)
}

func (h *Handler) setMedicineToPatient(c *gin.Context) {
	patientID, err := strconv.Atoi(c.Param("patient_id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid patient id"})
		return
	}

	medicineID, err := strconv.Atoi(c.Param("medicine_id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid medicine id"})
		return
	}

	var input struct {
		Schedule string `json:"schedule" binding:"required"`
	}
	if err := c.BindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = h.services.PatientMedicineAction.SetMedicineToPatient(patientID, medicineID, input.Schedule)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"status": "medicine set to patient"})
}

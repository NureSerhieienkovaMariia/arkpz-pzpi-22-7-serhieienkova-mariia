package handler

import (
	"clinic/server/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.Use(CORSMiddleware())

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
		auth.POST("/refresh-token", h.refreshToken)
		auth.GET("/current-user", h.currentUser)
	}

	api := router.Group("/api", h.userIdentity)
	{
		users := api.Group("/user")
		{
			users.GET("/email/:email", h.getUserByEmail)
		}

		patients := api.Group("/patient")
		{
			patients.POST("/", h.createPatient)
			patients.GET("/", h.getAllPatients)
			patients.GET("/:id", h.getPatientById)
			patients.PUT("/:id", h.updatePatient)
			patients.DELETE("/:id", h.deletePatient)
		}

		medicines := api.Group("/medicine")
		{
			medicines.POST("/", h.createMedicine)
			medicines.GET("/", h.getAllMedicines)
			medicines.GET("/:id", h.getMedicineById)
			medicines.PUT("/:id", h.updateMedicine)
			medicines.DELETE("/:id", h.deleteMedicine)
		}

		diagnoses := api.Group("/diagnosis")
		{
			diagnoses.POST("/", h.createDiagnosis)
			diagnoses.GET("/", h.getAllDiagnoses)
			diagnoses.GET("/:id", h.getDiagnosisById)
			diagnoses.PUT("/:id", h.updateDiagnosis)
			diagnoses.DELETE("/:id", h.deleteDiagnosis)
		}

		patientMedicines := api.Group("/patients-medicine")
		{
			patientMedicines.POST("/", h.createPatientMedicine)
			patientMedicines.GET("/", h.getAllPatientMedicines)
			patientMedicines.GET("/:id", h.getPatientMedicineById)
			patientMedicines.PUT("/:id", h.updatePatientMedicine)
			patientMedicines.DELETE("/:id", h.deletePatientMedicine)
		}

		userPatients := api.Group("/user-patient")
		{
			userPatients.POST("/", h.createUserPatient)
			userPatients.GET("/", h.getAllUserPatients)
			userPatients.GET("/:id", h.getUserPatientById)
			userPatients.PUT("/:id", h.updateUserPatient)
			userPatients.DELETE("/:id", h.deleteUserPatient)
		}

		devices := api.Group("/device")
		{
			devices.POST("/", h.createDevice)
			devices.GET("/", h.getAllDevices)
			devices.GET("/:id", h.getDeviceById)
			devices.PUT("/:id", h.updateDevice)
			devices.DELETE("/:id", h.deleteDevice)
		}

		notifications := api.Group("/notification")
		{
			notifications.POST("/", h.createNotification)
			notifications.GET("/", h.getAllNotifications)
			notifications.GET("/:id", h.getNotificationById)
			notifications.PUT("/:id", h.updateNotification)
			notifications.DELETE("/:id", h.deleteNotification)
		}
	}

	return router
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "*")
		c.Header("Access-Control-Allow-Methods", "GET, HEAD, POST, PUT, DELETE, OPTIONS, PATCH")
		c.Header("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

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
			users.POST("/", h.checkAccessLevel(adminAccessLevel), h.createUser)
			users.GET("/", h.checkAccessLevel(adminAccessLevel), h.getAllUsers)
			users.GET("/:id", h.getUserById)
			users.PUT("/:id", h.checkAdminOrSelf(), h.updateUser)
			users.DELETE("/:id", h.checkAccessLevel(adminAccessLevel), h.deleteUser)
		}

		patients := api.Group("/patient")
		{
			patients.POST("/", h.checkAccessLevel(doctorAccessLevel), h.createPatient)
			patients.GET("/", h.checkAccessLevel(doctorAccessLevel), h.getAllPatients)
			patients.GET("/:id", h.getPatientById)
			patients.GET("/:id/full-info", h.getPatientFullInfo)
			patients.PUT("/:id", h.checkAccessLevel(doctorAccessLevel), h.updatePatient)
			patients.DELETE("/:id", h.checkAccessLevel(doctorAccessLevel), h.deletePatient)
			patients.POST("/:patient_id/medicine/:medicine_id", h.checkAccessLevel(doctorAccessLevel), h.setMedicineToPatient)
		}

		medicines := api.Group("/medicine")
		{
			medicines.POST("/", h.checkAccessLevel(doctorAccessLevel), h.createMedicine)
			medicines.GET("/", h.checkAccessLevel(doctorAccessLevel), h.getAllMedicines)
			medicines.GET("/:id", h.getMedicineById)
			medicines.PUT("/:id", h.checkAccessLevel(doctorAccessLevel), h.updateMedicine)
			medicines.DELETE("/:id", h.checkAccessLevel(doctorAccessLevel), h.deleteMedicine)
		}

		diagnoses := api.Group("/diagnosis")
		{
			diagnoses.POST("/", h.checkAccessLevel(doctorAccessLevel), h.createDiagnosis)
			diagnoses.GET("/", h.checkAccessLevel(doctorAccessLevel), h.getAllDiagnoses)
			diagnoses.GET("/:id", h.getDiagnosisById)
			diagnoses.PUT("/:id", h.checkAccessLevel(doctorAccessLevel), h.updateDiagnosis)
			diagnoses.DELETE("/:id", h.checkAccessLevel(doctorAccessLevel), h.deleteDiagnosis)
		}

		devices := api.Group("/device")
		{
			devices.POST("/", h.checkAccessLevel(adminAccessLevel), h.createDevice)
			devices.GET("/", h.checkAccessLevel(adminAccessLevel), h.getAllDevices)
			devices.GET("/:id", h.checkAccessLevel(adminAccessLevel), h.getDeviceById)
			devices.PUT("/:id", h.checkAccessLevel(adminAccessLevel), h.updateDevice)
			devices.DELETE("/:id", h.checkAccessLevel(adminAccessLevel), h.deleteDevice)
		}

		notifications := api.Group("/notification")
		{
			notifications.GET("/", h.getAllNotifications)
			notifications.GET("/:id", h.getNotificationById)
			notifications.GET("/patient/:patient_id", h.checkDoctorOrRelative(), h.getAllNotificationsByPatientID)
		}

		indicators := api.Group("/indicators")
		{
			indicators.POST("/", h.createIndicatorsStamp)
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

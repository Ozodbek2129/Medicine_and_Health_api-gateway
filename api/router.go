package api

import (
	"api-gateway/api/handler"
	"api-gateway/api/middleware"

	_ "api-gateway/api/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title User Service API
// @version 1.0
// @description This is a sample server for a user service.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:50053
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func NewRouter(h *handler.Handler) *gin.Engine {
	router := gin.Default()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Use(middleware.Check)

	router.Use(middleware.CheckPermissionMiddleware(h.Enforcer))

	health := router.Group("/health")
	{
		health.POST("/medical_recordsAdd", h.AddMedicalRecord)
		health.GET("/medical_recordsGet/:id", h.GetMedicalRecord)
		health.PUT("/medical_recordsUp", h.UpdateMedicalRecord)
		health.DELETE("/medical_recordsDel/:id", h.DeleteMedicalRecord)
		health.GET("/medical_records/user/:userId", h.ListMedicalRecords)

		health.POST("/lifestyleAdd", h.AddLifestyleData)
		health.GET("/getalllifestyledata/:limit/:page", h.GetAllLifestyleData)
		health.GET("/lifestyleGet/:id", h.GetLifestyleData)
		health.PUT("/lifestyleUp", h.UpdateLifestyleData)
		health.DELETE("/lifestyleDel/:id", h.DeleteLifestyleData)

		health.POST("/wearable-dataAdd", h.AddWearableData)
		health.GET("/wearabledata/:limit/:page", h.GetAllWearableData)
		health.GET("/wearable-dataGet/:id", h.GetWearableData)
		health.PUT("/wearable-dataUp", h.UpdateWearableData)
		health.DELETE("/wearable-dataDel/:id", h.DeleteWearableData)

		health.POST("/recommendationsAdd", h.GenerateHealthRecommendations)
		health.GET("/recommendations/:id",h.GenerateHealthRecommendationsId)
		health.GET("/monitoring/:user_id/realtime", h.GetRealtimeHealthMonitoring)
		health.GET("/summary/:user_id/daily", h.GetDailyHealthSummary)
		health.GET("/summary/:user_id/weekly   ", h.GetWeeklyHealthSummary)
	}

	user:=router.Group("/user")
	{
		user.POST("/notifications",h.NotificationsAdd)
		user.GET("/notifications/:user_id",h.NotificationsGet)
		user.PUT("/notificationsPut")
	}
	return router
}

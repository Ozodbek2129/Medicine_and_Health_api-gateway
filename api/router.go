package api

import (
	"api-gateway/api/handler"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "api-gateway/api/docs"
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
func NewRouter(h *handler.Handler) *gin.Engine{
	router:=gin.Default()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	health:=router.Group("/health")
	{
		health.POST("/medical_records",h.AddMedicalRecord)
		health.GET("/medical_records/:id",h.GetMedicalRecord)
		health.PUT("/medical_records",h.UpdateMedicalRecord)
		health.DELETE("/medical_records/:id",h.DeleteMedicalRecord)
		health.GET("/medical_records/user/:userId",h.ListMedicalRecords)

		health.POST("/lifestyle",h.AddLifestyleData)
		health.GET("/lifestyle/:id",h.GetLifestyleData)
		health.PUT("/lifestyle",h.UpdateLifestyleData)
		health.DELETE("/lifestyle/:id",h.DeleteLifestyleData)
	}
	return router
}
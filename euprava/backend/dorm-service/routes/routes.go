package routes

import (
	"dorm-service/controllers"

	"github.com/gin-gonic/gin"
)

func MainRoutes(routes *gin.Engine, dc controllers.DormController) {
	routes.GET("/applications", dc.GetAllApplications())
	routes.POST("/applications/create/:_id", dc.InsertApplication())

}

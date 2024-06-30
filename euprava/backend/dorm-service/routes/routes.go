package routes

import (
	"dorm-service/controllers"
	"dorm-service/middleware"

	"github.com/gin-gonic/gin"
)

func MainRoutes(routes *gin.Engine, dc controllers.DormController) {
	routes.Use(middleware.Authentication())
	routes.GET("/applications", middleware.AuthorizeRoles([]string{"ADMIN"}), dc.GetAllApplications())
	routes.GET("/application", middleware.AuthorizeRoles([]string{"ADMIN,STUDENT"}), dc.GetApplication())
	routes.POST("/applications/create", middleware.AuthorizeRoles([]string{"ADMIN,STUDENT"}), dc.InsertApplication())

}

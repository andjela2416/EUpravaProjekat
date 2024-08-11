package routes

import (
	"dorm-service/controllers"
	"dorm-service/middleware"

	"github.com/gin-gonic/gin"
)

func MainRoutes(routes *gin.Engine, dc controllers.DormController) {
	routes.Use(middleware.Authentication())

	routes.GET("/applications", middleware.AuthorizeRoles([]string{"ADMIN"}), dc.GetAllApplications())
	routes.GET("/application", middleware.AuthorizeRoles([]string{"ADMIN", "STUDENT"}), dc.GetApplication())
	routes.POST("/applications/create", middleware.AuthorizeRoles([]string{"ADMIN", "STUDENT"}), dc.InsertApplication())

	routes.GET("/building/:id", middleware.AuthorizeRoles([]string{"ADMIN", "STUDENT"}), dc.GetBuilding())
	routes.POST("/building", middleware.AuthorizeRoles([]string{"ADMIN", "STUDENT"}), dc.InsertBuilding())
	routes.DELETE("/building/:id", middleware.AuthorizeRoles([]string{"ADMIN"}), dc.DeleteBuilding())
	//	routes.PUT("/buildings/:id", middleware.AuthorizeRoles([]string{"ADMIN"}), dc.UpdateBuilding())

	routes.GET("building/:id/room/:number", middleware.AuthorizeRoles([]string{"ADMIN", "STUDENT"}), dc.GetRoom())
	routes.POST("building/:id/room", middleware.AuthorizeRoles([]string{"ADMIN", "STUDENT"}), dc.InsertRoom())
}

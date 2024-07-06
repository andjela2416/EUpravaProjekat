package routes

import (
	"university-service/controllers"
	database "university-service/database"

	"github.com/gin-gonic/gin"
)

func ProfileRoutes(router *gin.Engine) {
	router.POST("/students/create", controllers.CreateStudentHandler(database.Client))
	router.GET("/students/:id", controllers.GetStudentByIDHandler(database.Client))
	router.POST("/notificationsByHealthcare", controllers.CreateNotificationByHealthcareHandler(database.Client))
	router.POST("/notifications", controllers.CreateNotificationHandler(database.Client))
	router.GET("/notifications/:id", controllers.GetNotificationByIDHandler(database.Client))
	router.GET("/notifications", controllers.GetAllNotificationsHandler(database.Client))
	router.DELETE("/notifications/:id", controllers.DeleteNotificationHandler(database.Client))
}

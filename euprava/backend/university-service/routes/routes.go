package routes

import (
	"university-service/controllers"
	database "university-service/database"

	"github.com/gin-gonic/gin"
)

func ProfileRoutes(router *gin.Engine) {
	router.POST("/students/create", controllers.CreateStudentHandler(database.Client))
	router.GET("/students/:id", controllers.GetStudentByIDHandler(database.Client))
}

package routes

import (
	"university-service/controllers"

	"github.com/gin-gonic/gin"
)

func MainRoutes(routes *gin.Engine, uc controllers.UniversityController) {
	routes.GET("/student/studentID/:student_id", uc.GetStudentByStudentID())
	routes.POST("/student", uc.InsertStudent())
}

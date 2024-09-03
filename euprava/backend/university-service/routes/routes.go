package routes

import (
	"university-service/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, ctrl *controllers.Controllers) {
	router.POST("/students/create", ctrl.CreateStudent)
	router.GET("/students/:id", ctrl.GetStudentByID)
	router.PUT("/students/:id", ctrl.UpdateStudent)
	router.DELETE("/students/:id", ctrl.DeleteStudent)

	router.POST("/professors/create", ctrl.CreateProfessor)
	router.GET("/professors/:id", ctrl.GetProfessorByID)
	router.PUT("/professors/:id", ctrl.UpdateProfessor)
	router.DELETE("/professors/:id", ctrl.DeleteProfessor)

	router.POST("/courses/create", ctrl.CreateCourse)
	router.GET("/courses/:id", ctrl.GetCourseByID)
	router.PUT("/courses/:id", ctrl.UpdateCourse)
	router.DELETE("/courses/:id", ctrl.DeleteCourse)

	router.POST("/departments/create", ctrl.CreateDepartment)
	router.GET("/departments/:id", ctrl.GetDepartmentByID)
	router.PUT("/departments/:id", ctrl.UpdateDepartment)
	router.DELETE("/departments/:id", ctrl.DeleteDepartment)

	router.POST("/universities/create", ctrl.CreateUniversity)
	router.GET("/universities/:id", ctrl.GetUniversityByID)
	router.PUT("/universities/:id", ctrl.UpdateUniversity)
	router.DELETE("/universities/:id", ctrl.DeleteUniversity)

	router.POST("/exams/create", ctrl.CreateExam)
	router.GET("/exams/:id", ctrl.GetExamByID)
	router.PUT("/exams/:id", ctrl.UpdateExam)
	router.DELETE("/exams/:id", ctrl.DeleteExam)

	router.POST("/administrators/create", ctrl.CreateAdministrator)
	router.GET("/administrators/:id", ctrl.GetAdministratorByID)
	router.PUT("/administrators/:id", ctrl.UpdateAdministrator)
	router.DELETE("/administrators/:id", ctrl.DeleteAdministrator)

	router.POST("/assistants/create", ctrl.CreateAssistant)
	router.GET("/assistants/:id", ctrl.GetAssistantByID)
	router.PUT("/assistants/:id", ctrl.UpdateAssistant)
	router.DELETE("/assistants/:id", ctrl.DeleteAssistant)
}

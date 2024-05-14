package controllers

import (
	"log"
	"net/http"
	"university-service/data"

	"github.com/gin-gonic/gin"
)

type UniversityController struct {
	logger *log.Logger
	repo   *data.UniRepo
}

func (uc UniversityController) GetStudentByStudentID() gin.HandlerFunc {
	return func(c *gin.Context) {

		studentId := c.Param("student_id")
		var student *data.Student

		student, err := uc.repo.GetStudentByStudentID(studentId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			uc.logger.Println(err.Error())
			return
		}

		c.JSON(http.StatusOK, gin.H{"Success, student found": student})
		uc.logger.Printf("Student sent")

	}
}
func (dc *UniversityController) InsertStudent() gin.HandlerFunc {
	return func(c *gin.Context) {
		var student data.Student
		if err := c.BindJSON(&student); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		err := dc.repo.InsertStudent(&student)
		if err != nil {
			dc.logger.Print("Database exception: ", err)
			c.JSON(http.StatusInternalServerError, gin.H{"database exception": err.Error()})
		}
		c.JSON(http.StatusOK, gin.H{"success": "student inserted"})
	}
}

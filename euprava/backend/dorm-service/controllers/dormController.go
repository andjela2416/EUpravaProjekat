package controllers

import (
	"dorm-service/data"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DormController struct {
	logger *log.Logger
	repo   *data.DormRepo
}

func (dc DormController) GetStudentByID(studentId string) (*data.Student, error) {

	uniUrl := fmt.Sprintf("http://university-service:8081/get/student", studentId)
	uniResponse, err := http.Get(uniUrl)
	if err != nil {
		dc.logger.Printf("Error making GET request for student: %v", err)
		return nil, fmt.Errorf("error making GET request for student: %v", err)
	}
	defer uniResponse.Body.Close()

	if uniResponse.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(uniResponse.Body)
		dc.logger.Println("error: ", string(body))
		return nil, fmt.Errorf("uni service returned error: %s", string(body))
	}
	var returnedStudent *data.Student
	if err := json.NewDecoder(uniResponse.Body).Decode(&returnedStudent); err != nil {
		dc.logger.Printf("error parsing auth response body: %v\n", err)
		return nil, fmt.Errorf("error parsing uni response body")
	}
	return returnedStudent, nil

}
func (dc *DormController) InsertApplication() gin.HandlerFunc {
	return func(c *gin.Context) {
		studentId := c.Param("_id")
		var application *data.Application
		application.Status = "Pending"

		student, err := dc.GetStudentByID(studentId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		application.Student = student
		if err = dc.repo.Insertapplications(application); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"database exception": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"Application created": application})
		return
	}
}

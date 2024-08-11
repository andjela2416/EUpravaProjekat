package controllers

import (
	"dorm-service/data"
	"dorm-service/models"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type DormController struct {
	logger *log.Logger
	repo   *data.DormRepo
}

var validate = validator.New()

func NewDormController(l *log.Logger, r *data.DormRepo) *DormController {
	return &DormController{l, r}
}
func (dc DormController) GetStudentByID(studentId string) (*models.User, error) {

	uniUrl := fmt.Sprintf("http://auth-service:8080/users/%v", studentId)
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
	var returnedStudent *models.User
	if err := json.NewDecoder(uniResponse.Body).Decode(&returnedStudent); err != nil {
		dc.logger.Printf("error parsing auth response body: %v\n", err)
		return nil, fmt.Errorf("error parsing uni response body")
	}
	return returnedStudent, nil

}
func (dc *DormController) InsertApplication() gin.HandlerFunc {
	return func(c *gin.Context) {

		id, exists := c.Get("uid")
		studentId := id.(string)
		if !exists {
			c.JSON(http.StatusBadRequest, gin.H{"error: student id not found ": studentId})
			return
		}
		var application models.Application

		student, err := dc.GetStudentByID(studentId)
		if student == nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Student not found"})
			return
		}
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		application.User = student
		if err = dc.repo.Insertapplications(&application); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"database exception": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"Application created": application})
	}
}

func (dc *DormController) GetAllApplications() gin.HandlerFunc {
	return func(c *gin.Context) {

		apps, err := dc.repo.GetAllapplications()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"database exception": err.Error()})
			return
		}
		c.JSON(http.StatusOK, apps)
		return
	}
}
func (dc *DormController) GetApplication() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, exists := c.Get("uid")
		studentId := id.(string)
		if !exists {
			c.JSON(http.StatusNotFound, gin.H{"error: student id not found in token": studentId})
			return
		}

		app, err := dc.repo.GetApplication(studentId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"database exception": err.Error()})
			return
		}
		c.JSON(http.StatusOK, app)
		return

	}
}

//func (dc DormController) ProcessApplications() error {
// select all pending applications, rank them,
//accept the first n number of them based on how many spaces there are
//assign students to random non-full rooms
//}

func (dc *DormController) InsertBuilding() gin.HandlerFunc {
	return func(c *gin.Context) {

		var building models.Building
		if err := c.BindJSON(&building); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "parsing failed"})
			return
		}
		validationErr := validate.Struct(building)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}
		err, buildingId := dc.repo.InsertBuilding(building)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		building.Id = buildingId
		c.JSON(http.StatusOK, gin.H{"Building created": building})

	}
}

func (dc *DormController) GetBuilding() gin.HandlerFunc {
	return func(c *gin.Context) {
		buildingId := c.Param("id")
		building, err := dc.repo.GetBuilding(buildingId)
		if err != nil {
			c.JSON(http.StatusNotFound, err.Error())
			return
		}

		c.JSON(http.StatusOK, building)
		return

	}
}

func (dc *DormController) GetBuildingLocal(buildingId string) (models.Building, error) {

	building, err := dc.repo.GetBuilding(buildingId)
	if err != nil {
		return *building, err
	}
	return *building, nil

}
func (dc *DormController) DeleteBuilding() gin.HandlerFunc {
	return func(c *gin.Context) {

		buildingId := c.Param("id")
		building, err := dc.repo.GetBuilding(buildingId)
		if err != nil {
			c.JSON(http.StatusNotFound, err.Error())
		}

		c.JSON(http.StatusOK, building)
		return

	}
}

// number is auto assigned
// insert capacity
// insert buildingid
// Todo: insert room into building based on id
// building id is not being inserted
func (dc *DormController) InsertRoom() gin.HandlerFunc {
	return func(c *gin.Context) {
		type RoomInfo struct {
			Capacity int `json:"capacity"`
		}
		buildingIdParam := c.Param("id")
		var room RoomInfo

		if err := c.BindJSON(&room); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "parsing failed", "details": err.Error()})
			return
		}

		err := dc.repo.InsertRoom(room.Capacity, buildingIdParam)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"Room created": room})
	}
}

func (dc *DormController) GetRoom() gin.HandlerFunc {
	return func(c *gin.Context) {
		roomNumberParam := c.Param("number")
		buildingIdParam := c.Param("id")
		roomNumber, err := strconv.Atoi(roomNumberParam)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		room, err := dc.repo.GetRoom(roomNumber, buildingIdParam)
		if err != nil {
			c.JSON(http.StatusNotFound, err.Error())
		}

		c.JSON(http.StatusOK, room)
		return

	}
}

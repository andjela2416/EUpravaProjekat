package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"

	repositories "university-service/repository"

	"go.mongodb.org/mongo-driver/mongo"
)

func CreateStudentHandler(client *mongo.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		var newStudent repositories.Student
		if err := c.ShouldBindJSON(&newStudent); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := repositories.Createstudent(client, &newStudent); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.Status(http.StatusCreated)
	}
}

func GetStudentByIDHandler(client *mongo.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid student ID"})
			return
		}

		student, err := repositories.GetstudentByID(client, userID.Hex())
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "student not found"})
			return
		}

		c.JSON(http.StatusOK, student)
	}
}

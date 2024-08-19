package controllers

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"time"

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

func CreateNotificationHandler(client *mongo.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		var newNotification repositories.Notification
		if err := c.ShouldBindJSON(&newNotification); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := repositories.CreateNotification(client, &newNotification); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.Status(http.StatusCreated)
	}
}

func CreateNotificationByHealthcareHandler(client *mongo.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		var appointmentData map[string]interface{}

		if err := c.ShouldBindJSON(&appointmentData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		description := "Systematic check appointment details: "

		if date, ok := appointmentData["date"].(string); ok {
			description += "Date: " + date + ", "
		}
		if facultyName, ok := appointmentData["faculty_name"].(string); ok {
			description += "Faculty: " + facultyName + ", "
		}
		if fieldOfStudy, ok := appointmentData["field_of_study"].(string); ok {
			description += "Field of Study: " + fieldOfStudy + ", "
		}
		if descriptionText, ok := appointmentData["description"].(string); ok {
			description += "Description: " + descriptionText
		}

		notification := repositories.Notification{
			Title:     "New Appointment for Systematic Check Notification",
			Content:   description,
			CreatedAt: time.Now(),
		}

		if err := repositories.CreateNotification(client, &notification); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.Status(http.StatusOK)
	}
}

func GetNotificationByIDHandler(client *mongo.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		notificationID, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid notification ID"})
			return
		}

		notification, err := repositories.GetNotificationByID(client, notificationID.Hex())
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Notification not found"})
			return
		}

		c.JSON(http.StatusOK, notification)
	}
}

func GetAllNotificationsHandler(client *mongo.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		notifications, err := repositories.GetAllNotifications(client)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, notifications)
	}
}
func DeleteNotificationHandler(client *mongo.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		notificationID := c.Param("id")

		err := repositories.DeleteNotification(client, notificationID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.Status(http.StatusOK)
	}
}

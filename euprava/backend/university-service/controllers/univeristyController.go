package controllers

import (
	"net/http"
	repositories "university-service/repository"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Controllers struct {
	Repo *repositories.Repository
}

func NewControllers(repo *repositories.Repository) *Controllers {
	return &Controllers{Repo: repo}
}

func (ctrl *Controllers) CreateStudent(c *gin.Context) {
	var student repositories.Student
	if err := c.BindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := ctrl.Repo.CreateStudent(&student)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, student)
}

func (ctrl *Controllers) GetStudentByID(c *gin.Context) {
	id := c.Param("id")

	student, err := ctrl.Repo.GetStudentByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
		return
	}

	c.JSON(http.StatusOK, student)
}

func (ctrl *Controllers) UpdateStudent(c *gin.Context) {
	id := c.Param("id")
	var student repositories.Student
	if err := c.BindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	student.ID = objectID

	err = ctrl.Repo.UpdateStudent(&student)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, student)
}

func (ctrl *Controllers) DeleteStudent(c *gin.Context) {
	id := c.Param("id")

	err := ctrl.Repo.DeleteStudent(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

func (ctrl *Controllers) CreateProfessor(c *gin.Context) {
	var professor repositories.Professor
	if err := c.BindJSON(&professor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := ctrl.Repo.CreateProfessor(&professor)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, professor)
}

func (ctrl *Controllers) GetProfessorByID(c *gin.Context) {
	id := c.Param("id")

	professor, err := ctrl.Repo.GetProfessorByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Professor not found"})
		return
	}

	c.JSON(http.StatusOK, professor)
}

func (ctrl *Controllers) UpdateProfessor(c *gin.Context) {
	id := c.Param("id")
	var professor repositories.Professor
	if err := c.BindJSON(&professor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	professor.ID = objectID

	err = ctrl.Repo.UpdateProfessor(&professor)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, professor)
}

func (ctrl *Controllers) DeleteProfessor(c *gin.Context) {
	id := c.Param("id")

	err := ctrl.Repo.DeleteProfessor(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

func (ctrl *Controllers) CreateDepartment(c *gin.Context) {
	var department repositories.Department
	if err := c.BindJSON(&department); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := ctrl.Repo.CreateDepartment(&department)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, department)
}

func (ctrl *Controllers) GetDepartmentByID(c *gin.Context) {
	id := c.Param("id")

	department, err := ctrl.Repo.GetDepartmentByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Department not found"})
		return
	}

	c.JSON(http.StatusOK, department)
}

func (ctrl *Controllers) UpdateDepartment(c *gin.Context) {
	id := c.Param("id")
	var department repositories.Department
	if err := c.BindJSON(&department); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	department.ID = objectID

	err = ctrl.Repo.UpdateDepartment(&department)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, department)
}

func (ctrl *Controllers) DeleteDepartment(c *gin.Context) {
	id := c.Param("id")

	err := ctrl.Repo.DeleteDepartment(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

func (ctrl *Controllers) CreateUniversity(c *gin.Context) {
	var university repositories.University
	if err := c.BindJSON(&university); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := ctrl.Repo.CreateUniversity(&university)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, university)
}

func (ctrl *Controllers) GetUniversityByID(c *gin.Context) {
	id := c.Param("id")

	university, err := ctrl.Repo.GetUniversityByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "University not found"})
		return
	}

	c.JSON(http.StatusOK, university)
}

func (ctrl *Controllers) UpdateUniversity(c *gin.Context) {
	id := c.Param("id")
	var university repositories.University
	if err := c.BindJSON(&university); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	university.ID = objectID

	err = ctrl.Repo.UpdateUniversity(&university)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, university)
}

func (ctrl *Controllers) DeleteUniversity(c *gin.Context) {
	id := c.Param("id")

	err := ctrl.Repo.DeleteUniversity(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

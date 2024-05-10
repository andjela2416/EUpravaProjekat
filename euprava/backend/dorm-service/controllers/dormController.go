package controllers

import (
	"backend/dorm-service/data"
	"log"
)

type DormController struct {
	logger *log.Logger
	repo   *data.DormRepo
}

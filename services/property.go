package services

import (
	"../domain"
	"../utils"
)

func GetProperty(propertyId int64) (*domain.Property, *utils.ApplicationError) {
	return domain.GetProperty(propertyId)
}

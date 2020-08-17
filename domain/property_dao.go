package domain

import (
	"../utils"
	"fmt"
	"net/http"
)

var property []*Property

func GetProperty(propertyId int64) (*Property, *utils.ApplicationError) {
	if property := property[propertyId]; property != nil {
		return property, nil
	}

	return nil, &utils.ApplicationError{
		Message:    fmt.Sprintf("property #{propertyId} was not found"),
		StatusCode: http.StatusNotFound,
		Code:       "not_found",
	}
}

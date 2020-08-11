package domain

import (
	"../utils"
	"fmt"
	"net/http"
)

var (
	property = map[int64]*Property{
		1: {},
	}
)

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

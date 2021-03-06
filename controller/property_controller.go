package controller

import (
	"../services"
	"../utils"
	"encoding/json"
	"net/http"
	"strconv"
)

func GetProperty(resp http.ResponseWriter, req *http.Request) {
	propertyId, err := strconv.ParseInt(req.URL.Query().Get("id"), 10, 64)
	if err != nil {
		apiErr := &utils.ApplicationError{
			Message:    "property_id must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}
		jsonValue, _ := json.Marshal(apiErr)
		resp.WriteHeader(apiErr.StatusCode)
		resp.Write(jsonValue)
		return
	}
	property, apiErr := services.GetProperty(propertyId)
	if apiErr != nil {
		jsonValue, _ := json.Marshal(apiErr)
		resp.WriteHeader(apiErr.StatusCode)
		resp.Write([]byte(jsonValue))
		return
	}
	jsonValue, _ := json.Marshal(property)
	resp.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if err := json.NewEncoder(resp).Encode(propertyId); err != nil {
		panic(err)
	}
	resp.Write(jsonValue)
}

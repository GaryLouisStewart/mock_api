package routing

import (
	"../controller"
	"../utils"
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Property_id",
		"GET",
		"/api/properties/{id}",
		controller.GetProperty,
	},
	Route{
		"Healthcheck",
		"GET",
		"/api/healthcheck",
		utils.HealthCheckEndpoint,
	},
}

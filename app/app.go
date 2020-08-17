package app

import (
	"../routing"
	"../utils"
	"net/http"
	"os"
)

func StartApp() {
	// initialise database
	d := utils.Database{}
	d.Initialise(os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"))
	// initialise the router from our router struct in the routing package.

	router := routing.InitNewRouter()
	if err := http.ListenAndServe(":9000", router); err != nil {
		panic(err)
	}
}

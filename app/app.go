package app

import (
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
	// initialise the router from our app struct in utils package
	if err := http.ListenAndServe(":9000", nil); err != nil {
		panic(err)
	}
}

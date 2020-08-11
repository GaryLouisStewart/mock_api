package utils

import (
	"database/sql"
	"github.com/gorilla/mux"
)

type ApplicationError struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status"`
	Code       string `json:"code"`
}

type healthCheckResponse struct {
	Status string `json:"status"`
}

type Database struct {
	DB *sql.DB
}

type App struct {
	Router *mux.Router
}

func (a *Database) Initialise(user, password, dbname string) {}

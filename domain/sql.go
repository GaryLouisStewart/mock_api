package domain

import (
	"database/sql"
	"errors"
)

func (p *Property) getProperty(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (p *Property) updateProperty(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (p *Property) deleteProperty(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (p *Property) createProperty(db *sql.DB) error {
	return errors.New("Not implemented")
}

func GetListOfProperties(db *sql.DB, start, count int) ([]*Property, error) {
	return nil, errors.New("Not implemented")
}

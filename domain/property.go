package domain

type Property struct {
	ID           uint    `json:"id" gorm:"primary_key"`
	PropertyType string  `json:"type,omitempty"`
	Bedrooms     int     `json:"bedrooms,omitempty"`
	Price        float64 `json:"price,omitempty"`
	Location     *Area   `json:"area,omitempty"`
}

type Area struct {
	Street   string `json:"street,omitempty"`
	City     string `json:"city,omitempty"`
	Province string `json:"province,omitempty"`
	Country  string `json:"country:omitempty"`
}

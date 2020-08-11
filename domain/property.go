package domain

type Property struct {
	ID       string `json:"id,omitempty"`
	Type     string `json:"type,omitempty"`
	Bedrooms int    `json:"bedrooms,omitempty"`
	Price    int    `json:"price,omitempty"`
	Area     *Area  `json:"area,omitempty"`
}

type Area struct {
	Street   string `json:"street,omitempty"`
	City     string `json:"city,omitempty"`
	Province string `json:"province,omitempty"`
}

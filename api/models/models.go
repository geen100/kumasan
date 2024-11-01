package models

type Sighting struct {
	ID             int     `json:"id"`
	Latitude       float64 `json:"latitude"`
	Longitude      float64 `json:"longitude"`
	City           string  `json:"city"`
	Area           string  `json:"area"`
	Date           string  `json:"date"`
	Time           string  `json:"time"`
	Situation      string  `json:"situation"`
	Details        string  `json:"details"`
	Classification string  `json:"classification"`
}

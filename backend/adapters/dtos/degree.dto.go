package dtos

type DegreeFilter struct {
	ID        *int    `json:"id,string,omitempty"`
	Alias     string  `json:"name,string,omitempty"`
	StartRate float64 `json:"start_rate,omitempty"`
	EndRate   float64 `json:"end_rate,omitempty"`
}

type AddDegree struct {
	Alias string   `json:"alias" validate:"required"`
	Rate  *float64 `json:"rate,omitempty"`
	ID    uint     `json:"degree_id,omitempty"`
}

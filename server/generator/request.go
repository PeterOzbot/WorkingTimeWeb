package generator

// Request : Request data as input for generator.
type Request struct {
	A_Hours int `json:"a_hours" binding:"required"`
	B_Hours int `json:"b_hours" binding:"required"`
	Month   int `json:"month" binding:"required"`
	Year    int `json:"year" binding:"required"`
}

package generator

//Request : Request data as input for generator.
type Request struct {
	TotalHours int `json:"totalHours" binding:"required"`
}

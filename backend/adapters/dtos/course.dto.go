package dtos

type CourseFilter struct {
	ID   *int   `json:"id,string,omitempty"`
	Alias string `json:"name,string,omitempty"`
}

type AddCourse struct {
	Alias string `json:"alias" validate:"required"`
}

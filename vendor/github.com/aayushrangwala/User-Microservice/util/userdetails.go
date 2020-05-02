package util

// UserDetails is the struct contains complete details of a user profile
type UserDetails struct {
	Name  string `json:"name"`
	Dob   string `json:"dob"`
	Age   int32  `json:"age"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

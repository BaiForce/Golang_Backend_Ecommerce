package models

// User represents the user model in the system
type User struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	JwtToken  string `json:"jwt_token"` // Add this field for storing JWT token
}

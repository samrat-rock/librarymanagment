package models

type Claims struct {
    Email string `json:"email"`
    Role  string `json:"role"` // Added Role field
    // Add other fields here
}
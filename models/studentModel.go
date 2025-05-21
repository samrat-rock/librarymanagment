package models

type Student struct {
	ID        uint   `gorm:"primaryKey"`
	FullName  string `json:"full_name" binding:"required"`
	Email     string `gorm:"unique" json:"email" binding:"required,email"`
	Phone     string `json:"phone" binding:"required"`
	Department string `json:"department"`
}

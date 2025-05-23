package models

type Student struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	FirstName string `json:"first_name" gorm:"not null"`
	LastName  string `json:"last_name" gorm:"not null"`
	Email     string `json:"email" gorm:"unique;not null"`
	Password  string `json:"password" gorm:"not null"`
	Role      string `json:"role" gorm:"default:'student'"`
	CreatedAt string `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt string `json:"updated_at" gorm:"autoUpdateTime"`
}
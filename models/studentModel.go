package models

import "time"

type Student struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	FirstName string    `json:"first_name" binding:"required"`
	LastName  string    `json:"last_name" binding:"required"`
	Email     string    `json:"email" gorm:"unique;not null"`
	Password  string    `json:"-" gorm:"not null"`
	Phone     string    `json:"phone" binding:"required"`
	Class     string    `json:"class" binding:"required"`
	RollNo    string    `json:"roll_no" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	ClassNumber int    `json:"class_number"`
}


type StudentRegister struct {
	FirstName   string `json:"first_name" binding:"required"`
	LastName    string `json:"last_name" binding:"required"`
	Email       string `json:"email" binding:"required,email"`
	Password    string `json:"password" binding:"required"`
	Phone       string `json:"phone" binding:"required"` 
	ClassNumber int    `json:"class_number" binding:"required"`
	RollNo      string `json:"roll_no" binding:"required"`
}



type StudentLogin struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type StudentUpdate struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Class     string `json:"class"`
	RollNo    string `json:"roll_no"`
	ClassNumber int    `json:"class_number"`

}

package models

import "time"

type Admin struct {
    ID        uint      `json:"id" gorm:"primaryKey"`
    FirstName string    `json:"first_name" gorm:"not null"`
    LastName  string    `json:"last_name" gorm:"not null"`
    Username  string    `json:"username" gorm:"unique;not null"`
    Email     string    `json:"email" gorm:"unique;not null"`
    Password  string    `json:"-" gorm:"not null"`
    Role      string    `json:"role" gorm:"default:'admin'"`
    CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
    UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

type AdminLogin struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type AdminRegister struct {
    FirstName string `json:"first_name" binding:"required"`
    LastName  string `json:"last_name" binding:"required"`
    Username  string `json:"username" binding:"required"`
    Email     string `json:"email" binding:"required,email"`
    Password  string `json:"password" binding:"required"`
    Role      string `json:"role"`
}


type AdminUpdate struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
	Role     string `json:"role" binding:"required"`
}

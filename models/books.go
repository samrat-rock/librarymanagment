package models

type Book struct {
	ID          uint   `gorm:"primaryKey"`
	Title       string `json:"title" binding:"required"`
	Author      string `json:"author" binding:"required"`
	Publisher   string `json:"publisher"`
	ISBN        string `gorm:"unique" json:"isbn" binding:"required"`
	Availability bool   `json:"availability" gorm:"default:true"`
}

package utils

import (
	"fmt"
	"librarymanagement/config"
	"librarymanagement/models"
)


func IsTokenBlacklisted(token string) bool {
	var blacklisted models.TokenBlacklist
	err := config.DB.Where("token = ?", token).First(&blacklisted).Error

	if err != nil {
		
		if err.Error() == "record not found" {
			return false
		}
	
		fmt.Println("Error checking token blacklist:", err)
		return true
	}

	return true 
}

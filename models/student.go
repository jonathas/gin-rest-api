package models

import "gorm.io/gorm"

type Student struct {
	gorm.Model // This inserts ID, CreatedAt, UpdatedAt, DeletedAt
	Name string `json:"name"`
	Document string `json:"document"`
}

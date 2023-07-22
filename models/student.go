package models

import (
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model // This inserts ID, CreatedAt, UpdatedAt, DeletedAt
	Name string `json:"name" validate:"nonzero"`
	Age int `json:"age" validate:"min=1,max=100"`
	Document string `json:"document" validate:"len=11,regexp=^[0-9]*$"`
}

func ValidateStudent(student *Student) error {
	if err := validator.Validate(student); err != nil {
		return err
	}
	return nil
}

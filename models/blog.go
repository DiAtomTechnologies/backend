package models

import (
	"time"
	"github.com/go-playground/validator/v10"
)

type Blog struct {
	ID          string    `json:"id"`
	ImgURL      string    `json:"imgurl" validate:"required,url"`
	Heading     string    `json:"heading" validate:"required"`
	Tag         string    `json:"tag" validate:"required"`
	Description string    `json:"description" validate:"required"`
	Content     string    `json:"content,omitempty" validate:"required"`
	Author      string    `json:"author" validate:"required"`
	CreatedAt   time.Time `json:"created_at"`
}

func ValidateBlog(s interface{}) error {
    validate := validator.New()
    return validate.Struct(s)
}

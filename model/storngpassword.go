package model

import "gorm.io/gorm"

type StorngPasswordRequest struct {
	Password string `json:"init_password"`
}

type StorngPasswordResponse struct {
	Steps int `json:"num_of_steps"`
}

type ErrorResponse struct {
	TimeStamp string `json:"timestamp"`
	Status    int    `json:"status"`
	Error     string `json:"error"`
	Path      string `json:"path"`
}

type ConfigStrongPassword struct {
	MinLowerCase int
	MinUpperCase int
	MinDigit     int
	Repeat       int
	MinLength    int
	MaxLength    int
}

type StorngPasswordLog struct {
	gorm.Model
	Ip       string `gorm:"column:ip"`
	Request  string `gorm:"column:request"`
	Response string `gorm:"column:response"`
}

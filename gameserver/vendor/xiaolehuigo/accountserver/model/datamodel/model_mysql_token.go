package datamodel

import (
	"time"
)

type Token struct {
	UserId int `gorm:"primary_key" json:"userId"`

	Token string `grom:"not null" json:"token"`

	LoginTime time.Time `json:"loginTime"`
}

func (Token) TableName() string {
	return "token_info"
}

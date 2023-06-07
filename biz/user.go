package main

import (
	"time"
)

type User struct {
	ID        uint64    `json:"id" gorm:"column:Id"`
	Name      string    `json:"name" gorm:"column:name"`
	Family    string    `json:"family" gorm:"column:family"`
	Age       uint8     `json:"age" gorm:"column:age"`
	Sex       string    `json:"sex" gorm:"column:sex"`
	CreatedAt time.Time `json:"createdAt" gorm:"column:create_at"`
}

type GetUserRequest struct {
	Id        uint64 `json:"id"`
	MessageID uint64 `json:"messageID"`
	AuthKey   string `json:"authKey"`
}

func (User) TableName() string {
	return "User"
}

package model

import (
    "time"
    "github.com/google/uuid"
)

type User struct {
    ID        uuid.UUID `gorm:"column:id;primaryKey"`
    Email     string    `gorm:"column:email;uniqueIndex;not null"`
    Password  string    `gorm:"column:password;not null"`
    CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
    UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime"`
}

func (User) TableName() string {
    return "users"
}
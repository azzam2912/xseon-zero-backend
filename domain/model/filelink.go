package model

import (
	"time"

	"github.com/google/uuid"
)

type FileLink struct {
	ID        uuid.UUID `gorm:"column:id;primaryKey"`
	Link      string    `gorm:"column:link;not null"`
	Caption   string    `gorm:"column:caption"`
	Category  string    `gorm:"column:category"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime"`
}

func (FileLink) TableName() string {
	return "file_links"
}

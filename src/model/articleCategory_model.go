package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ArticleCategory struct {
	ID        uuid.UUID `gorm:"primaryKey;default:uuid_generate_v4()" json:"id"`
	UserID    uuid.UUID `gorm:"not null" json:"user_id"`
	Name      string    `gorm:"not null" json:"name"`
	CreatedAt time.Time `gorm:"autoCreateTime:milli" json:"-"`
	UpdatedAt time.Time `gorm:"autoCreateTime:milli;autoUpdateTime:milli" json:"-"`
}

func (articleCategory *ArticleCategory) BeforeCreate(_ *gorm.DB) error {
	articleCategory.ID = uuid.New()
	return nil
}

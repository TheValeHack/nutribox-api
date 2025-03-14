package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Article struct {
	ID          uuid.UUID  `gorm:"primaryKey;default:uuid_generate_v4()" json:"id"`
	UserID      uuid.UUID  `gorm:"not null" json:"user_id"`
	Title       string     `gorm:"not null" json:"title"`
	CategoryID  *uuid.UUID `json:"category_id,omitempty"`
	Slug        string     `gorm:"unique;not null" json:"slug"`
	Image       *string    `json:"image,omitempty"`
	Content     string     `gorm:"type:text;not null" json:"content"`
	PublishedAt *time.Time `json:"published_at,omitempty"`
	CreatedAt   time.Time  `gorm:"autoCreateTime:milli" json:"-"`
	UpdatedAt   time.Time  `gorm:"autoCreateTime:milli;autoUpdateTime:milli" json:"-"`
}

func (article *Article) BeforeCreate(_ *gorm.DB) error {
	article.ID = uuid.New()
	return nil
}

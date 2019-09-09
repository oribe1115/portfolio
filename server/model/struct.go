package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type Base struct {
	ID        uuid.UUID  `gorm:"type:char(36);primary_key;"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"update_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`
}

func (base *Base) BeforeCreate(scope *gorm.Scope) error {
	id := uuid.New()
	return scope.SetColumn("ID", id)
}

type MainCategory struct {
	Base
	Name        string `gorm:"type:char(60) not null;"`
	Description string `gorm:"type:TEXT;"`

	SubCategories []SubCategory
}

type SubCategory struct {
	Base
	MainCategoryID uuid.UUID `gorm:"type:char(36);not null;"`
	Name           string    `gorm:"type:char(60);not null;"`
	Description    string    `gorm:"type:TEXT;"`
}

type Content struct {
	Base
	CategoryID     uuid.UUID `gorm:"type:char(36);not null;"`
	Title          string    `gorm:"type:char(60) not null;"`
	Image          string    `gorm:"type:char(200);"`
	Description    string    `gorm:"type:TEXT;"`
	Date           time.Time `json:"date"`
	MainImage      *MainImage
	SubImages      []*SubImage
	TaggedContents []*TaggedContent
	Tags           []*Tag
}

type MainImage struct {
	Base
	Name      string    `gorm:"type:char(60) not null;"`
	ContentID uuid.UUID `gorm:"type:char(36);not null;"`
	URL       string    `gorm:"type:char(200);"`
}

type SubImage struct {
	Base
	Name      string    `gorm:"type:char(60) not null;"`
	ContentID uuid.UUID `gorm:"type:char(36);not null;"`
	URL       string    `gorm:"type:char(200);"`
}

type Tag struct {
	Base
	Name        string `gorm:"type:char(60) not null;"`
	Description string `gorm:"type:TEXT;"`
}

type TaggedContent struct {
	Base
	TagID     uuid.UUID `gorm:"type:char(36);not null;"`
	ContentID uuid.UUID `gorm:"type:char(36);not null;"`
	Tag       *Tag
}

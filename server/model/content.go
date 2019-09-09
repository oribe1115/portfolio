package model

import (
	"github.com/google/uuid"
)

func NewContent(content *Content) (*Content, error) {
	if err := db.Create(content).Error; err != nil {
		return nil, err
	}
	return content, nil
}

func GetContentList() ([]*Content, error) {
	contentList := []*Content{}

	sub := db.Table("tagged_contents").Where("content_id = ?", "contents.id").Select("tag_id").SubQuery()

	if err := db.Preload("MainImage").Preload("SubImages").Preload("TaggedContents").Joins("LEFT JOIN tags ON tags.id = ?", sub).Find(&contentList).Error; err != nil {
		return nil, err
	}

	return contentList, nil
}

func GetContentByID(id uuid.UUID) (*Content, error) {
	content := &Content{}
	if err := db.Preload("MainImage").Preload("SubImages").Preload("TaggedContents").Find(&content).Error; err != nil {
		return nil, err
	}

	return content, nil
}

func SaveContent(content *Content) (*Content, error) {
	if err := db.Save(&content).Error; err != nil {
		return nil, err
	}
	return content, nil
}

func IsExistContentID(contentID uuid.UUID) bool {
	count := 0
	if err := db.Table("contents").Where("id = ?", contentID).Count(&count).Error; err != nil {
		return false
	}
	return count > 0
}

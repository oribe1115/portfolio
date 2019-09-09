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

	if err := db.Preload("MainImage").Find(&contentList).Error; err != nil {
		return nil, err
	}

	return contentList, nil
}

func GetContentByID(id uuid.UUID) (*Content, error) {
	content := &Content{}
	if err := db.Preload("MainImage").Preload("SubImages").Joins("LEFT JOIN sub_categories ON sub_categories.id = contents.category_id").Joins("LEFT JOIN main_categories ON main_categories.id = sub_categories.main_category_id").Find(&content).Error; err != nil {
		return nil, err
	}
	if err := db.Preload("Tag").Where("content_id = ?", content.ID).Find(&content.TaggedContents).Error; err != nil {
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

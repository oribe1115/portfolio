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

	if err := db.Find(&contentList).Error; err != nil {
		return nil, err
	}

	return contentList, nil
}

func GetContentByID(id uuid.UUID) (*Content, error) {
	content := &Content{}
	if err := db.Preload("TaggedContents").Find(&content).Error; err != nil {
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

func IsExistCategoryID(categoryID uuid.UUID) bool {
	mainCount := 0
	subCount := 0
	if err := db.Table("main_categories").Where("id = ?", categoryID).Count(&mainCount).Error; err != nil {
		return false
	}
	if err := db.Table("sub_categories").Where("id = ?", categoryID).Count(&subCount).Error; err != nil {
		return false
	}

	return mainCount+subCount > 0
}

func IsExistContentID(contentID uuid.UUID) bool {
	count := 0
	if err := db.Table("contents").Where("id = ?", contentID).Count(&count).Error; err != nil {
		return false
	}
	return count > 0
}

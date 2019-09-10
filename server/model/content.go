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
	if err := db.Preload("MainImage").Preload("SubImages").Find(&content).Error; err != nil {
		return nil, err
	}

	if err := db.Preload("Tag").Where("content_id = ?", content.ID).Find(&content.TaggedContents).Error; err != nil {
		return nil, err
	}

	if err := db.Where("id = ?", content.CategoryID).Find(&content.SubCategory).Error; err != nil {
		return nil, err
	}

	if err := db.Where("id = ?", content.SubCategory.MainCategoryID).Find(&content.MainCategory).Error; err != nil {
		return nil, err
	}

	return content, nil
}

func GetContentListByMainCategory(mainID uuid.UUID) ([]*Content, error) {
	contentList := []*Content{}
	sub := db.Table("sub_categories").Where("main_category_id = ?", mainID).Select("id").SubQuery()
	if err := db.Preload("MainImage").Where("category_id IN ?", sub).Find(&contentList).Error; err != nil {
		return nil, err
	}

	return contentList, nil
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

package model

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

func NewContent(content *Content) (*Content, error) {
	if err := db.Create(content).Error; err != nil {
		return nil, err
	}
	return content, nil
}

func GetContentList() ([]*Content, error) {
	contentList := []*Content{}

	if err := db.Preload("MainImage").Order("date desc").Find(&contentList).Error; err != nil {
		return nil, err
	}

	return contentList, nil
}

func GetContentByID(id uuid.UUID) (*Content, error) {
	content := &Content{}
	if err := db.Preload("MainImage").Preload("SubImages", func(db *gorm.DB) *gorm.DB {
		return db.Order("sub_images.created_at ASC")
	}).Where("id = ?", id).Find(&content).Error; err != nil {
		return nil, err
	}

	if err := db.Preload("Tag").Where("content_id = ?", id).Find(&content.TaggedContents).Error; err != nil {
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
	if err := db.Preload("MainImage").Where("category_id IN ?", sub).Order("date desc").Find(&contentList).Error; err != nil {
		return nil, err
	}

	return contentList, nil
}

func GetContentListBySubCategory(subID uuid.UUID) ([]*Content, error) {
	contentList := []*Content{}
	if err := db.Preload("MainImage").Where("category_id = ?", subID).Order("date desc").Find(&contentList).Error; err != nil {
		return nil, err
	}

	return contentList, nil
}

func GetContentListByTag(tagID uuid.UUID) ([]*Content, error) {
	contentList := []*Content{}
	sub := db.Table("tagged_contents").Where("tag_id = ?", tagID).Select("content_id").SubQuery()
	if err := db.Preload("MainImage").Where("id IN ?", sub).Order("date desc").Find(&contentList).Error; err != nil {
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

func IGetContentList() ([]*Content, error) {
	contentList := []*Content{}

	sub1 := db.Table("main_categories").Where("name LIKE ?", ".%").Select("id").SubQuery()
	sub2 := db.Table("sub_categories").Where("main_category_id IN ?", sub1).Select("id").SubQuery()

	if err := db.Preload("MainImage").Not("category_id IN ?", sub2).Order("date desc").Find(&contentList).Error; err != nil {
		return nil, err
	}

	return contentList, nil
}

func IGetContentListByTag(tagID uuid.UUID) ([]*Content, error) {
	contentList := []*Content{}
	sub1 := db.Table("tagged_contents").Where("tag_id = ?", tagID).Select("content_id").SubQuery()
	sub2 := db.Table("main_categories").Where("name LIKE ?", ".%").Select("id").SubQuery()
	sub3 := db.Table("sub_categories").Where("main_category_id IN ?", sub2).Select("id").SubQuery()
	if err := db.Preload("MainImage").Where("id IN ?", sub1).Not("category_id IN ?", sub3).Order("date desc").Find(&contentList).Error; err != nil {
		return nil, err
	}

	return contentList, nil
}

func IsNotIgnoredContent(contentID uuid.UUID) bool {
	count := 0
	sub1 := db.Table("main_categories").Where("name LIKE ?", ".%").Select("id").SubQuery()
	sub2 := db.Table("sub_categories").Where("main_category_id IN ?", sub1).Select("id").SubQuery()
	if err := db.Table("contents").Where("id = ?", contentID).Not("category_id IN ?", sub2).Count(&count).Error; err != nil {
		return false
	}

	return count > 0
}

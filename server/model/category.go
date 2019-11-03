package model

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

func NewMainCategory(mainCategory *MainCategory) (*MainCategory, error) {
	if err := db.Create(mainCategory).Error; err != nil {
		return nil, err
	}
	return mainCategory, nil
}

func GetMainCategories() ([]*MainCategory, error) {
	mainCategories := []*MainCategory{}

	if err := db.Preload("SubCategories", func(db *gorm.DB) *gorm.DB {
		return db.Order("sub_categories.created_at ASC")
	}).Order("created_at").Find(&mainCategories).Error; err != nil {
		return nil, err
	}

	return mainCategories, nil
}

func GetMainCategoryByID(mainID uuid.UUID) (*MainCategory, error) {
	mainCategory := &MainCategory{}

	if err := db.Preload("SubCategories").Where("id = ?", mainID).First(&mainCategory).Error; err != nil {
		return nil, err
	}

	return mainCategory, nil
}

func SaveMainCategory(mainCategory *MainCategory) (*MainCategory, error) {
	if err := db.Save(mainCategory).Error; err != nil {
		return nil, err
	}

	return mainCategory, nil
}

func NewSubCategory(subCategory *SubCategory) (*SubCategory, error) {
	if err := db.Create(subCategory).Error; err != nil {
		return nil, err
	}
	return subCategory, nil
}

func GetSubCategory(subID uuid.UUID) (*SubCategory, error) {
	subCategory := &SubCategory{}
	if err := db.Where("id = ?", subID).Find(subCategory).Error; err != nil {
		return nil, err
	}
	return subCategory, nil
}

func SaveSubCategory(subCategory *SubCategory) (*SubCategory, error) {
	if err := db.Save(subCategory).Error; err != nil {
		return nil, err
	}
	return subCategory, nil
}

func GetSubCategories() ([]*SubCategory, error) {
	subCategories := []*SubCategory{}

	if err := db.Table("sub_categories").Find(&subCategories).Error; err != nil {
		return nil, err
	}

	return subCategories, nil
}

func IsMainCategory(mainID uuid.UUID) bool {
	count := 0
	if err := db.Table("main_categories").Where("id = ?", mainID).Count(&count).Error; err != nil {
		return false
	}
	return count > 0
}

func IsExistSubCategoryID(categoryID uuid.UUID) bool {
	count := 0
	if err := db.Table("sub_categories").Where("id = ?", categoryID).Count(&count).Error; err != nil {
		return false
	}

	return count > 0
}

func IGetMainCategories() ([]*MainCategory, error) {
	mainCategories := []*MainCategory{}

	if err := db.Preload("SubCategories", "name NOT LIKE ?", ".%", func(db *gorm.DB) *gorm.DB {
		return db.Order("sub_categories.created_at ASC")
	}).Not("name LIKE ?", ".%").Order("created_at").Find(&mainCategories).Error; err != nil {
		return nil, err
	}

	return mainCategories, nil
}

func IGetSubCategories() ([]*SubCategory, error) {
	subCategories := []*SubCategory{}

	sub := db.Table("main_categories").Where("name LIKE ?", ".%").Select("id").SubQuery()

	if err := db.Table("sub_categories").Not("main_category_id IN ?", sub).Find(&subCategories).Error; err != nil {
		return nil, err
	}

	return subCategories, nil
}

func IsNotIgnoredMainCategory(mainID uuid.UUID) bool {
	count := 0
	if err := db.Table("main_categories").Where("id = ?", mainID).Not("name LIKE ?", ".%").Count(&count).Error; err != nil {
		return false
	}

	return count > 0
}

func IsNotIgnoredSubCategory(subID uuid.UUID) bool {
	count := 0
	sub := db.Table("main_categories").Where("name LIKE ?", ".%").Select("id").SubQuery()
	if err := db.Table("sub_categories").Where("id = ?", subID).Not("main_category_id IN ?", sub).Count(&count).Error; err != nil {
		return false
	}

	return count > 0
}

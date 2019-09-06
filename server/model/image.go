package model

import (
	"github.com/google/uuid"
)

func NewSubImage(subImage *SubImage) (*SubImage, error) {
	if err := db.Create(subImage).Error; err != nil {
		return nil, err
	}
	return subImage, nil
}

func GetSubImage(subImageID uuid.UUID) (*SubImage, error) {
	subImage := &SubImage{}
	if err := db.Where("id = ?", subImageID).First(subImage).Error; err != nil {
		return nil, err
	}
	return subImage, nil
}

func GetSubImageList() ([]*SubImage, error) {
	subImageList := []*SubImage{}
	if err := db.Find(&subImageList).Error; err != nil {
		return nil, err
	}
	return subImageList, nil
}

func IsExistSubImageID(subImageID uuid.UUID) bool {
	count := 0
	if err := db.Table("sub_images").Where("id = ?", subImageID).Count(&count).Error; err != nil {
		return false
	}
	return count > 0
}

func DeleteSubImage(subImage *SubImage) error {
	if err := db.Delete(subImage).Error; err != nil {
		return err
	}
	return nil
}

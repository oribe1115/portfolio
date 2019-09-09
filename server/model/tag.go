package model

import (
	"github.com/google/uuid"
)

func NewTag(tag *Tag) (*Tag, error) {
	if err := db.Create(tag).Error; err != nil {
		return nil, err
	}
	return tag, nil
}

func SaveTag(tag *Tag) (*Tag, error) {
	if err := db.Save(tag).Error; err != nil {
		return nil, err
	}
	return tag, nil
}

func GetTag(tagID uuid.UUID) (*Tag, error) {
	tag := &Tag{}
	if err := db.Where("id = ?", tagID).Find(tag).Error; err != nil {
		return nil, err
	}
	return tag, nil
}

func GetTagList() ([]*Tag, error) {
	tagList := []*Tag{}
	if err := db.Find(&tagList).Error; err != nil {
		return nil, err
	}
	return tagList, nil
}

func NewTaggedContent(taggedContent *TaggedContent) (*TaggedContent, error) {
	if err := db.Create(taggedContent).Error; err != nil {
		return nil, err
	}
	return taggedContent, nil
}

func GetTaggedContent(taggedContentID uuid.UUID) (*TaggedContent, error) {
	taggedContent := &TaggedContent{}
	if err := db.Where("id = ?", taggedContentID).Find(&taggedContent).Error; err != nil {
		return nil, err
	}
	return taggedContent, nil
}

func GetTaggedContentList() ([]*TaggedContent, error) {
	taggedContentList := []*TaggedContent{}
	if err := db.Find(&taggedContentList).Error; err != nil {
		return nil, err
	}
	return taggedContentList, nil
}

func DeleteTaggedContent(taggedContent *TaggedContent) error {
	if err := db.Delete(taggedContent).Error; err != nil {
		return err
	}
	return nil
}

func IsExistTagID(tagID uuid.UUID) bool {
	count := 0
	if err := db.Table("tags").Where("id = ?", tagID).Count(&count).Error; err != nil {
		return false
	}
	return count > 0
}

func IsExistTaggedContentID(taggedContentID uuid.UUID) bool {
	count := 0
	if err := db.Table("tagged_contents").Where("id = ?", taggedContentID).Count(&count).Error; err != nil {
		return false
	}
	return count > 0
}

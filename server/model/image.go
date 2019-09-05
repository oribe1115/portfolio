package model

func NewSubImage(subImage *SubImage) (*SubImage, error) {
	if err := db.Create(subImage).Error; err != nil {
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

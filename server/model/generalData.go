package model

func NewGeneralData(generalData *GeneralData) (*GeneralData, error) {
	if err := db.Create(generalData).Error; err != nil {
		return nil, err
	}

	return generalData, nil
}

func GetGeneralDataBySubject(subject string) (*GeneralData, error) {
	generalData := &GeneralData{}
	if err := db.Where("subject = ?", subject).Order("created_at desc").First(&generalData).Error; err != nil {
		return nil, err
	}

	return generalData, nil
}

func IsExistSubject(subject string) bool {
	count := 0
	if err := db.Table("general_data").Where("subject = ?", subject).Count(&count).Error; err != nil {
		return false
	}

	return count > 0
}

func GetAllGeneralData() ([]*GeneralData, error) {
	generalDataList := []*GeneralData{}
	if err := db.Table("general_data").Select("*").Group("subject").Having("created_at = max(created_at)").Find(&generalDataList).Error; err != nil {
		return nil, err
	}

	return generalDataList, nil
}

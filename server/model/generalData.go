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
	rows, err := db.Raw("SELECT * FROM general_data WHERE created_at = (SELECT MAX(created_at) FROM general_data AS gd WHERE general_data.subject = gd.subject)").Rows()
	defer rows.Close()
	if err != nil {
		return nil, err
	}

	generalDataList := make([]*GeneralData, 0)
	for rows.Next() {
		generalData := &GeneralData{}
		db.ScanRows(rows, &generalData)

		generalDataList = append(generalDataList, generalData)
	}

	return generalDataList, nil
}

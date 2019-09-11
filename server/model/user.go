package model

func NewUser(user *User) (*User, error) {
	if err := db.Create(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func GetUser(userName string) (*User, error) {
	user := &User{}
	if err := db.Where("user_name = ?", userName).Find(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func IsExistUserName(userName string) bool {
	count := 0
	if err := db.Table("users").Where("user_name = ?", userName).Count(&count).Error; err != nil {
		return false
	}

	return count > 0
}

func IsNotExistUser() bool {
	count := 0
	if err := db.Table("users").Count(&count).Error; err != nil {
		return false
	}

	return count == 0
}

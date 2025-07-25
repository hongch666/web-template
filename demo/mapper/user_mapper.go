package mapper

import (
	"demo/model"

	"gorm.io/gorm"
)

type UserMapper struct {
	DB *gorm.DB
}

func (m *UserMapper) GetAll() ([]model.User, error) {
	var users []model.User
	err := m.DB.Find(&users).Error
	return users, err
}

func (m *UserMapper) Create(user *model.User) error {
	return m.DB.Create(user).Error
}

package service

import (
	"demo/mapper"
	"demo/model"
)

type UserService struct {
	Mapper *mapper.UserMapper
}

func (s *UserService) GetAllUsers() ([]model.User, error) {
	return s.Mapper.GetAll()
}

func (s *UserService) CreateUser(user *model.User) error {
	return s.Mapper.Create(user)
}

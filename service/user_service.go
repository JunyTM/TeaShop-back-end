package service

import (
	"teashop/model"
	"teashop/repository"
)

type userService struct {
	userRepository model.UserRepository
}

type UserService interface {
	GetById(id int) (*model.User, error)
	GetAll() ([]*model.User, error)
	Create(newRecord *model.User) (*model.User, error)
	Update(id int, newRecord *model.User) (*model.User, error)
	Delete(id int) error
}

func (s *userService) GetById(id int) (*model.User, error) {
	record, err := s.userRepository.GetById(id)
	if err != nil {
		return nil, err
	}
	return record, nil
}

func (s *userService) GetAll() ([]*model.User, error) {
	records, err := s.userRepository.GetAll()
	if err != nil {
		return nil, err
	}
	return records, nil
}

func (s *userService) Create(newRecord *model.User) (*model.User, error) {
	record, err := s.userRepository.Create(newRecord)
	if err != nil {
		return nil, err
	}
	return record, nil
}

func (s *userService) Update(id int, newRecord *model.User) (*model.User, error) {
	record, err := s.userRepository.Update(id, newRecord)
	if err != nil {
		return nil, err
	}
	return record, nil
}

func (s *userService) Delete(id int) error {
	if err := s.userRepository.Delete(id); err != nil {
		return err
	}
	return nil
}

func NewUserService() UserService {
	userRepository := repository.NewUserRepository()
	return &userService{
		userRepository: userRepository,
	}
}


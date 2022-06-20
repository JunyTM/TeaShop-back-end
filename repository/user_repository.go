package repository

import (
	"teashop/infrastructure"
	"teashop/model"
)

type userRepository struct {}

func (r *userRepository) GetById(id int) (*model.User, error) {
	db := infrastructure.GetDB()
	var record model.User
	if err := db.Where("id = ?", id).First(&record).Error; err != nil {
		return nil, err
	}
	return &record, nil
}

func (r *userRepository) GetAll() ([]*model.User, error) {
	db := infrastructure.GetDB()
	var records []*model.User
	if err := db	.Find(&records).Error; err != nil {
		return nil, err
	}
	return records, nil
}

func (r *userRepository) Create(newRecord *model.User) (*model.User, error) {
	db := infrastructure.GetDB()
	var record *model.User
	if err := db.Create(newRecord).First(&record).Error; err != nil {
		return nil, err
	}
	return record, nil
}

func (r *userRepository) Update(id int, newRecord *model.User) (*model.User, error) {
	db := infrastructure.GetDB()

	var record *model.User
	if err := db.Where("id = ?", id).Updates(newRecord).First(&record).Error; err != nil {
		return nil, err
	}
	return record, nil
}

func (r *userRepository) Delete(id int) error {
	db := infrastructure.GetDB()
	var record model.User
	if err := db.Where("id = ?", id).Delete(&record).Error; err != nil {
		return err
	}
	return nil
}

func NewUserRepository() model.UserRepository {
	return &userRepository{}
}
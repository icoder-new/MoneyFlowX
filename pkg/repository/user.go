package repository

import (
	"fr33d0mz/moneyflowx/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (u *UserRepository) FindAll() ([]*models.User, error) {
	var users []*models.User

	err := u.db.Find(&users).Error
	if err != nil {
		return users, err
	}

	return users, nil
}

func (u *UserRepository) FindById(id string) (*models.User, error) {
	var user *models.User

	err := u.db.Where("id =?", id).Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (u *UserRepository) FindByUsernameName(username string) (*models.User, error) {
	var user *models.User

	err := u.db.Where("username =?", username).Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (u *UserRepository) FindByName(name string) ([]*models.User, error) {
	var users []*models.User

	err := u.db.Where("firstname ILIKE ?", "%"+name+"%").Find(&users).Error
	if err != nil {
		return users, err
	}

	return users, nil
}

func (u *UserRepository) FindByEmail(email string) (*models.User, error) {
	var user *models.User

	err := u.db.Where("email =?", email).Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (u *UserRepository) Update(user *models.User) (*models.User, error) {
	err := u.db.Save(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

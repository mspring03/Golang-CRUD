package Repository

import "gorm.io/gorm"

type userRepo struct {
	db *gorm.DB
}

func UserRepo (db *gorm.DB) *userRepo {
	return &userRepo{db}
}

func (r *userRepo) CreateUser() () {

}
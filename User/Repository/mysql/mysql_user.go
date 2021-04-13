package mysql

import (
	"fmt"
	"github.com/mspring03/Golang-CRUD/Models"
	"gorm.io/gorm"
)

type userRepo struct {
	database *gorm.DB
	userModel *Models.User
}

func UserRepo (db *gorm.DB, um *Models.User) *userRepo {
	return &userRepo{db, um}
}

func (ur *userRepo) CreateUser(id string, pw string, age uint8) {
	user := Models.User{ID: id, Password: pw, Age: age}
	fmt.Println(user)
	ur.database.Create(&user)
}

func (ur *userRepo) FindOneId(id string) *Models.User {
	ur.database.Where("ID = ?", id).Find(ur.userModel)

	return ur.userModel
}
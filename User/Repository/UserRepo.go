package Repository

import (
	"fmt"
	"github.com/mspring03/Golang-CRUD/Models"
	"gorm.io/gorm"
)

type userRepo struct {
	db *gorm.DB
	um *Models.User
}

func UserRepo (db *gorm.DB, um *Models.User) *userRepo {
	return &userRepo{db, um}
}

func (ur *userRepo) CreateUser(id string, pw string, age uint8) {
	user := Models.User{ID: id, Password: pw, Age: age}
	fmt.Println(user)
	ur.db.Create(&user)
}

func (ur *userRepo) FindOneId(id string) *Models.User {
	ur.db.Where("ID = ?", id).Find(ur.um)

	return ur.um
}
package mysql

import (
	"context"
	"database/sql"
	"github.com/mspring03/Golang-CRUD/domain"

	//"fmt"

	"github.com/mspring03/Golang-CRUD/models"
	"github.com/sirupsen/logrus"
)

type userRepo struct {
	database *sql.DB
}

func UserRepo (db *sql.DB) *userRepo {
	return &userRepo{
		database: db,
	}
}

func (ur *userRepo) fetch(ctx context.Context, query string, args ...interface{}) (result []domain.User, err error) {
	rows, err := ur.database.QueryContext(ctx, query, args...)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	defer func() {
		errRow := rows.Close()
		if errRow != nil {
			logrus.Error(errRow)
		}
	}()

	result = make([]domain.User, 0)
	for rows.Next() {
		t := domain.User{}
		err = rows.Scan(
			&t.Id,
			&t.Password,
			&t.Age,
			&t.UpdatedAt,
			&t.CreatedAt,
		)

		if err != nil {
			logrus.Error(err)
			return nil, err
		}
		result = append(result, t)
	}

	return result, nil
}

func (ur *userRepo) CreateUser(id string, pw string, age uint8) {
	//user := models.User{ID: id, Password: pw, Age: age}
	//fmt.Println(user)
	//ur.database.Create(&user)
}

func (ur *userRepo) IdConflictCheck(ctx context.Context, Id string) (res *domain.User, err error) {
	query := `SELECT ID, Password, Age, updated_at, created_at
  						FROM user WHERE ID = ? ORDER BY created_at`

	user, err := ur.fetch(ctx, query, Id)
	if err != nil {
		return &domain.User{}, err
	}

	if len(user) > 0 {
		res = &user[0]
	} else {
		return res, domain.ErrNotFound
	}


	return
}
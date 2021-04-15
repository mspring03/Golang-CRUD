package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/mspring03/Golang-CRUD/domain"

	//"fmt"

	"github.com/sirupsen/logrus"
)

type userRepo struct {
	database *sql.DB
}

func UserRepo (db *sql.DB) (ur *userRepo) {
	ur = &userRepo{
		database: db,
	}

	_, _ = db.Exec(`
		CREATE TABLE user ( 
		id VARCHAR(30) PRIMARY KEY,
		password VARCHAR(30) NOT NULL, 
		age INT DEFAULT 0, 
		created_at timestamp not null default current_timestamp,
		updated_at timestamp not null default current_timestamp on update current_timestamp,
	
		index (created_at),
		index (updated_at)
	) ENGINE=INNODB`)

	return
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

func (ur *userRepo) CreateUser(ctx context.Context, user *domain.User) (err error) {
	query := `INSERT user SET Id=? , Password=? , Age=?`
	_, err = ur.fetch(ctx, query, user.Id, user.Password, user.Age)
	if err != nil {
		fmt.Println("hello")
		fmt.Println(err)
	}
	return
}

func (ur *userRepo) IdConflictCheck(ctx context.Context, Id string) (res *domain.User, err error) {
	query := `SELECT ID, Password, Age, updated_at, created_at FROM user WHERE id = ? ORDER BY created_at`

	user, err := ur.fetch(ctx, query, Id)

	if err != nil {
		fmt.Println("hello")
		fmt.Println(err)
	}

	if len(user) > 0 {
		return res, domain.ErrConflict
	} else {
		return &domain.User{}, err
	}
}
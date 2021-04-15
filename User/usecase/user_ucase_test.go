package usecase_test

import (
	//"context"
	//"errors"
	"testing"
	//"time"
	//
	//"github.com/stretchr/testify/assert"
	//"github.com/stretchr/testify/mock"
	//
	//ucase "github.com/mspring03/Golang-CRUD/user/usecase"
	"github.com/mspring03/Golang-CRUD/domain"
	//"github.com/mspring03/Golang-CRUD/domain/mocks"
)

func TestSignup(t *testing.T) {
	//mockUserRepo := new(mocks.UserRepository)
	mockUser := domain.User{
		Id: "test_id",
		Password: "test_password",
		Age: 19,
	}

	mockListUser := make([]domain.User, 0)
	mockListUser = append(mockListUser, mockUser)

	t.Run("success", func(t *testing.T) {

		t.Log(mockListUser)
	})

	t.Run("error-failed", func(t *testing.T) {
		t.Log(mockListUser)
	})
}
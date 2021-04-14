package mocks

import (
	"context"

	"github.com/stretchr/testify/mock"
)

type ArticleRepository struct {
	mock.Mock
}

func (_m *ArticleRepository) Delete(ctx context.Context, id int64) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}



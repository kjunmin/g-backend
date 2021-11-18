package mocks

import (
	"context"

	"github.com/google/uuid"
	"github.com/kjunmin/g-backend/model"
	"github.com/stretchr/testify/mock"
)

type MockUserService struct {
	mock.Mock
}

func (m *MockUserService) Get(ctx context.Context, uid uuid.UUID) (r0 *model.User, r1 error) {
	ret := m.Called(ctx, uid)

	if ret.Get(0) != nil {
		r0 = ret.Get(0).(*model.User)
	}

	if ret.Get(1) != nil {
		r1 = ret.Get(1).(error)
	}
	return
}

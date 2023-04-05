package service

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"myapp/ent"
	"myapp/internal/applications/user/dto"
	mockRoleRepo "myapp/mocks/role/repository"
	mockRoleUserRepo "myapp/mocks/role_user/repository"
	mockTrx "myapp/mocks/transaction"
	"myapp/mocks/user/repository"
	"testing"
	"time"
)

var mockUserRepository = new(mock_repository.UserRepository)
var mockRoleRepository = new(mockRoleRepo.RoleRepository)
var mockRoleUserRepository = new(mockRoleUserRepo.RoleUserRepository)
var mockTransaction = new(mockTrx.TrxService)
var service = NewUserServiceImpl(mockUserRepository, mockRoleRepository, mockRoleUserRepository, mockTransaction)

func getUserMock(id uint64, name string, email string, password string) ent.User {
	return ent.User{
		ID:               id,
		Name:             name,
		Email:            email,
		IsVerified:       true,
		EmailVerifiedAt:  time.Time{},
		Password:         password,
		RememberToken:    "",
		SocialMediaID:    "",
		Avatar:           "",
		CreatedAt:        time.Time{},
		UpdatedAt:        time.Time{},
		RoleID:           0,
		LoginType:        "",
		SubSpecialist:    "",
		FirebaseToken:    "",
		Info:             "",
		Description:      "",
		Specialist:       "",
		Phone:            "",
		LastAccessAt:     time.Time{},
		PregnancyMode:    false,
		DeletedAt:        time.Time{},
		LatestSkipUpdate: time.Time{},
		LatestDeletedAt:  time.Time{},
	}
}

func TestUserServiceImpl_Create(t *testing.T) {
	ctx := context.Background()
	request := dto.UserRequest{
		RoleId:   0,
		Name:     "Admin",
		Email:    "admin@tentanganak.id",
		Password: "12345",
	}

	userMocks := []struct {
		name         string
		request      dto.UserRequest
		userRequest  ent.User
		userResponse ent.User
		scenario     bool
	}{
		{
			request:      request,
			name:         "Create_User_Success-1",
			userRequest:  getUserMock(uint64(0), "Admin", "admin@tentanganak.id", "12345"),
			userResponse: getUserMock(uint64(123000), "Admin", "admin@tentanganak.id", "12345"),
			scenario:     true,
		},
		{
			request:      request,
			name:         "Create_User_Success-2",
			userRequest:  getUserMock(uint64(0), "Admin", "admin@tentanganak.id", "12345"),
			userResponse: getUserMock(uint64(123001), "Admin", "admin@tentanganak.id", "12345"),
			scenario:     true,
		},
		{
			request:      request,
			name:         "Create_User_Failed-1",
			userRequest:  getUserMock(uint64(0), "Admin", "admin@tentanganak.id", "12345"),
			userResponse: getUserMock(uint64(123001), "Admin", "admin@tentanganak.id", "12345"),
			scenario:     false,
		},
	}

	//table test:
	for _, userMock := range userMocks {
		t.Run(userMock.name, func(t *testing.T) {

			if userMock.scenario {
				mockUserRepository.On("Create", ctx, mock.Anything, userMock.userRequest).Return(&userMock.userResponse, nil)
				mockRoleUserRepository.On("Create", ctx, mock.Anything, mock.Anything).Return(&ent.RoleUser{}, nil)
				mockTransaction.On("WithTx", ctx, mock.Anything).
					Run(func(args mock.Arguments) {
						f := args.Get(1).(func(tx *ent.Tx) error)
						f(ent.TxFromContext(ctx))
					}).Return(nil).Once()

				result, err := service.Create(ctx, &userMock.request)
				assert.NoError(t, err)
				assert.NotNil(t, result)
				assert.Equal(t, &userMock.userResponse, result)
			} else {
				mockUserRepository.On("Create", ctx, mock.Anything, userMock.userRequest).Return(&userMock.userResponse, nil)
				mockRoleUserRepository.On("Create", ctx, mock.Anything, mock.Anything).Panic("failed saved")
				mockTransaction.On("WithTx", ctx, mock.Anything).
					Run(func(args mock.Arguments) {
						f := args.Get(1).(func(tx *ent.Tx) error)
						f(ent.TxFromContext(ctx))
					}).Return(errors.New("failed saved")).Once()

				result, err := service.Create(ctx, &request)
				assert.NotNil(t, err)
				assert.Nil(t, result)
			}

		})
	}

}

func TestUserServiceImpl_Update(t *testing.T) {
	ctx := context.Background()
	request := dto.UserRequest{
		RoleId:   0,
		Name:     "User",
		Email:    "user@tentanganak.id",
		Password: "12345",
	}

	id := uint64(123000)
	userRequest := getUserMock(uint64(0), "User", "user@tentanganak.id", "12345")
	userResponse := getUserMock(uint64(123000), "User", "user@tentanganak.id", "12345")

	t.Run("Update_User_Success", func(t *testing.T) {

		mockUserRepository.On("Update", ctx, mock.Anything, userRequest, id).Return(&userResponse, nil)
		mockRoleUserRepository.On("Update", ctx, mock.Anything, mock.Anything, id).Return(&ent.RoleUser{}, nil)
		mockTransaction.On("WithTx", ctx, mock.Anything).
			Run(func(args mock.Arguments) {
				f := args.Get(1).(func(tx *ent.Tx) error)
				f(ent.TxFromContext(ctx))
			}).Return(nil).Once()

		result, err := service.Update(ctx, id, &request)
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, &userResponse, result)
	})

	t.Run("Update_User_Failed_User", func(t *testing.T) {
		err := errors.New("failed saved user")
		mockUserRepository.On("Update", ctx, mock.Anything, userRequest, id).Return(nil, err)
		mockTransaction.On("WithTx", ctx, mock.Anything).
			Run(func(args mock.Arguments) {
				f := args.Get(1).(func(tx *ent.Tx) error)
				f(ent.TxFromContext(ctx))
			}).Return(err).Once()

		result, err := service.Update(ctx, id, &request)
		assert.NotNil(t, err)
		assert.Nil(t, result)
	})

	t.Run("Update_User_Failed_Role_User", func(t *testing.T) {
		err := errors.New("failed saved role")
		mockUserRepository.On("Update", ctx, mock.Anything, userRequest, id).Return(&userResponse, nil)
		mockRoleUserRepository.On("Update", ctx, mock.Anything, mock.Anything, id).Return(nil, err)
		mockTransaction.On("WithTx", ctx, mock.Anything).
			Run(func(args mock.Arguments) {
				f := args.Get(1).(func(tx *ent.Tx) error)
				f(ent.TxFromContext(ctx))
			}).Return(err).Once()

		result, err := service.Update(ctx, id, &request)
		assert.NotNil(t, err)
		assert.Nil(t, result)
	})

}

func TestUserServiceImpl_Delete(t *testing.T) {
	ctx := context.Background()
	userMock := getUserMock(uint64(123000), "User-1", "user1@email.com", "12345")
	t.Run("Delete_success", func(t *testing.T) {
		mockUserRepository.On("SoftDelete", ctx, uint64(123000)).Return(&userMock, nil).Once()
		result, err := service.Delete(context.Background(), userMock.ID)
		assert.NoError(t, err)
		assert.NotNil(t, result)
	})

	t.Run("Delete_failed", func(t *testing.T) {
		errorMessage := errors.New("failed got user")
		mockUserRepository.On("SoftDelete", ctx, uint64(123000)).Return(nil, errorMessage).Once()
		result, err := service.Delete(context.Background(), userMock.ID)
		assert.NotNil(t, err)
		assert.Nil(t, result)
	})

}

func TestUserServiceImpl_GetById(t *testing.T) {
	ctx := context.Background()
	userMocks := []struct {
		id       uint64
		name     string
		expected ent.User
	}{
		{
			id:       uint64(10),
			name:     "GetById_User-1",
			expected: getUserMock(uint64(123000), "User-1", "user1@email.com", "12345"),
		},
		{
			id:       uint64(10),
			name:     "GetById_User-2",
			expected: getUserMock(uint64(11000), "User-2", "user2@email.com", "12345"),
		},
	}

	//table test:
	for _, userMock := range userMocks {
		t.Run(userMock.name, func(t *testing.T) {
			mockUserRepository.On("GetById", ctx, uint64(10)).Return(&userMock.expected, nil).Once()
			result, err := service.GetById(ctx, userMock.id)
			assert.NoError(t, err)
			assert.NotNil(t, result)
		})
	}

	//subtest failed:
	t.Run("GetById_failed", func(t *testing.T) {
		errorMessage := errors.New("failed got user")
		mockUserRepository.On("GetById", ctx, uint64(10)).Return(nil, errorMessage).Once()
		result, err := service.GetById(ctx, uint64(10))
		assert.NotNil(t, err)
		assert.Nil(t, result)
	})
}

func TestUserServiceImpl_GetAll(t *testing.T) {
	ctx := context.Background()
	//subtest success:
	t.Run("GetAll_success", func(t *testing.T) {
		user := getUserMock(uint64(13000), "user example", "example@email.com", "12345")

		mockListUser := make([]*ent.User, 0)
		mockListUser = append(mockListUser, &user)

		mockUserRepository.On("GetAll", ctx).Return(mockListUser, nil).Once()
		result, err := service.GetAll(ctx)
		assert.NoError(t, err)
		assert.NotNil(t, result)
	})

	//subtest failed:
	t.Run("GetAll_failed", func(t *testing.T) {
		errorMessage := errors.New("failed get all user")
		mockUserRepository.On("GetAll", ctx).Return(nil, errorMessage).Once()
		result, err := service.GetAll(ctx)
		assert.NotNil(t, err)
		assert.Nil(t, result)
	})

}

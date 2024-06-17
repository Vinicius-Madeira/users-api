package service

import (
	"github.com/Vinicius-Madeira/go-web-app/src/configuration/rest_err"
	"github.com/Vinicius-Madeira/go-web-app/src/model"
	"github.com/Vinicius-Madeira/go-web-app/src/tests/mocks"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/mock/gomock"
	"testing"
)

func TestUserDomainService_CreateUserServices(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mocks.NewMockUserRepository(ctrl)
	service := NewUserDomainService(repository)

	t.Run("when_user_already_exists_returns_error", func(t *testing.T) {
		id := primitive.NewObjectID().Hex()

		userDomain := model.NewUserDomain("test@test.com", "testing123$", "test", 30)
		userDomain.SetID(id)

		repository.EXPECT().FindUserByEmail(userDomain.GetEmail()).Return(userDomain, nil)

		user, err := service.CreateUserServices(userDomain)

		assert.Nil(t, user)
		assert.NotNil(t, err)
		assert.Equal(t, err.Message, "Email is already being used")
	})

	t.Run("when_user_is_not_registered_returns_error", func(t *testing.T) {
		id := primitive.NewObjectID().Hex()

		userDomain := model.NewUserDomain("test@test.com", "testing123$", "test", 30)
		userDomain.SetID(id)

		repository.EXPECT().FindUserByEmail(userDomain.GetEmail()).Return(nil, nil)
		repository.EXPECT().CreateUser(userDomain).Return(nil, rest_err.NewInternalServerError("Email is already being used"))

		user, err := service.CreateUserServices(userDomain)

		assert.Nil(t, user)
		assert.NotNil(t, err)
		assert.Equal(t, err.Message, "Email is already being used")
	})

	t.Run("when_user_is_not_registered_returns_success", func(t *testing.T) {
		id := primitive.NewObjectID().Hex()

		userDomain := model.NewUserDomain("test@test.com", "testing123$", "test", 30)
		userDomain.SetID(id)

		repository.EXPECT().FindUserByEmail(userDomain.GetEmail()).Return(nil, nil)
		repository.EXPECT().CreateUser(userDomain).Return(userDomain, nil)

		user, err := service.CreateUserServices(userDomain)

		assert.Nil(t, err)
		assert.EqualValues(t, user.GetID(), userDomain.GetID())
		assert.EqualValues(t, user.GetEmail(), userDomain.GetEmail())
		assert.EqualValues(t, user.GetName(), userDomain.GetName())
		assert.EqualValues(t, user.GetPassword(), userDomain.GetPassword())
		assert.EqualValues(t, user.GetAge(), userDomain.GetAge())
	})
}

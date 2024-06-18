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

func TestUserDomainService_UpdateUserServices(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mocks.NewMockUserRepository(ctrl)
	service := NewUserDomainService(repository)

	t.Run("when_sending_a_valid_user_and_userId_returns_success", func(t *testing.T) {

		id := primitive.NewObjectID().Hex()

		userDomain := model.NewUserDomain("test@test.com", "testing123$", "test", 30)
		userDomain.SetID(id)

		repository.EXPECT().UpdateUser(id, userDomain).Return(nil)

		err := service.UpdateUserServices(id, userDomain)

		assert.Nil(t, err)
	})

	t.Run("when_sending_a_valid_user_and_userId_returns_success", func(t *testing.T) {

		id := primitive.NewObjectID().Hex()

		userDomain := model.NewUserDomain("test@test.com", "testing123$", "test", 30)
		userDomain.SetID(id)

		repository.EXPECT().UpdateUser(id, userDomain).Return(rest_err.NewInternalServerError("error trying to update user"))

		err := service.UpdateUserServices(id, userDomain)

		assert.NotNil(t, err)
		assert.Equal(t, err.Message, "error trying to update user")
	})

}

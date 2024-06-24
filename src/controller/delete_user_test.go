package controller

import (
	"github.com/Vinicius-Madeira/go-web-app/src/configuration/rest_err"
	"github.com/Vinicius-Madeira/go-web-app/src/tests/mocks"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestUserControllerInterface_DeleteUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := mocks.NewMockUserDomainService(ctrl)
	controller := NewUserControllerInterface(service)

	t.Run("userId_is_invalid_returns_error", func(t *testing.T) {

		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)

		param := []gin.Param{
			{
				Key:   "userId",
				Value: "test",
			},
		}

		MakeRequest(context, param, url.Values{}, "DELETE", nil)
		controller.DeleteUser(context)

		assert.EqualValues(t, http.StatusBadRequest, recorder.Code)
	})

	t.Run("id_is_valid_service_returns_error", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)
		id := primitive.NewObjectID().Hex()

		param := []gin.Param{
			{
				Key:   "userId",
				Value: id,
			},
		}

		service.EXPECT().DeleteUserServices(id).Return(rest_err.NewInternalServerError("Error trying to delete user from database"))

		MakeRequest(context, param, url.Values{}, "DELETE", nil)
		controller.DeleteUser(context)

		assert.EqualValues(t, http.StatusInternalServerError, recorder.Code)
		assert.Contains(t, recorder.Body.String(), "Error trying to delete user from database")
	})

	t.Run("id_is_valid_service_returns_success", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)
		id := primitive.NewObjectID().Hex()

		param := []gin.Param{
			{
				Key:   "userId",
				Value: id,
			},
		}

		service.EXPECT().DeleteUserServices(id).Return(nil)

		MakeRequest(context, param, url.Values{}, "DELETE", nil)
		controller.DeleteUser(context)

		assert.EqualValues(t, http.StatusOK, recorder.Code)
	})
}

package controller

import (
	"encoding/json"
	"github.com/Vinicius-Madeira/go-web-app/src/configuration/rest_err"
	"github.com/Vinicius-Madeira/go-web-app/src/controller/model/request"
	"github.com/Vinicius-Madeira/go-web-app/src/model"
	"github.com/Vinicius-Madeira/go-web-app/src/tests/mocks"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestUserControllerInterface_UpdateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := mocks.NewMockUserDomainService(ctrl)
	controller := NewUserControllerInterface(service)

	t.Run("invalid_body_fields_returns_error", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)
		id := primitive.NewObjectID().Hex()

		param := []gin.Param{
			{
				Key:   "userId",
				Value: id,
			},
		}

		updateUserRequest := request.UserUpdateRequest{
			Name: "tst",
			Age:  20,
		}
		b, _ := json.Marshal(updateUserRequest)
		stringReader := io.NopCloser(strings.NewReader(string(b)))

		MakeRequest(context, param, url.Values{}, "PUT", stringReader)
		controller.UpdateUser(context)

		assert.EqualValues(t, http.StatusBadRequest, recorder.Code)
		assert.Contains(t, recorder.Body.String(), "Some fields are invalid")
	})

	t.Run("valid_body_but_invalid_userid_returns_error", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)
		id := "test"

		param := []gin.Param{
			{
				Key:   "userId",
				Value: id,
			},
		}

		updateUserRequest := request.UserUpdateRequest{
			Name: "test",
			Age:  40,
		}
		b, _ := json.Marshal(updateUserRequest)
		stringReader := io.NopCloser(strings.NewReader(string(b)))

		MakeRequest(context, param, url.Values{}, "PUT", stringReader)
		controller.UpdateUser(context)

		assert.EqualValues(t, http.StatusBadRequest, recorder.Code)
		assert.Contains(t, recorder.Body.String(), "Invalid userId, must be a hex value")
	})

	t.Run("valid_body_and_userid_but_service_returns_error", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)
		id := primitive.NewObjectID().Hex()

		param := []gin.Param{
			{
				Key:   "userId",
				Value: id,
			},
		}

		updateUserRequest := request.UserUpdateRequest{
			Name: "test",
			Age:  40,
		}

		domain := model.NewUserUpdateDomain(updateUserRequest.Name, updateUserRequest.Age)

		b, _ := json.Marshal(updateUserRequest)
		stringReader := io.NopCloser(strings.NewReader(string(b)))

		service.EXPECT().UpdateUserServices(id, domain).Return(rest_err.NewBadRequestError("Error trying to call repository"))

		MakeRequest(context, param, url.Values{}, "PUT", stringReader)
		controller.UpdateUser(context)

		assert.EqualValues(t, http.StatusBadRequest, recorder.Code)
		assert.Contains(t, recorder.Body.String(), "Error trying to call repository")
	})

	t.Run("valid_body_and_userid_returns_success", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)
		id := primitive.NewObjectID().Hex()

		param := []gin.Param{
			{
				Key:   "userId",
				Value: id,
			},
		}

		updateUserRequest := request.UserUpdateRequest{
			Name: "test",
			Age:  40,
		}

		domain := model.NewUserUpdateDomain(updateUserRequest.Name, updateUserRequest.Age)

		b, _ := json.Marshal(updateUserRequest)
		stringReader := io.NopCloser(strings.NewReader(string(b)))

		service.EXPECT().UpdateUserServices(id, domain).Return(nil)

		MakeRequest(context, param, url.Values{}, "PUT", stringReader)
		controller.UpdateUser(context)

		assert.EqualValues(t, http.StatusOK, recorder.Code)
	})
}

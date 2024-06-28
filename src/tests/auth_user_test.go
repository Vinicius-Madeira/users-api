package tests

import (
	"encoding/json"
	"fmt"
	"github.com/Vinicius-Madeira/go-web-app/src/controller/model/request"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"io"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestAuthUser(t *testing.T) {
	t.Run("user_and_password_is_not_valid", func(t *testing.T) {
		recorderCreateUser := httptest.NewRecorder()
		ctxCreateUser := GetTestGinContext(recorderCreateUser)

		recorderAuthUser := httptest.NewRecorder()
		ctxAuthUser := GetTestGinContext(recorderAuthUser)

		email := fmt.Sprintf("%d@test.com", rand.Int())
		password := fmt.Sprintf("%d$", rand.Int())

		userCreateRequest := request.UserRequest{
			Email:    email,
			Password: password,
			Name:     "testAuth",
			Age:      23,
		}
		bCreate, _ := json.Marshal(userCreateRequest)
		stringReaderCreate := io.NopCloser(strings.NewReader(string(bCreate)))

		MakeRequest(ctxCreateUser, []gin.Param{}, url.Values{}, "POST", stringReaderCreate)
		UserController.CreateUser(ctxCreateUser)

		userAuthRequest := request.UserAuth{
			Email:    "test@invalid.com",
			Password: "abc123$",
		}

		bAuth, _ := json.Marshal(userAuthRequest)
		stringReaderAuth := io.NopCloser(strings.NewReader(string(bAuth)))

		MakeRequest(ctxAuthUser, []gin.Param{}, url.Values{}, "POST", stringReaderAuth)
		UserController.AuthUser(ctxAuthUser)

		assert.EqualValues(t, http.StatusForbidden, recorderAuthUser.Result().StatusCode)
		assert.Empty(t, recorderAuthUser.Result().Header.Get("Authorization"))
	})

	t.Run("user_and_password_is_valid", func(t *testing.T) {
		recorderCreateUser := httptest.NewRecorder()
		ctxCreateUser := GetTestGinContext(recorderCreateUser)

		recorderAuthUser := httptest.NewRecorder()
		ctxAuthUser := GetTestGinContext(recorderAuthUser)

		email := fmt.Sprintf("%d@test.com", rand.Int())
		password := fmt.Sprintf("%d$", rand.Int())

		userCreateRequest := request.UserRequest{
			Email:    email,
			Password: password,
			Name:     "testAuth",
			Age:      23,
		}
		bCreate, _ := json.Marshal(userCreateRequest)
		stringReaderCreate := io.NopCloser(strings.NewReader(string(bCreate)))

		MakeRequest(ctxCreateUser, []gin.Param{}, url.Values{}, "POST", stringReaderCreate)
		UserController.CreateUser(ctxCreateUser)

		userAuthRequest := request.UserAuth{
			Email:    email,
			Password: password,
		}

		bAuth, _ := json.Marshal(userAuthRequest)
		stringReaderAuth := io.NopCloser(strings.NewReader(string(bAuth)))

		MakeRequest(ctxAuthUser, []gin.Param{}, url.Values{}, "POST", stringReaderAuth)
		UserController.AuthUser(ctxAuthUser)

		assert.EqualValues(t, http.StatusOK, recorderAuthUser.Result().StatusCode)
		assert.NotEmpty(t, recorderAuthUser.Result().Header.Get("Authorization"))

	})
}

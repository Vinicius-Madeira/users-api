package tests

import (
	"encoding/json"
	"github.com/Vinicius-Madeira/go-web-app/src/controller/model/request"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestCreateUser(t *testing.T) {
	t.Run("valid_request_body_returns_success", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		ctx := GetTestGinContext(recorder)

		userRequest := request.UserRequest{
			Email:    "viniciussmadeira@hotmail.com",
			Password: "testing123$",
			Name:     "testCreate",
			Age:      30,
		}
		b, _ := json.Marshal(userRequest)
		stringReader := io.NopCloser(strings.NewReader(string(b)))

		MakeRequest(ctx, []gin.Param{}, url.Values{}, "POST", stringReader)
		UserController.CreateUser(ctx)

		assert.EqualValues(t, http.StatusCreated, recorder.Code)
	})
}

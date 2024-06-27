package tests

import (
	"context"
	"fmt"
	"github.com/Vinicius-Madeira/go-web-app/src/controller"
	"github.com/Vinicius-Madeira/go-web-app/src/model/repository"
	"github.com/Vinicius-Madeira/go-web-app/src/model/service"
	"github.com/Vinicius-Madeira/go-web-app/src/tests/connection"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"
)

var (
	UserController controller.UserControllerInterface
	Database       *mongo.Database
)

func TestMain(m *testing.M) {
	err := os.Setenv("MONGODB_COLLECTION_NAME", "test_user")
	if err != nil {
		log.Fatalf(err.Error())
	}

	closeConnection := func() {}
	Database, closeConnection = connection.OpenConnection()

	repo := repository.NewUserRepository(Database)
	userService := service.NewUserDomainService(repo)
	UserController = controller.NewUserControllerInterface(userService)

	code := func() int {
		code := m.Run()
		fmt.Println("Clearing env...")
		os.Clearenv()
		closeConnection()
		return code
	}()

	os.Exit(code)
}

func TestFindUserByEmail(t *testing.T) {
	t.Run("user_not_found_returns_error", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		ctx := GetTestGinContext(recorder)

		param := []gin.Param{
			{
				Key:   "userEmail",
				Value: "test@test.com",
			},
		}

		MakeRequest(ctx, param, url.Values{}, "GET", nil)
		UserController.FindUserByEmail(ctx)

		assert.EqualValues(t, http.StatusNotFound, recorder.Code)
	})

	t.Run("user_found_returns_success", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		ctx := GetTestGinContext(recorder)
		id := primitive.NewObjectID().Hex()

		_, err := Database.
			Collection("test_user").
			InsertOne(context.Background(), bson.M{"_id": id, "name": t.Name(), "email": "test@test.com"})
		if err != nil {
			t.Fatal(err)
			return
		}

		param := []gin.Param{
			{
				Key:   "userEmail",
				Value: "test@test.com",
			},
		}

		MakeRequest(ctx, param, url.Values{}, "GET", nil)
		UserController.FindUserByEmail(ctx)

		assert.EqualValues(t, http.StatusOK, recorder.Code)
	})
}

func TestFindUserById(t *testing.T) {
	t.Run("userId_not_found_returns_error", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		ctx := GetTestGinContext(recorder)
		id := primitive.NewObjectID().Hex()

		param := []gin.Param{
			{
				Key:   "userId",
				Value: id,
			},
		}

		MakeRequest(ctx, param, url.Values{}, "GET", nil)
		UserController.FindUserById(ctx)

		assert.EqualValues(t, http.StatusNotFound, recorder.Code)
	})

	t.Run("userId_found_returns_success", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		ctx := GetTestGinContext(recorder)
		id := primitive.NewObjectID()

		_, err := Database.
			Collection("test_user").
			InsertOne(context.Background(), bson.M{"_id": id, "name": t.Name(), "email": "test@test.com"})
		if err != nil {
			t.Fatal(err)
			return
		}

		param := []gin.Param{
			{
				Key:   "userId",
				Value: id.Hex(),
			},
		}

		MakeRequest(ctx, param, url.Values{}, "GET", nil)
		UserController.FindUserById(ctx)

		assert.EqualValues(t, http.StatusOK, recorder.Code)
	})
}

func GetTestGinContext(recorder *httptest.ResponseRecorder) *gin.Context {
	gin.SetMode(gin.TestMode)

	ctx, _ := gin.CreateTestContext(recorder)
	ctx.Request = &http.Request{
		Header: make(http.Header),
		URL:    &url.URL{},
	}

	return ctx
}

func MakeRequest(
	c *gin.Context,
	param gin.Params,
	u url.Values,
	method string,
	body io.ReadCloser,
) {
	c.Request.Method = method
	c.Request.Header.Set("Content-Type", "application/json")
	c.Params = param

	c.Request.Body = body
	c.Request.URL.RawQuery = u.Encode()
}

package tests

import (
	"context"
	"encoding/json"
	"github.com/Vinicius-Madeira/go-web-app/src/controller/model/request"
	"github.com/Vinicius-Madeira/go-web-app/src/model/repository/entity"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestUpdateUser(t *testing.T) {
	t.Run("user_already_registered_with_this_email", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		ctx := GetTestGinContext(recorder)
		id := primitive.NewObjectID()

		_, err := Database.
			Collection("test_user").
			InsertOne(context.Background(), bson.M{"_id": id, "name": "OLD_NAME", "email": "testUpdate@test.com", "age": 30})
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

		userRequest := request.UserUpdateRequest{
			Name: "testUpdate",
			Age:  20,
		}

		b, _ := json.Marshal(userRequest)
		stringReader := io.NopCloser(strings.NewReader(string(b)))

		MakeRequest(ctx, param, url.Values{}, "PUT", stringReader)
		UserController.UpdateUser(ctx)

		assert.EqualValues(t, http.StatusOK, recorder.Result().StatusCode)

		var userEntity = entity.UserEntity{}

		filter := bson.M{"_id": id}
		_ = Database.
			Collection("test_user").
			FindOne(context.Background(), filter).
			Decode(&userEntity)

		assert.EqualValues(t, userRequest.Name, userEntity.Name)
		assert.EqualValues(t, userRequest.Age, userEntity.Age)
	})
}

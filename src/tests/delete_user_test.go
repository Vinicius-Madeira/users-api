package tests

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestDeleteUser(t *testing.T) {
	t.Run("user_found_returns_success", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		ctx := GetTestGinContext(recorder)
		id := primitive.NewObjectID()

		_, err := Database.
			Collection("test_user").
			InsertOne(context.Background(), bson.M{"_id": id, "name": t.Name(), "email": "testDelete@test.com"})
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

		MakeRequest(ctx, param, url.Values{}, "DELETE", nil)
		UserController.DeleteUser(ctx)

		assert.EqualValues(t, http.StatusNoContent, recorder.Code)

		filter := bson.M{"_id": id}
		result := Database.
			Collection("test_user").
			FindOne(context.Background(), filter)

		assert.NotNil(t, result.Err())
	})
}

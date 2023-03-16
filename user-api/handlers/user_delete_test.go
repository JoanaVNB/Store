package handlers_test

import (
	"BE-JoanaVidon/user-api/handlers"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestDeleteUseCase_Execute(t *testing.T) {

	controller := gomock.NewController(t)
	defer controller.Finish()

	wantStatusCode := 200
	wantBody := "deleted"

	mockRepository := handlers.NewMockDeleteHandler(controller)

		mockRepository.
			EXPECT().
			Execute("1").
			Return(nil)

		handler := handlers.NewDeleteHandler(mockRepository)

		r := gin.Default()
		r.DELETE("/users/:id", handler.Delete)

		req, err := http.NewRequest(http.MethodDelete,
			"/users/1", nil)
		if err != nil {
			t.Errorf(err.Error())
		}

		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)
		gotBody := w.Body.String()

		assert.EqualValues(t, wantStatusCode, w.Code)
		assert.EqualValues(t, wantBody, strings.Trim(gotBody, "\""))
}
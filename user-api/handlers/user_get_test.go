package handlers_test

import (
	"BE-JoanaVidon/user-api/domain"
	"BE-JoanaVidon/user-api/handlers"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGetUseCase_Execute(t *testing.T) {

	controller := gomock.NewController(t)
	defer controller.Finish()

	userMock := domain.User{
		ID: 						"1",
		Name:        "some name",
		CPF:         "some cpf",
		Email:       "some@gmail.com",
		PhoneNumber: "some phone",
	}

	wantStatusCode := 200
	wantBody := `{"id":"1","name":"some name","cpf":"some cpf","email":"some@gmail.com","phone_number":"some phone"}`

	mockRepository := handlers.NewMockGetHandler(controller)

		mockRepository.
			EXPECT().
			Execute("1", domain.User{}).
			Return(userMock, nil)

		handler := handlers.NewGetHandler(mockRepository)

		r := gin.Default()
		r.GET("/users/:id", handler.Get)

		req, err := http.NewRequest(http.MethodGet,
			"/users/1", nil)
		if err != nil {
			t.Errorf(err.Error())
		}

		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)
		gotBody := w.Body.String()

		assert.EqualValues(t, wantStatusCode, w.Code)
		assert.EqualValues(t, wantBody, gotBody)
}
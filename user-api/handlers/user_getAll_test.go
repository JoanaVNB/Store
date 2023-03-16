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

func TestGetAllUseCase_Execute(t *testing.T) {

	controller := gomock.NewController(t)
	defer controller.Finish()

	wantStatusCode := 200
	wantBody := `{"id":"1","name":"some name","cpf":"some cpf","email":"some@gmail.com","phone_number":"some phone"}`

	mockRepository := handlers.NewMockGetAllHandler(controller)

	mockRepository.
  EXPECT().
	Execute(gomock.Any()).
  Return([]domain.User{
    {ID: "1", Name: "some name", CPF: "some cpf", Email: "some@gmail.com", PhoneNumber: "some phone"},
  }, nil)



		handler := handlers.NewGetAllHandler(mockRepository)

		r := gin.Default()
		r.GET("/users", handler.GetAll)

		req, err := http.NewRequest(http.MethodGet,
			"/users", nil)
		if err != nil {
			t.Errorf(err.Error())
		}

		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)
		gotBody := w.Body.String()


		assert.EqualValues(t, wantStatusCode, w.Code)
		assert.EqualValues(t, wantBody, gotBody)
}
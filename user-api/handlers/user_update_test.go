package handlers_test

import (
	"BE-JoanaVidon/user-api/domain"
	"BE-JoanaVidon/user-api/handlers"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestUpdateUseCase_Execute(t *testing.T) {

	controller := gomock.NewController(t)
	defer controller.Finish()

	userMock := domain.User{
		ID: 						"",
		Name:        "some name",
		CPF:         "some cpf",
		Email:       "some@gmail.com",
		PhoneNumber: "phone",
	}

	wantStatusCode := 200
	wantBody := `{"id":"","name":"some name","cpf":"some cpf","email":"some@gmail.com","phone_number":"phone"}`

	mockRepository := handlers.NewMockUpdateHandler(controller)

		mockRepository.
			EXPECT().
			Execute(gomock.Eq("1"), gomock.Eq("phone"), gomock.Eq(userMock)).
			Return(userMock, nil)

	handler := handlers.NewUpdateHandler(mockRepository)

	r := gin.Default()
	r.PUT("/users/:id/:phone", handler.Update)

	
	payload := strings.NewReader(`{"name":"some name","cpf":"some cpf","email":"some@gmail.com","phone_number":"phone"}`)
//Essa estrutura representa os dados do usuário que serão atualizados. Lembre-se de que o phone_number no payload deve ser igual ao valor informado na URL (o :phone), pois é esse valor que será atualizado no usuário.
	req, err := http.NewRequest(http.MethodPut,
		"/users/1/phone", payload)
	if err != nil {
		t.Errorf(err.Error())
	}

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)
	gotRes := w.Body.String()

	assert.EqualValues(t, wantStatusCode, w.Code)
	assert.EqualValues(t, wantBody, gotRes)
}
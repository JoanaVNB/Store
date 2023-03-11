package service_test

import (
	"BE-JoanaVidon/user-api/domain"
	"BE-JoanaVidon/user-api/service"
	"testing"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGetAllUseCase_Execute(t *testing.T) {

	t.Run("Success repository", func(t *testing.T) {
		controller := gomock.NewController(t)
		defer controller.Finish()

		mockRepository := service.NewMockGetAllRepository(controller)
		mockRepository.
			EXPECT().
			GetAll(gomock.Any()).
			Return([]domain.User{{
				ID:						"1",
				Name:        "Joana Vidon",
				CPF:         "112.625.xxx-xx",
				Email:       "joanavidon@gmail.com",
				PhoneNumber: "(21)98108-xxxx",
			}}, nil)

		uc := service.NewGetAllUseCase(mockRepository)
		_, err := uc.Execute([]domain.User{})
		
		assert.Nil(t, err)
})
}
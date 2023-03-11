package service_test

import (
	"BE-JoanaVidon/user-api/domain"
	"BE-JoanaVidon/user-api/service"
	"testing"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestUpdateUseCase_Execute(t *testing.T) {

	t.Run("Success repository", func(t *testing.T) {
		controller := gomock.NewController(t)
		defer controller.Finish()

		mockRepository := service.NewMockUpdateRepository(controller)
		mockRepository.
			EXPECT().
			Update(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(domain.User{
				ID:						"1",
				Name:        "Joana Vidon",
				CPF:         "112.625.xxx-xx",
				Email:       "joanavidon@gmail.com",
				PhoneNumber: "(21)5555-xxxx",
			}, nil)

		uc := service.NewUpdateUseCase(mockRepository)
		_, err := uc.Execute("1", "(21)5555-xxxx", domain.User{
			Name:        "Joana Bloise",
			CPF:         "112.625.xxx-xx",
			Email:       "joanavidon@gmail.com",
			PhoneNumber: "(21)98108-xxxx",
		})

		assert.Nil(t, err)
})
}
package service_test

import (
	"BE-JoanaVidon/user-api/domain"
	"BE-JoanaVidon/user-api/service"
	"testing"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGetUseCase_Execute(t *testing.T) {

	t.Run("Success repository", func(t *testing.T) {
		controller := gomock.NewController(t)
		defer controller.Finish()
//Execute(string, domain.User) (domain.User, error)
		mockRepository := service.NewMockGetRepository(controller)
		mockRepository.
			EXPECT().
			Get(gomock.Any(), gomock.Any()).
			Return(domain.User{
				ID:						"1",
				Name:        "Joana Vidon",
				CPF:         "112.625.xxx-xx",
				Email:       "joanavidon@gmail.com",
				PhoneNumber: "(21)98108-xxxx",
			}, nil)

		uc := service.NewGetUseCase(mockRepository)
		_, err := uc.Execute("1", domain.User{})

		assert.Nil(t, err)
})
}
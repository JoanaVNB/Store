package service_test

import (
	"BE-JoanaVidon/user-api/domain"
	"BE-JoanaVidon/user-api/service"
	"testing"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestUpdateUseCase_Execute(t *testing.T) {

	controller := gomock.NewController(t)
	defer controller.Finish()

		mockRepository := service.NewMockUpdateRepository(controller)

		userMock := domain.User{
			Name:        "some name",
			CPF:         "some cpf",
			Email:       "some@gmail.com",
			PhoneNumber: "2721-4762",
		}

		mockRepository.
			EXPECT().
			Update("1", "2721-4762", domain.User{}).
			Return(userMock, nil)
	
		service := service.NewUpdateUseCase(mockRepository)
	
		user, err := service.Execute("1", "2721-4762", domain.User{})	

		assert.EqualValues(t, userMock, user)
		assert.Nil(t, err)
	}
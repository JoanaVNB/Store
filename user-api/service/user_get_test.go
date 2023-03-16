package service_test

import (
	"BE-JoanaVidon/user-api/domain"
	"BE-JoanaVidon/user-api/service"
	"testing"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGetUseCase_Execute(t *testing.T) {

	controller := gomock.NewController(t)
	defer controller.Finish()

		mockRepository := service.NewMockGetRepository(controller)

		userMock := domain.User{
			Name:        "some name",
			CPF:         "some cpf",
			Email:       "some@gmail.com",
			PhoneNumber: "some phone",
		}

		mockRepository.
			EXPECT().
			Get("1", domain.User{}).
			Return(userMock, nil)
	
		service := service.NewGetUseCase(mockRepository)
	
		user, err := service.Execute("1",domain.User{})	

		assert.EqualValues(t, userMock, user)
		assert.Nil(t, err)
	}
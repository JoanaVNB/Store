package service_test

import (
	"BE-JoanaVidon/user-api/domain"
	"BE-JoanaVidon/user-api/service"
	"testing"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGetAllUseCase_Execute(t *testing.T) {

	controller := gomock.NewController(t)
	defer controller.Finish()

		mockRepository := service.NewMockGetAllRepository(controller)

		userMock := []domain.User{{
			Name:        "some name",
			CPF:         "some cpf",
			Email:       "some@gmail.com",
			PhoneNumber: "some phone",
		}}

		mockRepository.
			EXPECT().
			GetAll([]domain.User{}).
			Return(userMock, nil)
	
		service := service.NewGetAllUseCase(mockRepository)
	
		user, err := service.Execute([]domain.User{})	

		assert.EqualValues(t, userMock, user)
		assert.Nil(t, err)
	}
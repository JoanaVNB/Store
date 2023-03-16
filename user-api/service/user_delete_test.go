package service_test

import (
	"BE-JoanaVidon/user-api/service"
	"testing"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestDeleteUseCase_Execute(t *testing.T) {

	controller := gomock.NewController(t)
	defer controller.Finish()

		mockRepository := service.NewMockDeleteRepository(controller)

		mockRepository.
			EXPECT().
			Delete("1").
			Return(nil)
	
		service := service.NewDeleteUseCase(mockRepository)
	
		err := service.Execute("1")	

		assert.Nil(t, err)
	}
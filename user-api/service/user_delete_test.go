package service_test

import (
	"BE-JoanaVidon/user-api/service"
	"testing"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestDeleteUseCase_Execute(t *testing.T) {

	t.Run("Success repository", func(t *testing.T) {
		controller := gomock.NewController(t)
		defer controller.Finish()

		mockRepository := service.NewMockDeleteRepository(controller)
		mockRepository.
			EXPECT().
			Delete(gomock.Any()).
			Return(nil)

		uc := service.NewDeleteUseCase(mockRepository)
		err := uc.Execute("1")

		assert.Nil(t, err)
})
}
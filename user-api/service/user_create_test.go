package service_test

import (
	"BE-JoanaVidon/user-api/domain"
	"BE-JoanaVidon/user-api/repository/inMemo"
	"BE-JoanaVidon/user-api/presenter"
	"BE-JoanaVidon/user-api/service"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCreateUseCase_Execute(t *testing.T) {

	controller := gomock.NewController(t)
	defer controller.Finish()

	type fields struct {
		createUC func() service.CreateUseCase
		created func() inMemo.InMemoRepository
	}
	tests := []struct {
		name          string
		fields         fields
		want				 presenter.PresentUser
	}{
		
	{
		name: "should successfully create a user",
		fields: fields{
			createUC: func() service.CreateUseCase {
				mockRepository := service.NewMockCreateUseCase(controller)

				mockRepository.
					EXPECT().
					Execute(domain.User{
						Name:        "some name",
						CPF:         "somecpf",
						Email:       "some@gmail.com",
						PhoneNumber: "somenumber",
					}).
					Return(presenter.PresentUser{
						ID:          "",
						Name:        "some name",
						CPF:         "somecpf",
						Email:       "some@gmail.com",
						PhoneNumber: "somenumber",
					}, nil)

				return mockRepository
			},
			created: func() inMemo.InMemoRepository {
				return *inMemo.NewInMemoRepository()
			},
		},
				want: presenter.PresentUser{
					ID:          "1",
					Name:        "Joana Vidon",
					CPF:         "112.625.xxx-xx",
					Email:       "joanavidon@gmail.com",
					PhoneNumber: "(21)98108-xxxx",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
				uc := service.NewCreateUseCase(tt.fields.created())
					created, err := uc.Execute(domain.User{
						Name:        "Joana Vidon",
						CPF:         "112.625.xxx-xx",
						Email:       "joanavidon@gmail.com",
						PhoneNumber: "(21)98108-xxxx",
					})
					created.ID = "1"
				
				assert.EqualValues(t, tt.want, created)
				assert.Nil(t, err)
		})
	}
}
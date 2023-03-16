package handlers_test

import (
	"BE-JoanaVidon/user-api/domain"
	"BE-JoanaVidon/user-api/presenter"
	"BE-JoanaVidon/user-api/handlers"
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCreateUseCase_Execute(t *testing.T) {

	controller := gomock.NewController(t)
	defer controller.Finish()

	type fields struct {
		createHandler func() handlers.CreateHandler
	}
	type args struct {
		requestBody    func() io.Reader
	}
	tests := []struct {
		name           string
		fields         fields
		args           args
		wantStatusCode int
		wantBody       string
	}{

		{
			name: "should fail because email is invalid",
			fields: fields{
				createHandler: func() handlers.CreateHandler { return nil },
			},
			args: args{
				requestBody: func() io.Reader {
					raw, _ := json.Marshal(presenter.PresentUser{
						Name:        "some name",
						CPF:         "some cpf",
						Email:       "some email",
						PhoneNumber: "some phone",
					})

					return bytes.NewReader(raw)
				},
			},
			wantStatusCode: 400,
			wantBody:       `{"errors":[{"field":"Email","message":"E-mail invalid"}]}`,
		},
		{
			name: "should successfully create a user",
			fields: fields{
				createHandler: func() handlers.CreateHandler {
					mockRepository := handlers.NewMockCreateHandler(controller)

					mockRepository.
						EXPECT().
						Execute(domain.User{
							ID:						"1",
							Name:        "some name",
							CPF:         "somecpf",
							Email:       "some@gmail.com",
							PhoneNumber: "somenumber",
						}).
						Return(domain.User{
							ID:						"1",
							Name:        "some name",
							CPF:         "somecpf",
							Email:       "some@gmail.com",
							PhoneNumber: "somenumber",
						}, nil)

					return mockRepository
				},
			},
			args: args{
				requestBody: func() io.Reader {
					raw, _ := json.Marshal(presenter.PresentUser{
					ID:							"1",
						Name:        "some name",
						CPF:         "somecpf",
						Email:       "some@gmail.com",
						PhoneNumber: "somenumber",
					})

					return bytes.NewReader(raw)
				},
			},
			wantStatusCode: 201,
			wantBody:       `{"id":"1","name":"some name","cpf":"somecpf","email":"some@gmail.com","phone_number":"somenumber"}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			handler := handlers.NewCreateHandler(tt.fields.createHandler())

			r := gin.Default()
			r.POST("/users", handler.Create)

			req, err := http.NewRequest(http.MethodPost,
				"/users",
				tt.args.requestBody())

			if err != nil {
				t.Errorf(err.Error())
			}

			w := httptest.NewRecorder()

			r.ServeHTTP(w, req)
			gotBody := w.Body.String()

			assert.EqualValues(t, tt.wantStatusCode, w.Code)
			assert.EqualValues(t, tt.wantBody, gotBody)
		})
	}
}

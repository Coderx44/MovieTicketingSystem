package booking_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Coderx44/MovieTicketingPortal/booking"
	"github.com/Coderx44/MovieTicketingPortal/booking/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type HandlerTestSuite struct {
	suite.Suite
	service *mocks.Service
}

func (suite *HandlerTestSuite) SetupTest() {
	suite.service = &mocks.Service{}
}

func (suite *HandlerTestSuite) TearDownTest() {
	t := suite.T()
	suite.service.AssertExpectations(t)
}

func TestHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(HandlerTestSuite))
}

func (suite *HandlerTestSuite) TestCreateUserHandler() {
	t := suite.T()

	type args struct {
		reqBody booking.NewUser
		service *mocks.Service
	}

	tests := []struct {
		name       string
		args       args
		bodyReader *strings.Reader
		resp       string
		statusCode int
		prepare    func(args, *mocks.Service)
	}{
		{
			name: "Success",
			args: args{
				reqBody: booking.NewUser{Name: "test1",
					Email:       "test1@gmail.com",
					Password:    "pass1",
					PhoneNumber: "1234567890",
					Role:        "user",
				},

				service: suite.service,
			},
			bodyReader: strings.NewReader(`{
				"name":"test1",
				"email" : "test1@gmail.com",
				"password":"pass1",
				"phone_number":"1234567890"
			}`),
			resp:       "1",
			statusCode: http.StatusOK,
			prepare: func(a args, s *mocks.Service) {
				suite.service.On("CreateNewUser", mock.Anything, a.reqBody).Return(uint(1), nil).Once()
			},
		},
		{
			name: "Failure (Acoount exists)",
			args: args{
				reqBody: booking.NewUser{Name: "test1",
					Email:       "test1@gmail.com",
					Password:    "pass1",
					PhoneNumber: "1234567890",
					Role:        "user",
				},

				service: suite.service,
			},
			bodyReader: strings.NewReader(`{
				"name":"test1",
				"email" : "test1@gmail.com",
				"password":"pass1",
				"phone_number":"1234567890"
			}`),
			resp:       "Err: User already exits for given email",
			statusCode: http.StatusOK,
			prepare: func(a args, s *mocks.Service) {
				suite.service.On("CreateNewUser", mock.Anything, a.reqBody).Return(uint(0), errors.New("account exists for the given email")).Once()

			},
		},
		{
			name: "Failure (Empty email)",
			args: args{
				reqBody: booking.NewUser{Name: "test1",
					Email:       "",
					Password:    "pass1",
					PhoneNumber: "1234567890",
					Role:        "user",
				},

				service: suite.service,
			},
			bodyReader: strings.NewReader(`{
				"name":"test1",
				"email" : "",
				"password":"pass1",
				"phone_number":"1234567890"
			}`),
			resp:       "Provide the required parameters",
			statusCode: http.StatusBadRequest,
			prepare: func(a args, s *mocks.Service) {
			},
		},
		{
			name: "Failure (Invalid email)",
			args: args{
				reqBody: booking.NewUser{Name: "test1",
					Email:       "test1gmail.com",
					Password:    "pass1",
					PhoneNumber: "1234567890",
					Role:        "user",
				},

				service: suite.service,
			},
			bodyReader: strings.NewReader(`{
				"name":"test1",
				"email" : "test1gmail.com",
				"password":"pass1",
				"phone_number":"1234567890"
			}`),
			resp:       "Err: Invalid email address",
			statusCode: http.StatusBadRequest,
			prepare: func(a args, s *mocks.Service) {
			},
		},
		{
			name: "Failure (Invalid contact)",
			args: args{
				reqBody: booking.NewUser{Name: "test1",
					Email:       "test1@gmail.com",
					Password:    "pass1",
					PhoneNumber: "123456780",
					Role:        "user",
				},

				service: suite.service,
			},
			bodyReader: strings.NewReader(`{
				"name":"test1",
				"email" : "test1@gmail.com",
				"password":"pass1",
				"phone_number":"123456780"
			}`),
			resp:       "Err: Phone must contain 10 digits",
			statusCode: http.StatusBadRequest,
			prepare: func(a args, s *mocks.Service) {
			},
		},
		{
			name: "Failure (Unkown err)",
			args: args{
				reqBody: booking.NewUser{Name: "test1",
					Email:       "test1@gmail.com",
					Password:    "pass1",
					PhoneNumber: "1234567890",
					Role:        "user",
				},

				service: suite.service,
			},
			bodyReader: strings.NewReader(`{
				"name":"test1",
				"email" : "test1@gmail.com",
				"password":"pass1",
				"phone_number":"1234567890"
			}`),
			resp:       "Err - Internal Server Error - Failure creating user account",
			statusCode: http.StatusInternalServerError,
			prepare: func(a args, s *mocks.Service) {
				suite.service.On("CreateNewUser", mock.Anything, a.reqBody).Return(uint(0), errors.New("err creating user")).Once()

			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.prepare(tt.args, suite.service)
			r := httptest.NewRequest(http.MethodGet, "/create/user", tt.bodyReader)
			w := httptest.NewRecorder()
			// tt.args.service.On("CreateNewUser", mock.Anything, tt.args.reqBody).Return(uint(0), errors.New("account exists for the given email")).Once()
			router := booking.CreateNewUser(suite.service)
			router.ServeHTTP(w, r)
			assert.Equal(t, tt.statusCode, w.Code)
			assert.Equal(t, tt.resp, w.Body.String())
		})
	}

}

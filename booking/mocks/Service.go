// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	context "context"

	booking "github.com/Coderx44/MovieTicketingPortal/booking"

	db "github.com/Coderx44/MovieTicketingPortal/db"

	mock "github.com/stretchr/testify/mock"
)

// Service is an autogenerated mock type for the Service type
type Service struct {
	mock.Mock
}

// AddBookingsBySeatId provides a mock function with given fields: ctx, seats, email, show_id
func (_m *Service) AddBookingsBySeatId(ctx context.Context, seats []int, email string, show_id int) (booking.Invoice, error) {
	ret := _m.Called(ctx, seats, email, show_id)

	var r0 booking.Invoice
	if rf, ok := ret.Get(0).(func(context.Context, []int, string, int) booking.Invoice); ok {
		r0 = rf(ctx, seats, email, show_id)
	} else {
		r0 = ret.Get(0).(booking.Invoice)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, []int, string, int) error); ok {
		r1 = rf(ctx, seats, email, show_id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AddMovie provides a mock function with given fields: ctx, m
func (_m *Service) AddMovie(ctx context.Context, m booking.NewMovie) (uint, error) {
	ret := _m.Called(ctx, m)

	var r0 uint
	if rf, ok := ret.Get(0).(func(context.Context, booking.NewMovie) uint); ok {
		r0 = rf(ctx, m)
	} else {
		r0 = ret.Get(0).(uint)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, booking.NewMovie) error); ok {
		r1 = rf(ctx, m)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AddMultiplex provides a mock function with given fields: ctx, m
func (_m *Service) AddMultiplex(ctx context.Context, m booking.NewMultiplex) (uint, error) {
	ret := _m.Called(ctx, m)

	var r0 uint
	if rf, ok := ret.Get(0).(func(context.Context, booking.NewMultiplex) uint); ok {
		r0 = rf(ctx, m)
	} else {
		r0 = ret.Get(0).(uint)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, booking.NewMultiplex) error); ok {
		r1 = rf(ctx, m)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AddScreen provides a mock function with given fields: ctx, s
func (_m *Service) AddScreen(ctx context.Context, s booking.NewScreen) (uint, error) {
	ret := _m.Called(ctx, s)

	var r0 uint
	if rf, ok := ret.Get(0).(func(context.Context, booking.NewScreen) uint); ok {
		r0 = rf(ctx, s)
	} else {
		r0 = ret.Get(0).(uint)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, booking.NewScreen) error); ok {
		r1 = rf(ctx, s)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AddShow provides a mock function with given fields: ctx, s
func (_m *Service) AddShow(ctx context.Context, s booking.NewShow) (uint, error) {
	ret := _m.Called(ctx, s)

	var r0 uint
	if rf, ok := ret.Get(0).(func(context.Context, booking.NewShow) uint); ok {
		r0 = rf(ctx, s)
	} else {
		r0 = ret.Get(0).(uint)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, booking.NewShow) error); ok {
		r1 = rf(ctx, s)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CancelBooking provides a mock function with given fields: ctx, id
func (_m *Service) CancelBooking(ctx context.Context, id int) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateNewUser provides a mock function with given fields: ctx, u
func (_m *Service) CreateNewUser(ctx context.Context, u booking.NewUser) (uint, error) {
	ret := _m.Called(ctx, u)

	var r0 uint
	if rf, ok := ret.Get(0).(func(context.Context, booking.NewUser) uint); ok {
		r0 = rf(ctx, u)
	} else {
		r0 = ret.Get(0).(uint)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, booking.NewUser) error); ok {
		r1 = rf(ctx, u)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllBookings provides a mock function with given fields: ctx, email
func (_m *Service) GetAllBookings(ctx context.Context, email string) ([]db.Booking, error) {
	ret := _m.Called(ctx, email)

	var r0 []db.Booking
	if rf, ok := ret.Get(0).(func(context.Context, string) []db.Booking); ok {
		r0 = rf(ctx, email)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]db.Booking)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllMultiplexesByCity provides a mock function with given fields: ctx, city
func (_m *Service) GetAllMultiplexesByCity(ctx context.Context, city string) ([]booking.NewMultiplex, error) {
	ret := _m.Called(ctx, city)

	var r0 []booking.NewMultiplex
	if rf, ok := ret.Get(0).(func(context.Context, string) []booking.NewMultiplex); ok {
		r0 = rf(ctx, city)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]booking.NewMultiplex)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, city)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllSeatsByShowID provides a mock function with given fields: ctx, show_id
func (_m *Service) GetAllSeatsByShowID(ctx context.Context, show_id int) (map[int][]booking.Seats, error) {
	ret := _m.Called(ctx, show_id)

	var r0 map[int][]booking.Seats
	if rf, ok := ret.Get(0).(func(context.Context, int) map[int][]booking.Seats); ok {
		r0 = rf(ctx, show_id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[int][]booking.Seats)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, show_id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllShowsByDateAndMultiplexId provides a mock function with given fields: ctx, date, multiplex_id
func (_m *Service) GetAllShowsByDateAndMultiplexId(ctx context.Context, date string, multiplex_id int) (map[string][]booking.MultiplexShow, error) {
	ret := _m.Called(ctx, date, multiplex_id)

	var r0 map[string][]booking.MultiplexShow
	if rf, ok := ret.Get(0).(func(context.Context, string, int) map[string][]booking.MultiplexShow); ok {
		r0 = rf(ctx, date, multiplex_id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string][]booking.MultiplexShow)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, int) error); ok {
		r1 = rf(ctx, date, multiplex_id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllShowsByMovieAndDate provides a mock function with given fields: ctx, date, title, city
func (_m *Service) GetAllShowsByMovieAndDate(ctx context.Context, date string, title string, city string) (map[string][]booking.MultiplexShow, error) {
	ret := _m.Called(ctx, date, title, city)

	var r0 map[string][]booking.MultiplexShow
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string) map[string][]booking.MultiplexShow); ok {
		r0 = rf(ctx, date, title, city)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string][]booking.MultiplexShow)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, string) error); ok {
		r1 = rf(ctx, date, title, city)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetMovieByTitle provides a mock function with given fields: ctx, title
func (_m *Service) GetMovieByTitle(ctx context.Context, title string) (booking.NewMovie, error) {
	ret := _m.Called(ctx, title)

	var r0 booking.NewMovie
	if rf, ok := ret.Get(0).(func(context.Context, string) booking.NewMovie); ok {
		r0 = rf(ctx, title)
	} else {
		r0 = ret.Get(0).(booking.NewMovie)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, title)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUpcomingMovies provides a mock function with given fields: ctx, date
func (_m *Service) GetUpcomingMovies(ctx context.Context, date string) ([]booking.NewMovie, error) {
	ret := _m.Called(ctx, date)

	var r0 []booking.NewMovie
	if rf, ok := ret.Get(0).(func(context.Context, string) []booking.NewMovie); ok {
		r0 = rf(ctx, date)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]booking.NewMovie)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, date)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Login provides a mock function with given fields: ctx, authU
func (_m *Service) Login(ctx context.Context, authU booking.Authentication) (string, error) {
	ret := _m.Called(ctx, authU)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, booking.Authentication) string); ok {
		r0 = rf(ctx, authU)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, booking.Authentication) error); ok {
		r1 = rf(ctx, authU)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewService interface {
	mock.TestingT
	Cleanup(func())
}

// NewService creates a new instance of Service. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewService(t mockConstructorTestingTNewService) *Service {
	mock := &Service{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
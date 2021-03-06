// Code generated by mockery v1.0.0. DO NOT EDIT.

// Regenerate this file using `make mocks`.

package mocks

import io "io"
import mock "github.com/stretchr/testify/mock"

// Server is an autogenerated mock type for the Server type
type Server struct {
	mock.Mock
}

// Convert provides a mock function with given fields: name, extension, file
func (_m *Server) Convert(name string, extension string, file io.Reader) (io.ReadCloser, error) {
	ret := _m.Called(name, extension, file)

	var r0 io.ReadCloser
	if rf, ok := ret.Get(0).(func(string, string, io.Reader) io.ReadCloser); ok {
		r0 = rf(name, extension, file)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(io.ReadCloser)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, io.Reader) error); ok {
		r1 = rf(name, extension, file)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Status provides a mock function with given fields:
func (_m *Server) Status() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

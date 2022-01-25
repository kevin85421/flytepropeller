// Code generated by mockery v1.0.1. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	protoiface "google.golang.org/protobuf/runtime/protoiface"
)

// EventSink is an autogenerated mock type for the EventSink type
type EventSink struct {
	mock.Mock
}

type EventSink_Close struct {
	*mock.Call
}

func (_m EventSink_Close) Return(_a0 error) *EventSink_Close {
	return &EventSink_Close{Call: _m.Call.Return(_a0)}
}

func (_m *EventSink) OnClose() *EventSink_Close {
	c := _m.On("Close")
	return &EventSink_Close{Call: c}
}

func (_m *EventSink) OnCloseMatch(matchers ...interface{}) *EventSink_Close {
	c := _m.On("Close", matchers...)
	return &EventSink_Close{Call: c}
}

// Close provides a mock function with given fields:
func (_m *EventSink) Close() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type EventSink_Sink struct {
	*mock.Call
}

func (_m EventSink_Sink) Return(_a0 error) *EventSink_Sink {
	return &EventSink_Sink{Call: _m.Call.Return(_a0)}
}

func (_m *EventSink) OnSink(ctx context.Context, message protoiface.MessageV1) *EventSink_Sink {
	c := _m.On("Sink", ctx, message)
	return &EventSink_Sink{Call: c}
}

func (_m *EventSink) OnSinkMatch(matchers ...interface{}) *EventSink_Sink {
	c := _m.On("Sink", matchers...)
	return &EventSink_Sink{Call: c}
}

// Sink provides a mock function with given fields: ctx, message
func (_m *EventSink) Sink(ctx context.Context, message protoiface.MessageV1) error {
	ret := _m.Called(ctx, message)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, protoiface.MessageV1) error); ok {
		r0 = rf(ctx, message)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
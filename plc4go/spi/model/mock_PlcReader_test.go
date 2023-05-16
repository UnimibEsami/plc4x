/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

// Code generated by mockery v2.27.1. DO NOT EDIT.

package model

import (
	context "context"

	apimodel "github.com/apache/plc4x/plc4go/pkg/api/model"

	mock "github.com/stretchr/testify/mock"
)

// MockPlcReader is an autogenerated mock type for the PlcReader type
type MockPlcReader struct {
	mock.Mock
}

type MockPlcReader_Expecter struct {
	mock *mock.Mock
}

func (_m *MockPlcReader) EXPECT() *MockPlcReader_Expecter {
	return &MockPlcReader_Expecter{mock: &_m.Mock}
}

// Read provides a mock function with given fields: ctx, readRequest
func (_m *MockPlcReader) Read(ctx context.Context, readRequest apimodel.PlcReadRequest) <-chan apimodel.PlcReadRequestResult {
	ret := _m.Called(ctx, readRequest)

	var r0 <-chan apimodel.PlcReadRequestResult
	if rf, ok := ret.Get(0).(func(context.Context, apimodel.PlcReadRequest) <-chan apimodel.PlcReadRequestResult); ok {
		r0 = rf(ctx, readRequest)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan apimodel.PlcReadRequestResult)
		}
	}

	return r0
}

// MockPlcReader_Read_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Read'
type MockPlcReader_Read_Call struct {
	*mock.Call
}

// Read is a helper method to define mock.On call
//   - ctx context.Context
//   - readRequest apimodel.PlcReadRequest
func (_e *MockPlcReader_Expecter) Read(ctx interface{}, readRequest interface{}) *MockPlcReader_Read_Call {
	return &MockPlcReader_Read_Call{Call: _e.mock.On("Read", ctx, readRequest)}
}

func (_c *MockPlcReader_Read_Call) Run(run func(ctx context.Context, readRequest apimodel.PlcReadRequest)) *MockPlcReader_Read_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(apimodel.PlcReadRequest))
	})
	return _c
}

func (_c *MockPlcReader_Read_Call) Return(_a0 <-chan apimodel.PlcReadRequestResult) *MockPlcReader_Read_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPlcReader_Read_Call) RunAndReturn(run func(context.Context, apimodel.PlcReadRequest) <-chan apimodel.PlcReadRequestResult) *MockPlcReader_Read_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewMockPlcReader interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockPlcReader creates a new instance of MockPlcReader. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockPlcReader(t mockConstructorTestingTNewMockPlcReader) *MockPlcReader {
	mock := &MockPlcReader{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

/*
Copyright 2023-2024 IONOS Cloud.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
// Code generated by mockery. DO NOT EDIT.

package proxmoxtest

import (
	context "context"

	go_proxmox "github.com/luthermonson/go-proxmox"
	mock "github.com/stretchr/testify/mock"

	proxmox "github.com/ionos-cloud/cluster-api-provider-proxmox/pkg/proxmox"
)

// MockClient is an autogenerated mock type for the Client type
type MockClient struct {
	mock.Mock
}

type MockClient_Expecter struct {
	mock *mock.Mock
}

func (_m *MockClient) EXPECT() *MockClient_Expecter {
	return &MockClient_Expecter{mock: &_m.Mock}
}

// CloneVM provides a mock function with given fields: ctx, templateID, clone
func (_m *MockClient) CloneVM(ctx context.Context, templateID int, clone proxmox.VMCloneRequest) (proxmox.VMCloneResponse, error) {
	ret := _m.Called(ctx, templateID, clone)

	var r0 proxmox.VMCloneResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int, proxmox.VMCloneRequest) (proxmox.VMCloneResponse, error)); ok {
		return rf(ctx, templateID, clone)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int, proxmox.VMCloneRequest) proxmox.VMCloneResponse); ok {
		r0 = rf(ctx, templateID, clone)
	} else {
		r0 = ret.Get(0).(proxmox.VMCloneResponse)
	}

	if rf, ok := ret.Get(1).(func(context.Context, int, proxmox.VMCloneRequest) error); ok {
		r1 = rf(ctx, templateID, clone)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockClient_CloneVM_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CloneVM'
type MockClient_CloneVM_Call struct {
	*mock.Call
}

// CloneVM is a helper method to define mock.On call
//   - ctx context.Context
//   - templateID int
//   - clone proxmox.VMCloneRequest
func (_e *MockClient_Expecter) CloneVM(ctx interface{}, templateID interface{}, clone interface{}) *MockClient_CloneVM_Call {
	return &MockClient_CloneVM_Call{Call: _e.mock.On("CloneVM", ctx, templateID, clone)}
}

func (_c *MockClient_CloneVM_Call) Run(run func(ctx context.Context, templateID int, clone proxmox.VMCloneRequest)) *MockClient_CloneVM_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int), args[2].(proxmox.VMCloneRequest))
	})
	return _c
}

func (_c *MockClient_CloneVM_Call) Return(_a0 proxmox.VMCloneResponse, _a1 error) *MockClient_CloneVM_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockClient_CloneVM_Call) RunAndReturn(run func(context.Context, int, proxmox.VMCloneRequest) (proxmox.VMCloneResponse, error)) *MockClient_CloneVM_Call {
	_c.Call.Return(run)
	return _c
}

// CloudInitStatus provides a mock function with given fields: ctx, vm
func (_m *MockClient) CloudInitStatus(ctx context.Context, vm *go_proxmox.VirtualMachine) (bool, error) {
	ret := _m.Called(ctx, vm)

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *go_proxmox.VirtualMachine) (bool, error)); ok {
		return rf(ctx, vm)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *go_proxmox.VirtualMachine) bool); ok {
		r0 = rf(ctx, vm)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(context.Context, *go_proxmox.VirtualMachine) error); ok {
		r1 = rf(ctx, vm)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockClient_CloudInitStatus_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CloudInitStatus'
type MockClient_CloudInitStatus_Call struct {
	*mock.Call
}

// CloudInitStatus is a helper method to define mock.On call
//   - ctx context.Context
//   - vm *go_proxmox.VirtualMachine
func (_e *MockClient_Expecter) CloudInitStatus(ctx interface{}, vm interface{}) *MockClient_CloudInitStatus_Call {
	return &MockClient_CloudInitStatus_Call{Call: _e.mock.On("CloudInitStatus", ctx, vm)}
}

func (_c *MockClient_CloudInitStatus_Call) Run(run func(ctx context.Context, vm *go_proxmox.VirtualMachine)) *MockClient_CloudInitStatus_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*go_proxmox.VirtualMachine))
	})
	return _c
}

func (_c *MockClient_CloudInitStatus_Call) Return(_a0 bool, _a1 error) *MockClient_CloudInitStatus_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockClient_CloudInitStatus_Call) RunAndReturn(run func(context.Context, *go_proxmox.VirtualMachine) (bool, error)) *MockClient_CloudInitStatus_Call {
	_c.Call.Return(run)
	return _c
}

// ConfigureVM provides a mock function with given fields: ctx, vm, options
func (_m *MockClient) ConfigureVM(ctx context.Context, vm *go_proxmox.VirtualMachine, options ...go_proxmox.VirtualMachineOption) (*go_proxmox.Task, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, vm)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *go_proxmox.Task
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *go_proxmox.VirtualMachine, ...go_proxmox.VirtualMachineOption) (*go_proxmox.Task, error)); ok {
		return rf(ctx, vm, options...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *go_proxmox.VirtualMachine, ...go_proxmox.VirtualMachineOption) *go_proxmox.Task); ok {
		r0 = rf(ctx, vm, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*go_proxmox.Task)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *go_proxmox.VirtualMachine, ...go_proxmox.VirtualMachineOption) error); ok {
		r1 = rf(ctx, vm, options...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockClient_ConfigureVM_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ConfigureVM'
type MockClient_ConfigureVM_Call struct {
	*mock.Call
}

// ConfigureVM is a helper method to define mock.On call
//   - ctx context.Context
//   - vm *go_proxmox.VirtualMachine
//   - options ...go_proxmox.VirtualMachineOption
func (_e *MockClient_Expecter) ConfigureVM(ctx interface{}, vm interface{}, options ...interface{}) *MockClient_ConfigureVM_Call {
	return &MockClient_ConfigureVM_Call{Call: _e.mock.On("ConfigureVM",
		append([]interface{}{ctx, vm}, options...)...)}
}

func (_c *MockClient_ConfigureVM_Call) Run(run func(ctx context.Context, vm *go_proxmox.VirtualMachine, options ...go_proxmox.VirtualMachineOption)) *MockClient_ConfigureVM_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]go_proxmox.VirtualMachineOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(go_proxmox.VirtualMachineOption)
			}
		}
		run(args[0].(context.Context), args[1].(*go_proxmox.VirtualMachine), variadicArgs...)
	})
	return _c
}

func (_c *MockClient_ConfigureVM_Call) Return(_a0 *go_proxmox.Task, _a1 error) *MockClient_ConfigureVM_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockClient_ConfigureVM_Call) RunAndReturn(run func(context.Context, *go_proxmox.VirtualMachine, ...go_proxmox.VirtualMachineOption) (*go_proxmox.Task, error)) *MockClient_ConfigureVM_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteVM provides a mock function with given fields: ctx, nodeName, vmID
func (_m *MockClient) DeleteVM(ctx context.Context, nodeName string, vmID int64) (*go_proxmox.Task, error) {
	ret := _m.Called(ctx, nodeName, vmID)

	var r0 *go_proxmox.Task
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, int64) (*go_proxmox.Task, error)); ok {
		return rf(ctx, nodeName, vmID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, int64) *go_proxmox.Task); ok {
		r0 = rf(ctx, nodeName, vmID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*go_proxmox.Task)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, int64) error); ok {
		r1 = rf(ctx, nodeName, vmID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockClient_DeleteVM_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteVM'
type MockClient_DeleteVM_Call struct {
	*mock.Call
}

// DeleteVM is a helper method to define mock.On call
//   - ctx context.Context
//   - nodeName string
//   - vmID int64
func (_e *MockClient_Expecter) DeleteVM(ctx interface{}, nodeName interface{}, vmID interface{}) *MockClient_DeleteVM_Call {
	return &MockClient_DeleteVM_Call{Call: _e.mock.On("DeleteVM", ctx, nodeName, vmID)}
}

func (_c *MockClient_DeleteVM_Call) Run(run func(ctx context.Context, nodeName string, vmID int64)) *MockClient_DeleteVM_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(int64))
	})
	return _c
}

func (_c *MockClient_DeleteVM_Call) Return(_a0 *go_proxmox.Task, _a1 error) *MockClient_DeleteVM_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockClient_DeleteVM_Call) RunAndReturn(run func(context.Context, string, int64) (*go_proxmox.Task, error)) *MockClient_DeleteVM_Call {
	_c.Call.Return(run)
	return _c
}

// FindVMResource provides a mock function with given fields: ctx, vmID
func (_m *MockClient) FindVMResource(ctx context.Context, vmID uint64) (*go_proxmox.ClusterResource, error) {
	ret := _m.Called(ctx, vmID)

	var r0 *go_proxmox.ClusterResource
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64) (*go_proxmox.ClusterResource, error)); ok {
		return rf(ctx, vmID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint64) *go_proxmox.ClusterResource); ok {
		r0 = rf(ctx, vmID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*go_proxmox.ClusterResource)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint64) error); ok {
		r1 = rf(ctx, vmID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockClient_FindVMResource_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindVMResource'
type MockClient_FindVMResource_Call struct {
	*mock.Call
}

// FindVMResource is a helper method to define mock.On call
//   - ctx context.Context
//   - vmID uint64
func (_e *MockClient_Expecter) FindVMResource(ctx interface{}, vmID interface{}) *MockClient_FindVMResource_Call {
	return &MockClient_FindVMResource_Call{Call: _e.mock.On("FindVMResource", ctx, vmID)}
}

func (_c *MockClient_FindVMResource_Call) Run(run func(ctx context.Context, vmID uint64)) *MockClient_FindVMResource_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint64))
	})
	return _c
}

func (_c *MockClient_FindVMResource_Call) Return(_a0 *go_proxmox.ClusterResource, _a1 error) *MockClient_FindVMResource_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockClient_FindVMResource_Call) RunAndReturn(run func(context.Context, uint64) (*go_proxmox.ClusterResource, error)) *MockClient_FindVMResource_Call {
	_c.Call.Return(run)
	return _c
}

// GetReservableMemoryBytes provides a mock function with given fields: ctx, nodeName, nodeMemoryAdjustment
func (_m *MockClient) GetReservableMemoryBytes(ctx context.Context, nodeName string, nodeMemoryAdjustment uint64) (uint64, error) {
	ret := _m.Called(ctx, nodeName, nodeMemoryAdjustment)

	var r0 uint64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, uint64) (uint64, error)); ok {
		return rf(ctx, nodeName, nodeMemoryAdjustment)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, uint64) uint64); ok {
		r0 = rf(ctx, nodeName, nodeMemoryAdjustment)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, uint64) error); ok {
		r1 = rf(ctx, nodeName, nodeMemoryAdjustment)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockClient_GetReservableMemoryBytes_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetReservableMemoryBytes'
type MockClient_GetReservableMemoryBytes_Call struct {
	*mock.Call
}

// GetReservableMemoryBytes is a helper method to define mock.On call
//   - ctx context.Context
//   - nodeName string
//   - nodeMemoryAdjustment uint64
func (_e *MockClient_Expecter) GetReservableMemoryBytes(ctx interface{}, nodeName interface{}, nodeMemoryAdjustment interface{}) *MockClient_GetReservableMemoryBytes_Call {
	return &MockClient_GetReservableMemoryBytes_Call{Call: _e.mock.On("GetReservableMemoryBytes", ctx, nodeName, nodeMemoryAdjustment)}
}

func (_c *MockClient_GetReservableMemoryBytes_Call) Run(run func(ctx context.Context, nodeName string, nodeMemoryAdjustment uint64)) *MockClient_GetReservableMemoryBytes_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(uint64))
	})
	return _c
}

func (_c *MockClient_GetReservableMemoryBytes_Call) Return(_a0 uint64, _a1 error) *MockClient_GetReservableMemoryBytes_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockClient_GetReservableMemoryBytes_Call) RunAndReturn(run func(context.Context, string, uint64) (uint64, error)) *MockClient_GetReservableMemoryBytes_Call {
	_c.Call.Return(run)
	return _c
}

// GetTask provides a mock function with given fields: ctx, upID
func (_m *MockClient) GetTask(ctx context.Context, upID string) (*go_proxmox.Task, error) {
	ret := _m.Called(ctx, upID)

	var r0 *go_proxmox.Task
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*go_proxmox.Task, error)); ok {
		return rf(ctx, upID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *go_proxmox.Task); ok {
		r0 = rf(ctx, upID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*go_proxmox.Task)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, upID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockClient_GetTask_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetTask'
type MockClient_GetTask_Call struct {
	*mock.Call
}

// GetTask is a helper method to define mock.On call
//   - ctx context.Context
//   - upID string
func (_e *MockClient_Expecter) GetTask(ctx interface{}, upID interface{}) *MockClient_GetTask_Call {
	return &MockClient_GetTask_Call{Call: _e.mock.On("GetTask", ctx, upID)}
}

func (_c *MockClient_GetTask_Call) Run(run func(ctx context.Context, upID string)) *MockClient_GetTask_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockClient_GetTask_Call) Return(_a0 *go_proxmox.Task, _a1 error) *MockClient_GetTask_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockClient_GetTask_Call) RunAndReturn(run func(context.Context, string) (*go_proxmox.Task, error)) *MockClient_GetTask_Call {
	_c.Call.Return(run)
	return _c
}

// GetVM provides a mock function with given fields: ctx, nodeName, vmID
func (_m *MockClient) GetVM(ctx context.Context, nodeName string, vmID int64) (*go_proxmox.VirtualMachine, error) {
	ret := _m.Called(ctx, nodeName, vmID)

	var r0 *go_proxmox.VirtualMachine
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, int64) (*go_proxmox.VirtualMachine, error)); ok {
		return rf(ctx, nodeName, vmID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, int64) *go_proxmox.VirtualMachine); ok {
		r0 = rf(ctx, nodeName, vmID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*go_proxmox.VirtualMachine)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, int64) error); ok {
		r1 = rf(ctx, nodeName, vmID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockClient_GetVM_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetVM'
type MockClient_GetVM_Call struct {
	*mock.Call
}

// GetVM is a helper method to define mock.On call
//   - ctx context.Context
//   - nodeName string
//   - vmID int64
func (_e *MockClient_Expecter) GetVM(ctx interface{}, nodeName interface{}, vmID interface{}) *MockClient_GetVM_Call {
	return &MockClient_GetVM_Call{Call: _e.mock.On("GetVM", ctx, nodeName, vmID)}
}

func (_c *MockClient_GetVM_Call) Run(run func(ctx context.Context, nodeName string, vmID int64)) *MockClient_GetVM_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(int64))
	})
	return _c
}

func (_c *MockClient_GetVM_Call) Return(_a0 *go_proxmox.VirtualMachine, _a1 error) *MockClient_GetVM_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockClient_GetVM_Call) RunAndReturn(run func(context.Context, string, int64) (*go_proxmox.VirtualMachine, error)) *MockClient_GetVM_Call {
	_c.Call.Return(run)
	return _c
}

// ResizeDisk provides a mock function with given fields: ctx, vm, disk, size
func (_m *MockClient) ResizeDisk(ctx context.Context, vm *go_proxmox.VirtualMachine, disk string, size string) error {
	ret := _m.Called(ctx, vm, disk, size)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *go_proxmox.VirtualMachine, string, string) error); ok {
		r0 = rf(ctx, vm, disk, size)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockClient_ResizeDisk_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ResizeDisk'
type MockClient_ResizeDisk_Call struct {
	*mock.Call
}

// ResizeDisk is a helper method to define mock.On call
//   - ctx context.Context
//   - vm *go_proxmox.VirtualMachine
//   - disk string
//   - size string
func (_e *MockClient_Expecter) ResizeDisk(ctx interface{}, vm interface{}, disk interface{}, size interface{}) *MockClient_ResizeDisk_Call {
	return &MockClient_ResizeDisk_Call{Call: _e.mock.On("ResizeDisk", ctx, vm, disk, size)}
}

func (_c *MockClient_ResizeDisk_Call) Run(run func(ctx context.Context, vm *go_proxmox.VirtualMachine, disk string, size string)) *MockClient_ResizeDisk_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*go_proxmox.VirtualMachine), args[2].(string), args[3].(string))
	})
	return _c
}

func (_c *MockClient_ResizeDisk_Call) Return(_a0 error) *MockClient_ResizeDisk_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockClient_ResizeDisk_Call) RunAndReturn(run func(context.Context, *go_proxmox.VirtualMachine, string, string) error) *MockClient_ResizeDisk_Call {
	_c.Call.Return(run)
	return _c
}

// ResumeVM provides a mock function with given fields: ctx, vm
func (_m *MockClient) ResumeVM(ctx context.Context, vm *go_proxmox.VirtualMachine) (*go_proxmox.Task, error) {
	ret := _m.Called(ctx, vm)

	var r0 *go_proxmox.Task
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *go_proxmox.VirtualMachine) (*go_proxmox.Task, error)); ok {
		return rf(ctx, vm)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *go_proxmox.VirtualMachine) *go_proxmox.Task); ok {
		r0 = rf(ctx, vm)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*go_proxmox.Task)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *go_proxmox.VirtualMachine) error); ok {
		r1 = rf(ctx, vm)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockClient_ResumeVM_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ResumeVM'
type MockClient_ResumeVM_Call struct {
	*mock.Call
}

// ResumeVM is a helper method to define mock.On call
//   - ctx context.Context
//   - vm *go_proxmox.VirtualMachine
func (_e *MockClient_Expecter) ResumeVM(ctx interface{}, vm interface{}) *MockClient_ResumeVM_Call {
	return &MockClient_ResumeVM_Call{Call: _e.mock.On("ResumeVM", ctx, vm)}
}

func (_c *MockClient_ResumeVM_Call) Run(run func(ctx context.Context, vm *go_proxmox.VirtualMachine)) *MockClient_ResumeVM_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*go_proxmox.VirtualMachine))
	})
	return _c
}

func (_c *MockClient_ResumeVM_Call) Return(_a0 *go_proxmox.Task, _a1 error) *MockClient_ResumeVM_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockClient_ResumeVM_Call) RunAndReturn(run func(context.Context, *go_proxmox.VirtualMachine) (*go_proxmox.Task, error)) *MockClient_ResumeVM_Call {
	_c.Call.Return(run)
	return _c
}

// StartVM provides a mock function with given fields: ctx, vm
func (_m *MockClient) StartVM(ctx context.Context, vm *go_proxmox.VirtualMachine) (*go_proxmox.Task, error) {
	ret := _m.Called(ctx, vm)

	var r0 *go_proxmox.Task
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *go_proxmox.VirtualMachine) (*go_proxmox.Task, error)); ok {
		return rf(ctx, vm)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *go_proxmox.VirtualMachine) *go_proxmox.Task); ok {
		r0 = rf(ctx, vm)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*go_proxmox.Task)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *go_proxmox.VirtualMachine) error); ok {
		r1 = rf(ctx, vm)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockClient_StartVM_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'StartVM'
type MockClient_StartVM_Call struct {
	*mock.Call
}

// StartVM is a helper method to define mock.On call
//   - ctx context.Context
//   - vm *go_proxmox.VirtualMachine
func (_e *MockClient_Expecter) StartVM(ctx interface{}, vm interface{}) *MockClient_StartVM_Call {
	return &MockClient_StartVM_Call{Call: _e.mock.On("StartVM", ctx, vm)}
}

func (_c *MockClient_StartVM_Call) Run(run func(ctx context.Context, vm *go_proxmox.VirtualMachine)) *MockClient_StartVM_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*go_proxmox.VirtualMachine))
	})
	return _c
}

func (_c *MockClient_StartVM_Call) Return(_a0 *go_proxmox.Task, _a1 error) *MockClient_StartVM_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockClient_StartVM_Call) RunAndReturn(run func(context.Context, *go_proxmox.VirtualMachine) (*go_proxmox.Task, error)) *MockClient_StartVM_Call {
	_c.Call.Return(run)
	return _c
}

// TagVM provides a mock function with given fields: ctx, vm, tag
func (_m *MockClient) TagVM(ctx context.Context, vm *go_proxmox.VirtualMachine, tag string) (*go_proxmox.Task, error) {
	ret := _m.Called(ctx, vm, tag)

	var r0 *go_proxmox.Task
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *go_proxmox.VirtualMachine, string) (*go_proxmox.Task, error)); ok {
		return rf(ctx, vm, tag)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *go_proxmox.VirtualMachine, string) *go_proxmox.Task); ok {
		r0 = rf(ctx, vm, tag)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*go_proxmox.Task)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *go_proxmox.VirtualMachine, string) error); ok {
		r1 = rf(ctx, vm, tag)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockClient_TagVM_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'TagVM'
type MockClient_TagVM_Call struct {
	*mock.Call
}

// TagVM is a helper method to define mock.On call
//   - ctx context.Context
//   - vm *go_proxmox.VirtualMachine
//   - tag string
func (_e *MockClient_Expecter) TagVM(ctx interface{}, vm interface{}, tag interface{}) *MockClient_TagVM_Call {
	return &MockClient_TagVM_Call{Call: _e.mock.On("TagVM", ctx, vm, tag)}
}

func (_c *MockClient_TagVM_Call) Run(run func(ctx context.Context, vm *go_proxmox.VirtualMachine, tag string)) *MockClient_TagVM_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*go_proxmox.VirtualMachine), args[2].(string))
	})
	return _c
}

func (_c *MockClient_TagVM_Call) Return(_a0 *go_proxmox.Task, _a1 error) *MockClient_TagVM_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockClient_TagVM_Call) RunAndReturn(run func(context.Context, *go_proxmox.VirtualMachine, string) (*go_proxmox.Task, error)) *MockClient_TagVM_Call {
	_c.Call.Return(run)
	return _c
}

// UnmountCloudInitISO provides a mock function with given fields: ctx, vm, device
func (_m *MockClient) UnmountCloudInitISO(ctx context.Context, vm *go_proxmox.VirtualMachine, device string) error {
	ret := _m.Called(ctx, vm, device)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *go_proxmox.VirtualMachine, string) error); ok {
		r0 = rf(ctx, vm, device)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockClient_UnmountCloudInitISO_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UnmountCloudInitISO'
type MockClient_UnmountCloudInitISO_Call struct {
	*mock.Call
}

// UnmountCloudInitISO is a helper method to define mock.On call
//   - ctx context.Context
//   - vm *go_proxmox.VirtualMachine
//   - device string
func (_e *MockClient_Expecter) UnmountCloudInitISO(ctx interface{}, vm interface{}, device interface{}) *MockClient_UnmountCloudInitISO_Call {
	return &MockClient_UnmountCloudInitISO_Call{Call: _e.mock.On("UnmountCloudInitISO", ctx, vm, device)}
}

func (_c *MockClient_UnmountCloudInitISO_Call) Run(run func(ctx context.Context, vm *go_proxmox.VirtualMachine, device string)) *MockClient_UnmountCloudInitISO_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*go_proxmox.VirtualMachine), args[2].(string))
	})
	return _c
}

func (_c *MockClient_UnmountCloudInitISO_Call) Return(_a0 error) *MockClient_UnmountCloudInitISO_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockClient_UnmountCloudInitISO_Call) RunAndReturn(run func(context.Context, *go_proxmox.VirtualMachine, string) error) *MockClient_UnmountCloudInitISO_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockClient creates a new instance of MockClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockClient {
	mock := &MockClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

// Code generated by mockery v2.32.4. DO NOT EDIT.

package mocks

import (
	context "context"

	serviceuser "github.com/raystack/frontier/core/serviceuser"
	mock "github.com/stretchr/testify/mock"
)

// ServiceUserService is an autogenerated mock type for the ServiceUserService type
type ServiceUserService struct {
	mock.Mock
}

type ServiceUserService_Expecter struct {
	mock *mock.Mock
}

func (_m *ServiceUserService) EXPECT() *ServiceUserService_Expecter {
	return &ServiceUserService_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: ctx, serviceUser
func (_m *ServiceUserService) Create(ctx context.Context, serviceUser serviceuser.ServiceUser) (serviceuser.ServiceUser, error) {
	ret := _m.Called(ctx, serviceUser)

	var r0 serviceuser.ServiceUser
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, serviceuser.ServiceUser) (serviceuser.ServiceUser, error)); ok {
		return rf(ctx, serviceUser)
	}
	if rf, ok := ret.Get(0).(func(context.Context, serviceuser.ServiceUser) serviceuser.ServiceUser); ok {
		r0 = rf(ctx, serviceUser)
	} else {
		r0 = ret.Get(0).(serviceuser.ServiceUser)
	}

	if rf, ok := ret.Get(1).(func(context.Context, serviceuser.ServiceUser) error); ok {
		r1 = rf(ctx, serviceUser)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ServiceUserService_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type ServiceUserService_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - ctx context.Context
//   - serviceUser serviceuser.ServiceUser
func (_e *ServiceUserService_Expecter) Create(ctx interface{}, serviceUser interface{}) *ServiceUserService_Create_Call {
	return &ServiceUserService_Create_Call{Call: _e.mock.On("Create", ctx, serviceUser)}
}

func (_c *ServiceUserService_Create_Call) Run(run func(ctx context.Context, serviceUser serviceuser.ServiceUser)) *ServiceUserService_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(serviceuser.ServiceUser))
	})
	return _c
}

func (_c *ServiceUserService_Create_Call) Return(_a0 serviceuser.ServiceUser, _a1 error) *ServiceUserService_Create_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ServiceUserService_Create_Call) RunAndReturn(run func(context.Context, serviceuser.ServiceUser) (serviceuser.ServiceUser, error)) *ServiceUserService_Create_Call {
	_c.Call.Return(run)
	return _c
}

// CreateKey provides a mock function with given fields: ctx, cred
func (_m *ServiceUserService) CreateKey(ctx context.Context, cred serviceuser.Credential) (serviceuser.Credential, error) {
	ret := _m.Called(ctx, cred)

	var r0 serviceuser.Credential
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, serviceuser.Credential) (serviceuser.Credential, error)); ok {
		return rf(ctx, cred)
	}
	if rf, ok := ret.Get(0).(func(context.Context, serviceuser.Credential) serviceuser.Credential); ok {
		r0 = rf(ctx, cred)
	} else {
		r0 = ret.Get(0).(serviceuser.Credential)
	}

	if rf, ok := ret.Get(1).(func(context.Context, serviceuser.Credential) error); ok {
		r1 = rf(ctx, cred)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ServiceUserService_CreateKey_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateKey'
type ServiceUserService_CreateKey_Call struct {
	*mock.Call
}

// CreateKey is a helper method to define mock.On call
//   - ctx context.Context
//   - cred serviceuser.Credential
func (_e *ServiceUserService_Expecter) CreateKey(ctx interface{}, cred interface{}) *ServiceUserService_CreateKey_Call {
	return &ServiceUserService_CreateKey_Call{Call: _e.mock.On("CreateKey", ctx, cred)}
}

func (_c *ServiceUserService_CreateKey_Call) Run(run func(ctx context.Context, cred serviceuser.Credential)) *ServiceUserService_CreateKey_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(serviceuser.Credential))
	})
	return _c
}

func (_c *ServiceUserService_CreateKey_Call) Return(_a0 serviceuser.Credential, _a1 error) *ServiceUserService_CreateKey_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ServiceUserService_CreateKey_Call) RunAndReturn(run func(context.Context, serviceuser.Credential) (serviceuser.Credential, error)) *ServiceUserService_CreateKey_Call {
	_c.Call.Return(run)
	return _c
}

// CreateSecret provides a mock function with given fields: ctx, credential
func (_m *ServiceUserService) CreateSecret(ctx context.Context, credential serviceuser.Credential) (serviceuser.Secret, error) {
	ret := _m.Called(ctx, credential)

	var r0 serviceuser.Secret
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, serviceuser.Credential) (serviceuser.Secret, error)); ok {
		return rf(ctx, credential)
	}
	if rf, ok := ret.Get(0).(func(context.Context, serviceuser.Credential) serviceuser.Secret); ok {
		r0 = rf(ctx, credential)
	} else {
		r0 = ret.Get(0).(serviceuser.Secret)
	}

	if rf, ok := ret.Get(1).(func(context.Context, serviceuser.Credential) error); ok {
		r1 = rf(ctx, credential)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ServiceUserService_CreateSecret_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateSecret'
type ServiceUserService_CreateSecret_Call struct {
	*mock.Call
}

// CreateSecret is a helper method to define mock.On call
//   - ctx context.Context
//   - credential serviceuser.Credential
func (_e *ServiceUserService_Expecter) CreateSecret(ctx interface{}, credential interface{}) *ServiceUserService_CreateSecret_Call {
	return &ServiceUserService_CreateSecret_Call{Call: _e.mock.On("CreateSecret", ctx, credential)}
}

func (_c *ServiceUserService_CreateSecret_Call) Run(run func(ctx context.Context, credential serviceuser.Credential)) *ServiceUserService_CreateSecret_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(serviceuser.Credential))
	})
	return _c
}

func (_c *ServiceUserService_CreateSecret_Call) Return(_a0 serviceuser.Secret, _a1 error) *ServiceUserService_CreateSecret_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ServiceUserService_CreateSecret_Call) RunAndReturn(run func(context.Context, serviceuser.Credential) (serviceuser.Secret, error)) *ServiceUserService_CreateSecret_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: ctx, id
func (_m *ServiceUserService) Delete(ctx context.Context, id string) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ServiceUserService_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type ServiceUserService_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - ctx context.Context
//   - id string
func (_e *ServiceUserService_Expecter) Delete(ctx interface{}, id interface{}) *ServiceUserService_Delete_Call {
	return &ServiceUserService_Delete_Call{Call: _e.mock.On("Delete", ctx, id)}
}

func (_c *ServiceUserService_Delete_Call) Run(run func(ctx context.Context, id string)) *ServiceUserService_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *ServiceUserService_Delete_Call) Return(_a0 error) *ServiceUserService_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ServiceUserService_Delete_Call) RunAndReturn(run func(context.Context, string) error) *ServiceUserService_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteKey provides a mock function with given fields: ctx, credID
func (_m *ServiceUserService) DeleteKey(ctx context.Context, credID string) error {
	ret := _m.Called(ctx, credID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, credID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ServiceUserService_DeleteKey_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteKey'
type ServiceUserService_DeleteKey_Call struct {
	*mock.Call
}

// DeleteKey is a helper method to define mock.On call
//   - ctx context.Context
//   - credID string
func (_e *ServiceUserService_Expecter) DeleteKey(ctx interface{}, credID interface{}) *ServiceUserService_DeleteKey_Call {
	return &ServiceUserService_DeleteKey_Call{Call: _e.mock.On("DeleteKey", ctx, credID)}
}

func (_c *ServiceUserService_DeleteKey_Call) Run(run func(ctx context.Context, credID string)) *ServiceUserService_DeleteKey_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *ServiceUserService_DeleteKey_Call) Return(_a0 error) *ServiceUserService_DeleteKey_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ServiceUserService_DeleteKey_Call) RunAndReturn(run func(context.Context, string) error) *ServiceUserService_DeleteKey_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteSecret provides a mock function with given fields: ctx, credID
func (_m *ServiceUserService) DeleteSecret(ctx context.Context, credID string) error {
	ret := _m.Called(ctx, credID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, credID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ServiceUserService_DeleteSecret_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteSecret'
type ServiceUserService_DeleteSecret_Call struct {
	*mock.Call
}

// DeleteSecret is a helper method to define mock.On call
//   - ctx context.Context
//   - credID string
func (_e *ServiceUserService_Expecter) DeleteSecret(ctx interface{}, credID interface{}) *ServiceUserService_DeleteSecret_Call {
	return &ServiceUserService_DeleteSecret_Call{Call: _e.mock.On("DeleteSecret", ctx, credID)}
}

func (_c *ServiceUserService_DeleteSecret_Call) Run(run func(ctx context.Context, credID string)) *ServiceUserService_DeleteSecret_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *ServiceUserService_DeleteSecret_Call) Return(_a0 error) *ServiceUserService_DeleteSecret_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ServiceUserService_DeleteSecret_Call) RunAndReturn(run func(context.Context, string) error) *ServiceUserService_DeleteSecret_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields: ctx, id
func (_m *ServiceUserService) Get(ctx context.Context, id string) (serviceuser.ServiceUser, error) {
	ret := _m.Called(ctx, id)

	var r0 serviceuser.ServiceUser
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (serviceuser.ServiceUser, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) serviceuser.ServiceUser); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(serviceuser.ServiceUser)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ServiceUserService_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type ServiceUserService_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - ctx context.Context
//   - id string
func (_e *ServiceUserService_Expecter) Get(ctx interface{}, id interface{}) *ServiceUserService_Get_Call {
	return &ServiceUserService_Get_Call{Call: _e.mock.On("Get", ctx, id)}
}

func (_c *ServiceUserService_Get_Call) Run(run func(ctx context.Context, id string)) *ServiceUserService_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *ServiceUserService_Get_Call) Return(_a0 serviceuser.ServiceUser, _a1 error) *ServiceUserService_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ServiceUserService_Get_Call) RunAndReturn(run func(context.Context, string) (serviceuser.ServiceUser, error)) *ServiceUserService_Get_Call {
	_c.Call.Return(run)
	return _c
}

// GetKey provides a mock function with given fields: ctx, credID
func (_m *ServiceUserService) GetKey(ctx context.Context, credID string) (serviceuser.Credential, error) {
	ret := _m.Called(ctx, credID)

	var r0 serviceuser.Credential
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (serviceuser.Credential, error)); ok {
		return rf(ctx, credID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) serviceuser.Credential); ok {
		r0 = rf(ctx, credID)
	} else {
		r0 = ret.Get(0).(serviceuser.Credential)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, credID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ServiceUserService_GetKey_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetKey'
type ServiceUserService_GetKey_Call struct {
	*mock.Call
}

// GetKey is a helper method to define mock.On call
//   - ctx context.Context
//   - credID string
func (_e *ServiceUserService_Expecter) GetKey(ctx interface{}, credID interface{}) *ServiceUserService_GetKey_Call {
	return &ServiceUserService_GetKey_Call{Call: _e.mock.On("GetKey", ctx, credID)}
}

func (_c *ServiceUserService_GetKey_Call) Run(run func(ctx context.Context, credID string)) *ServiceUserService_GetKey_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *ServiceUserService_GetKey_Call) Return(_a0 serviceuser.Credential, _a1 error) *ServiceUserService_GetKey_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ServiceUserService_GetKey_Call) RunAndReturn(run func(context.Context, string) (serviceuser.Credential, error)) *ServiceUserService_GetKey_Call {
	_c.Call.Return(run)
	return _c
}

// List provides a mock function with given fields: ctx, flt
func (_m *ServiceUserService) List(ctx context.Context, flt serviceuser.Filter) ([]serviceuser.ServiceUser, error) {
	ret := _m.Called(ctx, flt)

	var r0 []serviceuser.ServiceUser
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, serviceuser.Filter) ([]serviceuser.ServiceUser, error)); ok {
		return rf(ctx, flt)
	}
	if rf, ok := ret.Get(0).(func(context.Context, serviceuser.Filter) []serviceuser.ServiceUser); ok {
		r0 = rf(ctx, flt)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]serviceuser.ServiceUser)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, serviceuser.Filter) error); ok {
		r1 = rf(ctx, flt)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ServiceUserService_List_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'List'
type ServiceUserService_List_Call struct {
	*mock.Call
}

// List is a helper method to define mock.On call
//   - ctx context.Context
//   - flt serviceuser.Filter
func (_e *ServiceUserService_Expecter) List(ctx interface{}, flt interface{}) *ServiceUserService_List_Call {
	return &ServiceUserService_List_Call{Call: _e.mock.On("List", ctx, flt)}
}

func (_c *ServiceUserService_List_Call) Run(run func(ctx context.Context, flt serviceuser.Filter)) *ServiceUserService_List_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(serviceuser.Filter))
	})
	return _c
}

func (_c *ServiceUserService_List_Call) Return(_a0 []serviceuser.ServiceUser, _a1 error) *ServiceUserService_List_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ServiceUserService_List_Call) RunAndReturn(run func(context.Context, serviceuser.Filter) ([]serviceuser.ServiceUser, error)) *ServiceUserService_List_Call {
	_c.Call.Return(run)
	return _c
}

// ListByOrg provides a mock function with given fields: ctx, orgID
func (_m *ServiceUserService) ListByOrg(ctx context.Context, orgID string) ([]serviceuser.ServiceUser, error) {
	ret := _m.Called(ctx, orgID)

	var r0 []serviceuser.ServiceUser
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) ([]serviceuser.ServiceUser, error)); ok {
		return rf(ctx, orgID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) []serviceuser.ServiceUser); ok {
		r0 = rf(ctx, orgID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]serviceuser.ServiceUser)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, orgID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ServiceUserService_ListByOrg_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListByOrg'
type ServiceUserService_ListByOrg_Call struct {
	*mock.Call
}

// ListByOrg is a helper method to define mock.On call
//   - ctx context.Context
//   - orgID string
func (_e *ServiceUserService_Expecter) ListByOrg(ctx interface{}, orgID interface{}) *ServiceUserService_ListByOrg_Call {
	return &ServiceUserService_ListByOrg_Call{Call: _e.mock.On("ListByOrg", ctx, orgID)}
}

func (_c *ServiceUserService_ListByOrg_Call) Run(run func(ctx context.Context, orgID string)) *ServiceUserService_ListByOrg_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *ServiceUserService_ListByOrg_Call) Return(_a0 []serviceuser.ServiceUser, _a1 error) *ServiceUserService_ListByOrg_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ServiceUserService_ListByOrg_Call) RunAndReturn(run func(context.Context, string) ([]serviceuser.ServiceUser, error)) *ServiceUserService_ListByOrg_Call {
	_c.Call.Return(run)
	return _c
}

// ListKeys provides a mock function with given fields: ctx, serviceUserID
func (_m *ServiceUserService) ListKeys(ctx context.Context, serviceUserID string) ([]serviceuser.Credential, error) {
	ret := _m.Called(ctx, serviceUserID)

	var r0 []serviceuser.Credential
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) ([]serviceuser.Credential, error)); ok {
		return rf(ctx, serviceUserID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) []serviceuser.Credential); ok {
		r0 = rf(ctx, serviceUserID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]serviceuser.Credential)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, serviceUserID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ServiceUserService_ListKeys_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListKeys'
type ServiceUserService_ListKeys_Call struct {
	*mock.Call
}

// ListKeys is a helper method to define mock.On call
//   - ctx context.Context
//   - serviceUserID string
func (_e *ServiceUserService_Expecter) ListKeys(ctx interface{}, serviceUserID interface{}) *ServiceUserService_ListKeys_Call {
	return &ServiceUserService_ListKeys_Call{Call: _e.mock.On("ListKeys", ctx, serviceUserID)}
}

func (_c *ServiceUserService_ListKeys_Call) Run(run func(ctx context.Context, serviceUserID string)) *ServiceUserService_ListKeys_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *ServiceUserService_ListKeys_Call) Return(_a0 []serviceuser.Credential, _a1 error) *ServiceUserService_ListKeys_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ServiceUserService_ListKeys_Call) RunAndReturn(run func(context.Context, string) ([]serviceuser.Credential, error)) *ServiceUserService_ListKeys_Call {
	_c.Call.Return(run)
	return _c
}

// ListSecret provides a mock function with given fields: ctx, serviceUserID
func (_m *ServiceUserService) ListSecret(ctx context.Context, serviceUserID string) ([]serviceuser.Credential, error) {
	ret := _m.Called(ctx, serviceUserID)

	var r0 []serviceuser.Credential
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) ([]serviceuser.Credential, error)); ok {
		return rf(ctx, serviceUserID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) []serviceuser.Credential); ok {
		r0 = rf(ctx, serviceUserID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]serviceuser.Credential)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, serviceUserID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ServiceUserService_ListSecret_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListSecret'
type ServiceUserService_ListSecret_Call struct {
	*mock.Call
}

// ListSecret is a helper method to define mock.On call
//   - ctx context.Context
//   - serviceUserID string
func (_e *ServiceUserService_Expecter) ListSecret(ctx interface{}, serviceUserID interface{}) *ServiceUserService_ListSecret_Call {
	return &ServiceUserService_ListSecret_Call{Call: _e.mock.On("ListSecret", ctx, serviceUserID)}
}

func (_c *ServiceUserService_ListSecret_Call) Run(run func(ctx context.Context, serviceUserID string)) *ServiceUserService_ListSecret_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *ServiceUserService_ListSecret_Call) Return(_a0 []serviceuser.Credential, _a1 error) *ServiceUserService_ListSecret_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ServiceUserService_ListSecret_Call) RunAndReturn(run func(context.Context, string) ([]serviceuser.Credential, error)) *ServiceUserService_ListSecret_Call {
	_c.Call.Return(run)
	return _c
}

// NewServiceUserService creates a new instance of ServiceUserService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewServiceUserService(t interface {
	mock.TestingT
	Cleanup(func())
}) *ServiceUserService {
	mock := &ServiceUserService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

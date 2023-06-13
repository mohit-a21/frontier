// Code generated by mockery v2.20.2. DO NOT EDIT.

package mocks

import (
	context "context"

	user "github.com/raystack/shield/core/user"
	mock "github.com/stretchr/testify/mock"
)

// UserService is an autogenerated mock type for the UserService type
type UserService struct {
	mock.Mock
}

type UserService_Expecter struct {
	mock *mock.Mock
}

func (_m *UserService) EXPECT() *UserService_Expecter {
	return &UserService_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: ctx, _a1
func (_m *UserService) Create(ctx context.Context, _a1 user.User) (user.User, error) {
	ret := _m.Called(ctx, _a1)

	var r0 user.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, user.User) (user.User, error)); ok {
		return rf(ctx, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, user.User) user.User); ok {
		r0 = rf(ctx, _a1)
	} else {
		r0 = ret.Get(0).(user.User)
	}

	if rf, ok := ret.Get(1).(func(context.Context, user.User) error); ok {
		r1 = rf(ctx, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UserService_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type UserService_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - ctx context.Context
//   - _a1 user.User
func (_e *UserService_Expecter) Create(ctx interface{}, _a1 interface{}) *UserService_Create_Call {
	return &UserService_Create_Call{Call: _e.mock.On("Create", ctx, _a1)}
}

func (_c *UserService_Create_Call) Run(run func(ctx context.Context, _a1 user.User)) *UserService_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(user.User))
	})
	return _c
}

func (_c *UserService_Create_Call) Return(_a0 user.User, _a1 error) *UserService_Create_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *UserService_Create_Call) RunAndReturn(run func(context.Context, user.User) (user.User, error)) *UserService_Create_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: ctx, id
func (_m *UserService) Delete(ctx context.Context, id string) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UserService_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type UserService_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - ctx context.Context
//   - id string
func (_e *UserService_Expecter) Delete(ctx interface{}, id interface{}) *UserService_Delete_Call {
	return &UserService_Delete_Call{Call: _e.mock.On("Delete", ctx, id)}
}

func (_c *UserService_Delete_Call) Run(run func(ctx context.Context, id string)) *UserService_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *UserService_Delete_Call) Return(_a0 error) *UserService_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *UserService_Delete_Call) RunAndReturn(run func(context.Context, string) error) *UserService_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// Disable provides a mock function with given fields: ctx, id
func (_m *UserService) Disable(ctx context.Context, id string) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UserService_Disable_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Disable'
type UserService_Disable_Call struct {
	*mock.Call
}

// Disable is a helper method to define mock.On call
//   - ctx context.Context
//   - id string
func (_e *UserService_Expecter) Disable(ctx interface{}, id interface{}) *UserService_Disable_Call {
	return &UserService_Disable_Call{Call: _e.mock.On("Disable", ctx, id)}
}

func (_c *UserService_Disable_Call) Run(run func(ctx context.Context, id string)) *UserService_Disable_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *UserService_Disable_Call) Return(_a0 error) *UserService_Disable_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *UserService_Disable_Call) RunAndReturn(run func(context.Context, string) error) *UserService_Disable_Call {
	_c.Call.Return(run)
	return _c
}

// Enable provides a mock function with given fields: ctx, id
func (_m *UserService) Enable(ctx context.Context, id string) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UserService_Enable_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Enable'
type UserService_Enable_Call struct {
	*mock.Call
}

// Enable is a helper method to define mock.On call
//   - ctx context.Context
//   - id string
func (_e *UserService_Expecter) Enable(ctx interface{}, id interface{}) *UserService_Enable_Call {
	return &UserService_Enable_Call{Call: _e.mock.On("Enable", ctx, id)}
}

func (_c *UserService_Enable_Call) Run(run func(ctx context.Context, id string)) *UserService_Enable_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *UserService_Enable_Call) Return(_a0 error) *UserService_Enable_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *UserService_Enable_Call) RunAndReturn(run func(context.Context, string) error) *UserService_Enable_Call {
	_c.Call.Return(run)
	return _c
}

// FetchCurrentUser provides a mock function with given fields: ctx
func (_m *UserService) FetchCurrentUser(ctx context.Context) (user.User, error) {
	ret := _m.Called(ctx)

	var r0 user.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (user.User, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) user.User); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(user.User)
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UserService_FetchCurrentUser_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FetchCurrentUser'
type UserService_FetchCurrentUser_Call struct {
	*mock.Call
}

// FetchCurrentUser is a helper method to define mock.On call
//   - ctx context.Context
func (_e *UserService_Expecter) FetchCurrentUser(ctx interface{}) *UserService_FetchCurrentUser_Call {
	return &UserService_FetchCurrentUser_Call{Call: _e.mock.On("FetchCurrentUser", ctx)}
}

func (_c *UserService_FetchCurrentUser_Call) Run(run func(ctx context.Context)) *UserService_FetchCurrentUser_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *UserService_FetchCurrentUser_Call) Return(_a0 user.User, _a1 error) *UserService_FetchCurrentUser_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *UserService_FetchCurrentUser_Call) RunAndReturn(run func(context.Context) (user.User, error)) *UserService_FetchCurrentUser_Call {
	_c.Call.Return(run)
	return _c
}

// GetByEmail provides a mock function with given fields: ctx, email
func (_m *UserService) GetByEmail(ctx context.Context, email string) (user.User, error) {
	ret := _m.Called(ctx, email)

	var r0 user.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (user.User, error)); ok {
		return rf(ctx, email)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) user.User); ok {
		r0 = rf(ctx, email)
	} else {
		r0 = ret.Get(0).(user.User)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UserService_GetByEmail_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetByEmail'
type UserService_GetByEmail_Call struct {
	*mock.Call
}

// GetByEmail is a helper method to define mock.On call
//   - ctx context.Context
//   - email string
func (_e *UserService_Expecter) GetByEmail(ctx interface{}, email interface{}) *UserService_GetByEmail_Call {
	return &UserService_GetByEmail_Call{Call: _e.mock.On("GetByEmail", ctx, email)}
}

func (_c *UserService_GetByEmail_Call) Run(run func(ctx context.Context, email string)) *UserService_GetByEmail_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *UserService_GetByEmail_Call) Return(_a0 user.User, _a1 error) *UserService_GetByEmail_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *UserService_GetByEmail_Call) RunAndReturn(run func(context.Context, string) (user.User, error)) *UserService_GetByEmail_Call {
	_c.Call.Return(run)
	return _c
}

// GetByID provides a mock function with given fields: ctx, id
func (_m *UserService) GetByID(ctx context.Context, id string) (user.User, error) {
	ret := _m.Called(ctx, id)

	var r0 user.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (user.User, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) user.User); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(user.User)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UserService_GetByID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetByID'
type UserService_GetByID_Call struct {
	*mock.Call
}

// GetByID is a helper method to define mock.On call
//   - ctx context.Context
//   - id string
func (_e *UserService_Expecter) GetByID(ctx interface{}, id interface{}) *UserService_GetByID_Call {
	return &UserService_GetByID_Call{Call: _e.mock.On("GetByID", ctx, id)}
}

func (_c *UserService_GetByID_Call) Run(run func(ctx context.Context, id string)) *UserService_GetByID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *UserService_GetByID_Call) Return(_a0 user.User, _a1 error) *UserService_GetByID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *UserService_GetByID_Call) RunAndReturn(run func(context.Context, string) (user.User, error)) *UserService_GetByID_Call {
	_c.Call.Return(run)
	return _c
}

// GetByIDs provides a mock function with given fields: ctx, userIDs
func (_m *UserService) GetByIDs(ctx context.Context, userIDs []string) ([]user.User, error) {
	ret := _m.Called(ctx, userIDs)

	var r0 []user.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []string) ([]user.User, error)); ok {
		return rf(ctx, userIDs)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []string) []user.User); ok {
		r0 = rf(ctx, userIDs)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]user.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, []string) error); ok {
		r1 = rf(ctx, userIDs)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UserService_GetByIDs_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetByIDs'
type UserService_GetByIDs_Call struct {
	*mock.Call
}

// GetByIDs is a helper method to define mock.On call
//   - ctx context.Context
//   - userIDs []string
func (_e *UserService_Expecter) GetByIDs(ctx interface{}, userIDs interface{}) *UserService_GetByIDs_Call {
	return &UserService_GetByIDs_Call{Call: _e.mock.On("GetByIDs", ctx, userIDs)}
}

func (_c *UserService_GetByIDs_Call) Run(run func(ctx context.Context, userIDs []string)) *UserService_GetByIDs_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].([]string))
	})
	return _c
}

func (_c *UserService_GetByIDs_Call) Return(_a0 []user.User, _a1 error) *UserService_GetByIDs_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *UserService_GetByIDs_Call) RunAndReturn(run func(context.Context, []string) ([]user.User, error)) *UserService_GetByIDs_Call {
	_c.Call.Return(run)
	return _c
}

// IsSudo provides a mock function with given fields: ctx, id
func (_m *UserService) IsSudo(ctx context.Context, id string) (bool, error) {
	ret := _m.Called(ctx, id)

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (bool, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) bool); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UserService_IsSudo_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsSudo'
type UserService_IsSudo_Call struct {
	*mock.Call
}

// IsSudo is a helper method to define mock.On call
//   - ctx context.Context
//   - id string
func (_e *UserService_Expecter) IsSudo(ctx interface{}, id interface{}) *UserService_IsSudo_Call {
	return &UserService_IsSudo_Call{Call: _e.mock.On("IsSudo", ctx, id)}
}

func (_c *UserService_IsSudo_Call) Run(run func(ctx context.Context, id string)) *UserService_IsSudo_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *UserService_IsSudo_Call) Return(_a0 bool, _a1 error) *UserService_IsSudo_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *UserService_IsSudo_Call) RunAndReturn(run func(context.Context, string) (bool, error)) *UserService_IsSudo_Call {
	_c.Call.Return(run)
	return _c
}

// List provides a mock function with given fields: ctx, flt
func (_m *UserService) List(ctx context.Context, flt user.Filter) ([]user.User, error) {
	ret := _m.Called(ctx, flt)

	var r0 []user.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, user.Filter) ([]user.User, error)); ok {
		return rf(ctx, flt)
	}
	if rf, ok := ret.Get(0).(func(context.Context, user.Filter) []user.User); ok {
		r0 = rf(ctx, flt)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]user.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, user.Filter) error); ok {
		r1 = rf(ctx, flt)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UserService_List_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'List'
type UserService_List_Call struct {
	*mock.Call
}

// List is a helper method to define mock.On call
//   - ctx context.Context
//   - flt user.Filter
func (_e *UserService_Expecter) List(ctx interface{}, flt interface{}) *UserService_List_Call {
	return &UserService_List_Call{Call: _e.mock.On("List", ctx, flt)}
}

func (_c *UserService_List_Call) Run(run func(ctx context.Context, flt user.Filter)) *UserService_List_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(user.Filter))
	})
	return _c
}

func (_c *UserService_List_Call) Return(_a0 []user.User, _a1 error) *UserService_List_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *UserService_List_Call) RunAndReturn(run func(context.Context, user.Filter) ([]user.User, error)) *UserService_List_Call {
	_c.Call.Return(run)
	return _c
}

// ListByOrg provides a mock function with given fields: ctx, orgID, permissionFilter
func (_m *UserService) ListByOrg(ctx context.Context, orgID string, permissionFilter string) ([]user.User, error) {
	ret := _m.Called(ctx, orgID, permissionFilter)

	var r0 []user.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) ([]user.User, error)); ok {
		return rf(ctx, orgID, permissionFilter)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) []user.User); ok {
		r0 = rf(ctx, orgID, permissionFilter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]user.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, orgID, permissionFilter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UserService_ListByOrg_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListByOrg'
type UserService_ListByOrg_Call struct {
	*mock.Call
}

// ListByOrg is a helper method to define mock.On call
//   - ctx context.Context
//   - orgID string
//   - permissionFilter string
func (_e *UserService_Expecter) ListByOrg(ctx interface{}, orgID interface{}, permissionFilter interface{}) *UserService_ListByOrg_Call {
	return &UserService_ListByOrg_Call{Call: _e.mock.On("ListByOrg", ctx, orgID, permissionFilter)}
}

func (_c *UserService_ListByOrg_Call) Run(run func(ctx context.Context, orgID string, permissionFilter string)) *UserService_ListByOrg_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *UserService_ListByOrg_Call) Return(_a0 []user.User, _a1 error) *UserService_ListByOrg_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *UserService_ListByOrg_Call) RunAndReturn(run func(context.Context, string, string) ([]user.User, error)) *UserService_ListByOrg_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: ctx, toUpdate
func (_m *UserService) Update(ctx context.Context, toUpdate user.User) (user.User, error) {
	ret := _m.Called(ctx, toUpdate)

	var r0 user.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, user.User) (user.User, error)); ok {
		return rf(ctx, toUpdate)
	}
	if rf, ok := ret.Get(0).(func(context.Context, user.User) user.User); ok {
		r0 = rf(ctx, toUpdate)
	} else {
		r0 = ret.Get(0).(user.User)
	}

	if rf, ok := ret.Get(1).(func(context.Context, user.User) error); ok {
		r1 = rf(ctx, toUpdate)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UserService_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type UserService_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - ctx context.Context
//   - toUpdate user.User
func (_e *UserService_Expecter) Update(ctx interface{}, toUpdate interface{}) *UserService_Update_Call {
	return &UserService_Update_Call{Call: _e.mock.On("Update", ctx, toUpdate)}
}

func (_c *UserService_Update_Call) Run(run func(ctx context.Context, toUpdate user.User)) *UserService_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(user.User))
	})
	return _c
}

func (_c *UserService_Update_Call) Return(_a0 user.User, _a1 error) *UserService_Update_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *UserService_Update_Call) RunAndReturn(run func(context.Context, user.User) (user.User, error)) *UserService_Update_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateByEmail provides a mock function with given fields: ctx, toUpdate
func (_m *UserService) UpdateByEmail(ctx context.Context, toUpdate user.User) (user.User, error) {
	ret := _m.Called(ctx, toUpdate)

	var r0 user.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, user.User) (user.User, error)); ok {
		return rf(ctx, toUpdate)
	}
	if rf, ok := ret.Get(0).(func(context.Context, user.User) user.User); ok {
		r0 = rf(ctx, toUpdate)
	} else {
		r0 = ret.Get(0).(user.User)
	}

	if rf, ok := ret.Get(1).(func(context.Context, user.User) error); ok {
		r1 = rf(ctx, toUpdate)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UserService_UpdateByEmail_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateByEmail'
type UserService_UpdateByEmail_Call struct {
	*mock.Call
}

// UpdateByEmail is a helper method to define mock.On call
//   - ctx context.Context
//   - toUpdate user.User
func (_e *UserService_Expecter) UpdateByEmail(ctx interface{}, toUpdate interface{}) *UserService_UpdateByEmail_Call {
	return &UserService_UpdateByEmail_Call{Call: _e.mock.On("UpdateByEmail", ctx, toUpdate)}
}

func (_c *UserService_UpdateByEmail_Call) Run(run func(ctx context.Context, toUpdate user.User)) *UserService_UpdateByEmail_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(user.User))
	})
	return _c
}

func (_c *UserService_UpdateByEmail_Call) Return(_a0 user.User, _a1 error) *UserService_UpdateByEmail_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *UserService_UpdateByEmail_Call) RunAndReturn(run func(context.Context, user.User) (user.User, error)) *UserService_UpdateByEmail_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewUserService interface {
	mock.TestingT
	Cleanup(func())
}

// NewUserService creates a new instance of UserService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUserService(t mockConstructorTestingTNewUserService) *UserService {
	mock := &UserService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

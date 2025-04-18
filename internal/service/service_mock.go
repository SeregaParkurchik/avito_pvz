// Code generated by mockery v2.53.3. DO NOT EDIT.

package service

import (
	models "avitopvz/internal/models"
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// MockInterface is an autogenerated mock type for the Interface type
type MockInterface struct {
	mock.Mock
}

type MockInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *MockInterface) EXPECT() *MockInterface_Expecter {
	return &MockInterface_Expecter{mock: &_m.Mock}
}

// AddProduct provides a mock function with given fields: ctx, product
func (_m *MockInterface) AddProduct(ctx context.Context, product *models.Product) (*models.Product, error) {
	ret := _m.Called(ctx, product)

	if len(ret) == 0 {
		panic("no return value specified for AddProduct")
	}

	var r0 *models.Product
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *models.Product) (*models.Product, error)); ok {
		return rf(ctx, product)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *models.Product) *models.Product); ok {
		r0 = rf(ctx, product)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Product)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *models.Product) error); ok {
		r1 = rf(ctx, product)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockInterface_AddProduct_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddProduct'
type MockInterface_AddProduct_Call struct {
	*mock.Call
}

// AddProduct is a helper method to define mock.On call
//   - ctx context.Context
//   - product *models.Product
func (_e *MockInterface_Expecter) AddProduct(ctx interface{}, product interface{}) *MockInterface_AddProduct_Call {
	return &MockInterface_AddProduct_Call{Call: _e.mock.On("AddProduct", ctx, product)}
}

func (_c *MockInterface_AddProduct_Call) Run(run func(ctx context.Context, product *models.Product)) *MockInterface_AddProduct_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*models.Product))
	})
	return _c
}

func (_c *MockInterface_AddProduct_Call) Return(_a0 *models.Product, _a1 error) *MockInterface_AddProduct_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockInterface_AddProduct_Call) RunAndReturn(run func(context.Context, *models.Product) (*models.Product, error)) *MockInterface_AddProduct_Call {
	_c.Call.Return(run)
	return _c
}

// CloseLastReception provides a mock function with given fields: ctx, pvzIDStr
func (_m *MockInterface) CloseLastReception(ctx context.Context, pvzIDStr string) (*models.Receptions, error) {
	ret := _m.Called(ctx, pvzIDStr)

	if len(ret) == 0 {
		panic("no return value specified for CloseLastReception")
	}

	var r0 *models.Receptions
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*models.Receptions, error)); ok {
		return rf(ctx, pvzIDStr)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *models.Receptions); ok {
		r0 = rf(ctx, pvzIDStr)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Receptions)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, pvzIDStr)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockInterface_CloseLastReception_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CloseLastReception'
type MockInterface_CloseLastReception_Call struct {
	*mock.Call
}

// CloseLastReception is a helper method to define mock.On call
//   - ctx context.Context
//   - pvzIDStr string
func (_e *MockInterface_Expecter) CloseLastReception(ctx interface{}, pvzIDStr interface{}) *MockInterface_CloseLastReception_Call {
	return &MockInterface_CloseLastReception_Call{Call: _e.mock.On("CloseLastReception", ctx, pvzIDStr)}
}

func (_c *MockInterface_CloseLastReception_Call) Run(run func(ctx context.Context, pvzIDStr string)) *MockInterface_CloseLastReception_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockInterface_CloseLastReception_Call) Return(_a0 *models.Receptions, _a1 error) *MockInterface_CloseLastReception_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockInterface_CloseLastReception_Call) RunAndReturn(run func(context.Context, string) (*models.Receptions, error)) *MockInterface_CloseLastReception_Call {
	_c.Call.Return(run)
	return _c
}

// CreatePVZ provides a mock function with given fields: ctx, newPVZ
func (_m *MockInterface) CreatePVZ(ctx context.Context, newPVZ *models.PVZ) (*models.PVZ, error) {
	ret := _m.Called(ctx, newPVZ)

	if len(ret) == 0 {
		panic("no return value specified for CreatePVZ")
	}

	var r0 *models.PVZ
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *models.PVZ) (*models.PVZ, error)); ok {
		return rf(ctx, newPVZ)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *models.PVZ) *models.PVZ); ok {
		r0 = rf(ctx, newPVZ)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.PVZ)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *models.PVZ) error); ok {
		r1 = rf(ctx, newPVZ)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockInterface_CreatePVZ_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreatePVZ'
type MockInterface_CreatePVZ_Call struct {
	*mock.Call
}

// CreatePVZ is a helper method to define mock.On call
//   - ctx context.Context
//   - newPVZ *models.PVZ
func (_e *MockInterface_Expecter) CreatePVZ(ctx interface{}, newPVZ interface{}) *MockInterface_CreatePVZ_Call {
	return &MockInterface_CreatePVZ_Call{Call: _e.mock.On("CreatePVZ", ctx, newPVZ)}
}

func (_c *MockInterface_CreatePVZ_Call) Run(run func(ctx context.Context, newPVZ *models.PVZ)) *MockInterface_CreatePVZ_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*models.PVZ))
	})
	return _c
}

func (_c *MockInterface_CreatePVZ_Call) Return(_a0 *models.PVZ, _a1 error) *MockInterface_CreatePVZ_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockInterface_CreatePVZ_Call) RunAndReturn(run func(context.Context, *models.PVZ) (*models.PVZ, error)) *MockInterface_CreatePVZ_Call {
	_c.Call.Return(run)
	return _c
}

// CreateReceptions provides a mock function with given fields: ctx, newReceptions
func (_m *MockInterface) CreateReceptions(ctx context.Context, newReceptions *models.Receptions) (*models.Receptions, error) {
	ret := _m.Called(ctx, newReceptions)

	if len(ret) == 0 {
		panic("no return value specified for CreateReceptions")
	}

	var r0 *models.Receptions
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *models.Receptions) (*models.Receptions, error)); ok {
		return rf(ctx, newReceptions)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *models.Receptions) *models.Receptions); ok {
		r0 = rf(ctx, newReceptions)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Receptions)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *models.Receptions) error); ok {
		r1 = rf(ctx, newReceptions)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockInterface_CreateReceptions_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateReceptions'
type MockInterface_CreateReceptions_Call struct {
	*mock.Call
}

// CreateReceptions is a helper method to define mock.On call
//   - ctx context.Context
//   - newReceptions *models.Receptions
func (_e *MockInterface_Expecter) CreateReceptions(ctx interface{}, newReceptions interface{}) *MockInterface_CreateReceptions_Call {
	return &MockInterface_CreateReceptions_Call{Call: _e.mock.On("CreateReceptions", ctx, newReceptions)}
}

func (_c *MockInterface_CreateReceptions_Call) Run(run func(ctx context.Context, newReceptions *models.Receptions)) *MockInterface_CreateReceptions_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*models.Receptions))
	})
	return _c
}

func (_c *MockInterface_CreateReceptions_Call) Return(_a0 *models.Receptions, _a1 error) *MockInterface_CreateReceptions_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockInterface_CreateReceptions_Call) RunAndReturn(run func(context.Context, *models.Receptions) (*models.Receptions, error)) *MockInterface_CreateReceptions_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteLastProduct provides a mock function with given fields: ctx, pvzIDStr
func (_m *MockInterface) DeleteLastProduct(ctx context.Context, pvzIDStr string) error {
	ret := _m.Called(ctx, pvzIDStr)

	if len(ret) == 0 {
		panic("no return value specified for DeleteLastProduct")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, pvzIDStr)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockInterface_DeleteLastProduct_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteLastProduct'
type MockInterface_DeleteLastProduct_Call struct {
	*mock.Call
}

// DeleteLastProduct is a helper method to define mock.On call
//   - ctx context.Context
//   - pvzIDStr string
func (_e *MockInterface_Expecter) DeleteLastProduct(ctx interface{}, pvzIDStr interface{}) *MockInterface_DeleteLastProduct_Call {
	return &MockInterface_DeleteLastProduct_Call{Call: _e.mock.On("DeleteLastProduct", ctx, pvzIDStr)}
}

func (_c *MockInterface_DeleteLastProduct_Call) Run(run func(ctx context.Context, pvzIDStr string)) *MockInterface_DeleteLastProduct_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockInterface_DeleteLastProduct_Call) Return(_a0 error) *MockInterface_DeleteLastProduct_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockInterface_DeleteLastProduct_Call) RunAndReturn(run func(context.Context, string) error) *MockInterface_DeleteLastProduct_Call {
	_c.Call.Return(run)
	return _c
}

// GetAllPVZ provides a mock function with given fields: ctx, listInfo
func (_m *MockInterface) GetAllPVZ(ctx context.Context, listInfo models.GetAllPVZRequest) ([]models.PVZWithReceptions, error) {
	ret := _m.Called(ctx, listInfo)

	if len(ret) == 0 {
		panic("no return value specified for GetAllPVZ")
	}

	var r0 []models.PVZWithReceptions
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, models.GetAllPVZRequest) ([]models.PVZWithReceptions, error)); ok {
		return rf(ctx, listInfo)
	}
	if rf, ok := ret.Get(0).(func(context.Context, models.GetAllPVZRequest) []models.PVZWithReceptions); ok {
		r0 = rf(ctx, listInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.PVZWithReceptions)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, models.GetAllPVZRequest) error); ok {
		r1 = rf(ctx, listInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockInterface_GetAllPVZ_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAllPVZ'
type MockInterface_GetAllPVZ_Call struct {
	*mock.Call
}

// GetAllPVZ is a helper method to define mock.On call
//   - ctx context.Context
//   - listInfo models.GetAllPVZRequest
func (_e *MockInterface_Expecter) GetAllPVZ(ctx interface{}, listInfo interface{}) *MockInterface_GetAllPVZ_Call {
	return &MockInterface_GetAllPVZ_Call{Call: _e.mock.On("GetAllPVZ", ctx, listInfo)}
}

func (_c *MockInterface_GetAllPVZ_Call) Run(run func(ctx context.Context, listInfo models.GetAllPVZRequest)) *MockInterface_GetAllPVZ_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(models.GetAllPVZRequest))
	})
	return _c
}

func (_c *MockInterface_GetAllPVZ_Call) Return(_a0 []models.PVZWithReceptions, _a1 error) *MockInterface_GetAllPVZ_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockInterface_GetAllPVZ_Call) RunAndReturn(run func(context.Context, models.GetAllPVZRequest) ([]models.PVZWithReceptions, error)) *MockInterface_GetAllPVZ_Call {
	_c.Call.Return(run)
	return _c
}

// Login provides a mock function with given fields: ctx, newUser
func (_m *MockInterface) Login(ctx context.Context, newUser *models.User) (string, error) {
	ret := _m.Called(ctx, newUser)

	if len(ret) == 0 {
		panic("no return value specified for Login")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *models.User) (string, error)); ok {
		return rf(ctx, newUser)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *models.User) string); ok {
		r0 = rf(ctx, newUser)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, *models.User) error); ok {
		r1 = rf(ctx, newUser)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockInterface_Login_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Login'
type MockInterface_Login_Call struct {
	*mock.Call
}

// Login is a helper method to define mock.On call
//   - ctx context.Context
//   - newUser *models.User
func (_e *MockInterface_Expecter) Login(ctx interface{}, newUser interface{}) *MockInterface_Login_Call {
	return &MockInterface_Login_Call{Call: _e.mock.On("Login", ctx, newUser)}
}

func (_c *MockInterface_Login_Call) Run(run func(ctx context.Context, newUser *models.User)) *MockInterface_Login_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*models.User))
	})
	return _c
}

func (_c *MockInterface_Login_Call) Return(_a0 string, _a1 error) *MockInterface_Login_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockInterface_Login_Call) RunAndReturn(run func(context.Context, *models.User) (string, error)) *MockInterface_Login_Call {
	_c.Call.Return(run)
	return _c
}

// Register provides a mock function with given fields: ctx, newUser
func (_m *MockInterface) Register(ctx context.Context, newUser *models.User) (string, error) {
	ret := _m.Called(ctx, newUser)

	if len(ret) == 0 {
		panic("no return value specified for Register")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *models.User) (string, error)); ok {
		return rf(ctx, newUser)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *models.User) string); ok {
		r0 = rf(ctx, newUser)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, *models.User) error); ok {
		r1 = rf(ctx, newUser)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockInterface_Register_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Register'
type MockInterface_Register_Call struct {
	*mock.Call
}

// Register is a helper method to define mock.On call
//   - ctx context.Context
//   - newUser *models.User
func (_e *MockInterface_Expecter) Register(ctx interface{}, newUser interface{}) *MockInterface_Register_Call {
	return &MockInterface_Register_Call{Call: _e.mock.On("Register", ctx, newUser)}
}

func (_c *MockInterface_Register_Call) Run(run func(ctx context.Context, newUser *models.User)) *MockInterface_Register_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*models.User))
	})
	return _c
}

func (_c *MockInterface_Register_Call) Return(_a0 string, _a1 error) *MockInterface_Register_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockInterface_Register_Call) RunAndReturn(run func(context.Context, *models.User) (string, error)) *MockInterface_Register_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockInterface creates a new instance of MockInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockInterface {
	mock := &MockInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

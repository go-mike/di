package di

import (
	"errors"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testStructWithFields struct {
	Field1 int
	Field2 string
	Field3 []bool
}
func (*testStructWithFields) String() string {
	return "I have fields"
}

type testStructWithFieldsProvider struct{}
var testBoolSlice = []bool{true, false}
func (*testStructWithFieldsProvider) GetService(serviceType reflect.Type) (any, error) {
	if serviceType == reflect.TypeOf(int(0)) {
		return 42, nil
	} else if serviceType == reflect.TypeOf(string("")) {
		return "hello", nil
	} else if serviceType == reflect.TypeOf(testBoolSlice) {
		return testBoolSlice, nil
	}
	panic("unexpected")
}

var expectedStructWithFields testStructWithFields = testStructWithFields{
	Field1: 42,
	Field2: "hello",
	Field3: testBoolSlice,
}
var expectedFieldRequirements []reflect.Type = []reflect.Type{
	reflect.TypeOf(0),
	reflect.TypeOf(""),
	reflect.TypeOf(testBoolSlice),
}

type testStructWithFailProvider struct{}

var errTestFailProvider = errors.New("test fail provider")

func (*testStructWithFailProvider) GetService(serviceType reflect.Type) (any, error) {
	return nil, errTestFailProvider
}

func TestActivateStructFactoryForType_OnNil(t *testing.T) {
	actual, err := ActivateStructFactoryForType(nil)
	assert.Nil(t, actual)
	assert.Equal(t, ErrInvalidStructType, err)
}

func TestActivateStructFactoryForType_OnNonStruct(t *testing.T) {
	actual, err := ActivateStructFactoryForType(reflect.TypeOf(1))
	assert.Nil(t, actual)
	assert.Equal(t, ErrInvalidStructType, err)
}

func TestActivateStructFactoryForType(t *testing.T) {
	factory, err := ActivateStructFactoryForType(reflect.TypeOf(testStructWithFields{}))
	assert.Nil(t, err)
	assert.NotNil(t, factory)
	service, err := factory(&testStructWithFieldsProvider{})
	assert.Nil(t, err)
	assert.NotNil(t, service.Instance)
	assert.NotNil(t, service.Disposable)
	assert.IsType(t, &testStructWithFields{}, service.Instance)
	assert.Equal(t, &expectedStructWithFields, service.Instance)
}

func TestActivateStructFactoryForType_OnFailingProvider(t *testing.T) {
	factory, err := ActivateStructFactoryForType(reflect.TypeOf(testStructWithFields{}))
	assert.Nil(t, err)
	assert.NotNil(t, factory)
	service, err := factory(&testStructWithFailProvider{})
	assert.Equal(t, errTestFailProvider, err)
	assert.Nil(t, service.Instance)
	assert.Nil(t, service.Disposable)
}

func TestActivateStructSimpleFactoryForType_OnNil(t *testing.T) {
	actual, err := ActivateStructSimpleFactoryForType(nil)
	assert.Nil(t, actual)
	assert.Equal(t, ErrInvalidStructType, err)
}

func TestActivateStructSimpleFactoryForType_OnNonStruct(t *testing.T) {
	actual, err := ActivateStructSimpleFactoryForType(reflect.TypeOf(1))
	assert.Nil(t, actual)
	assert.Equal(t, ErrInvalidStructType, err)
}

func TestActivateStructSimpleFactoryForType(t *testing.T) {
	factory, err := ActivateStructSimpleFactoryForType(reflect.TypeOf(testStructWithFields{}))
	assert.Nil(t, err)
	assert.NotNil(t, factory)
	service, err := factory(&testStructWithFieldsProvider{})
	assert.Nil(t, err)
	assert.NotNil(t, service)
	assert.IsType(t, &testStructWithFields{}, service)
	assert.Equal(t, &expectedStructWithFields, service)
}

func TestActivateStructSimpleFactoryForType_OnFailingProvider(t *testing.T) {
	factory, err := ActivateStructSimpleFactoryForType(reflect.TypeOf(testStructWithFields{}))
	assert.Nil(t, err)
	assert.NotNil(t, factory)
	service, err := factory(&testStructWithFailProvider{})
	assert.Equal(t, errTestFailProvider, err)
	assert.Nil(t, service)
}

func TestActivateStructFactory_OnNonStruct(t *testing.T) {
	actual, err := ActivateStructFactory[int64]()
	assert.Nil(t, actual)
	assert.Equal(t, ErrInvalidStructType, err)
}

func TestActivateStructFactory(t *testing.T) {
	factory, err := ActivateStructFactory[testStructWithFields]()
	assert.Nil(t, err)
	assert.NotNil(t, factory)
	service, err := factory(&testStructWithFieldsProvider{})
	assert.Nil(t, err)
	assert.NotNil(t, service)
	assert.IsType(t, &testStructWithFields{}, service)
	assert.Equal(t, &expectedStructWithFields, service)
}

func TestActivateStructFactory_OnFailingProvider(t *testing.T) {
	factory, err := ActivateStructFactory[testStructWithFields]()
	assert.Nil(t, err)
	assert.NotNil(t, factory)
	service, err := factory(&testStructWithFailProvider{})
	assert.Equal(t, errTestFailProvider, err)
	assert.Nil(t, service)
}

func TestActivateStructForType_OnNil(t *testing.T) {
	service, err := ActivateStructForType(nil, nil)
	assert.Nil(t, service.Instance)
	assert.Equal(t, ErrInvalidStructType, err)
}

func TestActivateStructForType_OnNonStruct(t *testing.T) {
	service, err := ActivateStructForType(reflect.TypeOf(1), nil)
	assert.Nil(t, service.Instance)
	assert.Equal(t, ErrInvalidStructType, err)
}

func TestActivateStructForType(t *testing.T) {
	service, err := ActivateStructForType(
		reflect.TypeOf(testStructWithFields{}),
		&testStructWithFieldsProvider{})
	assert.Nil(t, err)
	assert.NotNil(t, service.Instance)
	assert.NotNil(t, service.Disposable)
	assert.IsType(t, &testStructWithFields{}, service.Instance)
	assert.Equal(t, &expectedStructWithFields, service.Instance)
}

func TestActivateStructForType_OnFailingProvider(t *testing.T) {
	service, err := ActivateStructForType(
		reflect.TypeOf(testStructWithFields{}),
		&testStructWithFailProvider{})
	assert.Equal(t, errTestFailProvider, err)
	assert.Nil(t, service.Instance)
	assert.Nil(t, service.Disposable)
}

func TestActivateStructSimple_OnNil(t *testing.T) {
	service, err := ActivateStructSimple(nil, nil)
	assert.Nil(t, service)
	assert.Equal(t, ErrInvalidStructType, err)
}

func TestActivateStructSimple_OnNonStruct(t *testing.T) {
	service, err := ActivateStructSimple(reflect.TypeOf(1), nil)
	assert.Nil(t, service)
	assert.Equal(t, ErrInvalidStructType, err)
}

func TestActivateStructSimple(t *testing.T) {
	service, err := ActivateStructSimple(
		reflect.TypeOf(testStructWithFields{}),
		&testStructWithFieldsProvider{})
	assert.Nil(t, err)
	assert.NotNil(t, service)
	assert.IsType(t, &testStructWithFields{}, service)
	assert.Equal(t, &expectedStructWithFields, service)
}

func TestActivateStructSimple_OnFailingProvider(t *testing.T) {
	service, err := ActivateStructSimple(
		reflect.TypeOf(testStructWithFields{}),
		&testStructWithFailProvider{})
	assert.Equal(t, errTestFailProvider, err)
	assert.Nil(t, service)
	assert.Nil(t, service)
}

func TestActivateStruct_OnNonStruct(t *testing.T) {
	service, err := ActivateStruct[int64](nil)
	assert.Nil(t, service)
	assert.Equal(t, ErrInvalidStructType, err)
}

func TestActivateStruct(t *testing.T) {
	service, err := ActivateStruct[testStructWithFields](&testStructWithFieldsProvider{})
	assert.Nil(t, err)
	assert.NotNil(t, service)
	assert.IsType(t, &testStructWithFields{}, service)
	assert.Equal(t, &expectedStructWithFields, service)
}

func TestActivateStruct_OnFailingProvider(t *testing.T) {
	service, err := ActivateStruct[testStructWithFields](&testStructWithFailProvider{})
	assert.Equal(t, errTestFailProvider, err)
	assert.Nil(t, service)
	assert.Nil(t, service)
}

func testFuncFactoryNoResult(field1 int, field2 string, field3 []bool) {
}
func testFuncFactoryNoError(field1 int, field2 string, field3 []bool) *testStructWithFields {
	return &testStructWithFields{
		Field1: field1,
		Field2: field2,
		Field3: field3,
	}
}
func testFuncFactoryWithError(field1 int, field2 string, field3 []bool) (*testStructWithFields, error) {
	return &testStructWithFields{
		Field1: field1,
		Field2: field2,
		Field3: field3,
	}, nil
}
func testFuncFactoryWithWrongResults(field1 int, field2 string, field3 []bool) (error, *testStructWithFields) {
	return nil, &testStructWithFields{
		Field1: field1,
		Field2: field2,
		Field3: field3,
	}
}
func testFuncFactoryWithManyResults(field1 int, field2 string, field3 []bool) (*testStructWithFields, string, error) {
	return &testStructWithFields{
		Field1: field1,
		Field2: field2,
		Field3: field3,
	}, "", nil
}

var errTestFailFactory = errors.New("test fail factory")

func testFuncFactoryWithFail(field1 int, field2 string, field3 []bool) (*testStructWithFields, error) {
	return nil, errTestFailFactory
}

func TestActivateFuncFactoryForType_OnNil(t *testing.T) {
	factory, err := ActivateFuncFactoryForType(nil)
	assert.Nil(t, factory)
	assert.Equal(t, ErrInvalidFuncType, err)
}

func TestActivateFuncFactoryForType_OnNonFunction(t *testing.T) {
	factory, err := ActivateFuncFactoryForType(&testStructWithFields{})
	assert.Nil(t, factory)
	assert.Equal(t, ErrInvalidFuncType, err)
}

func TestActivateFuncFactoryForType_OnNoResults(t *testing.T) {
	factory, err := ActivateFuncFactoryForType(testFuncFactoryNoResult)
	assert.Nil(t, factory)
	assert.Equal(t, ErrInvalidFuncResults, err)
}

func TestActivateFuncFactoryForType_OnTooManyResults(t *testing.T) {
	factory, err := ActivateFuncFactoryForType(testFuncFactoryWithManyResults)
	assert.Nil(t, factory)
	assert.Equal(t, ErrInvalidFuncResults, err)
}

func TestActivateFuncFactoryForType_OnWrongResults(t *testing.T) {
	factory, err := ActivateFuncFactoryForType(testFuncFactoryWithWrongResults)
	assert.Nil(t, factory)
	assert.Equal(t, ErrInvalidFuncResults, err)
}

func TestActivateFuncFactoryForType_OnNoErrorResult(t *testing.T) {
	factory, err := ActivateFuncFactoryForType(testFuncFactoryNoError)
	assert.Nil(t, err)
	assert.NotNil(t, factory)
	service, err := factory(&testStructWithFieldsProvider{})
	assert.Nil(t, err)
	assert.NotNil(t, service.Instance)
	assert.NotNil(t, service.Disposable)
	assert.IsType(t, &testStructWithFields{}, service.Instance)
	assert.Equal(t, &expectedStructWithFields, service.Instance)
}

func TestActivateFuncFactoryForType_OnErrorResult(t *testing.T) {
	factory, err := ActivateFuncFactoryForType(testFuncFactoryWithError)
	assert.Nil(t, err)
	assert.NotNil(t, factory)
	service, err := factory(&testStructWithFieldsProvider{})
	assert.Nil(t, err)
	assert.NotNil(t, service.Instance)
	assert.NotNil(t, service.Disposable)
	assert.IsType(t, &testStructWithFields{}, service.Instance)
	assert.Equal(t, &expectedStructWithFields, service.Instance)
}

func TestActivateFuncFactoryForType_WithFailedResult(t *testing.T) {
	factory, err := ActivateFuncFactoryForType(testFuncFactoryWithFail)
	assert.Nil(t, err)
	assert.NotNil(t, factory)
	service, err := factory(&testStructWithFieldsProvider{})
	assert.Equal(t, errTestFailFactory, err)
	assert.Nil(t, service.Instance)
	assert.Nil(t, service.Disposable)
}

func TestActivateFuncFactoryForType_WithFailedProvider(t *testing.T) {
	factory, err := ActivateFuncFactoryForType(testFuncFactoryWithFail)
	assert.Nil(t, err)
	assert.NotNil(t, factory)
	service, err := factory(&testStructWithFailProvider{})
	assert.Equal(t, errTestFailProvider, err)
	assert.Nil(t, service.Instance)
	assert.Nil(t, service.Disposable)
}

func TestActivateFuncSimpleFactoryForType_OnNil(t *testing.T) {
	factory, err := ActivateFuncSimpleFactoryForType(nil)
	assert.Nil(t, factory)
	assert.Equal(t, ErrInvalidFuncType, err)
}

func TestActivateFuncSimpleFactoryForType(t *testing.T) {
	factory, err := ActivateFuncSimpleFactoryForType(testFuncFactoryWithError)
	assert.Nil(t, err)
	assert.NotNil(t, factory)
	service, err := factory(&testStructWithFieldsProvider{})
	assert.Nil(t, err)
	assert.NotNil(t, service)
	assert.IsType(t, &testStructWithFields{}, service)
	assert.Equal(t, &expectedStructWithFields, service)
}

func TestActivateFuncSimpleFactoryForType_WithFailedProvider(t *testing.T) {
	factory, err := ActivateFuncSimpleFactoryForType(testFuncFactoryWithFail)
	assert.Nil(t, err)
	assert.NotNil(t, factory)
	service, err := factory(&testStructWithFailProvider{})
	assert.Equal(t, errTestFailProvider, err)
	assert.Nil(t, service)
}

func TestActivateFuncFactory_OnNil(t *testing.T) {
	factory, err := ActivateFuncFactory[*testStructWithFields](nil)
	assert.Nil(t, factory)
	assert.Equal(t, ErrInvalidFuncType, err)
}

func TestActivateFuncFactory(t *testing.T) {
	factory, err := ActivateFuncFactory[*testStructWithFields](testFuncFactoryWithError)
	assert.Nil(t, err)
	assert.NotNil(t, factory)
	service, err := factory(&testStructWithFieldsProvider{})
	assert.Nil(t, err)
	assert.NotNil(t, service)
	assert.IsType(t, &testStructWithFields{}, service)
	assert.Equal(t, &expectedStructWithFields, service)
}

func TestActivateFuncFactory_WithFailedProvider(t *testing.T) {
	factory, err := ActivateFuncFactory[*testStructWithFields](testFuncFactoryWithFail)
	assert.Nil(t, err)
	assert.NotNil(t, factory)
	service, err := factory(&testStructWithFailProvider{})
	assert.Equal(t, errTestFailProvider, err)
	assert.Nil(t, service)
}

func TestActivateFuncForType_OnNil(t *testing.T) {
	service, err := ActivateFuncForType(nil, nil)
	assert.Nil(t, service.Instance)
	assert.Nil(t, service.Disposable)
	assert.Equal(t, ErrInvalidFuncType, err)
}

func TestActivateFuncForType(t *testing.T) {
	service, err := ActivateFuncForType(
		testFuncFactoryWithError,
		&testStructWithFieldsProvider{})
	assert.Nil(t, err)
	assert.NotNil(t, service.Instance)
	assert.NotNil(t, service.Disposable)
	assert.IsType(t, &testStructWithFields{}, service.Instance)
	assert.Equal(t, &expectedStructWithFields, service.Instance)
}

func TestActivateFuncForType_WithFailedProvider(t *testing.T) {
	service, err := ActivateFuncForType(
		testFuncFactoryWithFail,
		&testStructWithFailProvider{})
	assert.Equal(t, errTestFailProvider, err)
	assert.Nil(t, service.Instance)
	assert.Nil(t, service.Disposable)
}

func TestActivateFuncSimple_OnNil(t *testing.T) {
	service, err := ActivateFuncSimple(nil, nil)
	assert.Nil(t, service)
	assert.Equal(t, ErrInvalidFuncType, err)
}

func TestActivateFuncSimple(t *testing.T) {
	service, err := ActivateFuncSimple(
		testFuncFactoryWithError,
		&testStructWithFieldsProvider{})
	assert.Nil(t, err)
	assert.NotNil(t, service)
	assert.IsType(t, &testStructWithFields{}, service)
	assert.Equal(t, &expectedStructWithFields, service)
}

func TestActivateFuncSimple_WithFailedProvider(t *testing.T) {
	service, err := ActivateFuncSimple(
		testFuncFactoryWithFail,
		&testStructWithFailProvider{})
	assert.Equal(t, errTestFailProvider, err)
	assert.Nil(t, service)
}

func TestActivateFunc_OnNil(t *testing.T) {
	service, err := ActivateFunc[*testStructWithFields](nil, nil)
	assert.Nil(t, service)
	assert.Equal(t, ErrInvalidFuncType, err)
}

func TestActivateFunc(t *testing.T) {
	service, err := ActivateFunc[*testStructWithFields](
		testFuncFactoryWithError,
		&testStructWithFieldsProvider{})
	assert.Nil(t, err)
	assert.NotNil(t, service)
	assert.IsType(t, &testStructWithFields{}, service)
	assert.Equal(t, &expectedStructWithFields, service)
}

func TestActivateFunc_WithFailedProvider(t *testing.T) {
	service, err := ActivateFunc[*testStructWithFields](
		testFuncFactoryWithFail,
		&testStructWithFailProvider{})
	assert.Equal(t, errTestFailProvider, err)
	assert.Nil(t, service)
}

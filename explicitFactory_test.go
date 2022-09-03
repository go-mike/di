package di

import (
	"errors"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrInvalidStructType(t *testing.T) {
	// Given
	expected := "invalid struct type"
	// Then
	assert.Equal(t, expected, ErrInvalidStructType.Error())
}

func TestErrInvalidFuncType(t *testing.T) {
	// Given
	expected := "invalid function type"
	// Then
	assert.Equal(t, expected, ErrInvalidFuncType.Error())
}

func TestErrInvalidFuncResults(t *testing.T) {
	// Given
	expected := "invalid function results"
	// Then
	assert.Equal(t, expected, ErrInvalidFuncResults.Error())
}

type testServiceStruct struct{}

func testFactoryFunc(provider ServiceProvider) (any, error) {
	return &testServiceStruct{}, nil
}

func TestNewExplicitFactory(t *testing.T) {
	// Given
	factoryFunc := testFactoryFunc
	// Given
	requirements := []reflect.Type{
		reflect.TypeOf(int(0)),
		reflect.TypeOf(string("")),
		reflect.TypeOf((*error)(nil)).Elem(),
	}
	// Given
	displayName := "some service name"
	// When
	factory := NewExplicitFactory(factoryFunc, requirements, displayName)
	// Then
	assert.NotNil(t, factory)
	// Then
	assert.Equal(t, displayName, factory.DisplayName())
	// Then
	actualRequirements := factory.Requirements()
	assert.NotSame(t, requirements, actualRequirements)
	assert.Equal(t, requirements, actualRequirements)
	// When
	service, err := factory.Create(nil)
	// Then
	assert.Nil(t, err)
	// Then
	assert.NotNil(t, service.Instance)
	// Then
	assert.NotNil(t, service.Disposable)
	// Then
	assert.IsType(t, &testServiceStruct{}, service.Instance)
}

func TestNewFactory(t *testing.T) {
	// Given
	factoryFunc := testFactoryFunc
	// When
	factory := NewFactory(factoryFunc)
	// Then
	assert.NotNil(t, factory)
	// Then
	assert.Equal(t, "<Factory>", factory.DisplayName())
	// Then
	actualRequirements := factory.Requirements()
	assert.Empty(t, actualRequirements)
	// When
	service, err := factory.Create(nil)
	// Then
	assert.Nil(t, err)
	// Then
	assert.NotNil(t, service.Instance)
	// Then
	assert.NotNil(t, service.Disposable)
	// Then
	assert.IsType(t, &testServiceStruct{}, service.Instance)
}

type testStructWithFields struct {
	Field1 int
	Field2 string
	Field3 []bool
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
type testStructWithFailProvider struct{}
var ErrTestFailProvider = errors.New("test fail provider")
func (*testStructWithFailProvider) GetService(serviceType reflect.Type) (any, error) {
	return nil, ErrTestFailProvider
}

func TestNewStructFactoryForTypeOnNilType(t *testing.T) {
	// When
	factory, err := NewStructFactoryForType(nil)
	// Then
	assert.Equal(t, ErrInvalidStructType, err)
	// Then
	assert.Nil(t, factory)
}

func TestNewStructFactoryForTypeOnOtherType(t *testing.T) {
	// When
	factory, err := NewStructFactoryForType(reflect.TypeOf(int(0)))
	// Then
	assert.Equal(t, ErrInvalidStructType, err)
	// Then
	assert.Nil(t, factory)
}

func TestNewStructFactoryForType(t *testing.T) {
	// Given
	displayName := "testStructWithFields"
	// Given
	expectedRequirements := []reflect.Type{
		reflect.TypeOf(int(0)),
		reflect.TypeOf(string("")),
		reflect.TypeOf(testBoolSlice),
	}
	// Given
	expectedService := testStructWithFields{
		Field1: 42,
		Field2: "hello",
		Field3: testBoolSlice,
	}
	// When
	factory, err := NewStructFactoryForType(reflect.TypeOf(testStructWithFields{}))
	// Then
	assert.Nil(t, err)
	// Then
	assert.NotNil(t, factory)
	// Then
	assert.Equal(t, displayName, factory.DisplayName())
	// Then
	actualRequirements := factory.Requirements()
	assert.Equal(t, expectedRequirements, actualRequirements)
	// When
	service, err := factory.Create(&testStructWithFieldsProvider{})
	// Then
	assert.Nil(t, err)
	// Then
	assert.NotNil(t, service.Instance)
	// Then
	assert.NotNil(t, service.Disposable)
	// Then
	assert.IsType(t, &testStructWithFields{}, service.Instance)
	// Then
	assert.Equal(t, &expectedService, service.Instance)
}

func TestNewStructFactoryForTypeOnFailingProvider(t *testing.T) {
	// Given
	displayName := "testStructWithFields"
	// Given
	expectedRequirements := []reflect.Type{
		reflect.TypeOf(int(0)),
		reflect.TypeOf(string("")),
		reflect.TypeOf(testBoolSlice),
	}
	// When
	factory, err := NewStructFactoryForType(reflect.TypeOf(testStructWithFields{}))
	// Then
	assert.Nil(t, err)
	// Then
	assert.NotNil(t, factory)
	// Then
	assert.Equal(t, displayName, factory.DisplayName())
	// Then
	actualRequirements := factory.Requirements()
	assert.Equal(t, expectedRequirements, actualRequirements)
	// When
	service, err := factory.Create(&testStructWithFailProvider{})
	// Then
	assert.Equal(t, ErrTestFailProvider, err)
	// Then
	assert.Nil(t, service.Instance)
}

func TestNewStructFactory(t *testing.T) {
	// Given
	displayName := "testStructWithFields"
	// Given
	expectedRequirements := []reflect.Type{
		reflect.TypeOf(int(0)),
		reflect.TypeOf(string("")),
		reflect.TypeOf(testBoolSlice),
	}
	// Given
	expectedService := testStructWithFields{
		Field1: 42,
		Field2: "hello",
		Field3: testBoolSlice,
	}
	// When
	factory, err := NewStructFactory[testStructWithFields]()
	// Then
	assert.Nil(t, err)
	// Then
	assert.NotNil(t, factory)
	// Then
	assert.Equal(t, displayName, factory.DisplayName())
	// Then
	actualRequirements := factory.Requirements()
	assert.Equal(t, expectedRequirements, actualRequirements)
	// When
	service, err := factory.Create(&testStructWithFieldsProvider{})
	// Then
	assert.Nil(t, err)
	// Then
	assert.NotNil(t, service.Instance)
	// Then
	assert.NotNil(t, service.Disposable)
	// Then
	assert.IsType(t, &testStructWithFields{}, service.Instance)
	// Then
	assert.Equal(t, &expectedService, service.Instance)
}

func testFuncFactoryNoResult (field1 int, field2 string, field3 []bool) {
}
func testFuncFactoryNoError (field1 int, field2 string, field3 []bool) *testStructWithFields {
	return &testStructWithFields{
		Field1: field1,
		Field2: field2,
		Field3: field3,
	}
}
func testFuncFactoryWithError (field1 int, field2 string, field3 []bool) (*testStructWithFields, error) {
	return &testStructWithFields{
		Field1: field1,
		Field2: field2,
		Field3: field3,
	}, nil
}
func testFuncFactoryWithWrongResults (field1 int, field2 string, field3 []bool) (error, *testStructWithFields) {
	return nil, &testStructWithFields{
		Field1: field1,
		Field2: field2,
		Field3: field3,
	}
}
func testFuncFactoryWithManyResults (field1 int, field2 string, field3 []bool) (*testStructWithFields, string, error) {
	return &testStructWithFields{
		Field1: field1,
		Field2: field2,
		Field3: field3,
	}, "", nil
}
var ErrTestFailFactory = errors.New("test fail factory")
func testFuncFactoryWithFail (field1 int, field2 string, field3 []bool) (*testStructWithFields, error) {
	return nil, ErrTestFailFactory
}

func TestNewFuncFactoryOnNilType(t *testing.T) {
	// When
	factory, err := NewFuncFactory(nil)
	// Then
	assert.Equal(t, ErrInvalidFuncType, err)
	// Then
	assert.Nil(t, factory)
}

func TestNewFuncFactoryOnOtherType(t *testing.T) {
	// When
	factory, err := NewFuncFactory(&testStructWithFields{})
	// Then
	assert.Equal(t, ErrInvalidFuncType, err)
	// Then
	assert.Nil(t, factory)
}

func TestNewFuncFactoryOnNoResults(t *testing.T) {
	// When
	factory, err := NewFuncFactory(testFuncFactoryNoResult)
	// Then
	assert.Equal(t, ErrInvalidFuncResults, err)
	// Then
	assert.Nil(t, factory)
}

func TestNewFuncFactoryOnWrongResults(t *testing.T) {
	// When
	factory, err := NewFuncFactory(testFuncFactoryWithWrongResults)
	// Then
	assert.Equal(t, ErrInvalidFuncResults, err)
	// Then
	assert.Nil(t, factory)
}

func TestNewFuncFactoryOnManyResults(t *testing.T) {
	// When
	factory, err := NewFuncFactory(testFuncFactoryWithManyResults)
	// Then
	assert.Equal(t, ErrInvalidFuncResults, err)
	// Then
	assert.Nil(t, factory)
}

func TestNewFuncFactoryOnNoError(t *testing.T) {
	// Given
	displayName := "testFuncFactoryNoError"
	// Given
	expectedRequirements := []reflect.Type{
		reflect.TypeOf(int(0)),
		reflect.TypeOf(string("")),
		reflect.TypeOf(testBoolSlice),
	}
	// Given
	expectedService := testStructWithFields{
		Field1: 42,
		Field2: "hello",
		Field3: testBoolSlice,
	}
	// When
	factory, err := NewFuncFactory(testFuncFactoryNoError)
	// Then
	assert.Nil(t, err)
	// Then
	assert.NotNil(t, factory)
	// Then
	assert.Equal(t, displayName, factory.DisplayName())
	// Then
	actualRequirements := factory.Requirements()
	assert.Equal(t, expectedRequirements, actualRequirements)
	// When
	service, err := factory.Create(&testStructWithFieldsProvider{})
	// Then
	assert.Nil(t, err)
	// Then
	assert.NotNil(t, service.Instance)
	// Then
	assert.NotNil(t, service.Disposable)
	// Then
	assert.IsType(t, &testStructWithFields{}, service.Instance)
	// Then
	assert.Equal(t, &expectedService, service.Instance)
}

func TestNewFuncFactoryOnWithError(t *testing.T) {
	// Given
	displayName := "testFuncFactoryWithError"
	// Given
	expectedRequirements := []reflect.Type{
		reflect.TypeOf(int(0)),
		reflect.TypeOf(string("")),
		reflect.TypeOf(testBoolSlice),
	}
	// Given
	expectedService := testStructWithFields{
		Field1: 42,
		Field2: "hello",
		Field3: testBoolSlice,
	}
	// When
	factory, err := NewFuncFactory(testFuncFactoryWithError)
	// Then
	assert.Nil(t, err)
	// Then
	assert.NotNil(t, factory)
	// Then
	assert.Equal(t, displayName, factory.DisplayName())
	// Then
	actualRequirements := factory.Requirements()
	assert.Equal(t, expectedRequirements, actualRequirements)
	// When
	service, err := factory.Create(&testStructWithFieldsProvider{})
	// Then
	assert.Nil(t, err)
	// Then
	assert.NotNil(t, service.Instance)
	// Then
	assert.NotNil(t, service.Disposable)
	// Then
	assert.IsType(t, &testStructWithFields{}, service.Instance)
	// Then
	assert.Equal(t, &expectedService, service.Instance)
}

func TestNewFuncFactoryOnWithFail(t *testing.T) {
	// Given
	displayName := "testFuncFactoryWithFail"
	// Given
	expectedRequirements := []reflect.Type{
		reflect.TypeOf(int(0)),
		reflect.TypeOf(string("")),
		reflect.TypeOf(testBoolSlice),
	}
	// When
	factory, err := NewFuncFactory(testFuncFactoryWithFail)
	// Then
	assert.Nil(t, err)
	// Then
	assert.NotNil(t, factory)
	// Then
	assert.Equal(t, displayName, factory.DisplayName())
	// Then
	actualRequirements := factory.Requirements()
	assert.Equal(t, expectedRequirements, actualRequirements)
	// When
	service, err := factory.Create(&testStructWithFieldsProvider{})
	// Then
	assert.Equal(t, ErrTestFailFactory, err)
	// Then
	assert.Nil(t, service.Instance)
}

func TestNewFuncFactoryOnFailingProvider(t *testing.T) {
	// Given
	displayName := "testFuncFactoryNoError"
	// Given
	expectedRequirements := []reflect.Type{
		reflect.TypeOf(int(0)),
		reflect.TypeOf(string("")),
		reflect.TypeOf(testBoolSlice),
	}
	// When
	factory, err := NewFuncFactory(testFuncFactoryNoError)
	// Then
	assert.Nil(t, err)
	// Then
	assert.NotNil(t, factory)
	// Then
	assert.Equal(t, displayName, factory.DisplayName())
	// Then
	actualRequirements := factory.Requirements()
	assert.Equal(t, expectedRequirements, actualRequirements)
	// When
	service, err := factory.Create(&testStructWithFailProvider{})
	// Then
	assert.Equal(t, ErrTestFailProvider, err)
	// Then
	assert.Nil(t, service.Instance)
}

func TestNewInstanceFactoryOnNil(t *testing.T) {
	// When
	factory, err := newInstanceFactory(nil)
	// Then
	assert.Equal(t, ErrInvalidInstance, err)
	// Then
	assert.Nil(t, factory)
}

type testInstanceNoStringer struct {}
type testInstanceStringer struct {}
func (i *testInstanceStringer) String() string {
	return "ImStringer"
}

func TestNewInstanceFactoryOnNoPtr(t *testing.T) {
	// Given
	instance := testInstanceNoStringer{}
	// When
	factory, err := newInstanceFactory(instance)
	// Then
	assert.Equal(t, ErrInvalidInstance, err)
	// Then
	assert.Nil(t, factory)
}

func TestNewInstanceFactoryOnNoStringer(t *testing.T) {
	// Given
	displayName := "<testInstanceNoStringer Instance>"
	// Given
	expectedRequirements := []reflect.Type{}
	// Given
	instance := &testInstanceNoStringer{}
	// When
	factory, err := newInstanceFactory(instance)
	// Then
	assert.Nil(t, err)
	// Then
	assert.NotNil(t, factory)
	// Then
	assert.Equal(t, displayName, factory.DisplayName())
	// Then
	actualRequirements := factory.Requirements()
	assert.Equal(t, expectedRequirements, actualRequirements)
	// When
	service, err := factory.Create(&testStructWithFieldsProvider{})
	// Then
	assert.Nil(t, err)
	// Then
	assert.NotNil(t, service)
	// Then
	assert.Same(t, instance, service.Instance)
}

func TestNewInstanceFactoryOnStringer(t *testing.T) {
	// Given
	displayName := "ImStringer"
	// Given
	expectedRequirements := []reflect.Type{}
	// Given
	instance := &testInstanceStringer{}
	// When
	factory, err := newInstanceFactory(instance)
	// Then
	assert.Nil(t, err)
	// Then
	assert.NotNil(t, factory)
	// Then
	assert.Equal(t, displayName, factory.DisplayName())
	// Then
	actualRequirements := factory.Requirements()
	assert.Equal(t, expectedRequirements, actualRequirements)
	// When
	service, err := factory.Create(&testStructWithFieldsProvider{})
	// Then
	assert.Nil(t, err)
	// Then
	assert.NotNil(t, service)
	// Then
	assert.Same(t, instance, service.Instance)
}

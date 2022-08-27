package di

import (
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

func testFactoryFunc(provider ServiceProvider) (interface{}, error) {
	return &testServiceStruct{}, nil
}

// func testFactoryFunc(value1 int, value2 string, value3 error) (*testServiceStruct, error) {
// 	panic("unimplemented")
// }

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
	assert.NotNil(t, service)
	// Then
	assert.IsType(t, &testServiceStruct{}, service)
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
	assert.NotNil(t, service)
	// Then
	assert.IsType(t, &testServiceStruct{}, service)
}

type testStructWithFields struct {
	Field1 int
	Field2 string
	Field3 []bool
}
type testStructWithFieldsProvider struct{}
var testBoolSlice = []bool{true, false}
func (*testStructWithFieldsProvider) GetService(serviceType reflect.Type) (interface{}, error) {
	if serviceType == reflect.TypeOf(int(0)) {
		return 42, nil
	} else if serviceType == reflect.TypeOf(string("")) {
		return "hello", nil
	} else if serviceType == reflect.TypeOf(testBoolSlice) {
		return testBoolSlice, nil
	}
	panic("unexpected")
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
	assert.NotNil(t, service)
	// Then
	assert.IsType(t, &testStructWithFields{}, service)
	// Then
	assert.Equal(t, &expectedService, service)
}

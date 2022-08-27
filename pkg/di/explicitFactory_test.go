package di

import (
	"reflect"
	"testing"
)

func TestErrInvalidStructType(t *testing.T) {
	// Given
	expected := "invalid struct type"
	// Then
	if ErrInvalidStructType.Error() != expected {
		t.Error("ErrInvalidStructType.Error() is not ", expected, " but", ErrInvalidStructType.Error())
	}
}

func TestErrInvalidFuncType(t *testing.T) {
	// Given
	expected := "invalid function type"
	// Then
	if ErrInvalidFuncType.Error() != expected {
		t.Error("ErrInvalidFuncType.Error() is not ", expected, " but", ErrInvalidFuncType.Error())
	}
}

func TestErrInvalidFuncResults(t *testing.T) {
	// Given
	expected := "invalid function results"
	// Then
	if ErrInvalidFuncResults.Error() != expected {
		t.Error("ErrInvalidFuncResults.Error() is not ", expected, " but", ErrInvalidFuncResults.Error())
	}
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
	if factory == nil {
		t.Error("factory is nil")
	}
	// Then
	if factory.DisplayName() != displayName {
		t.Error("factory display name is not", displayName, "but", factory.DisplayName())
	}
	// Then
	actualRequirements := factory.Requirements()
	if &actualRequirements == &requirements {
		t.Error("factory requirements is not a copy of requirements")
	}
	if len(actualRequirements) != len(requirements) {
		t.Error("factory requirements length is not", len(requirements), "but", len(actualRequirements))
	}
	for i := range actualRequirements {
		if actualRequirements[i] != requirements[i] {
			t.Error("factory requirements[", i, "] is not", requirements[i], "but", actualRequirements[i])
		}
	}
	// When
	service, err := factory.Create(nil)
	// Then
	if err != nil {
		t.Error("factory.Create returned error", err)
	}
	// Then
	if service == nil {
		t.Error("factory.Create returned nil service")
	}
	// Then
	if service.(*testServiceStruct) == nil {
		t.Error("factory.Create returned nil *testServiceStruct")
	}
}
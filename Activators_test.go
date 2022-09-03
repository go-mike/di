package di

import (
	"errors"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testDummyDisposable struct{}
func (*testDummyDisposable) Dispose() {}
type testDummyNonDisposable struct{}

func TestToDisposable_OnActualDisposable(t *testing.T) {
	actual := &testDummyDisposable{}
	disposable := toDisposable(actual)
	assert.Equal(t, actual, disposable)
}

func TestToDisposable_OnNonDisposable(t *testing.T) {
	actual := &testDummyNonDisposable{}
	disposable := toDisposable(actual)
	assert.NotNil(t, disposable)
	assert.IsType(t, &noopDisposable{}, disposable)
}

func TestToSimpleFactory(t *testing.T) {
	instance := &testDummyDisposable{}
	fullFactory := func(provider ServiceProvider) (ServiceInstance, error) {
		return ServiceInstance{
			Instance: instance,
			Disposable: instance,
		}, nil
	}
	actualFactory := toSimpleFactory(fullFactory)
	assert.NotNil(t, actualFactory)
	actualInstance, err := actualFactory(&testStructWithFieldsProvider{})
	assert.Nil(t, err)
	assert.Equal(t, instance, actualInstance)
}

func TestToSimpleFactory_WithError(t *testing.T) {
	customError := errors.New("error")
	fullFactory := func(provider ServiceProvider) (ServiceInstance, error) {
		return ServiceInstance{}, customError
	}
	actualFactory := toSimpleFactory(fullFactory)
	assert.NotNil(t, actualFactory)
	actualInstance, err := actualFactory(&testStructWithFieldsProvider{})
	assert.Nil(t, actualInstance)
	assert.Equal(t, customError, err)
}

func TestToTypedFactory(t *testing.T) {
	instance := &testDummyDisposable{}
	simpleFactory := func(provider ServiceProvider) (any, error) {
		return instance, nil
	}
	actualFactory := toTypedFactory[testDummyDisposable](simpleFactory)
	assert.NotNil(t, actualFactory)
	actualInstance, err := actualFactory(&testStructWithFieldsProvider{})
	assert.Nil(t, err)
	assert.Equal(t, instance, actualInstance)
}

func TestToTypedFactory_WithError(t *testing.T) {
	customError := errors.New("error")
	simpleFactory := func(provider ServiceProvider) (any, error) {
		return nil, customError
	}
	actualFactory := toTypedFactory[testDummyDisposable](simpleFactory)
	assert.NotNil(t, actualFactory)
	actualInstance, err := actualFactory(&testStructWithFieldsProvider{})
	assert.Nil(t, actualInstance)
	assert.Equal(t, customError, err)
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

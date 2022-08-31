package di

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testDummyDisposable struct{}
func (*testDummyDisposable) Dispose() {}
type testDummyNonDisposable struct{}

func TestToDisposableOnActualDisposable(t *testing.T) {
	actual := &testDummyDisposable{}
	disposable := toDisposable(actual)
	assert.Equal(t, actual, disposable)
}

func TestToDisposableOnNonDisposable(t *testing.T) {
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

func TestToSimpleFactoryWithError(t *testing.T) {
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

func TestToTypedFactoryWithError(t *testing.T) {
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

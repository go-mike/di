package di

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsNil_OnNil(t *testing.T) {
	assert.True(t, isNil(nil))
}

func TestIsNil_OnNilChan(t *testing.T) {
	var channel chan int
	assert.True(t, isNil(channel))
}

func TestIsNil_OnChan(t *testing.T) {
	var channel chan int = make(chan int)
	assert.False(t, isNil(channel))
}

func TestIsNil_OnNilMap(t *testing.T) {
	var aMap map[int]string
	assert.True(t, isNil(aMap))
}

func TestIsNil_OnMap(t *testing.T) {
	var aMap map[int]string = make(map[int]string)
	assert.False(t, isNil(aMap))
}

func TestIsNil_OnNilSlice(t *testing.T) {
	var aSlice []int
	assert.True(t, isNil(aSlice))
}

func TestIsNil_OnSlice(t *testing.T) {
	var aSlice []int = make([]int, 5)
	assert.False(t, isNil(aSlice))
}

func TestIsNil_OnNilFunction(t *testing.T) {
	var function func()
	assert.True(t, isNil(function))
}

func TestIsNil_OnFunction(t *testing.T) {
	var function func() = func() {}
	assert.False(t, isNil(function))
}

func TestIsNil_OnNilInterface(t *testing.T) {
	var intr Disposable
	assert.True(t, isNil(intr))
}

func TestIsNil_OnInterface(t *testing.T) {
	var intr Disposable = &testDummyDisposable{}
	assert.False(t, isNil(intr))
}

func TestIsNil_OnNilAny(t *testing.T) {
	var intr any
	assert.True(t, isNil(intr))
}

func TestIsNil_OnAny(t *testing.T) {
	var intr any = &testDummyDisposable{}
	assert.False(t, isNil(intr))
}

func TestIsNil_OnStructAtInterface(t *testing.T) {
	var intr any = testDummyDisposable{}
	assert.False(t, isNil(intr))
}

func TestIsNil_OnStruct(t *testing.T) {
	var value testDummyDisposable
	assert.False(t, isNil(value))
}


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
			Instance:   instance,
			Disposable: instance,
		}, nil
	}
	actualFactory := toSimpleFactory(fullFactory)
	assert.NotNil(t, actualFactory)
	actualInstance, err := actualFactory(&testStructWithFieldsProvider{})
	assert.NoError(t, err)
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

func TestToTypedFactoryOfPtr(t *testing.T) {
	instance := &testDummyDisposable{}
	simpleFactory := func(provider ServiceProvider) (any, error) {
		return instance, nil
	}
	actualFactory := toTypedFactoryOfPtr[testDummyDisposable](simpleFactory)
	assert.NotNil(t, actualFactory)
	actualInstance, err := actualFactory(&testStructWithFieldsProvider{})
	assert.NoError(t, err)
	assert.Equal(t, instance, actualInstance)
}

func TestToTypedFactoryOfPtr_WithError(t *testing.T) {
	customError := errors.New("error")
	simpleFactory := func(provider ServiceProvider) (any, error) {
		return nil, customError
	}
	actualFactory := toTypedFactoryOfPtr[testDummyDisposable](simpleFactory)
	assert.NotNil(t, actualFactory)
	actualInstance, err := actualFactory(&testStructWithFieldsProvider{})
	assert.Nil(t, actualInstance)
	assert.Equal(t, customError, err)
}


func TestToSimpleFactoryOfPtr(t *testing.T) {
	instance := &testDummyDisposable{}
	fullFactory := func(provider ServiceProvider) (ServiceInstance, error) {
		return ServiceInstance{
			Instance:   instance,
			Disposable: instance,
		}, nil
	}
	actualFactory := toSimpleFactoryOfPtr[testDummyDisposable](fullFactory)
	assert.NotNil(t, actualFactory)
	actualInstance, err := actualFactory(&testStructWithFieldsProvider{})
	assert.NoError(t, err)
	assert.Equal(t, instance, actualInstance)
}

func TestToSimpleFactoryOfPtr_OnError(t *testing.T) {
	expectedError := errors.New("error")
	fullFactory := func(provider ServiceProvider) (ServiceInstance, error) {
		return ServiceInstance{}, expectedError
	}
	actualFactory := toSimpleFactoryOfPtr[testDummyDisposable](fullFactory)
	assert.NotNil(t, actualFactory)
	actualInstance, err := actualFactory(&testStructWithFieldsProvider{})
	assert.Nil(t, actualInstance)
	assert.Equal(t, expectedError, err)
}

func TestToSimpleFactoryOfPtr_OnOtherType(t *testing.T) {
	fullFactory := func(provider ServiceProvider) (ServiceInstance, error) {
		value := &testStructWithFields{}
		return ServiceInstance{
			Instance:   value,
			Disposable: NewNoopDisposable(),
		}, nil
	}
	actualFactory := toSimpleFactoryOfPtr[testDummyDisposable](fullFactory)
	assert.NotNil(t, actualFactory)
	actualInstance, err := actualFactory(&testStructWithFieldsProvider{})
	assert.Nil(t, actualInstance)
	assert.Equal(t, ErrInvalidFuncResultType, err)
}


func TestToSimpleFactoryOf(t *testing.T) {
	instance := &testDummyDisposable{}
	fullFactory := func(provider ServiceProvider) (ServiceInstance, error) {
		return ServiceInstance{
			Instance:   instance,
			Disposable: instance,
		}, nil
	}
	actualFactory := toSimpleFactoryOf[Disposable](fullFactory)
	assert.NotNil(t, actualFactory)
	actualInstance, err := actualFactory(&testStructWithFieldsProvider{})
	assert.NoError(t, err)
	assert.Equal(t, instance, actualInstance)
}

func TestToSimpleFactoryOf_OnError(t *testing.T) {
	expectedError := errors.New("error")
	fullFactory := func(provider ServiceProvider) (ServiceInstance, error) {
		return ServiceInstance{}, expectedError
	}
	actualFactory := toSimpleFactoryOf[Disposable](fullFactory)
	assert.NotNil(t, actualFactory)
	actualInstance, err := actualFactory(&testStructWithFieldsProvider{})
	assert.Nil(t, actualInstance)
	assert.Equal(t, expectedError, err)
}

func TestToSimpleFactoryOf_OnOtherType(t *testing.T) {
	fullFactory := func(provider ServiceProvider) (ServiceInstance, error) {
		value := &testStructWithFields{}
		return ServiceInstance{
			Instance:   value,
			Disposable: NewNoopDisposable(),
		}, nil
	}
	actualFactory := toSimpleFactoryOf[Disposable](fullFactory)
	assert.NotNil(t, actualFactory)
	actualInstance, err := actualFactory(&testStructWithFieldsProvider{})
	assert.Nil(t, actualInstance)
	assert.Equal(t, ErrInvalidFuncResultType, err)
}


func TestToServiceInstanceFactoryFunc_OnDisposable(t *testing.T) {
	instance := &testDummyDisposable{}
	provider := &testStructWithFieldsProvider{}
	factory := func(p ServiceProvider) (any, error) {
		assert.Equal(t, provider, p)
		return instance, nil
	}
	actualFactory := toServiceInstanceFactoryFunc(factory)
	assert.NotNil(t, actualFactory)
	actualInstance, err := actualFactory(provider)
	assert.NoError(t, err)
	assert.Equal(t, instance, actualInstance.Instance)
	assert.Equal(t, instance, actualInstance.Disposable)
}

func TestToServiceInstanceFactoryFunc_OnNonDisposable(t *testing.T) {
	instance := &testStructWithFields{}
	provider := &testStructWithFieldsProvider{}
	factory := func(p ServiceProvider) (any, error) {
		assert.Equal(t, provider, p)
		return instance, nil
	}
	actualFactory := toServiceInstanceFactoryFunc(factory)
	assert.NotNil(t, actualFactory)
	actualInstance, err := actualFactory(provider)
	assert.NoError(t, err)
	assert.Equal(t, instance, actualInstance.Instance)
	assert.NotEqual(t, instance, actualInstance.Disposable)
	assert.NotNil(t, actualInstance.Disposable)
}

func TestToServiceInstanceFactoryFunc_OnError(t *testing.T) {
	expectedError := errors.New("error")
	provider := &testStructWithFieldsProvider{}
	factory := func(p ServiceProvider) (any, error) {
		assert.Equal(t, provider, p)
		return nil, expectedError
	}
	actualFactory := toServiceInstanceFactoryFunc(factory)
	assert.NotNil(t, actualFactory)
	actualInstance, err := actualFactory(provider)
	assert.Nil(t, actualInstance.Instance)
	assert.Nil(t, actualInstance.Disposable)
	assert.Equal(t, expectedError, err)
}

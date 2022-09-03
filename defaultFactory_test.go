package di

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)


type testServiceInterface any
type testServiceStruct struct{}
func testSimpleFactoryFunc(provider ServiceProvider) (any, error) {
	return &testServiceStruct{}, nil
}
func testFactoryFunc(provider ServiceProvider) (ServiceInstance, error) {
	return ServiceInstance{
		Instance: &testServiceStruct{},
		Disposable: NewNoopDisposable(),
	}, nil
}

type testServiceFactoryFunc struct{}
func (*testServiceFactoryFunc) Factory() ServiceFactoryFunc {
	panic("unimplemented")
}
func (*testServiceFactoryFunc) DisplayName() string {
	panic("unimplemented")
}
func (*testServiceFactoryFunc) Requirements() []reflect.Type {
	panic("unimplemented")
}


func TestNewServiceInstanceFactoryWith(t *testing.T) {
	factoryFunc := testFactoryFunc
	displayName := "some service name"

	factory := NewServiceInstanceFactoryWith(displayName, expectedFieldRequirements, factoryFunc)

	assert.NotNil(t, factory)
	assert.Equal(t, displayName, factory.DisplayName())
	assert.Equal(t, expectedFieldRequirements, factory.Requirements())

	instance, err := factory.Factory()(&testStructWithFieldsProvider{})
	assert.Nil(t, err)
	assert.NotNil(t, instance.Instance)
	assert.NotNil(t, instance.Disposable)
	assert.IsType(t, &testServiceStruct{}, instance.Instance)
}


func TestNewFactoryWith(t *testing.T) {
	factoryFunc := testSimpleFactoryFunc
	displayName := "some service name"

	factory := NewFactoryWith(displayName, expectedFieldRequirements, factoryFunc)

	assert.NotNil(t, factory)
	assert.Equal(t, displayName, factory.DisplayName())
	assert.Equal(t, expectedFieldRequirements, factory.Requirements())

	instance, err := factory.Factory()(&testStructWithFieldsProvider{})
	assert.Nil(t, err)
	assert.NotNil(t, instance.Instance)
	assert.NotNil(t, instance.Disposable)
	assert.IsType(t, &testServiceStruct{}, instance.Instance)
}


func TestNewServiceInstanceFactory(t *testing.T) {
	factoryFunc := testFactoryFunc

	factory := NewServiceInstanceFactory(factoryFunc)

	assert.NotNil(t, factory)
	assert.Equal(t, "testFactoryFunc", factory.DisplayName())
	assert.Equal(t, []reflect.Type{}, factory.Requirements())

	instance, err := factory.Factory()(&testStructWithFieldsProvider{})
	assert.Nil(t, err)
	assert.NotNil(t, instance.Instance)
	assert.NotNil(t, instance.Disposable)
	assert.IsType(t, &testServiceStruct{}, instance.Instance)
}


func TestNewFactory(t *testing.T) {
	factoryFunc := testSimpleFactoryFunc

	factory := NewFactory(factoryFunc)

	assert.NotNil(t, factory)
	assert.Equal(t, "testSimpleFactoryFunc", factory.DisplayName())
	assert.Equal(t, []reflect.Type{}, factory.Requirements())

	instance, err := factory.Factory()(&testStructWithFieldsProvider{})
	assert.Nil(t, err)
	assert.NotNil(t, instance.Instance)
	assert.NotNil(t, instance.Disposable)
	assert.IsType(t, &testServiceStruct{}, instance.Instance)
}


func TestNewStructFactoryForType(t *testing.T) {
	factory, err := NewStructFactoryForType(reflect.TypeOf(testStructWithFields{}))

	assert.Nil(t, err)
	assert.NotNil(t, factory)
	assert.Equal(t, "testStructWithFields", factory.DisplayName())
	assert.Equal(t, expectedFieldRequirements, factory.Requirements())

	instance, err := factory.Factory()(&testStructWithFieldsProvider{})
	assert.Nil(t, err)
	assert.NotNil(t, instance.Instance)
	assert.NotNil(t, instance.Disposable)
	assert.Equal(t, &expectedStructWithFields, instance.Instance)
}

func TestNewStructFactoryForType_OnNil(t *testing.T) {
	factory, err := NewStructFactoryForType(nil)

	assert.NotNil(t, err)
	assert.Nil(t, factory)
}


func TestNewStructFactory(t *testing.T) {
	factory, err := NewStructFactory[testStructWithFields]()

	assert.Nil(t, err)
	assert.NotNil(t, factory)
	assert.Equal(t, "testStructWithFields", factory.DisplayName())
	assert.Equal(t, expectedFieldRequirements, factory.Requirements())

	instance, err := factory.Factory()(&testStructWithFieldsProvider{})
	assert.Nil(t, err)
	assert.NotNil(t, instance.Instance)
	assert.NotNil(t, instance.Disposable)
	assert.Equal(t, &expectedStructWithFields, instance.Instance)
}

func TestNewStructFactory_OnNil(t *testing.T) {
	factory, err := NewStructFactory[int64]()

	assert.NotNil(t, err)
	assert.Nil(t, factory)
}


func TestNewFuncFactory(t *testing.T) {
	factory, err := NewFuncFactory(testFuncFactoryWithError)

	assert.Nil(t, err)
	assert.NotNil(t, factory)
	assert.Equal(t, "testFuncFactoryWithError", factory.DisplayName())
	assert.Equal(t, expectedFieldRequirements, factory.Requirements())

	instance, err := factory.Factory()(&testStructWithFieldsProvider{})
	assert.Nil(t, err)
	assert.NotNil(t, instance.Instance)
	assert.NotNil(t, instance.Disposable)
	assert.Equal(t, &expectedStructWithFields, instance.Instance)
}

func TestNewFuncFactory_OnNil(t *testing.T) {
	factory, err := NewFuncFactory(nil)

	assert.NotNil(t, err)
	assert.Nil(t, factory)
}


func TestNewInstanceFactoryWith(t *testing.T) {
	instance := &expectedStructWithFields
	displayName := "some service name"

	factory, err := newInstanceFactoryWith(displayName, instance)

	assert.Nil(t, err)
	assert.NotNil(t, factory)
	assert.Equal(t, displayName, factory.DisplayName())
	assert.Equal(t, []reflect.Type{}, factory.Requirements())

	actualInstance, err := factory.Factory()(&testStructWithFieldsProvider{})
	assert.Nil(t, err)
	assert.NotNil(t, actualInstance.Instance)
	assert.NotNil(t, actualInstance.Disposable)
	assert.Equal(t, &expectedStructWithFields, actualInstance.Instance)
}

func TestNewInstanceFactoryWith_OnInterface(t *testing.T) {
	var instance Disposable = &testDummyDisposable{}
	displayName := "some service name"

	factory, err := newInstanceFactoryWith(displayName, instance)

	assert.Nil(t, err)
	assert.NotNil(t, factory)
	assert.Equal(t, displayName, factory.DisplayName())
	assert.Equal(t, []reflect.Type{}, factory.Requirements())

	actualInstance, err := factory.Factory()(&testStructWithFieldsProvider{})
	assert.Nil(t, err)
	assert.NotNil(t, actualInstance.Instance)
	assert.NotNil(t, actualInstance.Disposable)
	assert.Equal(t, instance, actualInstance.Instance)
}

func TestNewInstanceFactoryWith_OnStruct(t *testing.T) {
	instance := expectedStructWithFields
	displayName := "some service name"

	factory, err := newInstanceFactoryWith(displayName, instance)

	assert.NotNil(t, err)
	assert.Nil(t, factory)
}

func TestNewInstanceFactoryWith_OnNil(t *testing.T) {
	displayName := "some service name"

	factory, err := newInstanceFactoryWith(displayName, nil)

	assert.NotNil(t, err)
	assert.Nil(t, factory)
}


func TestNewInstanceFactory(t *testing.T) {
	instance := &expectedStructWithFields
	displayName := instance.String()

	factory, err := newInstanceFactory(instance)

	assert.Nil(t, err)
	assert.NotNil(t, factory)
	assert.Equal(t, displayName, factory.DisplayName())
	assert.Equal(t, []reflect.Type{}, factory.Requirements())

	actualInstance, err := factory.Factory()(&testStructWithFieldsProvider{})
	assert.Nil(t, err)
	assert.NotNil(t, actualInstance.Instance)
	assert.NotNil(t, actualInstance.Disposable)
	assert.Equal(t, instance, actualInstance.Instance)
}

func TestNewInstanceFactory_NoStringer(t *testing.T) {
	instance := &testDummyDisposable{}
	displayName := "<testDummyDisposable Instance>"

	factory, err := newInstanceFactory(instance)

	assert.Nil(t, err)
	assert.NotNil(t, factory)
	assert.Equal(t, displayName, factory.DisplayName())
	assert.Equal(t, []reflect.Type{}, factory.Requirements())

	actualInstance, err := factory.Factory()(&testStructWithFieldsProvider{})
	assert.Nil(t, err)
	assert.NotNil(t, actualInstance.Instance)
	assert.NotNil(t, actualInstance.Disposable)
	assert.Equal(t, instance, actualInstance.Instance)
}

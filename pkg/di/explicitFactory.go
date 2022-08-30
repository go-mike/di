package di

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"

	"golang.org/x/exp/slices"
)

type explicitFactory struct {
	factoryFunc  ServiceFactoryFunc
	requirements []reflect.Type
	displayName  string
}

// NewDisposableExplicitFactory creates a new service factory from the given factory function.
// parameters:
// 	factoryFunc - the factory function to create the service
// 	requirements - the service's requirements
// 	displayName - the service's display name
// returns:
// 	the new service factory
func NewDisposableExplicitFactory(
	factoryFunc ServiceFactoryFunc,
	requirements []reflect.Type,
	displayName string) ServiceFactory {
	return &explicitFactory {
		factoryFunc:  factoryFunc,
		requirements: requirements,
		displayName:  displayName,
	}
}

// NewExplicitFactory creates a new service factory from the given factory function.
// parameters:
// 	factoryFunc - the factory function to create the service
// 	requirements - the service's requirements
// 	displayName - the service's display name
// returns:
// 	the new service factory
func NewExplicitFactory(
	factoryFunc SimpleServiceFactoryFunc,
	requirements []reflect.Type,
	displayName string) ServiceFactory {
	return NewDisposableExplicitFactory(factoryFunc.toDisposable(), requirements, displayName)
}

// NewDisposableFactory creates a new service factory from the given factory function.
// parameters:
// 	factoryFunc - the factory function to create the service
// returns:
// 	the new service factory
func NewDisposableFactory(factoryFunc ServiceFactoryFunc) ServiceFactory {
	return NewDisposableExplicitFactory(factoryFunc, []reflect.Type{}, "<Factory>")
}

// NewFactory creates a new service factory from the given factory function.
// parameters:
// 	factoryFunc - the factory function to create the service
// returns:
// 	the new service factory
func NewFactory(factoryFunc SimpleServiceFactoryFunc) ServiceFactory {
	return NewDisposableFactory(factoryFunc.toDisposable())
}

// NewStructFactoryForType creates a new service factory from the given struct type.
// parameters:
// 	structType - the struct type to create the service from
// returns:
// 	the new service factory
func NewStructFactoryForType(structType reflect.Type) (ServiceFactory, error) {
	factory, err := ActivateStructFactoryForType(structType)
	if err != nil { return nil, err }

	displayName := structType.Name()

	requirements := rangeMapSlice(0, structType.NumField(),
		func (i int) reflect.Type {
			return structType.Field(i).Type
		})

	return NewDisposableExplicitFactory(factory, requirements, displayName), nil
}

// NewStructFactory creates a new service factory from the given struct type.
// parameters:
// 	structType - the struct type to create the service from
// returns:
// 	the new service factory
func NewStructFactory[T any]() (ServiceFactory, error) {
	return NewStructFactoryForType(reflect.TypeOf((*T)(nil)).Elem())
}

// NewFuncFactory creates a new service factory from the given function.
// parameters:
// 	function - the function to create the service from
// returns:
// 	the new service factory
func NewFuncFactory(function any) (ServiceFactory, error) {
	factory, err := ActivateFuncFactoryForType(function)
	if err != nil { return nil, err }

	funcType := reflect.TypeOf(function)

	valueOfFunc := reflect.ValueOf(function)
	fullName := runtime.FuncForPC(valueOfFunc.Pointer()).Name()
	fullNameSplit := strings.Split(fullName, ".")
	displayName := fullNameSplit[len(fullNameSplit)-1]

	requirements := rangeMapSlice(0, funcType.NumIn(),
		func (i int) reflect.Type {
			return funcType.In(i)
		})

	return NewDisposableExplicitFactory(factory, requirements, displayName), nil
}

// newInstanceFactory creates a new service factory from the given instance.
// Only singletons are supported.
// parameters:
// 	function - the function to create the service from
// returns:
// 	the new service factory
func newInstanceFactory(instance any) (ServiceFactory, error) {
	if instance == nil {
		return nil, ErrInvalidInstance
	}
	instanceType := reflect.TypeOf(instance)
	if instanceType.Kind() != reflect.Ptr {
		return nil, ErrInvalidInstance
	}

	var displayName string
	stringer, ok := instance.(fmt.Stringer)
	if ok {
		displayName = stringer.String()
	} else {
		displayName = fmt.Sprintf("<%s Instance>", instanceType.Elem().Name())
	}

	factory := func (provider ServiceProvider) (any, error) {
		return instance, nil
	}

	return NewExplicitFactory(factory, []reflect.Type{}, displayName), nil
}

// ServiceFactory interface implementation

// Create implements ServiceFactory.Create to create a new instance of the service.
func (fact *explicitFactory) Create(provider ServiceProvider) (ServiceInstance, error) {
	return fact.factoryFunc(provider)
}

// DisplayName implements ServiceFactory.DisplayName to return the service's display name.
func (fact *explicitFactory) DisplayName() string {
	return fact.displayName
}

// Requirements implements ServiceFactory.Requirements to return the service's requirements.
func (fact *explicitFactory) Requirements() []reflect.Type {
	return slices.Clone(fact.requirements)
}

package di

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"

	"golang.org/x/exp/slices"
)

type defaultFactory struct {
	factoryFunc  ServiceFactoryFunc
	requirements []reflect.Type
	displayName  string
}

// ServiceFactory interface implementation

// Create implements ServiceFactory.Create to create a new instance of the service.
func (fact *defaultFactory) Factory() ServiceFactoryFunc {
	return fact.factoryFunc
}

// DisplayName implements ServiceFactory.DisplayName to return the service's display name.
func (fact *defaultFactory) DisplayName() string {
	return fact.displayName
}

// Requirements implements ServiceFactory.Requirements to return the service's requirements.
func (fact *defaultFactory) Requirements() []reflect.Type {
	return slices.Clone(fact.requirements)
}

// NewServiceInstanceFactoryWith creates a new service factory from the given factory function.
// parameters:
// 	factoryFunc - the factory function to create the service
// 	requirements - the service's requirements
// 	displayName - the service's display name
// returns:
// 	the new service factory
func NewServiceInstanceFactoryWith(
	displayName string,
	requirements []reflect.Type,
	factoryFunc ServiceFactoryFunc) ServiceFactory {
	return &defaultFactory {
		factoryFunc:  factoryFunc,
		requirements: requirements,
		displayName:  displayName,
	}
}

// NewFactoryWith creates a new service factory from the given factory function.
// parameters:
// 	factoryFunc - the factory function to create the service
// 	requirements - the service's requirements
// 	displayName - the service's display name
// returns:
// 	the new service factory
func NewFactoryWith(
	displayName string,
	requirements []reflect.Type,
	factoryFunc SimpleServiceFactoryFunc) ServiceFactory {
	return NewServiceInstanceFactoryWith(
		displayName,
		requirements,
		toServiceInstanceFactoryFunc(factoryFunc))
}

func getFunctionName(function any) string {
	valueOfFunc := reflect.ValueOf(function)
	fullName := runtime.FuncForPC(valueOfFunc.Pointer()).Name()
	fullNameSplit := strings.Split(fullName, ".")
	return fullNameSplit[len(fullNameSplit)-1]
}

// NewServiceInstanceFactory creates a new service factory from the given factory function.
// parameters:
// 	factoryFunc - the factory function to create the service
// returns:
// 	the new service factory
func NewServiceInstanceFactory(factoryFunc ServiceFactoryFunc) ServiceFactory {
	return NewServiceInstanceFactoryWith(
		getFunctionName(factoryFunc),
		[]reflect.Type{},
		factoryFunc)
}

// NewFactory creates a new service factory from the given factory function.
// parameters:
// 	factoryFunc - the factory function to create the service
// returns:
// 	the new service factory
func NewFactory(factoryFunc SimpleServiceFactoryFunc) ServiceFactory {
	return NewServiceInstanceFactoryWith(
		getFunctionName(factoryFunc),
		[]reflect.Type{},
		toServiceInstanceFactoryFunc(factoryFunc))
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

	return NewServiceInstanceFactoryWith(displayName, requirements, factory), nil
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

	requirements := rangeMapSlice(0, funcType.NumIn(),
		func (i int) reflect.Type {
			return funcType.In(i)
		})

	return NewServiceInstanceFactoryWith(
		getFunctionName(function),
		requirements, 
		factory), nil
}

// newInstanceFactoryWith creates a new service factory from the given instance.
// Only singletons are supported.
// parameters:
//  displayName - the display name of the service
// 	function - the function to create the service from
// returns:
// 	the new service factory
func newInstanceFactoryWith(displayName string, instance any) (ServiceFactory, error) {
	if instance == nil {
		return nil, ErrInvalidInstance
	}
	instanceType := reflect.TypeOf(instance)
	if instanceType.Kind() != reflect.Ptr {
		return nil, ErrInvalidInstance
	}

	factory := func (provider ServiceProvider) (any, error) {
		return instance, nil
	}

	return NewFactoryWith(displayName, []reflect.Type{}, factory), nil
}

func getInstanceName(instance any) string {
	if isNil(instance) { return "<nil>" }
	stringer, ok := instance.(fmt.Stringer)
	if ok {
		return stringer.String()
	} else {
		instanceType := reflect.TypeOf(instance)
		return fmt.Sprintf("<%s Instance>", instanceType.Elem().Name())
	}
}

// newInstanceFactory creates a new service factory from the given instance.
// Only singletons are supported.
// parameters:
// 	function - the function to create the service from
// returns:
// 	the new service factory
func newInstanceFactory(instance any) (ServiceFactory, error) {
	return newInstanceFactoryWith(getInstanceName(instance), instance)
}

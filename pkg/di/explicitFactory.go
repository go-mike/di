package di

import (
	"errors"
	"fmt"
	"reflect"
	"runtime"
	"strings"

	"golang.org/x/exp/slices"
)

var ErrInvalidStructType = errors.New("invalid struct type")
var ErrInvalidFuncType = errors.New("invalid function type")
var ErrInvalidFuncResults = errors.New("invalid function results")
var ErrInvalidInstance = errors.New("invalid instance")

type explicitFactory struct {
	factoryFunc  ServiceFactoryFunc
	requirements []reflect.Type
	displayName  string
}

// NewExplicitFactory creates a new service factory from the given factory function.
// parameters:
// 	factoryFunc - the factory function to create the service
// 	requirements - the service's requirements
// 	displayName - the service's display name
// returns:
// 	the new service factory
func NewExplicitFactory(factoryFunc ServiceFactoryFunc, requirements []reflect.Type, displayName string) ServiceFactory {
	return &explicitFactory{
		factoryFunc:  factoryFunc,
		requirements: slices.Clone(requirements),
		displayName:  displayName,
	}
}

// NewFactory creates a new service factory from the given factory function.
// parameters:
// 	factoryFunc - the factory function to create the service
// returns:
// 	the new service factory
func NewFactory(factoryFunc ServiceFactoryFunc) ServiceFactory {
	return NewExplicitFactory(factoryFunc, []reflect.Type{}, "<Factory>")
}

// NewStructFactoryForType creates a new service factory from the given struct type.
// parameters:
// 	structType - the struct type to create the service from
// returns:
// 	the new service factory
func NewStructFactoryForType(structType reflect.Type) (ServiceFactory, error) {
	if structType == nil || structType.Kind() != reflect.Struct {
		return nil, ErrInvalidStructType
	}

	displayName := structType.Name()

	numField := structType.NumField()

	requirements := make([]reflect.Type, numField)

	for i := 0; i < numField; i++ {
		requirements[i] = structType.Field(i).Type
	}

	factory := func (provider ServiceProvider) (interface{}, error) {
		result := reflect.New(structType)
		elem := result.Elem()
		for i := 0; i < numField; i++ {
			service, err := provider.GetService(requirements[i])
			if err != nil {
				return nil, err
			}
			elem.Field(i).Set(reflect.ValueOf(service))
		}
		return result.Interface(), nil
	}

	return NewExplicitFactory(factory, requirements, displayName), nil
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
func NewFuncFactory(function interface{}) (ServiceFactory, error) {
	funcType := reflect.TypeOf(function)

	if funcType == nil || funcType.Kind() != reflect.Func {
		return nil, ErrInvalidFuncType
	}

	numResults := funcType.NumOut()

	if numResults < 1 || numResults > 2 {
		return nil, ErrInvalidFuncResults
	}

	if numResults == 2 && funcType.Out(1) != reflect.TypeOf((*error)(nil)).Elem() {
		return nil, ErrInvalidFuncResults
	}

	valueOfFunc := reflect.ValueOf(function)

	fullName := runtime.FuncForPC(valueOfFunc.Pointer()).Name()
	fullNameSplit := strings.Split(fullName, ".")
	displayName := fullNameSplit[len(fullNameSplit)-1]

	numParams := funcType.NumIn()
	requirements := make([]reflect.Type, numParams)

	for i := 0; i < numParams; i++ {
		requirements[i] = funcType.In(i)
	}

	factory := func (provider ServiceProvider) (interface{}, error) {
		args := make([]reflect.Value, numParams)
		for i := 0; i < numParams; i++ {
			service, err := provider.GetService(requirements[i])
			if err != nil {
				return nil, err
			}
			args[i] = reflect.ValueOf(service)
		}

		funcResult := valueOfFunc.Call(args)

		if len(funcResult) == 1 {
			return funcResult[0].Interface(), nil
		} else {
			err, _ := funcResult[1].Interface().(error)
			return funcResult[0].Interface(), err
		}
	}

	return NewExplicitFactory(factory, requirements, displayName), nil
}

// newInstanceFactory creates a new service factory from the given instance.
// Only singletons are supported.
// parameters:
// 	function - the function to create the service from
// returns:
// 	the new service factory
func newInstanceFactory(instance interface{}) (ServiceFactory, error) {
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

	factory := func (provider ServiceProvider) (interface{}, error) {
		return instance, nil
	}

	return NewExplicitFactory(factory, []reflect.Type{}, displayName), nil
}

// ServiceFactory interface implementation

// Create implements ServiceFactory.Create to create a new instance of the service.
func (fact *explicitFactory) Create(provider ServiceProvider) (interface{}, error) {
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

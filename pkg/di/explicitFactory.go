package di

import (
	"errors"
	"reflect"

	"golang.org/x/exp/slices"
)

var ErrInvalidStructType = errors.New("invalid struct type")
var ErrInvalidFuncType = errors.New("invalid function type")
var ErrInvalidFuncResults = errors.New("invalid function results")

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

// NewStructFactory creates a new service factory from the given struct type.
// parameters:
// 	structType - the struct type to create the service from
// returns:
// 	the new service factory
func NewStructFactory(structure interface{}) (ServiceFactory, error) {
	structType := reflect.TypeOf(structure)

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

// NewFuncFactory creates a new service factory from the given struct type.
// parameters:
// 	factoryFunc - the factory function to create the service
// 	requirements - the service's requirements
// 	displayName - the service's display name
// returns:
// 	the new service factory
func NewFuncFactory(function interface{}) (ServiceFactory, error) {
	funcType := reflect.TypeOf(function)

	if funcType == nil || funcType.Kind() != reflect.Func {
		return nil, ErrInvalidFuncType
	}

	displayName := funcType.Name()

	numParams := funcType.NumIn()
	numResults := funcType.NumOut()

	if numResults < 1 || numResults > 2 {
		return nil, ErrInvalidFuncResults
	}

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

		funcResult := reflect.ValueOf(function).Call(args)

		if len(funcResult) == 1 {
			return funcResult[0].Interface(), nil
		} else if len(funcResult) == 2 {
			return funcResult[0].Interface(), funcResult[1].Interface().(error)
		} else {
			return nil, ErrInvalidFuncResults
		}
	}

	return NewExplicitFactory(factory, requirements, displayName), nil
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

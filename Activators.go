package di

import (
	"errors"
	"reflect"
)

var ErrInvalidStructType = errors.New("invalid struct type")
var ErrInvalidFuncType = errors.New("invalid function type")
var ErrInvalidFuncResultType = errors.New("invalid function result type")
var ErrInvalidFuncResults = errors.New("invalid function results")
var ErrInvalidInstance = errors.New("invalid instance")

func typeOf[T any]() reflect.Type {
	return reflect.TypeOf((*T)(nil)).Elem()
}

func ActivateStructFactoryForType(structType reflect.Type) (ServiceFactoryFunc, error) {
	if structType == nil || structType.Kind() != reflect.Struct {
		return nil, ErrInvalidStructType
	}

	numFields := structType.NumField()

	requirements := rangeMapSlice(0, numFields,
		func(i int) reflect.Type {
			return structType.Field(i).Type
		})

	factory := func(provider ServiceProvider) (ServiceInstance, error) {
		result := reflect.New(structType)
		elem := result.Elem()
		for i := 0; i < numFields; i++ {
			service, err := provider.GetService(requirements[i])
			if err != nil {
				return ServiceInstance{}, err
			}
			elem.Field(i).Set(reflect.ValueOf(service))
		}
		instance := result.Interface()
		return ServiceInstance{
			Instance:   instance,
			Disposable: toDisposable(instance),
		}, nil
	}

	return factory, nil
}

func ActivateStructSimpleFactoryForType(structType reflect.Type) (SimpleServiceFactoryFunc, error) {
	fullFunc, err := ActivateStructFactoryForType(structType)
	if err != nil {
		return nil, err
	}
	return toSimpleFactory(fullFunc), nil
}

func ActivateStructFactory[T any]() (SimpleServiceFactoryFuncOfPtr[T], error) {
	simpleFunc, err := ActivateStructSimpleFactoryForType(typeOf[T]())
	if err != nil {
		return nil, err
	}
	return toTypedFactoryOfPtr[T](simpleFunc), nil
}

func ActivateStructForType(structType reflect.Type, provider ServiceProvider) (ServiceInstance, error) {
	factory, err := ActivateStructFactoryForType(structType)
	if err != nil {
		return ServiceInstance{}, err
	}
	return factory(provider)
}

func ActivateStructSimple(structType reflect.Type, provider ServiceProvider) (any, error) {
	factory, err := ActivateStructSimpleFactoryForType(structType)
	if err != nil {
		return nil, err
	}
	return factory(provider)
}

func ActivateStruct[T any](provider ServiceProvider) (*T, error) {
	factory, err := ActivateStructFactory[T]()
	if err != nil {
		return nil, err
	}
	return factory(provider)
}

func ActivateFuncFactoryForType(function any) (ServiceFactoryFunc, error) {
	if function == nil {
		return nil, ErrInvalidFuncType
	}

	funcType := reflect.TypeOf(function)

	if funcType == nil || funcType.Kind() != reflect.Func {
		return nil, ErrInvalidFuncType
	}

	numResults := funcType.NumOut()

	if numResults < 1 || numResults > 2 {
		return nil, ErrInvalidFuncResults
	}

	if numResults == 2 && funcType.Out(1) != typeOf[error]() {
		return nil, ErrInvalidFuncResults
	}

	valueOfFunc := reflect.ValueOf(function)

	numParams := funcType.NumIn()
	requirements := rangeMapSlice(0, numParams,
		func(i int) reflect.Type {
			return funcType.In(i)
		})

	factory := func(provider ServiceProvider) (ServiceInstance, error) {
		args := make([]reflect.Value, numParams)
		for i := 0; i < numParams; i++ {
			service, err := provider.GetService(requirements[i])
			if err != nil {
				return ServiceInstance{}, err
			}
			args[i] = reflect.ValueOf(service)
		}

		funcResult := valueOfFunc.Call(args)

		if len(funcResult) == 1 {
			instance := funcResult[0].Interface()
			return ServiceInstance{
				Instance:   instance,
				Disposable: toDisposable(instance),
			}, nil
		} else {
			err, _ := funcResult[1].Interface().(error)
			instance := funcResult[0].Interface()
			return ServiceInstance{
				Instance:   instance,
				Disposable: toDisposable(instance),
			}, err
		}
	}

	return factory, nil
}

func ActivateFuncSimpleFactoryForType(function any) (SimpleServiceFactoryFunc, error) {
	fullFunc, err := ActivateFuncFactoryForType(function)
	if err != nil {
		return nil, err
	}
	return toSimpleFactory(fullFunc), nil
}

func ActivateFuncFactory[T any](function any) (SimpleServiceFactoryFuncOf[T], error) {
	fullFunc, err := ActivateFuncFactoryForType(function)
	if err != nil {
		return nil, err
	}
	return toSimpleFactoryOf[T](fullFunc), nil
}

func ActivateFuncForType(function any, provider ServiceProvider) (ServiceInstance, error) {
	factory, err := ActivateFuncFactoryForType(function)
	if err != nil {
		return ServiceInstance{}, err
	}
	return factory(provider)
}

func ActivateFuncSimple(function any, provider ServiceProvider) (any, error) {
	factory, err := ActivateFuncSimpleFactoryForType(function)
	if err != nil {
		return nil, err
	}
	return factory(provider)
}

func ActivateFunc[T any](function any, provider ServiceProvider) (T, error) {
	factory, err := ActivateFuncFactory[T](function)
	if err != nil {
		var empty T
		return empty, err
	}
	return factory(provider)
}

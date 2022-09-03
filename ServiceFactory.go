package di

import (
	"reflect"
)

type ServiceInstance struct {
	Instance   any
	Disposable Disposable
}

type ServiceFactoryFunc func(provider ServiceProvider) (ServiceInstance, error)

type SimpleServiceFactoryFunc func(provider ServiceProvider) (any, error)

type SimpleServiceFactoryFuncOf[T any] func(provider ServiceProvider) (T, error)
type SimpleServiceFactoryFuncOfPtr[T any] func(provider ServiceProvider) (*T, error)

type ServiceFactory interface {
	Factory() ServiceFactoryFunc
	Requirements() []reflect.Type
	DisplayName() string
}

func isNil(i any) bool {
	if i == nil {
		return true
	}
	switch reflect.TypeOf(i).Kind() {
	case reflect.Chan:
		return reflect.ValueOf(i).IsNil()
	case reflect.Func:
		return reflect.ValueOf(i).IsNil()
	// case reflect.Interface: // Is this needed/possible?
	// 	return reflect.ValueOf(i).IsNil()
	case reflect.Map:
		return reflect.ValueOf(i).IsNil()
	case reflect.Ptr:
		return reflect.ValueOf(i).IsNil()
	case reflect.Slice:
		return reflect.ValueOf(i).IsNil()
	default:
		return false
	}
}

func toDisposable(service any) Disposable {
	if isNil(service) {
		return nil
	}
	disposable, ok := service.(Disposable)
	if !ok {
		disposable = NewNoopDisposable()
	}
	return disposable
}

func toSimpleFactory(fullFunc ServiceFactoryFunc) SimpleServiceFactoryFunc {
	return func(provider ServiceProvider) (any, error) {
		value, err := fullFunc(provider)
		if err != nil {
			return nil, err
		}
		return value.Instance, nil
	}
}

func toSimpleFactoryOfPtr[T any](fullFunc ServiceFactoryFunc) SimpleServiceFactoryFuncOfPtr[T] {
	return func(provider ServiceProvider) (*T, error) {
		value, err := fullFunc(provider)
		if err != nil {
			return nil, err
		}
		typed, ok := value.Instance.(*T)
		if !ok {
			return nil, ErrInvalidFuncResultType
		}
		return typed, nil
	}
}

func toSimpleFactoryOf[T any](fullFunc ServiceFactoryFunc) SimpleServiceFactoryFuncOf[T] {
	return func(provider ServiceProvider) (T, error) {
		value, err := fullFunc(provider)
		if err != nil {
			var empty T
			return empty, err
		}
		typed, ok := value.Instance.(T)
		if !ok {
			var empty T
			return empty, ErrInvalidFuncResultType
		}
		return typed, nil
	}
}

func toTypedFactoryOfPtr[T any](simpleFunc SimpleServiceFactoryFunc) SimpleServiceFactoryFuncOfPtr[T] {
	return func(provider ServiceProvider) (*T, error) {
		value, err := simpleFunc(provider)
		if err != nil {
			return nil, err
		}
		return value.(*T), nil
	}
}

func toServiceInstanceFactoryFunc(factory SimpleServiceFactoryFunc) ServiceFactoryFunc {
	return func(provider ServiceProvider) (ServiceInstance, error) {
		instance, err := factory(provider)
		if err != nil {
			return ServiceInstance{}, err
		}

		return ServiceInstance{
			Instance:   instance,
			Disposable: toDisposable(instance),
		}, nil
	}
}

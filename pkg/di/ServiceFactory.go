package di

import (
	"reflect"
)

type ServiceInstance struct {
	Instance interface{}
	Disposable Disposable
}

type ServiceFactoryFunc func(provider ServiceProvider) (ServiceInstance, error)

type SimpleServiceFactoryFunc func(provider ServiceProvider) (interface{}, error)

type ServiceFactory interface {
	Create(provider ServiceProvider) (ServiceInstance, error)
	Requirements() []reflect.Type
	DisplayName() string
}

func (f SimpleServiceFactoryFunc) toDisposable() (ServiceFactoryFunc) {
	return func(provider ServiceProvider) (ServiceInstance, error) {
		instance, err := f(provider)
		if err != nil {
			return ServiceInstance{}, err
		}

		disposable, ok := instance.(Disposable)
		if !ok {
			disposable = NewNoopDisposable()
		}

		return ServiceInstance{
			Instance: instance,
			Disposable: disposable,
		}, nil
	}
}

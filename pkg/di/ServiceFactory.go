package di

import (
	"reflect"
)

type ServiceInstance struct {
	Instance any
	Disposable Disposable
}

type ServiceFactoryFunc func(provider ServiceProvider) (ServiceInstance, error)

type SimpleServiceFactoryFunc func(provider ServiceProvider) (any, error)

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

		return ServiceInstance{
			Instance: instance,
			Disposable: toDisposable(instance),
		}, nil
	}
}

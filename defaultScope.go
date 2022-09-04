package di

import (
	"errors"
	"reflect"
)

var ErrServiceScopeDisposed = errors.New("service scope has been disposed")

type defaultScope struct {
	parent      ServiceProvider
	lifetime    Lifetime
	descriptors []ServiceDescriptor
	instances   []ServiceInstance
	isDisposed  bool
}

var _ ServiceScope = (*defaultScope)(nil)
var _ ServiceProvider = (*defaultScope)(nil)

// Provider implements ServiceScope
func (scope *defaultScope) Provider() ServiceProvider {
	return scope
}

// Lifetime implements ServiceScope
func (scope *defaultScope) Lifetime() Lifetime {
	return scope.lifetime
}

// Dispose implements ServiceScope
func (scope *defaultScope) Dispose() {
	if !scope.isDisposed {
		scope.isDisposed = true
		for _, disposable := range scope.instances {
			disposable.Disposable.Dispose()
		}
	}
}

// IsDisposed implements ServiceScope
func (scope *defaultScope) IsDisposed() bool {
	return scope.isDisposed
}

// GetService implements ServiceProvider
func (scope *defaultScope) GetService(serviceType reflect.Type) (any, error) {
	panic("unimplemented")
}

// GetServiceInfo implements ServiceProvider
func (scope *defaultScope) GetServiceInfo(serviceType reflect.Type) ServiceInfo {
	if scope.isDisposed {
		return newNotFoundServiceInfo(serviceType)
	}
	return newNotFoundServiceInfo(serviceType)
}

func newSingletonScope(descriptors []ServiceDescriptor) (*defaultScope, error) {
	err := findValidationErrors(nil, descriptors)
	if err != nil {
		return nil, err
	}

	return &defaultScope{
		lifetime:    Singleton,
		descriptors: descriptors,
		instances:   []ServiceInstance{},
	}, nil
}

func findValidationErrors(parent ServiceProvider, descriptors []ServiceDescriptor) error {
	// TODO: Validate
	return nil
}

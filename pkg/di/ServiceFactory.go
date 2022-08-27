package di

import (
	"reflect"
)

type ServiceFactoryFunc func(provider ServiceProvider) (interface{}, error)

type ServiceFactory interface {
	Create(provider ServiceProvider) (interface{}, error)
	Requirements() []reflect.Type
	DisplayName() string
}

package di

import "reflect"

type ServiceProvider interface {
	GetService(serviceType reflect.Type) (any, error)
}

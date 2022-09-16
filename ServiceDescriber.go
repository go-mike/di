package di

import "reflect"

type ServiceDescriber interface {
	GetServiceDescriptor(serviceType reflect.Type) ServiceDescriptor
}

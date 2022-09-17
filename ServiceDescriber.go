package di

import "reflect"

type ServiceDescriber interface {
	GetServiceDescriptor(serviceType reflect.Type) ServiceDescriptor
	GetServiceDescriptors(serviceType reflect.Type) []ServiceDescriptor
}

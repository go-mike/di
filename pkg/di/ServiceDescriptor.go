package di

import "reflect"

type ServiceDescriptor interface {
	Lifetime() Lifetime
	ServiceType() reflect.Type
	Factory() ServiceFactory
}

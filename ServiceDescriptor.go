package di

import (
	"fmt"
	"reflect"
)

type ServiceDescriptor interface {
	fmt.Stringer
	Lifetime() Lifetime
	ServiceType() reflect.Type
	Factory() ServiceFactory
}

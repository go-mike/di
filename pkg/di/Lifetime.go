package di

import "fmt"

type Lifetime int

const (
	Singleton Lifetime = iota
	Scoped
	Transient
)

var _ fmt.Stringer = Singleton

func (lifetime Lifetime) String() string {
	switch lifetime {
	case Singleton:
		return "Singleton"
	case Scoped:
		return "Scoped"
	case Transient:
		return "Transient"
	default:
		return "Unknown"
	}
}

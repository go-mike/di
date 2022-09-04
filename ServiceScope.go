package di

type ServiceScope interface {
	Provider() ServiceProvider
	Lifetime() Lifetime
	Dispose()
	IsDisposed() bool
}

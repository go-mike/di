package di

// Disposable represents objects that can be disposed.
type Disposable interface {
	Dispose()
}

// NewDisposable returns a new disposable that calls the given dispose function when disposed.
func NewDisposable(dispose func()) Disposable {
	return &disposable{
		dispose: dispose,
	}
}

// disposable is a disposable implementation backed by a dispose function
type disposable struct {
	dispose func()
	disposed bool
}

// Dispose implements Disposable
func (disp *disposable) Dispose() {
	if !disp.disposed {
		disp.dispose()
		disp.disposed = true
	}
}

// NewNoopDisposable returns a new disposable that does nothing when disposed.
func NewNoopDisposable() Disposable {
	return &noopDisposable{}
}

// noopDisposable is a disposable implementation that does nothing when disposed.
type noopDisposable struct{}

// Dispose implements Disposable
func (*noopDisposable) Dispose() {}

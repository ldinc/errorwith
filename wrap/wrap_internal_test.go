package wrap_test

import (
	"errors"

	"errorwith/wrap"
)

var errNotFound = errors.New("not found")

func a() error {
	return wrap.WithCaller(b(c))
}

func b(fn func() error) error {
	return wrap.WithCaller(fn())
}

func c() error {
	return wrap.WithCaller(d())
}

func d() error {
	return wrap.WithCaller(func() error {
		return wrap.WithCaller(errNotFound)
	}(),
	)
}

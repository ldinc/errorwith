package wrap_test

import (
	"errors"

	"errorwith/wrap"
)

var errNotFound = errors.New("not found")

func a() error {
	return wrap.With(b(c))
}

func b(fn func() error) error {
	return wrap.With(fn())
}

func c() error {
	return wrap.With(d())
}

func d() error {
	return wrap.With(func() error {
		return wrap.With(errNotFound)
	}(),
	)
}

package wrap_test

import (
	"errorwith/wrap"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_With(t *testing.T) {
	t.Parallel()

	baseErr := fmt.Errorf("base error")

	t.Run("simple call default", func(t *testing.T) {
		t.Parallel()

		// arrange
		fn := func() error {
			return wrap.With(baseErr)
		}

		// act
		actual := fn()

		// assert
		require.Error(t, actual)
		require.ErrorIs(t, actual, baseErr)
		require.EqualError(t, actual, "wrap/wrap_test.go:21 wrap_test.Test_With.func1.1: base error")
	},
	)

	t.Run("simple call without error source", func(t *testing.T) {
		t.Parallel()

		// arrange
		fn := func() error {
			return wrap.With(baseErr, wrap.NoErrorSource())
		}

		// act
		actual := fn()

		// assert
		require.Error(t, actual)
		require.ErrorIs(t, actual, baseErr)
		require.EqualError(t, actual, baseErr.Error())
	})

	t.Run("simple call with decorate", func(t *testing.T) {
		t.Parallel()

		// arrange
		fn := func() error {
			return wrap.With(
				baseErr,
				wrap.NoErrorSource(),
				wrap.Format("%s some test %d", "jib", 10),
			)
		}

		// act
		actual := fn()

		// assert
		require.Error(t, actual)
		require.ErrorIs(t, actual, baseErr)
		require.EqualError(t, actual, "jib some test 10: "+baseErr.Error())
	})

	t.Run("complex call default", func(t *testing.T) {
		t.Parallel()

		// act
		actual := a()

		// assert
		require.Error(t, actual)
		require.ErrorIs(t, actual, errNotFound)
		require.EqualError(t, actual,
			"wrap/wrap_internal_test.go:12 wrap_test.a: "+
				"wrap/wrap_internal_test.go:16 wrap_test.b: "+
				"wrap/wrap_internal_test.go:20 wrap_test.c: "+
				"wrap/wrap_internal_test.go:24 wrap_test.d: "+
				"wrap/wrap_internal_test.go:25 wrap_test.d.func1: "+
				"not found",
		)
	},
	)
}

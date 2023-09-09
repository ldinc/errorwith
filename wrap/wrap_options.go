package wrap

var _ iBuilder = (*wrapBuilder)(nil)

type WrapOption interface {
	apply(iBuilder)
}

type iBuilder interface {
	WithoutErrorSource()
	WithFormatMessage(string, ...interface{})
}

type noErrorSource struct{}

func (noErrorSource) apply(b iBuilder) {
	b.WithoutErrorSource()
}

func NoErrorSource() WrapOption {
	return noErrorSource{}
}

type formatMessage struct {
	format string
	args   []any
}

func (fm formatMessage) apply(b iBuilder) {
	b.WithFormatMessage(fm.format, fm.args...)
}

func Format(format string, a ...any) WrapOption {
	return &formatMessage{
		format: format,
		args:   a,
	}
}

// https://github.com/grpc/grpc-go/blob/master/dialoptions.go#L86

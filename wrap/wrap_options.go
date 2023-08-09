package wrap

type WrapOption interface {
	apply(*fnInfo)
}

// https://github.com/grpc/grpc-go/blob/master/dialoptions.go#L86

type defaultOptions struct{}

func (defaultOptions) apply(*fnInfo) {}

type dontShowSourceFile struct{}

func (dontShowSourceFile) apply(cfg *fnInfo) {
	cfg.skipSourceFile = true
}

type dontShowSourceFileLineNo struct{}

func (dontShowSourceFileLineNo) apply(cfg *fnInfo) {
	cfg.noLineNo = true
}

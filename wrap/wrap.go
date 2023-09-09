package wrap

import (
	"runtime"
	"strconv"
	"strings"
)

const (
	defaultFnLevel = 3
	dirLen         = 2
	pcLen          = 15
)

type fnInfo struct {
	FuncName string
	FileName string
	LineNum  int
}

func (info fnInfo) GetName() string {
	if info.FuncName == "" {
		return ""
	}

	names := strings.Split(info.FuncName, "/")
	target := names[len(names)-1]

	return target
}

func (info fnInfo) GetSource() string {
	elems := strings.Split(info.FileName, "/")

	if len(elems) > dirLen {
		return strings.Join(elems[len(elems)-dirLen:], "/")
	}

	return info.FileName
}

func (info fnInfo) GetLineNo() int {
	return info.LineNum
}

func (info fnInfo) ToString() string {
	var sb strings.Builder

	sb.WriteString(info.GetSource())
	sb.WriteRune(':')
	sb.WriteString(strconv.Itoa(info.LineNum))
	sb.WriteRune(' ')
	sb.WriteString(info.GetName())

	return sb.String()
}

func extractFnInfo() (fnInfo, bool) {
	var (
		pc       = make([]uintptr, pcLen)
		n        = runtime.Callers(defaultFnLevel, pc)
		frames   = runtime.CallersFrames(pc[:n])
		frame, _ = frames.Next()
	)

	if frame.Function == "" {
		return fnInfo{}, false
	}

	info := fnInfo{
		FuncName: frame.Function,
		FileName: frame.File,
		LineNum:  frame.Line,
	}

	return info, true
}

func With(err error, opts ...WrapOption) error {
	if err == nil {
		return nil
	}

	info, ok := extractFnInfo()

	if !ok {
		return err
	}

	builder := defaultWrapBuilder(err, info)

	for _, opt := range opts {
		opt.apply(builder)
	}

	return builder.Build()
}

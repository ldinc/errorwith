package wrap

import (
	"fmt"
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

	skipSourceFile bool
	noLineNo       bool
}

func (info fnInfo) getName() string {
	if info.FuncName == "" {
		return ""
	}

	names := strings.Split(info.FuncName, "/")
	target := names[len(names)-1]

	return target
}

func (info fnInfo) getSource() string {
	elems := strings.Split(info.FileName, "/")

	if len(elems) > dirLen {
		return strings.Join(elems[len(elems)-dirLen:], "/")
	}

	return info.FileName
}

func (info fnInfo) ToString() string {
	var sb strings.Builder

	if !info.skipSourceFile {
		sb.WriteString(info.getSource())

		if !info.noLineNo {
			sb.WriteRune(':')
			sb.WriteString(strconv.Itoa(info.LineNum))
		}

		sb.WriteRune(' ')
	}

	sb.WriteString(info.getName())

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

func WithCaller(err error, opts ...WrapOption) error {
	if err == nil {
		return nil
	}

	info, ok := extractFnInfo()

	if !ok {
		return err
	}

	for _, opt := range opts {
		opt.apply(&info)
	}

	return fmt.Errorf("%s: %w", info.ToString(), err)
	// TODO: check with bench wrap with https://github.com/go-faster/errors or uber code
	// return errors.Wrap(err, info.ToString())
}

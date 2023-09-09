package wrap

import "fmt"

type wrapBuilder struct {
	base error
	fn   fnInfo

	showErrorSource    bool
	decorateWithFormat bool

	format string
	args   []any
}

func (wb *wrapBuilder) WithoutErrorSource() {
	wb.showErrorSource = false
}

func (wb *wrapBuilder) WithFormatMessage(format string, args ...any) {
	wb.decorateWithFormat = true
	wb.format = format
	wb.args = args
}

func (wb *wrapBuilder) Build() error {
	result := wb.base

	if wb.decorateWithFormat {
		result = fmt.Errorf("%s: %w", fmt.Sprintf(wb.format, wb.args...), result)
	}

	if wb.showErrorSource {
		result = fmt.Errorf("%s: %w", wb.fn.ToString(), result)
	}

	// TODO: check with bench wrap with https://github.com/go-faster/errors or uber code
	// return errors.Wrap(err, info.ToString())

	return result
}

func defaultWrapBuilder(err error, info fnInfo) *wrapBuilder {
	return &wrapBuilder{
		base:            err,
		fn:              info,
		showErrorSource: true,
	}
}

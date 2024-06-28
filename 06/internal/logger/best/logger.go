package best

import (
	lg "06/internal/logger"
	"fmt"
	"io"
)

var _ lg.Logger = (*logger)(nil)

type logger struct {
	out io.Writer
}

func NewLogger(out io.Writer) *logger {
	return &logger{out: out}
}

func (l *logger) Print(args ...any) {
	newEvent(lg.DebugLevel, l.out).Msg(fmt.Sprint(args...))
}

func (l *logger) Debug() *event {
	return newEvent(lg.DebugLevel, l.out)
}

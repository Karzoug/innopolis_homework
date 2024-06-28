package simple

import (
	"fmt"
	"io"
	"time"

	lg "06/internal/logger"
)

var _ lg.Logger = (*logger)(nil)

type logger struct {
	out io.Writer
}

func NewLogger(out io.Writer) *logger {
	return &logger{out: out}
}

func (l *logger) Print(v ...any) {
	fmt.Fprintln(l.out,
		time.Now().Format(lg.TimeFieldFormat),
		lg.LevelDebugValue,
		fmt.Sprint(v...),
	)
}

package optim

import (
	lg "06/internal/logger"
	"fmt"
	"io"
	"sync"
	"time"
)

var _ lg.Logger = (*logger)(nil)

var bufferPool = sync.Pool{
	New: func() any {
		return new([]byte)
	},
}

type logger struct {
	//mu  sync.Mutex
	out io.Writer
}

func NewLogger(out io.Writer) *logger {
	return &logger{out: out}
}

func (l *logger) Print(v ...any) {
	t := time.Now()

	b := bufferPool.Get().(*[]byte)
	*b = (*b)[:0]
	defer bufferPool.Put(b)

	*b = t.AppendFormat(*b, lg.TimeFieldFormat)
	*b = append(*b, ',', ' ')
	*b = append(*b, lg.LevelFieldName...)
	*b = append(*b, '=')
	*b = append(*b, lg.LevelDebugValue...)
	*b = append(*b, ',', ' ')
	*b = append(*b, fmt.Sprintln(v...)...)

	//l.mu.Lock()
	//defer l.mu.Unlock()
	l.out.Write(*b)
}

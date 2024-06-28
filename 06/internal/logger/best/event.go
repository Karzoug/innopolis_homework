package best

import (
	lg "06/internal/logger"
	"io"
	"strconv"
	"sync"
	"time"
)

const messageFieldName = "message"

var TimeFieldFormat = time.RFC3339

var eventPool = sync.Pool{
	New: func() any {
		return &event{
			buf: make([]byte, 0, 64), // TODO: find a size better mb
		}
	},
}

type event struct {
	level lg.Level
	buf   []byte
	out   io.Writer
}

func newEvent(lvl lg.Level, out io.Writer) *event {
	e := eventPool.Get().(*event)
	e.buf = (e.buf)[:0]
	e.level = lvl
	e.out = out

	e.buf = time.Now().AppendFormat(e.buf, TimeFieldFormat)
	e.buf = append(e.buf, ',', ' ')

	e.buf = append(e.buf, lg.LevelFieldName...)
	e.buf = append(e.buf, '=')
	e.buf = append(e.buf, lvl.String()...)
	e.buf = append(e.buf, ',', ' ')

	return e
}

func (e *event) Msg(msg string) {
	e.buf = append(e.buf, messageFieldName...)
	e.buf = append(e.buf, '=', '"')
	e.buf = append(e.buf, msg...)
	e.buf = append(e.buf, '"')

	e.write()
}

func (e *event) Int(key string, value int) *event {
	e.buf = append(e.buf, key...)
	e.buf = append(e.buf, '=')
	e.buf = strconv.AppendInt(e.buf, int64(value), 10)
	e.buf = append(e.buf, ',', ' ')

	return e
}

func (e *event) Float64(key string, value float64) *event {
	e.buf = append(e.buf, key...)
	e.buf = append(e.buf, '=')
	e.buf = strconv.AppendFloat(e.buf, value, 'f', -1, 64)
	e.buf = append(e.buf, ',', ' ')

	return e
}

func (e *event) write() {
	e.buf = append(e.buf, '\n')
	e.out.Write(e.buf)
	putEvent(e)
}

func putEvent(e *event) {
	eventPool.Put(e)
}

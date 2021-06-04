package head

import (
	"encoding/json"
	"fmt"
	"os"
	"runtime"
)

// Logger ...Logger
type Logger struct {
	d string
	s int
}

// NewLogger ...NewLogger
func NewLogger() Logger {
	return Logger{d: "/var/log/go/cargo.log", s: 1}
}

// SetSkip ...SetSkip
func (l *Logger) SetSkip(i int) *Logger {
	l.s = i
	return l
}

// Info ...Info
func (l *Logger) Info(a ...interface{}) {
	_, file, no, ok := runtime.Caller(l.s)
	f, err := os.OpenFile(l.d, os.O_APPEND|os.O_WRONLY, 0600)
	c, err := f.WriteString(fmt.Sprintf("\n>> %s:%d\n   ", file, no))
	b, err := json.Marshal(a)
	if err == nil {
		if _, err := f.Write(b); err == nil {
		}
	} else {
		fmt.Println(err, ok, c)
	}
}

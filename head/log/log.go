package log

import (
	"encoding/json"
	"fmt"
	"os"
	"runtime"
)

// Log ...Log
type Log struct {
	d string
	s int
}

// New ...New
func New() Log {
	return Log{d: "/var/log/go/cargo.log", s: 1}
}

// SetSkip ...SetSkip
func (l *Log) SetSkip(i int) *Log {
	l.s = i
	return l
}

// Info ...Info
func (l *Log) Info(a ...interface{}) {
	_, file, no, ok := runtime.Caller(l.s)
	f, err := os.OpenFile(l.d, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		_, err := f.WriteString(fmt.Sprintf("\n>> %s:%d\n   ", file, no))
		if err != nil {
			b, err := json.Marshal(a)
			if err == nil {
				if _, err := f.Write(b); err == nil {
					return
				}
			}
		}
	}
	fmt.Println(err, ok)
}

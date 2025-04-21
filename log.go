package plusgorm

import (
	"fmt"
	"strings"
)

type Logger interface {
	// Errorf logs to the ERROR log. Arguments are handled in the manner of fmt.Printf.
	Errorf(format string, args ...interface{})
	Infof(format string, args ...interface{})
}

var defLog = &Log{}

type Log struct {
}

func (l *Log) Errorf(format string, args ...interface{}) {
	format = strings.ReplaceAll(format, "%v", "%+v")
	fmt.Println(fmt.Errorf(format, args...))
}

func (l *Log) Infof(format string, args ...interface{}) {
	format = strings.ReplaceAll(format, "%v", "%+v")
	fmt.Println(fmt.Sprintf(format, args...))
}

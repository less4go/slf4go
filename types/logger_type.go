package types

import (
	"github.com/less4go/slf4go/config"
)

type (
	// Logger interface
	Logger interface {
		Debug(msg string, fields ...*Field)
		Info(msg string, fields ...*Field)
		Warn(msg string, fields ...*Field)
		Error(msg string, fields ...*Field)
		Fatal(msg string, fields ...*Field)
		Panic(msg string, fields ...*Field)
		With(fields ...*Field) Logger
	}

	// Bootstrap interface
	Bootstrap interface {
		BootLogger(logConfig *config.LogConfig) Logger
	}

	// Field struct
	Field struct {
		Key string
		Val interface{}
	}
)

// NewField func
func NewField(key string, val interface{}) *Field {
	return &Field{key, val}
}

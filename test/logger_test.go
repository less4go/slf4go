package test

import (
	"testing"
	"time"

	"github.com/less4go/slf4go"
	"github.com/less4go/slf4go/types"
)

type MyTest struct {
	aaa string
	bbb int
	ccc bool
}

func TestLogger(t *testing.T) {

	slf4go.Logger.Debug("aaa",
		types.NewField("string", "string"),
		types.NewField("int", 100),
	)
	slf4go.Logger.Info("a fields",
		types.NewField("string", "string"),
		types.NewField("int", 100),
		types.NewField("time", time.Now()),
	)
	slf4go.Logger.Info("aaa")

	slf4go.Logger.With(
		types.NewField("a", "b"),
		types.NewField("int", 100),
		types.NewField("time", time.Now()),
	).Info("12345")
	slf4go.Logger.Info("bbb")

	slf4go.Logger.Warn("Warn log")
}

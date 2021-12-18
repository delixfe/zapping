package appender

import (
	"go.uber.org/zap/buffer"
	"go.uber.org/zap/zapcore"
	"zap_ing/appender/appendercore"
	"zap_ing/internal/bufferpool"
)

// EnvelopingFn Function to create the enveloped output.
// p contains the original content
// the enveloped content must be written into output
// entry by ref or pointer?
// -> benchmarks show that using a pointer creates one alloc (for the pointer)
//    but passing by value does not
type EnvelopingFn func(p []byte, ent zapcore.Entry, output *buffer.Buffer) error

var _ appendercore.SynchronizationAwareAppender = &Enveloping{}

type Enveloping struct {
	primary appendercore.Appender
	envFn   EnvelopingFn
}

func (a *Enveloping) Synchronized() bool {
	return appendercore.Synchronized(a.primary)
}

func NewEnveloping(inner appendercore.Appender, envFn EnvelopingFn) *Enveloping {
	return &Enveloping{
		primary: inner,
		envFn:   envFn,
	}
}

func NewEnvelopingPreSuffix(inner appendercore.Appender, prefix, suffix string) *Enveloping {
	envFn := func(p []byte, ent zapcore.Entry, output *buffer.Buffer) error {
		output.WriteString(prefix)
		output.Write(p)
		output.WriteString(suffix)
		return nil
	}
	return NewEnveloping(inner, envFn)
}

func (a *Enveloping) Write(p []byte, ent zapcore.Entry) (n int, err error) {
	buf := bufferpool.Get()
	defer buf.Free()
	err = a.envFn(p, ent, buf)
	if err != nil {
		return
	}
	n, err = a.primary.Write(buf.Bytes(), ent)
	return
}

func (a *Enveloping) Sync() error {
	return a.Sync()
}

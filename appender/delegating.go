package appender

import (
	"go.uber.org/zap/zapcore"
	"zap_ing/appender/appendercore"
)

var _ appendercore.SynchronizationAwareAppender = &Delegating{}

type Delegating struct {
	WriteFn           func(p []byte, ent zapcore.Entry) (n int, err error)
	SyncFn            func() error
	SynchronizedValue bool
}

func NewDelegating(writeFn func(p []byte, ent zapcore.Entry) (n int, err error), syncFn func() error, synchronized bool) *Delegating {
	return &Delegating{
		WriteFn:           writeFn,
		SyncFn:            syncFn,
		SynchronizedValue: synchronized,
	}
}

func (a *Delegating) Write(p []byte, ent zapcore.Entry) (int, error) {
	writeFn := a.WriteFn
	if writeFn == nil {
		return len(p), nil
	}
	return writeFn(p, ent)
}

func (a *Delegating) Sync() error {
	syncFn := a.SyncFn
	if syncFn == nil {
		return nil
	}
	return syncFn()
}

func (a *Delegating) Synchronized() bool {
	return a.SynchronizedValue
}

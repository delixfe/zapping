package chaos

import (
	"context"
	"go.uber.org/zap/zapcore"
	"zap_ing/appender"
)

var (
	_ appender.Appender = &BlockingSwitchable{}
	_ Switchable        = &BlockingSwitchable{}
)

type BlockingSwitchable struct {
	primary appender.Appender
	enabled bool
	waiting chan struct{}
	ctx     context.Context
}

func NewBlockingSwitchable(ctx context.Context, inner appender.Appender) *BlockingSwitchable {
	if ctx == nil {
		ctx = context.Background()
	}
	return &BlockingSwitchable{
		primary: inner,
		enabled: false,
		ctx:     ctx,
	}
}

func (a *BlockingSwitchable) Enabled() bool {
	return a.enabled
}

func (a *BlockingSwitchable) Enable() {
	if a.enabled {
		return
	}
	a.enabled = true
	a.waiting = make(chan struct{})
}

func (a *BlockingSwitchable) Disable() {
	a.enabled = false
	close(a.waiting)
}

func (a *BlockingSwitchable) Write(p []byte, ent zapcore.Entry) (n int, err error) {
	if a.enabled {
		select {
		case <-a.waiting:
		case <-a.ctx.Done():
			return 0, a.ctx.Err()
		}
	}
	n, err = a.primary.Write(p, ent)
	return

}

func (a *BlockingSwitchable) Sync() error {
	return a.Sync()
}
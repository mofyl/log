package log

import (
	"fmt"
	"os"
)

const (
	DefaultMessageKey = "msg"
)

type HelperOption func(*Helper)

type Helper struct {
	logger Logger
	msgKey string
}

func NewHelper(logger Logger, op ...HelperOption) *Helper {

	h := &Helper{
		logger: logger,
		msgKey: DefaultMessageKey,
	}

	for _, v := range op {
		v(h)
	}

	return h
}

func (h *Helper) Log(level Level, keyvals ...interface{}) {
	_ = h.logger.Log(level, keyvals...)
}

func (h *Helper) Debug(a ...interface{}) {
	h.Log(LevelDebug, h.msgKey, fmt.Sprint(a...))
}

func (h *Helper) Debugf(format string, a ...interface{}) {
	h.Log(LevelDebug, fmt.Sprintf(format, a...))
}

func (h *Helper) Debugw(keyvals ...interface{}) {
	h.Log(LevelDebug, keyvals...)
}

func (h *Helper) Info(a ...interface{}) {
	h.Log(LevelInfo, h.msgKey, fmt.Sprint(a...))
}

func (h *Helper) Infof(format string, a ...interface{}) {
	h.Log(LevelInfo, fmt.Sprintf(format, a...))
}

func (h *Helper) Infow(keyvals ...interface{}) {
	h.Log(LevelInfo, keyvals...)
}

func (h *Helper) Warn(a ...interface{}) {
	h.Log(LevelWarn, h.msgKey, fmt.Sprint(a...))
}

func (h *Helper) Warnf(format string, a ...interface{}) {
	h.Log(LevelWarn, fmt.Sprintf(format, a...))
}

func (h *Helper) Warnw(keyvals ...interface{}) {
	h.Log(LevelWarn, keyvals...)
}

func (h *Helper) Error(a ...interface{}) {
	h.Log(LevelError, h.msgKey, fmt.Sprint(a...))
}

func (h *Helper) Errorf(format string, a ...interface{}) {
	h.Log(LevelError, fmt.Sprintf(format, a...))
}

func (h *Helper) Errorw(keyvals ...interface{}) {
	h.Log(LevelError, keyvals...)
}

func (h *Helper) Fatal(a ...interface{}) {
	h.Log(LevelFatal, h.msgKey, fmt.Sprint(a...))
	os.Exit(1)
}

func (h *Helper) Fatalf(format string, a ...interface{}) {
	h.Log(LevelFatal, fmt.Sprintf(format, a...))
	os.Exit(1)
}

func (h *Helper) Fatalw(keyvals ...interface{}) {
	h.Log(LevelFatal, keyvals...)
	os.Exit(1)
}

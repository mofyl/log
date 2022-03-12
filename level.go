package log

import (
	"go.uber.org/zap/zapcore"
	"strings"
)

type Level int8

const LevelKey = "level"

const (
	LevelDebug = iota - 1

	LevelInfo

	LevelWarn

	LevelError

	LevelFatal

	LevelDebugStr = "DEBUG"
	LevelInfoStr  = "INFO"
	LevelWarnStr  = "WARN"
	LevelErrorStr = "ERROR"
	LevelFatalStr = "FATAL"
)

func (l Level) String() string {
	switch l {
	case LevelDebug:
		return LevelDebugStr
	case LevelInfo:
		return LevelInfoStr
	case LevelWarn:
		return LevelWarnStr
	case LevelError:
		return LevelErrorStr
	case LevelFatal:
		return LevelFatalStr
	}
	return ""
}

func (l Level) ToZapLevel() zapcore.Level {

	switch l {
	case LevelDebug:
		return zapcore.DebugLevel
	case LevelInfo:
		return zapcore.InfoLevel
	case LevelWarn:
		return zapcore.WarnLevel
	case LevelError:
		return zapcore.ErrorLevel
	case LevelFatal:
		return zapcore.FatalLevel
	}

	return zapcore.InfoLevel
}

func ParseLevel(level string) Level {

	switch strings.ToUpper(level) {
	case LevelDebugStr:
		return LevelDebug
	case LevelInfoStr:
		return LevelInfo
	case LevelWarnStr:
		return LevelWarn
	case LevelErrorStr:
		return LevelError
	case LevelFatalStr:
		return LevelFatal
	}
	return LevelInfo
}

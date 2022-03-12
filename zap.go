package log

import (
	"fmt"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"os"
	"strings"
	"time"
)

var (
	depth = 3
)

type ZapOption func(*ZapConfig)

type RotateCfg struct {
	FilePath   string
	PrintLevel Level
	MaxAge     time.Duration
	RotateTime time.Duration
}

type ZapConfig struct {
	zapCfg  *zapcore.EncoderConfig
	writers []RotateCfg
	Skip    int
}

type ZapLogger struct {
	logger *zap.Logger
}

var deZapConfig = zapcore.EncoderConfig{
	LevelKey:      "level",
	TimeKey:       "ts",
	StacktraceKey: "stack",
	EncodeLevel:   zapcore.CapitalLevelEncoder,
	EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		// 默认毫秒级别
		enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
	},
	CallerKey:    "file",
	EncodeCaller: zapcore.ShortCallerEncoder,
	EncodeDuration: func(d time.Duration, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendInt64(int64(d) / 1000000)
	},
}

func WithEncodeTime(layout string) ZapOption {
	return func(cfg *ZapConfig) {
		cfg.zapCfg.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format(layout))
		}
	}
}

func WithSkip(skip int) ZapOption {
	return func(cfg *ZapConfig) {
		cfg.Skip = skip
	}
}

func WithWriters(writer ...RotateCfg) ZapOption {
	return func(cfg *ZapConfig) {
		cfg.writers = append(cfg.writers, writer...)
	}
}

func NewZapLogger(ops ...ZapOption) Logger {
	defaultCfg := deZapConfig
	zCfg := &ZapConfig{
		zapCfg: &defaultCfg,
		Skip:   1,
	}

	for _, v := range ops {
		v(zCfg)
	}

	encoder := zapcore.NewJSONEncoder(*zCfg.zapCfg)

	zcs := make([]zapcore.Core, 0, len(zCfg.writers))

	for i := 0; i < len(zCfg.writers); i++ {
		curCfg := zCfg.writers[i]
		writer := getWriter(curCfg)
		level := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
			return lvl >= curCfg.PrintLevel.ToZapLevel()
		})
		zapCore := zapcore.NewCore(encoder, zapcore.AddSync(writer), level)
		zcs = append(zcs, zapCore)
	}

	// 添加 std
	zcs = append(zcs, zapcore.NewCore(encoder,
		zapcore.AddSync(os.Stdout),
		zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
			return lvl >= zapcore.DebugLevel
		}),
	))

	core := zapcore.NewTee(zcs...)

	log := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(depth))

	return &ZapLogger{
		logger: log,
	}
}

func getWriter(cfg RotateCfg) io.Writer {

	fileName := strings.Replace(cfg.FilePath, ".log", "", -1) + "-%Y%m%d%H.log"
	hook, err := rotatelogs.New(
		fileName,
		rotatelogs.WithLinkName(cfg.FilePath),
		rotatelogs.WithMaxAge(cfg.MaxAge),
		rotatelogs.WithRotationTime(cfg.RotateTime),
	)

	if err != nil {
		panic("new file rotate err" + err.Error())
	}
	return hook
}

// caller 不对的时候 就调整
func Caller(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
}

func (zl *ZapLogger) Log(level Level, kv ...interface{}) error {

	if len(kv) == 0 || len(kv)%2 != 0 {
		zl.logger.Warn(fmt.Sprint("kv must appear in pairs: ", kv))
		return nil
	}

	var data []zap.Field
	for i := 0; i < len(kv); i += 2 {
		data = append(data, zap.Any(fmt.Sprint(kv[i]), kv[i+1]))
	}

	switch level {
	case LevelDebug:
		zl.logger.Debug("", data...)
	case LevelInfo:
		zl.logger.Info("", data...)
	case LevelWarn:
		zl.logger.Warn("", data...)
	case LevelError:
		zl.logger.Error("", data...)
	}

	//zl.logger.Sync()
	return nil

}

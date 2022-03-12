package log

import (
	"testing"
	"time"
)

func TestLogger(t *testing.T) {
	//
	//rcs := make([]RotateCfg, 0)
	//
	//rcs = append(rcs, RotateCfg{
	//	FilePath:   "./logs/tmp.log",
	//	PrintLevel: LevelInfo,
	//	MaxAge:     time.Hour,
	//	RotateTime: time.Hour,
	//})
	//
	//rcs = append(rcs, RotateCfg{
	//	FilePath:   "./logs/tmp1.log",
	//	PrintLevel: LevelDebug,
	//	MaxAge:     time.Hour,
	//	RotateTime: time.Hour,
	//})
	//
	//logger := NewZapLogger(WithWriters(rcs...))
	//
	//h := NewHelper(logger)
	//
	//for i := 0; i < 10; i++ {
	//	//h.Debug(fmt.Sprintf("%s_%d", "asd", i))
	//	//h.Info(fmt.Sprintf("%s_%d", "asdf234", i))
	//	h.Infow(i, "asd")
	//	time.Sleep(2 * time.Second)
	//}

	NewZap()

}

func NewZap() {

	rcs := make([]RotateCfg, 0)

	rcs = append(rcs, RotateCfg{
		FilePath:   "./logs/tmp.log",
		PrintLevel: LevelInfo,
		MaxAge:     time.Hour,
		RotateTime: time.Hour,
	})

	rcs = append(rcs, RotateCfg{
		FilePath:   "./logs/tmp1.log",
		PrintLevel: LevelDebug,
		MaxAge:     time.Hour,
		RotateTime: time.Hour,
	})

	logger := NewZapLogger(WithWriters(rcs...))

	h := NewHelper(logger)

	for i := 0; i < 10; i++ {
		//h.Debug(fmt.Sprintf("%s_%d", "asd", i))
		//h.Info(fmt.Sprintf("%s_%d", "asdf234", i))
		h.Infow(i, "asd")
		time.Sleep(2 * time.Second)
	}
}

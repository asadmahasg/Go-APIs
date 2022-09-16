package common

import (
	"log"
	"time"

	"project/libgo/random"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type CommonC struct {
	L *zap.SugaredLogger

	App string

}

func Init(
	App string,
) *CommonC {
	// configure logger.
	logConf := zap.NewProductionConfig()
	logConf.EncoderConfig.LevelKey = "l"
	logConf.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.RFC3339)
	l, err := logConf.Build()
	if err != nil {
		log.Panicf("failed to build logger")
	}
	lFinal := l.Sugar().With(
		"app", App,
		"apphash", random.Base62Str(9),
	)

	return &CommonC{
		L:         lFinal,
		App:       App,
	}
}

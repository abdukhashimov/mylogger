package zaplogger

import (
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func LoadZapConfig(prod bool) zap.Config {
	var (
		zapConfig zap.Config
	)

	if prod {
		// 2019-08-13T04:39:11Z

		zapConfig = zap.NewProductionConfig()
		zapConfig.EncoderConfig.EncodeTime = zapcore.TimeEncoder(
			func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
				enc.AppendString(t.UTC().Format("2006-01-02T15:04:05Z0700"))
			},
		)
	} else {
		// Aug 13 00:38:11

		zapConfig = zap.NewDevelopmentConfig()
		zapConfig.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		zapConfig.EncoderConfig.EncodeTime = zapcore.TimeEncoder(
			func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
				enc.AppendString(t.Format(time.Stamp))
			},
		)
	}

	return zapConfig
}

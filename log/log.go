package log

import (
	"github.com/go-logr/logr"
	"github.com/go-logr/zapr"
	"go.uber.org/zap"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
)

import (
	"go.uber.org/zap/zapcore"
)

func New(w io.Writer, ops *Options) logr.Logger {
	validOptions(ops)
	syncer := zapcore.AddSync(w)
	encoder := zapcore.NewJSONEncoder(ops.EncoderConfig)
	core := zapcore.NewCore(encoder, syncer, ops.Level)
	zapLogger := zap.New(core)

	logger := zapr.NewLogger(zapLogger)
	return logger
}

// todo
func validOptions(ops *Options) {

}

func NewDefaultFileWriter() io.Writer {
	return &lumberjack.Logger{
		Filename:   "app.log",
		MaxSize:    128,
		MaxBackups: 3,
		Compress:   false,
	}
}

func NewDefaultFileWriterWithFileName(fileName string) io.Writer {
	return &lumberjack.Logger{
		Filename:   fileName,
		MaxSize:    128,
		MaxBackups: 3,
		Compress:   false,
	}
}

type Options struct {
	zapcore.EncoderConfig
	zapcore.Level
}

func NewDefaultOptions() *Options {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return &Options{
		EncoderConfig: encoderConfig,
		Level:         zapcore.DebugLevel,
	}
}

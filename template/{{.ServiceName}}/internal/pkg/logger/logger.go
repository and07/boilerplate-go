package logger

import (
	"os"
	"strconv"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type logStdout struct {
}

// Error ...
func (l *logStdout) Error(msg string) {
	log.Error(msg)
}

// Infof ...
func (l *logStdout) Infof(msg string, args ...interface{}) {
	log.Infof(msg, args...)
}

// Printf ...
func (l *logStdout) Printf(msg string, args ...interface{}) {
	log.Debugf(msg, args...)
}

// Logger ...

var log *zap.SugaredLogger

// StdLogger ...
var StdLogger = &logStdout{}

// New ...
func New() *zap.SugaredLogger {
	logLevel := os.Getenv("LOG_LEVEL")
	level, _ := strconv.ParseInt(logLevel, 10, 8)
	var zapLevel zapcore.Level
	switch level {
	case 0:
		zapLevel = zapcore.DebugLevel
	case 1:
		zapLevel = zapcore.InfoLevel
	case 2:
		zapLevel = zapcore.FatalLevel
	case 3:
		zapLevel = zapcore.PanicLevel
	case 4:
		zapLevel = zapcore.ErrorLevel
	case 5:
		zapLevel = zapcore.DPanicLevel
	case 6:
		zapLevel = zapcore.WarnLevel
	default:
		zapLevel = zapcore.DebugLevel
	}
	cfg := zap.Config{
		Encoding:         "json",
		Level:            zap.NewAtomicLevelAt(zapLevel),
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stdout"},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey: "message",

			LevelKey:    "level",
			EncodeLevel: zapcore.CapitalLevelEncoder,

			TimeKey:    "time",
			EncodeTime: zapcore.ISO8601TimeEncoder,
		},
	}
	logTemp, _ := cfg.Build()
	logger := logTemp.Sugar()
	if err := logger.Sync(); err != nil {
		logger.Error(err)
	} // flushes buffer, if any
	return logger
}

func init() {
	log = New()
}

// Debug logs a message at level Debug on the standard logger.
func Debug(args ...interface{}) {
	log.Debug(args...)
}

// Print logs a message at level Info on the standard logger.
func Print(args ...interface{}) {
	//log.Print(args...)
	log.Info(args...)
}

// Info logs a message at level Info on the standard logger.
func Info(args ...interface{}) {
	log.Info(args...)
}

// Error logs a message at level Error on the standard logger.
func Error(args ...interface{}) {
	log.Error(args...)
}

// Fatal logs a message at level Fatal on the standard logger then the process will exit with status set to 1.
func Fatal(args ...interface{}) {
	log.Fatal(args...)
}

// Debugf logs a message at level Debug on the standard logger.
func Debugf(format string, args ...interface{}) {
	log.Debugf(format, args...)
}

// Printf logs a message at level Info on the standard logger.
func Printf(format string, args ...interface{}) {
	log.Infof(format, args...)
}

// Infof logs a message at level Info on the standard logger.
func Infof(format string, args ...interface{}) {
	log.Infof(format, args...)
}

// Errorf logs a message at level Error on the standard logger.
func Errorf(format string, args ...interface{}) {
	log.Errorf(format, args...)
}

// Fatalf logs a message at level Fatal on the standard logger then the process will exit with status set to 1.
func Fatalf(format string, args ...interface{}) {
	log.Fatalf(format, args...)
}

// Debugln logs a message at level Debug on the standard logger.
func Debugln(args ...interface{}) {
	log.Debug(args...)
}

// Println logs a message at level Info on the standard logger.
func Println(args ...interface{}) {
	log.Info(args...)
}

// Infoln logs a message at level Info on the standard logger.
func Infoln(args ...interface{}) {
	log.Info(args...)
}

// Errorln logs a message at level Error on the standard logger.
func Errorln(args ...interface{}) {
	log.Error(args...)
}

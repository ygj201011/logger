package logger

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type Level uint32
type Logger interface {
	Print(args ...interface{})
	Printf(format string, args ...interface{})

	Trace(args ...interface{})
	Tracef(format string, args ...interface{})

	Debug(args ...interface{})
	Debugf(format string, args ...interface{})

	Info(args ...interface{})
	Infof(format string, args ...interface{})

	Warn(args ...interface{})
	Warnf(format string, args ...interface{})

	Error(args ...interface{})
	Errorf(format string, args ...interface{})

	Fatal(args ...interface{})
	Fatalf(format string, args ...interface{})

	Panic(args ...interface{})
	Panicf(format string, args ...interface{})

	WithPrefix(prefix string) Logger
	Prefix() string

	WithSection(prefix string) Logger
	Section() string

	WithFields(fields Fields) Logger
	Fields() Fields

	SetLevel(level Level)
}

type Loggable interface {
	Log() Logger
}

type Fields map[string]interface{}

func (fields Fields) String() string {
	str := make([]string, 0)

	for k, v := range fields {
		str = append(str, fmt.Sprintf("%s=%+v", k, v))
	}

	return strings.Join(str, " ")
}

func (fields Fields) WithFields(newFields Fields) Fields {
	allFields := make(Fields)

	for k, v := range fields {
		allFields[k] = v
	}

	for k, v := range newFields {
		allFields[k] = v
	}

	return allFields
}

func AddFieldsFrom(logger Logger, values ...interface{}) Logger {
	for _, value := range values {
		switch v := value.(type) {
		case Logger:
			logger = logger.WithFields(v.Fields())
		case Loggable:
			logger = logger.WithFields(v.Log().Fields())
		case interface{ Fields() Fields }:
			logger = logger.WithFields(v.Fields())
		}
	}
	return logger
}

var Log MyLogger

type MyLogger struct {
	zap.Logger
}

func (l MyLogger) Print(args ...interface{}) {
	l.Sugar().Info(args...)
}
func (l MyLogger) Printf(format string, args ...interface{}) {
	// _, file, line, _ := runtime.Caller(1)
	// fileList := strings.Split(file, "/")
	// format = fmt.Sprintf("[%s:%d] %s", fileList[len(fileList)-1], line, format)
	l.Sugar().Infof(format, args...)
}
func (l MyLogger) Println(args ...interface{}) {
	// _, file, line, _ := runtime.Caller(1)
	// fileList := strings.Split(file, "/")
	// args = append([]interface{}{fmt.Sprintf("[%s:%d] ", fileList[len(fileList)-1], line)}, args...)
	l.Sugar().Info(args...)
}
func (l MyLogger) Trace(args ...interface{}) {
	// _, file, line, _ := runtime.Caller(1)
	// fileList := strings.Split(file, "/")
	// args = append([]interface{}{fmt.Sprintf("[%s:%d] ", fileList[len(fileList)-1], line)}, args...)
	l.Sugar().Debug(args...)
}
func (l MyLogger) Tracef(format string, args ...interface{}) {
	// _, file, line, _ := runtime.Caller(1)
	// fileList := strings.Split(file, "/")
	// format = fmt.Sprintf("[%s:%d] %s", fileList[len(fileList)-1], line, format)
	l.Sugar().Debugf(format, args...)
}

func (l MyLogger) Debug(args ...interface{}) {
	// _, file, line, _ := runtime.Caller(1)
	// fileList := strings.Split(file, "/")
	// args = append([]interface{}{fmt.Sprintf("[%s:%d] ", fileList[len(fileList)-1], line)}, args...)
	l.Sugar().Debug(args...)
}
func (l MyLogger) Debugf(format string, args ...interface{}) {
	// _, file, line, _ := runtime.Caller(1)
	// fileList := strings.Split(file, "/")
	// format = fmt.Sprintf("[%s:%d] %s", fileList[len(fileList)-1], line, format)
	l.Sugar().Debugf(format, args...)
}
func (l MyLogger) Info(args ...interface{}) {
	// _, file, line, _ := runtime.Caller(1)
	// fileList := strings.Split(file, "/")
	// args = append([]interface{}{fmt.Sprintf("[%s:%d] ", fileList[len(fileList)-1], line)}, args...)
	l.Sugar().Info(args...)
}
func (l MyLogger) Infof(format string, args ...interface{}) {
	// _, file, line, _ := runtime.Caller(1)
	// fileList := strings.Split(file, "/")
	// format = fmt.Sprintf("[%s:%d] %s", fileList[len(fileList)-1], line, format)
	l.Sugar().Infof(format, args...)
}
func (l MyLogger) Warn(args ...interface{}) {
	// _, file, line, _ := runtime.Caller(1)
	// fileList := strings.Split(file, "/")
	// args = append([]interface{}{fmt.Sprintf("[%s:%d] ", fileList[len(fileList)-1], line)}, args...)
	l.Sugar().Warn(args...)
}
func (l MyLogger) Warnf(format string, args ...interface{}) {
	// _, file, line, _ := runtime.Caller(1)
	// fileList := strings.Split(file, "/")
	// format = fmt.Sprintf("[%s:%d] %s", fileList[len(fileList)-1], line, format)
	l.Sugar().Warnf(format, args...)
}
func (l MyLogger) Panic(args ...interface{}) {
	// _, file, line, _ := runtime.Caller(1)
	// fileList := strings.Split(file, "/")
	// args = append([]interface{}{fmt.Sprintf("[%s:%d] ", fileList[len(fileList)-1], line)}, args...)
	l.Sugar().Panic(args...)
}
func (l MyLogger) Panicf(format string, args ...interface{}) {
	// _, file, line, _ := runtime.Caller(1)
	// fileList := strings.Split(file, "/")
	// format = fmt.Sprintf("[%s:%d] %s", fileList[len(fileList)-1], line, format)
	l.Sugar().Panicf(format, args...)
}
func (l MyLogger) Fatal(args ...interface{}) {
	// _, file, line, _ := runtime.Caller(1)
	// fileList := strings.Split(file, "/")
	// args = append([]interface{}{fmt.Sprintf("[%s:%d] ", fileList[len(fileList)-1], line)}, args...)
	l.Sugar().Fatal(args...)
}
func (l MyLogger) Error(args ...interface{}) {
	// _, file, line, _ := runtime.Caller(1)
	// fileList := strings.Split(file, "/")
	// args = append([]interface{}{fmt.Sprintf("[%s:%d] ", fileList[len(fileList)-1], line)}, args...)
	l.Sugar().Error(args...)
}
func (l MyLogger) Errorf(format string, args ...interface{}) {
	// _, file, line, _ := runtime.Caller(1)
	// fileList := strings.Split(file, "/")
	// format = fmt.Sprintf("[%s:%d] %s", fileList[len(fileList)-1], line, format)
	l.Sugar().Errorf(format, args...)
}
func (l MyLogger) Fatalf(format string, args ...interface{}) {
	// _, file, line, _ := runtime.Caller(1)
	// fileList := strings.Split(file, "/")
	// format = fmt.Sprintf("[%s:%d] %s", fileList[len(fileList)-1], line, format)
	l.Sugar().Fatalf(format, args...)
}
func (l MyLogger) WithPrefix(prefix string) Logger {
	return MyLogger{*l.Named(prefix)}
}
func (l MyLogger) Prefix() string {
	return ""
}
func (l MyLogger) Fields() Fields {
	return make(map[string]interface{})
}
func (l MyLogger) WithFields(fields Fields) Logger {
	fs := make([]zap.Field, 0)

	for k, v := range fields {
		fs = append(fs, zap.Any(k, v))
	}
	return MyLogger{*(l.With(fs...))}
}

func (l MyLogger) SetLevel(level Level) {
	//can't change level
	return
}
func (l MyLogger) Section() (s string) {
	return
}
func (l MyLogger) WithSection(fileds string) (ws Logger) {
	ws = MyLogger{*(l.With(zap.String("fileds", fileds)))}
	return
}

func GetMyLogger() MyLogger {
	return Log
}

type LoggerConfig struct {
	Level          string `json:"log_level"`
	Path           string `json:"log_path"`
	LogName        string `json:"log_name"`
	MaxSize        int    `json:"log_max_size"`
	MaxBackups     int    `json:"log_max_backup"`
	MaxAge         int    `json:"log_max_age"`
	EnableCompress int    `json:"log_compress"`
}

func InitLogger(logger LoggerConfig) {
	logPathAccessabel := LogPathExists(logger.Path)
	logPath := "./"

	if logPathAccessabel {
		log.Println("Log Path not exist of not writable, use default path: ./")
	} else {
		logPath = logger.Path
	}
	if logPath[len(logPath)-1:] == "/" {
		logPath = logPath[0 : len(logPath)-1]
	}

	maxLogSize := logger.MaxSize
	maxBackups := logger.MaxBackups
	maxAge := logger.MaxAge

	var enableCompress bool
	if maxLogSize == 0 {
		maxLogSize = 100
	}
	if maxBackups == 0 {
		maxBackups = 10
	}
	if maxAge == 0 {
		maxAge = 7
	}
	if logger.EnableCompress == 1 {
		enableCompress = true
	} else {
		enableCompress = false
	}

	w := zapcore.AddSync(&lumberjack.Logger{
		Filename:   logPath + "/" + logger.LogName + ".log",
		MaxSize:    maxLogSize, // megabytes
		MaxBackups: maxBackups,
		MaxAge:     maxAge, // days
		Compress:   enableCompress,
	})

	encoderCfg := zapcore.EncoderConfig{
		// Keys can be anything except the empty string.
		TimeKey:        "T",
		LevelKey:       "L",
		NameKey:        "N",
		CallerKey:      "C",
		MessageKey:     "M",
		StacktraceKey:  "S",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
	level := zap.InfoLevel
	switch strings.ToLower(logger.Level) {
	case "debug":
		level = zap.DebugLevel
	case "warn":
		level = zap.WarnLevel
	case "info":
		level = zap.InfoLevel
	case "error":
		level = zap.ErrorLevel
	default:
		level = zap.InfoLevel
	}
	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderCfg),
		w,
		level,
	)
	Log = MyLogger{
		*zap.New(core),
	}

}

func LogPathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	return false
}

func Print(args ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	fileList := strings.Split(file, "/")
	args = append([]interface{}{fmt.Sprintf("[%s:%d] ", fileList[len(fileList)-1], line)}, args...)
	Log.Sugar().Info(args...)
}
func Printf(format string, args ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	fileList := strings.Split(file, "/")
	format = fmt.Sprintf("[%s:%d] %s", fileList[len(fileList)-1], line, format)
	Log.Sugar().Infof(format, args...)
}
func Println(args ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	fileList := strings.Split(file, "/")
	args = append([]interface{}{fmt.Sprintf("[%s:%d] ", fileList[len(fileList)-1], line)}, args...)
	Log.Info(args...)
}
func Trace(args ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	fileList := strings.Split(file, "/")
	args = append([]interface{}{fmt.Sprintf("[%s:%d] ", fileList[len(fileList)-1], line)}, args...)
	Log.Debug(args...)
}
func Tracef(format string, args ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	fileList := strings.Split(file, "/")
	format = fmt.Sprintf("[%s:%d] %s", fileList[len(fileList)-1], line, format)
	Log.Debugf(format, args...)
}

func Debug(args ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	fileList := strings.Split(file, "/")
	args = append([]interface{}{fmt.Sprintf("[%s:%d] ", fileList[len(fileList)-1], line)}, args...)
	Log.Debug(args...)
}
func Debugf(format string, args ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	fileList := strings.Split(file, "/")
	format = fmt.Sprintf("[%s:%d] %s", fileList[len(fileList)-1], line, format)
	Log.Debugf(format, args...)
}
func Info(args ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	fileList := strings.Split(file, "/")
	args = append([]interface{}{fmt.Sprintf("[%s:%d] ", fileList[len(fileList)-1], line)}, args...)
	Log.Info(args...)
}
func Infof(format string, args ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	fileList := strings.Split(file, "/")
	format = fmt.Sprintf("[%s:%d] %s", fileList[len(fileList)-1], line, format)
	Log.Infof(format, args...)
}
func Warn(args ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	fileList := strings.Split(file, "/")
	args = append([]interface{}{fmt.Sprintf("[%s:%d] ", fileList[len(fileList)-1], line)}, args...)
	Log.Warn(args...)
}
func Warnf(format string, args ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	fileList := strings.Split(file, "/")
	format = fmt.Sprintf("[%s:%d] %s", fileList[len(fileList)-1], line, format)
	Log.Warnf(format, args...)
}
func Panic(args ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	fileList := strings.Split(file, "/")
	args = append([]interface{}{fmt.Sprintf("[%s:%d] ", fileList[len(fileList)-1], line)}, args...)
	Log.Panic(args...)
}
func Panicf(format string, args ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	fileList := strings.Split(file, "/")
	format = fmt.Sprintf("[%s:%d] %s", fileList[len(fileList)-1], line, format)
	Log.Panicf(format, args...)
}
func Fatal(args ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	fileList := strings.Split(file, "/")
	args = append([]interface{}{fmt.Sprintf("[%s:%d] ", fileList[len(fileList)-1], line)}, args...)
	Log.Fatal(args...)
}
func Error(args ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	fileList := strings.Split(file, "/")
	args = append([]interface{}{fmt.Sprintf("[%s:%d] ", fileList[len(fileList)-1], line)}, args...)
	Log.Error(args...)
}
func Errorf(format string, args ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	fileList := strings.Split(file, "/")
	format = fmt.Sprintf("[%s:%d] %s", fileList[len(fileList)-1], line, format)
	Log.Errorf(format, args...)
}
func Fatalf(format string, args ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	fileList := strings.Split(file, "/")
	format = fmt.Sprintf("[%s:%d] %s", fileList[len(fileList)-1], line, format)
	Log.Fatalf(format, args...)
}

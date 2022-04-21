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
	IsInit() bool
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
	return newFields
}

type Level uint32

// These are the different logging levels. You can set the logging level to log
// on your instance of logger, obtained with `logrus.New()`.
const (
	// PanicLevel level, highest level of severity. Logs and then calls panic with the
	// message passed to Debug, Info, ...
	PanicLevel Level = iota
	// FatalLevel level. Logs and then calls `logger.Exit(1)`. It will exit even if the
	// logging level is set to Panic.
	FatalLevel
	// ErrorLevel level. Logs. Used for errors that should definitely be noted.
	// Commonly used for hooks to send errors to an error tracking service.
	ErrorLevel
	// WarnLevel level. Non-critical entries that deserve eyes.
	WarnLevel
	// InfoLevel level. General operational entries about what's going on inside the
	// application.
	InfoLevel
	// DebugLevel level. Usually only enabled when debugging. Very verbose logging.
	DebugLevel
	// TraceLevel level. Designates finer-grained informational events than the Debug.
	TraceLevel
)

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

var (
	mylog  MyLogger
	isInit bool
)

type MyLogger struct {
	log    *zap.Logger
	level  Level
	prefix []string
	fields Fields
}

func (l MyLogger) Print(args ...interface{}) {
	if isInit {
		l.log.Sugar().Info(args...)
	} else {
		fmt.Println(args...)
	}
}
func (l MyLogger) IsInit() bool {
	return isInit
}
func addCodeInfo(format string) (r string) {
	r = format
	// _, file, line, _ := runtime.Caller(2)
	// fileList := strings.Split(file, "/")
	// r = fmt.Sprintf("[%s:%d] %s", fileList[len(fileList)-1], line, format)
	return
}
func appendCodeInfo(args []interface{}) (r []interface{}) {
	r = args
	_, file, line, _ := runtime.Caller(2)
	fileList := strings.Split(file, "/")
	args = append([]interface{}{fmt.Sprintf("[%s:%d] ", fileList[len(fileList)-1], line)}, args)
	return
}
func (l MyLogger) Printf(format string, args ...interface{}) {
	if isInit {
		l.log.Sugar().Infof(addCodeInfo(format), args...)
	} else {
		fmt.Printf(addCodeInfo(format)+"\r\n", args...)
	}
}
func (l MyLogger) Println(args ...interface{}) {
	if isInit {
		l.log.Sugar().Info(appendCodeInfo(args)...)
	} else {
		fmt.Println(args...)
	}
}
func (l MyLogger) Trace(args ...interface{}) {
	if isInit {
		if l.level == TraceLevel {
			l.log.Sugar().Debug(appendCodeInfo(args)...)
		}
	} else {
		fmt.Println(args...)
	}
}
func (l MyLogger) Tracef(format string, args ...interface{}) {
	if isInit {
		if l.level == TraceLevel {
			l.log.Sugar().Debugf(addCodeInfo(format), args...)
		}
	} else {
		fmt.Printf(addCodeInfo(format)+"\r\n", args...)
	}
}

func (l MyLogger) Debug(args ...interface{}) {
	if isInit {
		l.log.Sugar().Debug(appendCodeInfo(args)...)
	} else {
		fmt.Println(args...)
	}
}
func (l MyLogger) Debugf(format string, args ...interface{}) {
	if isInit {
		l.log.Sugar().Debugf(addCodeInfo(format), args...)
	} else {
		fmt.Printf(addCodeInfo(format)+"\r\n", args...)
	}
}
func (l MyLogger) Info(args ...interface{}) {
	if isInit {
		l.log.Sugar().Info(appendCodeInfo(args)...)
	} else {
		fmt.Println(args...)
	}
}
func (l MyLogger) Infof(format string, args ...interface{}) {
	if isInit {
		l.log.Sugar().Infof(addCodeInfo(format), args...)
	} else {
		fmt.Printf(addCodeInfo(format)+"\r\n", args...)
	}
}
func (l MyLogger) Warn(args ...interface{}) {
	if isInit {
		l.log.Sugar().Warn(appendCodeInfo(args)...)
	} else {
		fmt.Println(args...)
	}
}
func (l MyLogger) Warnf(format string, args ...interface{}) {
	if isInit {
		l.log.Sugar().Warnf(addCodeInfo(format), args...)
	} else {
		fmt.Printf(addCodeInfo(format)+"\r\n", args...)
	}
}
func (l MyLogger) Panic(args ...interface{}) {
	if isInit {
		l.log.Sugar().Panic(appendCodeInfo(args)...)
	} else {
		fmt.Println(args...)
	}
}
func (l MyLogger) Panicf(format string, args ...interface{}) {
	if isInit {
		l.log.Sugar().Panicf(addCodeInfo(format), args...)
	} else {
		fmt.Printf(addCodeInfo(format)+"\r\n", args...)
	}
}
func (l MyLogger) Fatal(args ...interface{}) {
	if isInit {
		l.log.Sugar().Fatal(appendCodeInfo(args)...)
	} else {
		fmt.Println(args...)
	}
}
func (l MyLogger) Error(args ...interface{}) {
	if isInit {
		l.log.Sugar().Error(appendCodeInfo(args)...)
	} else {
		fmt.Println(args...)
	}
}
func (l MyLogger) Errorf(format string, args ...interface{}) {
	if isInit {
		l.log.Sugar().Errorf(addCodeInfo(format), args...)
	} else {
		fmt.Printf(addCodeInfo(format)+"\r\n", args...)
	}
}
func (l MyLogger) Fatalf(format string, args ...interface{}) {
	if isInit {
		l.log.Sugar().Fatalf(addCodeInfo(format), args...)
	} else {
		fmt.Printf(addCodeInfo(format)+"\r\n", args...)
	}
}
func (l MyLogger) WithPrefix(prefix string) (log Logger) {
	log = MyLogger{
		log:    l.log.Named(prefix),
		level:  l.level,
		prefix: append(l.prefix, prefix),
		fields: l.fields,
	}
	return
}
func (l MyLogger) Prefix() string {
	return strings.Join(l.prefix, ".")
}
func (l MyLogger) Fields() Fields {
	if l.fields == nil {
		return make(Fields)
	}
	return l.fields
}
func (l MyLogger) WithFields(fields Fields) (log Logger) {
	fs := make([]zap.Field, 0)
	for k, v := range fields {
		fs = append(fs, zap.Any(k, v))
	}
	log = MyLogger{
		log:    l.log.With(fs...),
		level:  l.level,
		prefix: l.prefix,
		fields: fields,
	}
	return
}

func (l MyLogger) SetLevel(level Level) {
	//can't change level
	return
}
func (l MyLogger) GetLevel() (level Level) {
	//can't change level
	return l.level
}
func (l MyLogger) Section() (s string) {
	return
}
func (l MyLogger) WithSection(sec string) (log Logger) {
	fs := make([]zap.Field, 0)
	l.fields["section"] = sec
	if l.fields == nil {
		l.fields = make(map[string]interface{})
	}
	for k, v := range l.fields {
		fs = append(fs, zap.Any(k, v))
	}
	log = MyLogger{
		log:    l.log.With(zap.String("section", sec)),
		level:  l.level,
		prefix: l.prefix,
		fields: l.fields,
	}
	return
}

func GetMyLogger() MyLogger {
	return mylog
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
		log.Println("mylog Path not exist of not writable, use default path: ./")
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
	myLevel := TraceLevel
	switch strings.ToLower(logger.Level) {
	case "trace":
		myLevel = TraceLevel
		level = zap.DebugLevel
	case "debug":
		myLevel = DebugLevel
		level = zap.DebugLevel
	case "warn":
		myLevel = WarnLevel
		level = zap.WarnLevel
	case "info":
		myLevel = InfoLevel
		level = zap.InfoLevel
	case "error":
		myLevel = ErrorLevel
		level = zap.ErrorLevel
	default:
		level = zap.InfoLevel
	}
	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderCfg),
		w,
		level,
	)
	mylog = MyLogger{
		log:    zap.New(core),
		level:  myLevel,
		fields: make(map[string]interface{}),
	}
	isInit = true
}

func LogPathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	return false
}

func Print(args ...interface{}) {
	if isInit {
		mylog.log.Sugar().Info(appendCodeInfo(args)...)
	} else {
		fmt.Println(args...)
	}
}
func Printf(format string, args ...interface{}) {
	if isInit {
		mylog.log.Sugar().Infof(addCodeInfo(format), args...)
	} else {
		fmt.Printf(addCodeInfo(format)+"\r\n", args...)
	}
}
func Println(args ...interface{}) {
	if isInit {
		mylog.log.Sugar().Info(appendCodeInfo(args)...)
	} else {
		fmt.Print(args...)
	}

}
func Trace(args ...interface{}) {
	if isInit {
		if mylog.level == TraceLevel {
			mylog.log.Sugar().Debug(appendCodeInfo(args)...)
		}
	} else {
		fmt.Println(args...)
	}
}
func Tracef(format string, args ...interface{}) {
	if isInit {
		if mylog.level == TraceLevel {
			mylog.log.Sugar().Debugf(addCodeInfo(format), args...)
		}
	} else {
		fmt.Printf(addCodeInfo(format)+"\r\n", args...)
	}
}

func Debug(args ...interface{}) {
	if isInit {
		mylog.log.Sugar().Debug(appendCodeInfo(args)...)
	} else {
		fmt.Println(args...)
	}
}
func Debugf(format string, args ...interface{}) {
	if isInit {
		mylog.log.Sugar().Debugf(addCodeInfo(format), args...)
	} else {
		fmt.Printf(addCodeInfo(format)+"\r\n", args...)
	}
}
func Info(args ...interface{}) {
	if isInit {
		mylog.log.Sugar().Info(appendCodeInfo(args)...)
	} else {
		fmt.Println(args...)
	}
}
func Infof(format string, args ...interface{}) {
	if isInit {
		mylog.log.Sugar().Infof(addCodeInfo(format), args...)
	} else {
		fmt.Printf(addCodeInfo(format)+"\r\n", args...)
	}
}
func Warn(args ...interface{}) {
	if isInit {
		mylog.log.Sugar().Warn(appendCodeInfo(args)...)
	} else {
		fmt.Println(args...)
	}
}
func Warnf(format string, args ...interface{}) {
	if isInit {
		mylog.log.Sugar().Warnf(addCodeInfo(format), args...)
	} else {
		fmt.Printf(addCodeInfo(format)+"\r\n", args...)
	}
}
func Panic(args ...interface{}) {
	if isInit {
		mylog.log.Sugar().Panic(appendCodeInfo(args)...)
	} else {
		fmt.Println(args...)
	}
}
func Panicf(format string, args ...interface{}) {
	if isInit {
		mylog.log.Sugar().Panicf(addCodeInfo(format), args...)
	} else {
		fmt.Printf(addCodeInfo(format)+"\r\n", args...)
	}
}
func Error(args ...interface{}) {
	if isInit {
		mylog.log.Sugar().Error(appendCodeInfo(args)...)
	} else {
		fmt.Println(args...)
	}
}
func Errorf(format string, args ...interface{}) {
	if isInit {
		mylog.log.Sugar().Errorf(addCodeInfo(format), args...)
	} else {
		fmt.Printf(addCodeInfo(format)+"\r\n", args...)
	}
}
func Fatal(args ...interface{}) {
	if isInit {
		mylog.log.Sugar().Fatal(appendCodeInfo(args)...)
	} else {
		fmt.Println(args...)
	}
}
func Fatalf(format string, args ...interface{}) {
	if isInit {
		mylog.log.Sugar().Fatalf(addCodeInfo(format), args...)
	} else {
		fmt.Printf(addCodeInfo(format)+"\r\n", args...)
	}
}

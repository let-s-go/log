package log

import (
	"os"
)

var stdLogger = NewLogger(os.Stdout)

func StdLogger() *Logger {
	return stdLogger
}

func SetOptions(level Level, fileLine bool) {
	stdLogger.SetOptions(level, fileLine)
}

func Debugf(format string, args ...interface{}) {
	stdLogger.Debugf(format, args...)
}

func Infof(format string, args ...interface{}) {
	stdLogger.Infof(format, args...)
}

func Printf(format string, args ...interface{}) {
	stdLogger.Printf(format, args...)
}

func Warnf(format string, args ...interface{}) {
	stdLogger.Warnf(format, args...)
}

func Errorf(format string, args ...interface{}) {
	stdLogger.Errorf(format, args...)
}

func Fatalf(format string, args ...interface{}) {
	stdLogger.Fatalf(format, args...)
	os.Exit(1)
}

func Panicf(format string, args ...interface{}) {
	stdLogger.Panicf(format, args...)
}

func Debug(args ...interface{}) {
	stdLogger.Debug(args...)
}

func Info(args ...interface{}) {
	stdLogger.Info(args...)
}

func Print(args ...interface{}) {
	stdLogger.Print(args...)
}

func Warn(args ...interface{}) {
	stdLogger.Warn(args...)
}

func Error(args ...interface{}) {
	stdLogger.Error(args...)
}

func Fatal(args ...interface{}) {
	stdLogger.Fatal(args...)
}

func Panic(args ...interface{}) {
	stdLogger.Panic(args...)
}

func Debugln(args ...interface{}) {
	stdLogger.Debugln(args...)
}

func Infoln(args ...interface{}) {
	stdLogger.Infoln(args...)
}

func Println(args ...interface{}) {
	stdLogger.Println(args...)
}

func Warnln(args ...interface{}) {
	stdLogger.Warnln(args...)
}

func Errorln(args ...interface{}) {
	stdLogger.Errorln(args...)
}

func Fatalln(args ...interface{}) {
	stdLogger.Fatalln(args...)
}

func Panicln(args ...interface{}) {
	stdLogger.Panicln(args...)
}

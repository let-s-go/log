package log

import (
	"fmt"
	"io"
	"os"
	"sync"
	"time"
)

type Logger struct {
	writer io.Writer
	level  Level
	mutex  sync.Mutex
}

func NewLogger(writer io.Writer) *Logger {
	return &Logger{
		writer: writer,
		level:  WarnLevel,
	}
}

func NewFileLogger(fileName string, fileSize int64) *Logger {
	return &Logger{
		writer: newFileWriter(fileName, fileSize),
	}
}

func (logger *Logger) SetLevel(level Level) {
	logger.level = level
}

func (logger *Logger) Debugf(format string, args ...interface{}) {
	logger.logf(DebugLevel, format, args...)
}

func (logger *Logger) Infof(format string, args ...interface{}) {
	logger.logf(InfoLevel, format, args...)
}

func (logger *Logger) Printf(format string, args ...interface{}) {
	logger.logf(InfoLevel, format, args...)
}

func (logger *Logger) Warnf(format string, args ...interface{}) {
	logger.logf(WarnLevel, format, args...)
}

func (logger *Logger) Errorf(format string, args ...interface{}) {
	logger.logf(ErrorLevel, format, args...)
}

func (logger *Logger) Fatalf(format string, args ...interface{}) {
	logger.logf(FatalLevel, format, args...)
	os.Exit(1)
}

func (logger *Logger) Panicf(format string, args ...interface{}) {
	logger.logf(PanicLevel, format, args...)
}

func (logger *Logger) Debug(args ...interface{}) {
	logger.log(DebugLevel, args...)
}

func (logger *Logger) Info(args ...interface{}) {
	logger.log(InfoLevel, args...)
}

func (logger *Logger) Print(args ...interface{}) {
	logger.log(InfoLevel, args...)
}

func (logger *Logger) Warn(args ...interface{}) {
	logger.log(WarnLevel, args...)
}

func (logger *Logger) Error(args ...interface{}) {
	logger.log(ErrorLevel, args...)
}

func (logger *Logger) Fatal(args ...interface{}) {
	logger.log(FatalLevel, args...)
	os.Exit(1)
}

func (logger *Logger) Panic(args ...interface{}) {
	logger.log(PanicLevel, args...)
}

func (logger *Logger) Debugln(args ...interface{}) {
	logger.logln(DebugLevel, args...)
}

func (logger *Logger) Infoln(args ...interface{}) {
	logger.logln(InfoLevel, args...)
}

func (logger *Logger) Println(args ...interface{}) {
	logger.logln(InfoLevel, args...)
}

func (logger *Logger) Warnln(args ...interface{}) {
	logger.logln(WarnLevel, args...)
}

func (logger *Logger) Errorln(args ...interface{}) {
	logger.logln(ErrorLevel, args...)
}

func (logger *Logger) Fatalln(args ...interface{}) {
	logger.logln(FatalLevel, args...)
	os.Exit(1)
}

func (logger *Logger) Panicln(args ...interface{}) {
	logger.logln(PanicLevel, args...)
}

func (logger *Logger) logf(level Level, format string, args ...interface{}) {
	if logger.level >= level {
		str := fmt.Sprintf(format, args...)
		logger.write(level, str)
	}
}

func (logger *Logger) log(level Level, args ...interface{}) {
	if logger.level >= level {
		str := fmt.Sprint(args...)
		logger.write(level, str)
	}
}

func (logger *Logger) logln(level Level, args ...interface{}) {
	if logger.level >= level {
		str := fmt.Sprintln(args...)
		logger.write(level, str)
	}
}

func (logger *Logger) write(level Level, str string) {
	logger.mutex.Lock()
	defer logger.mutex.Unlock()

	str = fmt.Sprintf("%s [%v] %s", time.Now().Format("2006-01-02 15:04:05"), level, str)
	logger.writer.Write([]byte(str))
}

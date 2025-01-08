package logs

import (
	"log"
	"os"
)

const (
	LevelDebug = iota
	LevelInfo
	LevelWarn
	LevelError
)

type Logger struct {
	debugLogger *log.Logger
	infoLogger  *log.Logger
	warnLogger  *log.Logger
	errorLogger *log.Logger
	level       int
}

var defaultLogger *Logger

func init() {
	// Create a default logger with INFO level, or whatever you like.
	defaultLogger = &Logger{
		debugLogger: log.New(os.Stderr, "[DEBUG] ", 0),
		infoLogger:  log.New(os.Stderr, "[INFO] ", 0),
		warnLogger:  log.New(os.Stderr, "[WARN] ", 0),
		errorLogger: log.New(os.Stderr, "[ERROR] ", 0),
		level:       LevelInfo,
	}
}

// SetLevel allows changing the global logger’s level at runtime.
func SetLevel(level int) {
	defaultLogger.level = level
}

func DebugOn() bool {
	return defaultLogger.level == LevelDebug
}

func Debugln(v ...any) {
	if defaultLogger.level <= LevelDebug {
		defaultLogger.debugLogger.Println(v...)
	}
}

func Debug(format string, v ...any) {
	if defaultLogger.level <= LevelDebug {
		defaultLogger.debugLogger.Printf(format, v...)
	}
}

func Infoln(v ...any) {
	if defaultLogger.level <= LevelInfo {
		defaultLogger.infoLogger.Println(v...)
	}
}

func Info(format string, v ...any) {
	if defaultLogger.level <= LevelInfo {
		defaultLogger.infoLogger.Printf(format, v...)
	}
}

func Warnln(v ...any) {
	if defaultLogger.level <= LevelWarn {
		defaultLogger.warnLogger.Println(v...)
	}
}

func Warn(format string, v ...any) {
	if defaultLogger.level <= LevelWarn {
		defaultLogger.warnLogger.Printf(format, v...)
	}
}

func Errorln(v ...any) {
	if defaultLogger.level <= LevelError {
		defaultLogger.errorLogger.Println(v...)
	}
}

func Error(format string, v ...any) {
	if defaultLogger.level <= LevelError {
		defaultLogger.errorLogger.Printf(format, v...)
	}
}

func Fatalln(v ...any) {
	if defaultLogger.level <= LevelError {
		defaultLogger.errorLogger.Println(v...)
	}
	os.Exit(1)
}

func Fatal(format string, v ...any) {
	if defaultLogger.level <= LevelError {
		defaultLogger.errorLogger.Printf(format, v...)
	}
	os.Exit(1)
}

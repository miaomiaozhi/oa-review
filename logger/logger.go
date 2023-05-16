package logger

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"sync"
)

type logLevel int

// 日志等级
const (
	DEBUG logLevel = iota
	INFO
	WARNING
	ERROR
	FATAL
)

var (
	logFile            *os.File                                              // 日志文件地址
	defaultPrefix      = ""                                                  // 默认前缀
	defaultCallerDepth = 2                                                   // 返回上两层函数的函数调用者信息
	logger             *log.Logger                                           // 日志
	mu                 sync.Mutex                                            // 锁
	logPrefix          = ""                                                  // 日志前缀
	levelFlags         = []string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"} // 日志等级标志
)

const flags = log.LstdFlags

type Settings struct {
	Path     string // 日志文件目录
	FileName string // 日志文件名
}

func init() {
	// TODO: add to runner
	logger = log.New(os.Stdout, defaultPrefix, flags)
}

// 初始化日志
func Setup(settings *Settings) {
	var err error
	dir := settings.Path
	fileName := settings.FileName
	logFile, err = mustOpen(fileName, dir)
	if err != nil {
		log.Fatalf("logging.Setup err: %s", err)
	}

	mw := io.MultiWriter(os.Stdout, logFile)
	logger = log.New(mw, defaultPrefix, flags)
}

func setPrefix(level logLevel) {
	_, file, line, ok := runtime.Caller(defaultCallerDepth)
	if ok {
		logPrefix = fmt.Sprintf("[%s][%s:%d] ", levelFlags[level], filepath.Base(file), line)
	} else {
		logPrefix = fmt.Sprintf("[%s] ", levelFlags[level])
	}

	logger.SetPrefix(logPrefix)
}

// Debug prints debug log
func Debug(v ...interface{}) {
	mu.Lock()
	defer mu.Unlock()
	setPrefix(DEBUG)
	logger.Println(v...)
}

// Debug prints debug log
func Debugf(format string, v ...interface{}) {
	mu.Lock()
	defer mu.Unlock()
	setPrefix(DEBUG)
	logger.Println(fmt.Sprintf(format, v...))
}

// Info prints normal log
func Info(v ...interface{}) {
	mu.Lock()
	defer mu.Unlock()
	setPrefix(INFO)
	logger.Println(v...)
}

// Warn prints warning log
func Warn(v ...interface{}) {
	mu.Lock()
	defer mu.Unlock()
	setPrefix(WARNING)
	logger.Println(v...)
}

// Error prints error log
func Error(v ...interface{}) {
	mu.Lock()
	defer mu.Unlock()
	setPrefix(ERROR)
	logger.Println(v...)
}

func Errorf(format string, v ...interface{}) {
	mu.Lock()
	defer mu.Unlock()
	setPrefix(ERROR)
	logger.Println(fmt.Sprintf(format, v...))
}

// Fatal prints error log then stop the program
func Fatal(v ...interface{}) {
	mu.Lock()
	defer mu.Unlock()
	setPrefix(FATAL)
	logger.Fatalln(v...)
}

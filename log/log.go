package log

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path"
	"time"
)

// Instance log实例
var Instance *log.Logger
var bufferWriter *bufio.Writer  // 缓冲区
var realtimeWrite bool          // 是否实时写入
var defaultPath string = "logs" //默认日志文件路径
var bufferSize = 256 * 1024     // 256KB 的缓冲区大小

func init() {
	Instance = log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)
	go flushDaemon() // 启动一个后台线程，定时刷新缓冲区
}

// SetRealtimeWriteLog 设置是否实时写入日志
func SetRealtimeWriteLog(realtime bool) {
	realtimeWrite = realtime
}

// SetBufferSize 设置缓冲区大小,默认256KB的缓冲区大小
func SetBufferSize(size int) {
	bufferSize = size
}

// SetPath 设置日志文件路径。如果为空，则使用默认路径:./logs
func SetPath(path string) {
	if path != "" {
		defaultPath = path
	}
}

// 设置日志输出方式: stdout和log file
// onlyStdout 为true时,日志只输出到标准输出;为false时,日志同时输出到标准输出和文件.
func SetOutput(onlyStdout bool) error {
	// 先刷新并关闭之前的缓冲区
	if bufferWriter != nil {
		Flush()
		bufferWriter = nil
	}

	// 标准输出
	Instance.SetOutput(os.Stdout)
	if onlyStdout {
		return nil
	}

	_, err := os.Stat(defaultPath)
	if os.IsNotExist(err) {
		err := os.MkdirAll(defaultPath, os.ModePerm)
		if err != nil {
			return fmt.Errorf("failed to create log directory: %v", err)
		}
	}

	// 输出到文件
	logpath := path.Join(defaultPath, fmt.Sprintf("%s.log", time.Now().Format("20060102")))
	f, err := os.OpenFile(logpath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return fmt.Errorf("failed to open log file: %v", err)
	}

	bufferWriter = bufio.NewWriterSize(f, bufferSize)
	Instance.SetOutput(bufferWriter)

	return nil
}

// Flush 刷新缓冲区，将日志写入文件
func Flush() {
	if bufferWriter != nil {
		err := bufferWriter.Flush()
		if err != nil {
			log.Printf("error flushing buffer: %v\n", err)
		}
	}
}

func flushDaemon() {
	tick := time.NewTicker(30 * time.Second)
	defer tick.Stop()
	for {
		<-tick.C
		Flush()
	}
}

// Info 信息
func Info(v ...interface{}) {
	Instance.SetPrefix("[INFO]")
	_ = Instance.Output(2, fmt.Sprint(v...))
	if realtimeWrite {
		Flush()
	}
}

// Infof 信息
func Infof(format string, v ...interface{}) {
	Instance.SetPrefix("[INFO]")
	_ = Instance.Output(2, fmt.Sprintf(format, v...))
	if realtimeWrite {
		Flush()
	}
}

// Infoln 信息
func Infoln(v ...interface{}) {
	Instance.SetPrefix("[INFO]")
	_ = Instance.Output(2, fmt.Sprintln(v...))
	if realtimeWrite {
		Flush()
	}
}

// Warn 提示
func Warn(v ...interface{}) {
	Instance.SetPrefix("[WARN]")
	_ = Instance.Output(2, fmt.Sprint(v...))
	if realtimeWrite {
		Flush()
	}
}

// Warnf 提示
func Warnf(format string, v ...interface{}) {
	Instance.SetPrefix("[WARN]")
	_ = Instance.Output(2, fmt.Sprintf(format, v...))
	if realtimeWrite {
		Flush()
	}
}

// Warnln 提示
func Warnln(v ...interface{}) {
	Instance.SetPrefix("[WARN]")
	_ = Instance.Output(2, fmt.Sprintln(v...))
	if realtimeWrite {
		Flush()
	}
}

// Error 错误
func Error(v ...interface{}) {
	Instance.SetPrefix("[ERRO]")
	_ = Instance.Output(2, fmt.Sprint(v...))
	if realtimeWrite {
		Flush()
	}
}

// Errorf 错误
func Errorf(format string, v ...interface{}) {
	Instance.SetPrefix("[ERRO]")
	_ = Instance.Output(2, fmt.Sprintf(format, v...))
	if realtimeWrite {
		Flush()
	}
}

// Errorln 错误
func Errorln(v ...interface{}) {
	Instance.SetPrefix("[ERRO]")
	_ = Instance.Output(2, fmt.Sprintln(v...))
	if realtimeWrite {
		Flush()
	}
}

// Debug 调试
func Debug(v ...interface{}) {
	Instance.SetPrefix("[DEBG]")
	_ = Instance.Output(2, fmt.Sprint(v...))
	if realtimeWrite {
		Flush()
	}
}

// Debugf 调试
func Debugf(format string, v ...interface{}) {
	Instance.SetPrefix("[DEBG]")
	_ = Instance.Output(2, fmt.Sprintf(format, v...))
	if realtimeWrite {
		Flush()
	}
}

// Debugln 调试
func Debugln(v ...interface{}) {
	Instance.SetPrefix("[DEBG]")
	_ = Instance.Output(2, fmt.Sprintln(v...))
	if realtimeWrite {
		Flush()
	}
}

// Fatal 致命信息
func Fatal(v ...interface{}) {
	Instance.SetPrefix("[FTAL]")
	_ = Instance.Output(2, fmt.Sprint(v...))
	Flush()
	os.Exit(1)
}

// Fatalf 致命信息
func Fatalf(format string, v ...interface{}) {
	Instance.SetPrefix("[FTAL]")
	_ = Instance.Output(2, fmt.Sprintf(format, v...))
	Flush()
	os.Exit(1)
}

// Fataln 致命信息
func Fataln(v ...interface{}) {
	Instance.SetPrefix("[FTAL]")
	_ = Instance.Output(2, fmt.Sprintln(v...))
	Flush()
	os.Exit(1)
}

// Painc Painc
func Painc(v ...interface{}) {
	Instance.SetPrefix("[PANC]")
	s := fmt.Sprint(v...)
	_ = Instance.Output(2, s)
	Flush()
	panic(s)
}

// Paincf Painc
func Paincf(format string, v ...interface{}) {
	Instance.SetPrefix("[PANC]")
	s := fmt.Sprintf(format, v...)
	_ = Instance.Output(2, s)
	Flush()
	panic(s)
}

// Paincln Painc
func Paincln(v ...interface{}) {
	Instance.SetPrefix("[PANC]")
	s := fmt.Sprintln(v...)
	_ = Instance.Output(2, s)
	Flush()
	panic(s)
}

/***************************/

// Printf Printf
func Printf(format string, v ...interface{}) {
	_ = Instance.Output(2, fmt.Sprintf(format, v...))
	if realtimeWrite {
		Flush()
	}
}

// Println Println
func Println(v ...interface{}) {
	_ = Instance.Output(2, fmt.Sprintln(v...))
	if realtimeWrite {
		Flush()
	}
}

// Print Print
func Print(v ...interface{}) {
	_ = Instance.Output(2, fmt.Sprint(v...))
	if realtimeWrite {
		Flush()
	}
}

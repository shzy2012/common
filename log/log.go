package log

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"time"
)

// Instance log实例
var Instance *log.Logger
var bufferWriter *bufio.Writer // 缓冲区

const bufferSize = 256 * 1024 // 256KB 的缓冲区大小
const defaultPath string = "logs"

func init() {
	Instance = log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)
	go flushDaemon() // 启动一个后台线程，定时刷新缓冲区
}

// SetOutput 设置log输出到文件
// useStdout 为true时，日志输出到标准输出，否则输出到文件;false时，日志输出到文件;
func SetOutput(useStdout bool) error {

	_, err := os.Stat(defaultPath)
	if os.IsNotExist(err) {
		err := os.MkdirAll(defaultPath, os.ModePerm)
		if err != nil {
			log.Print(err)
		}
	}

	logpath := path.Join(defaultPath, fmt.Sprintf("%s.log", time.Now().Format("20060102")))
	f, err := os.OpenFile(logpath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return err
	}

	if useStdout {
		bufferWriter = bufio.NewWriterSize(f, bufferSize)
		mw := io.MultiWriter(os.Stdout, bufferWriter)
		Instance.SetOutput(mw)
	} else {
		bufferWriter = bufio.NewWriterSize(f, bufferSize)
		Instance.SetOutput(bufferWriter)
	}

	return nil
}

func SetOutputWithPath(useStdout bool, yourPath string) error {

	_, err := os.Stat(yourPath)
	if os.IsNotExist(err) {
		err := os.MkdirAll(yourPath, os.ModePerm)
		if err != nil {
			log.Print(err)
		}
	}

	logpath := path.Join(yourPath, fmt.Sprintf("%s.log", time.Now().Format("20060102")))
	f, err := os.OpenFile(logpath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Printf("error opening file: %v\n", err)
		return err
	}

	if useStdout {
		bufferWriter = bufio.NewWriterSize(f, bufferSize)
		mw := io.MultiWriter(os.Stdout, bufferWriter)
		Instance.SetOutput(mw)
	} else {
		bufferWriter = bufio.NewWriterSize(f, bufferSize)
		Instance.SetOutput(bufferWriter)
	}

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
		select {
		case <-tick.C:
			Flush()
		}
	}
}

// Info 信息
func Info(v ...interface{}) {
	Instance.SetPrefix("[INFO]")
	_ = Instance.Output(2, fmt.Sprint(v...))
}

// Infof 信息
func Infof(format string, v ...interface{}) {
	Instance.SetPrefix("[INFO]")
	_ = Instance.Output(2, fmt.Sprintf(format, v...))
}

// Infoln 信息
func Infoln(v ...interface{}) {
	Instance.SetPrefix("[INFO]")
	_ = Instance.Output(2, fmt.Sprintln(v...))
}

// Warn 提示
func Warn(v ...interface{}) {
	Instance.SetPrefix("[WARN]")
	_ = Instance.Output(2, fmt.Sprint(v...))
}

// Warnf 提示
func Warnf(format string, v ...interface{}) {
	Instance.SetPrefix("[WARN]")
	_ = Instance.Output(2, fmt.Sprintf(format, v...))
}

// Warnln 提示
func Warnln(v ...interface{}) {
	Instance.SetPrefix("[WARN]")
	_ = Instance.Output(2, fmt.Sprintln(v...))
}

// Error 错误
func Error(v ...interface{}) {
	Instance.SetPrefix("[ERRO]")
	_ = Instance.Output(2, fmt.Sprint(v...))
}

// Errorf 错误
func Errorf(format string, v ...interface{}) {
	Instance.SetPrefix("[ERRO]")
	_ = Instance.Output(2, fmt.Sprintf(format, v...))
}

// Errorln 错误
func Errorln(v ...interface{}) {
	Instance.SetPrefix("[ERRO]")
	_ = Instance.Output(2, fmt.Sprintln(v...))
}

// Debug 调试
func Debug(v ...interface{}) {
	Instance.SetPrefix("[DEBG]")
	_ = Instance.Output(2, fmt.Sprint(v...))
}

// Debugf 调试
func Debugf(format string, v ...interface{}) {
	Instance.SetPrefix("[DEBG]")
	_ = Instance.Output(2, fmt.Sprintf(format, v...))
}

// Debugln 调试
func Debugln(v ...interface{}) {
	Instance.SetPrefix("[DEBG]")
	_ = Instance.Output(2, fmt.Sprintln(v...))
}

// Fatal 致命信息
func Fatal(v ...interface{}) {
	Instance.SetPrefix("[FTAL]")
	_ = Instance.Output(2, fmt.Sprint(v...))
	os.Exit(1)
}

// Fatalf 致命信息
func Fatalf(format string, v ...interface{}) {
	Instance.SetPrefix("[FTAL]")
	_ = Instance.Output(2, fmt.Sprintf(format, v...))
	os.Exit(1)
}

// Fataln 致命信息
func Fataln(v ...interface{}) {
	Instance.SetPrefix("[FTAL]")
	_ = Instance.Output(2, fmt.Sprintln(v...))
	os.Exit(1)
}

// Painc Painc
func Painc(v ...interface{}) {
	Instance.SetPrefix("[PANC]")
	s := fmt.Sprint(v...)
	_ = Instance.Output(2, s)
	panic(s)
}

// Paincf Painc
func Paincf(format string, v ...interface{}) {
	Instance.SetPrefix("[PANC]")
	s := fmt.Sprintf(format, v...)
	_ = Instance.Output(2, s)
	panic(s)
}

// Paincln Painc
func Paincln(v ...interface{}) {
	Instance.SetPrefix("[PANC]")
	s := fmt.Sprintln(v...)
	_ = Instance.Output(2, s)
	panic(s)
}

/***************************/

// Printf Printf
func Printf(format string, v ...interface{}) {
	_ = Instance.Output(2, fmt.Sprintf(format, v...))
}

// Println Println
func Println(v ...interface{}) {
	_ = Instance.Output(2, fmt.Sprintln(v...))
}

// Print Print
func Print(v ...interface{}) {
	_ = Instance.Output(2, fmt.Sprint(v...))
}

package log

import (
	"fmt"
	"log"
	"os"
)

// Instance log实例
var Instance *log.Logger
var defaultPath string = "logs" //默认日志文件路径

func init() {
	Instance = log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)
}

// SetPath 设置日志文件路径。如果为空，则使用默认路径:./logs
func SetPath(path string) {
	if path != "" {
		defaultPath = path
	}
}

// SetOutput 设置日志输出方式: stdout和log file
// onlyStdout 为true时,日志只输出到标准输出;为false时,日志同时输出到标准输出和文件.
func SetOutput(onlyStdout bool) error {
	if onlyStdout {
		// 标准输出
		Instance.SetOutput(os.Stdout)
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
	fileWriter, err := NewFileWriter(defaultPath)
	if err != nil {
		return fmt.Errorf("failed to create custom writer: %v", err)
	}

	Instance.SetOutput(NewCustomMultiWriter(fileWriter, os.Stdout))
	return nil
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

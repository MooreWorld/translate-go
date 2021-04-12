package tool

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

type logUtil struct {
}

var _logUtil *logUtil

// LogAPI 通用日志
func LogAPI() *logUtil {
	_logUtil = &logUtil{}
	return _logUtil
}

// 日志打印颜色
const (
	colorRed = iota + 91
	colorGreen
	colorYellow
	colorBlue
	colorMagenta
)

// Notice ...
func (api *logUtil) Notice(a ...interface{}) {
	msg := strings.TrimRight(fmt.Sprintln(a...), "\n")
	log.Println("\x1b[" + strconv.Itoa(colorGreen) + "m" + "▶ " + msg + "\x1b[0m")
}

// Debug ...
func (api *logUtil) Debug(a ...interface{}) {
	msg := strings.TrimRight(fmt.Sprintln(a...), "\n")
	log.Println("\x1b[" + strconv.Itoa(colorBlue) + "m" + "▶ " + msg + "\x1b[0m")
}

// Info ...
func (api *logUtil) Info(a ...interface{}) {
	msg := strings.TrimRight(fmt.Sprintln(a...), "\n")
	log.Println("\x1b[" + strconv.Itoa(colorYellow) + "m" + "▶ " + msg + "\x1b[0m")
}

// Warn ...
func (api *logUtil) Warn(a ...interface{}) {
	msg := strings.TrimRight(fmt.Sprintln(a...), "\n")
	log.Println("\x1b[" + strconv.Itoa(colorMagenta) + "m" + "▶ " + msg + "\x1b[0m")
}

// Error ...
func (api *logUtil) Error(a ...interface{}) {
	msg := strings.TrimRight(fmt.Sprintln(a...), "\n")
	log.Println("\x1b[" + strconv.Itoa(colorRed) + "m" + "▶ " + msg + "\x1b[0m")
}

package colorlog

import (
	"fmt"
	"github.com/LuckyWindsck/sfc-course-guide/sfc-course-guide-server/internal/pkg/ansicolor"
	"sync"
	"time"
)

// LogLevel represents the number of logger level
type LogLevel uint8

// Log Levels
const (
	ALL   = LogLevel(iota) // ALL: show all logs
	DEBUG                  // DEBUG: informational events that are most useful to debug an application.
	INFO                   // INFO : informational messages that highlight the progress of the application.
	WARN                   // WARN : potentially harmful situations.
	ERROR                  // ERROR: error events that might still allow the application to continue running.
	FATAL                  // FATAL: very severe error events that will presumably lead the application to abort.
	OFF                    // OFF  : Do not show any log
)

// ColorLogger represents an active logging object that generates lines of output in
// 16 ansicolors. A ColorLogger can be used simultaneously from multiple goroutines.
type ColorLogger struct {
	mu          sync.Mutex // ensures atomic writes; protects the following fields
	prefix      string
	fillPrefix  bool
	message     string
	fillMessage bool
	timeFormat  string
	color       ansicolor.Color
	level       LogLevel
}

// NewColorLogger creates a new Logger with default settings.
// To change settings, please use ColorLogger setter functions
func NewColorLogger() *ColorLogger {
	return &ColorLogger{
		prefix:      "",
		fillPrefix:  true,
		message:     "",
		fillMessage: true,
		timeFormat:  "2006/01/02 15:04:05",
		color:       ansicolor.White,
		level:       ALL,
	}
}

// Create standard ColorLogger.
var std = NewColorLogger()

// LogColor set the color of log for different log level.
var LogColor = map[LogLevel]ansicolor.Color{
	DEBUG: ansicolor.Blue,
	INFO:  ansicolor.Green,
	WARN:  ansicolor.Yellow,
	ERROR: ansicolor.Magenta,
	FATAL: ansicolor.Red,
}

// LogPrefix set the prefix of log for different log level.
var LogPrefix = map[LogLevel]string{
	DEBUG: "[DEBUG]",
	INFO:  "[INFO] ",
	WARN:  "[WARN] ",
	ERROR: "[ERROR]",
	FATAL: "[FATAL]",
}

// Print calls std.Output and log in the log level ALL.
// Arguments are handled in the manner of fmt.SPrint.
func Print(a ...interface{}) {
	std.Output(ALL, fmt.Sprint(a...))
}

// Printf calls std.Output and log in the log level ALL.
// Arguments are handled in the manner of fmt.SPrintf.
func Printf(format string, a ...interface{}) {
	std.Output(ALL, fmt.Sprintf(format, a...))
}

// Println calls std.Output and log in the log level ALL.
// Arguments are handled in the manner of fmt.SPrintln.
func Println(a ...interface{}) {
	std.Output(ALL, fmt.Sprintln(a...))
}

// Debug calls std.Output and log in the log level DEBUG.
// Arguments are handled in the manner of fmt.SPrint.
func Debug(a ...interface{}) {
	std.Output(DEBUG, fmt.Sprint(a...))
}

// Debugf calls std.Output and log in the log level DEBUG.
// Arguments are handled in the manner of fmt.SPrintf.
func Debugf(format string, a ...interface{}) {
	std.Output(DEBUG, fmt.Sprintf(format, a...))
}

// Debugln calls std.Output and log in the log level DEBUG.
// Arguments are handled in the manner of fmt.SPrintln.
func Debugln(a ...interface{}) {
	std.Output(DEBUG, fmt.Sprintln(a...))
}

// Info calls std.Output and log in the log level INFO.
// Arguments are handled in the manner of fmt.SPrint.
func Info(a ...interface{}) {
	std.Output(INFO, fmt.Sprint(a...))
}

// Infof calls std.Output and log in the log level INFO.
// Arguments are handled in the manner of fmt.SPrintf.
func Infof(format string, a ...interface{}) {
	std.Output(INFO, fmt.Sprintf(format, a...))
}

// Infoln calls std.Output and log in the log level INFO.
// Arguments are handled in the manner of fmt.SPrintln.
func Infoln(a ...interface{}) {
	std.Output(INFO, fmt.Sprintln(a...))
}

// Warn calls std.Output and log in the log level WARN.
// Arguments are handled in the manner of fmt.SPrint.
func Warn(a ...interface{}) {
	std.Output(WARN, fmt.Sprint(a...))
}

// Warnf calls std.Output and log in the log level WARN.
// Arguments are handled in the manner of fmt.SPrintf.
func Warnf(format string, a ...interface{}) {
	std.Output(WARN, fmt.Sprintf(format, a...))
}

// Warnln calls std.Output and log in the log level WARN.
// Arguments are handled in the manner of fmt.SPrintln.
func Warnln(a ...interface{}) {
	std.Output(WARN, fmt.Sprintln(a...))
}

// Error calls std.Output and log in the log level ERROR.
// Arguments are handled in the manner of fmt.SPrint.
func Error(a ...interface{}) {
	std.Output(ERROR, fmt.Sprint(a...))
}

// Errorf calls std.Output and log in the log level ERROR.
// Arguments are handled in the manner of fmt.SPrintf.
func Errorf(format string, a ...interface{}) {
	std.Output(ERROR, fmt.Sprintf(format, a...))
}

// Errorln calls std.Output and log in the log level ERROR.
// Arguments are handled in the manner of fmt.SPrintln.
func Errorln(a ...interface{}) {
	std.Output(ERROR, fmt.Sprintln(a...))
}

// Fatal calls std.Output and log in the log level FATAL.
// Arguments are handled in the manner of fmt.SPrint.
func Fatal(a ...interface{}) {
	std.Output(FATAL, fmt.Sprint(a...))
}

// Fatalf calls std.Output and log in the log level FATAL.
// Arguments are handled in the manner of fmt.SPrintf.
func Fatalf(format string, a ...interface{}) {
	std.Output(FATAL, fmt.Sprintf(format, a...))
}

// Fatalln calls std.Output and log in the log level FATAL.
// Arguments are handled in the manner of fmt.SPrintln.
func Fatalln(a ...interface{}) {
	std.Output(FATAL, fmt.Sprintln(a...))
}

// Log calls std.Output and log in the given log level.
// Arguments are handled in the manner of fmt.SPrint.
func Log(logLevel LogLevel, a ...interface{}) {
	std.Output(logLevel, fmt.Sprint(a...))
}

// Logf calls std.Output and log in the given log level.
// Arguments are handled in the manner of fmt.SPrintf.
func Logf(logLevel LogLevel, format string, a ...interface{}) {
	std.Output(logLevel, fmt.Sprintf(format, a...))
}

// Logln calls std.Output and log in the given log level.
// Arguments are handled in the manner of fmt.SPrintln.
func Logln(logLevel LogLevel, a ...interface{}) {
	std.Output(logLevel, fmt.Sprintln(a...))
}

// Output println the output for a logging event in the given log level with
// message. The format of log is "time prefix message"
func (l *ColorLogger) Output(logLevel LogLevel, message string) {
	if logLevel < l.level {
		return
	}

	color, exist := LogColor[logLevel]
	if exist {
		l.color = color
	}
	defer func() { l.color = ansicolor.White }()

	prefix, exist := LogPrefix[logLevel]
	if exist {
		l.prefix = prefix
	}
	defer func() { l.prefix = "" }()

	l.message = message
	defer func() { l.message = "" }()

	output := time.Now().Format(l.timeFormat)
	if len(l.prefix) != 0 {
		output += " " + l.Prefix()
	}
	output += " " + l.Message()

	fmt.Print(output)
}

//
// getter and setter of color logger
//

// Prefix returns the output prefix for the color logger.
func (l *ColorLogger) Prefix() string {
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.fillPrefix {
		return ansicolor.ToColor(l.color, l.prefix)
	}
	return l.prefix
}

// SetPrefix sets the output prefix for the color logger.
func (l *ColorLogger) SetPrefix(prefix string) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.prefix = prefix
}

// FillPrefix returns the logger configuration of whether to fill prefix with color or not.
func (l *ColorLogger) FillPrefix() bool {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.fillPrefix
}

// SetFillPrefix sets the logger configuration of whether to fill prefix with color or not.
func (l *ColorLogger) SetFillPrefix(fillPrefix bool) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.fillPrefix = fillPrefix
}

// Message returns the output message for the color logger.
func (l *ColorLogger) Message() string {
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.fillMessage {
		return ansicolor.ToColor(l.color, l.message)
	}
	return l.message
}

// SetMessage sets the output message for the color logger.
func (l *ColorLogger) SetMessage(message string) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.message = message
}

// FillMessage returns the logger configuration of whether to fill message with color or not.
func (l *ColorLogger) FillMessage() bool {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.fillMessage
}

// SetFillMessage sets the logger configuration of whether to fill message with color or not.
func (l *ColorLogger) SetFillMessage(fillMessage bool) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.fillMessage = fillMessage
}

// TimeFormat returns the output timeFormat for the color logger.
func (l *ColorLogger) TimeFormat() string {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.timeFormat
}

// SetTimeFormat sets the output timeFormat for the color logger.
func (l *ColorLogger) SetTimeFormat(timeFormat string) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.timeFormat = timeFormat
}

// Color returns the output color for the color logger.
func (l *ColorLogger) Color() ansicolor.Color {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.color
}

// SetColor sets the output color for the color logger.
func (l *ColorLogger) SetColor(color ansicolor.Color) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.color = color
}

// Level returns the log level for the color logger.
func (l *ColorLogger) Level() LogLevel {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.level
}

// SetLevel sets the log level for the color logger.
func (l *ColorLogger) SetLevel(logLevel LogLevel) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.level = logLevel
}

//
// getter and setter of standard color logger
//

// Prefix returns the output prefix for the standard color logger.
func Prefix() string {
	return std.Prefix()
}

// SetPrefix sets the output prefix for the standard color logger.
func SetPrefix(prefix string) {
	std.SetPrefix(prefix)
}

// FillPrefix returns the standard color logger configuration of whether to fill prefix with color or not.
func FillPrefix() bool {
	return std.FillPrefix()
}

// SetFillPrefix sets the standard color logger configuration of whether to fill prefix with color or not.
func SetFillPrefix(fillPrefix bool) {
	std.SetFillPrefix(fillPrefix)
}

// Message returns the output message for the standard color logger.
func Message() string {
	return std.Message()
}

// SetMessage sets the output message for the standard color logger.
func SetMessage(message string) {
	std.SetMessage(message)
}

// FillMessage returns the standard color logger configuration of whether to fill message with color or not.
func FillMessage() bool {
	return std.FillMessage()
}

// SetFillMessage sets the standard color logger configuration of whether to fill message with color or not.
func SetFillMessage(fillMessage bool) {
	std.SetFillMessage(fillMessage)
}

// TimeFormat returns the output timeFormat for the standard color logger.
func TimeFormat() string {
	return std.TimeFormat()
}

// SetTimeFormat sets the output timeFormat for the standard color logger.
func SetTimeFormat(timeFormat string) {
	std.SetTimeFormat(timeFormat)
}

// Color returns the output color for the standard color logger.
func Color() ansicolor.Color {
	return std.Color()
}

// SetColor sets the output color for the standard color logger.
func SetColor(color ansicolor.Color) {
	std.SetColor(color)
}

// Level returns the log level for the standard color logger.
func Level() LogLevel {
	return std.Level()
}

// SetLevel sets the log level for the standard color logger.
func SetLevel(logLevel LogLevel) {
	std.SetLevel(logLevel)
}

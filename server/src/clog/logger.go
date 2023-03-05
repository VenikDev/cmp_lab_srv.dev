package clog

import (
	"github.com/charmbracelet/log"
	"time"
)

var (
	// Logger
	// The code creates a logger instance using the log package in Go. The logger instance is created with three options:
	// - `log.WithTimestamp()` adds a timestamp to each log message.
	// - `log.WithTimeFormat(time.Kitchen)` sets the timestamp format to "3:04PM"
	//   (the format defined by the `time. Kitchen` constant).
	// - `log.WithCaller()` includes the filename and line number of the code
	//    that called the logger in each log message.
	// The resulting logger instance can be used to write log messages with the specified options.
	Logger = log.New(
		log.WithTimestamp(),
		log.WithTimeFormat(time.Kitchen),
		log.WithCaller(),
	)
)

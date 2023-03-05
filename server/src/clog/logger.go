package clog

import (
	"github.com/charmbracelet/log"
	"time"
)

var (
	// Custom logger
	Logger = log.New(
		log.WithTimestamp(),
		log.WithTimeFormat(time.Kitchen),
		log.WithCaller(),
	)
)

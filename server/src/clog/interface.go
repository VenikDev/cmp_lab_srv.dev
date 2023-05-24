package clog

func Info(msg interface{}, keyval ...interface{}) {
	Logger.Info(msg, keyval)
}

func Error(msg interface{}, keyval ...interface{}) {
	Logger.Error(msg, keyval)
}

func Fatal(msg interface{}, keyval ...interface{}) {
	Logger.Fatal(msg, keyval)
}

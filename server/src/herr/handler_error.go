package herr

import "cmp_lab/src/clog"

// HandlerError
// The code defines a function named `HandlerError` that takes two arguments: `err` of type `error` and `msg` of type
// `string`. If `err` is not `nil`,
// it logs the error message along with the given `msg` and the error object using a logger named `clog.Logger`.
func HandlerError(err error, msg string) {
	if err != nil {
		clog.Logger.Error(msg, "Error", err)
	}
}

// HandlerFatal
// This code defines a function named `HandlerFatal` that takes in two parameters - an `error` and a `string` message.
// If the error is not null, it logs a fatal message to a logger named `clog.Logger`,
// with the message and error found in the parameters.
func HandlerFatal(err error, msg string) {
	if err != nil {
		clog.Logger.Fatal(msg, "Error", err)
	}
}

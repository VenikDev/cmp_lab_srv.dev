package herr

import "comparisonLaboratories/src/clog"

func HandlerError(err error, msg string) {
	if err != nil {
		clog.Logger.Error(msg, "Error", err)
	}
}

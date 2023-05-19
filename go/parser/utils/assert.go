package utils

func AssertCallback(cnd bool, callback func()) {
	if cnd == true {
		callback()
	}
}

package pkg

import (
	"runtime"
)

func GetCurrentShell() (string, string) {
	currOS := runtime.GOOS
	if currOS == "windows" {
		return "cmd", "/C"
	}

	return "sh", "-c"
}

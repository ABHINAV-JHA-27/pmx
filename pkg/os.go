package pkg

import (
	"fmt"
	"runtime"
)

func GetCurrentShell() (string, string) {
	currOS := runtime.GOOS
	if currOS == "windows" {
		fmt.Println("** Windows OS detected **")
		return "cmd", "/C"
	}
	fmt.Println("** Unix-like OS detected **")
	return "sh", "-c"
}

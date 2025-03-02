package main

import (
	"github.com/ABHINAV-JHA-27/pmx/cmd"
	"github.com/ABHINAV-JHA-27/pmx/internal/logger"
)

func main() {
	logger.Init()
	cmd.Execute()
}

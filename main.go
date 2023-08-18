package main

import (
	"goftc/internal/checker"
	"goftc/internal/task"
)

func main() {
	in := checker.Configuration()
	if in != nil {
		task.CheckerTask(in)
	}

}

package main

import (
	"os"
	"runtime"

	"github.com/sanjid133/ksd/cmds"
	"k8s.io/apiserver/pkg/util/logs"
)

func main() {
	logs.InitLogs()
	defer logs.FlushLogs()

	if len(os.Getenv("GOMAXPROCS")) == 0 {
		runtime.GOMAXPROCS(runtime.NumCPU())
	}

	if err := cmds.NewRootCmd().Execute(); err != nil {
		os.Exit(1)
	}

	os.Exit(0)
}

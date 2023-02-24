package main

import (
	"math/rand"
	"os"
	"runtime"
	"time"

	"github.com/CSU-Apple-Lab/tinyTikTok/internal/apiserver"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	if len(os.Getenv("GOMAXPROCS")) == 0 {
		runtime.GOMAXPROCS(runtime.NumCPU())
	}

	apiserver.NewApp("apiserver").Run()
}

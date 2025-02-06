package main

import (
	"github.com/bhagyashriw777/gosc/gosh"
	"os"
)

func main() {
	for _, arg := range os.Args[1:] {
		gosh.RemoteShell(arg)
	}
}

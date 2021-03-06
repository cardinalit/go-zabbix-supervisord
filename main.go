package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/abrander/go-supervisord"
	"os"
	"strconv"
)

var SupervisordSock = "/tmp/supervisor.sock"

func main() {
	var zbxData *ZBXFormat
	var b []byte

	flag.StringVar(&SupervisordSock, "sock", SupervisordSock, "The full path to the socket\n")
	flag.Parse()

	cmd := flag.Arg(0)

	supervisor, err := supervisord.NewUnixSocketClient(SupervisordSock)
	if err != nil {
		panic(err.Error())
	}

	state, _ := supervisor.GetState()
	if state.Code != supervisord.StateCodeRunning && state.Name != supervisord.StateNameRunning  {
		fmt.Println("supervisord is not running or path to socket is incorrect.\n" +
			"You can type `-h | --help`")

		os.Exit(0)
	}

	switch cmd {
	case "discovery":
		zbxData = discovery(supervisor)

		b, err = json.Marshal(zbxData)
		if err != nil {
			panic(err.Error())
		}

		fmt.Println(string(b))
	case "state.name":
		ps := psStateName(supervisor, flag.Arg(1))

		fmt.Println(ps)
	case "process.ping":
		pid, err := strconv.ParseInt(flag.Arg(1), 10, 64)
		if err != nil {
			fmt.Printf("%v\n", err)
			os.Exit(1)
		}

		status := psPing(pid)
		fmt.Println(status)
	default:
		fmt.Println("No argument was passed or passed argument is incorrect.\n" +
			"Type `-h | --help`")
	}
}

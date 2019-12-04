package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/abrander/go-supervisord"
	"os"
)

type ZabbixJsonFormat struct {
	Data 	[]*SupervisorProcess	`json:"data"`
}

type SupervisorProcess struct {
	ProcessName		string			`json:"{#PROCESS.NAME}"`
	ProcessGroup	string			`json:"{#PROCESS.GROUP}"`
	ProcessPID		int				`json:"{#PROCESS.PID}"`
}

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
		os.Exit(1)
	}

	processList, err := supervisor.GetAllProcessInfo()
	if err != nil {
		panic(err.Error())
	}

	toZabbix := ZabbixJsonFormat{Data: []*SupervisorProcess{}}
	for _, process := range processList {
		n := &SupervisorProcess{
			ProcessName:  process.Name,
			ProcessGroup: process.Group,
			ProcessPID:   process.Pid,
		}

		toZabbix.Data = append(toZabbix.Data, n)
	}

	b, err := json.Marshal(toZabbix)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(string(b))
}

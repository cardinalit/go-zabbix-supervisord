package main

import (
	"github.com/abrander/go-supervisord"
	"os"
	"syscall"
)

func discovery(supervisor *supervisord.Client) *ZBXFormat {
	processList, err := supervisor.GetAllProcessInfo()
	if err != nil {
		panic(err.Error())
	}

	zbx := &ZBXFormat{Data: []*SupervisorProcess{}}
	for _, p := range processList {
		zbx.Data = append(zbx.Data, &SupervisorProcess{
			ProcessName:  p.Name,
			ProcessGroup: p.Group,
			ProcessPID:   p.Pid,
		})
	}

	return zbx
}

func psStateName(supervisor *supervisord.Client, processName string) string {
	info, err := supervisor.GetProcessInfo(processName)
	if err != nil {
		panic(err.Error())
	}

	return info.StateName
}

func psPing(pid int64) uint {
	ps, _ := os.FindProcess(int(pid))
	err   := ps.Signal(syscall.Signal(0))

	if err == nil {
		return 1
	}

	return 0
}

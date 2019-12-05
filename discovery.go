package main

import "github.com/abrander/go-supervisord"

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

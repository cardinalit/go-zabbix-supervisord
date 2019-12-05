package main

type ZBXFormat struct {
	Data 	[]*SupervisorProcess	`json:"data"`
}

type SupervisorProcess struct {
	ProcessName		string			`json:"{#PROCESS.NAME}"`
	ProcessGroup	string			`json:"{#PROCESS.GROUP}"`
	ProcessPID		int				`json:"{#PROCESS.PID}"`
}
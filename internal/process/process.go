package process

import (
	"os/exec"
	"time"

	"github.com/google/uuid"
)

type Process struct {
	ID          string
	Pid         int
	Name        string
	Cmd         string
	Restart     bool
	Status      string
	CPUUsage    float64
	MemoryUsage float64
	Uptime      time.Time
	CmdProcess  *exec.Cmd
}

func NewProcess(name, cmd string, restart bool) *Process {
	id := uuid.New().String()
	if name == "" {
		name = "process" + "_" + id[:4]
	}
	return &Process{
		ID:          id,
		Name:        name,
		Cmd:         cmd,
		Status:      "stopped",
		CmdProcess:  nil,
		CPUUsage:    0,
		MemoryUsage: 0,
		Restart:     restart,
		Uptime:      time.Time{},
	}
}

func (p *Process) StartProcess(shell, flag string) {
	proc := exec.Command(shell, flag, p.Cmd)
	if err := proc.Start(); err != nil {
		panic(err)
	}
	p.CmdProcess = proc
	p.Pid = proc.Process.Pid
	p.Status = "running"
	p.Uptime = time.Now()
}

func (p *Process) StopProcess() {
	if p.CmdProcess != nil && p.Status == "running" {
		if err := p.CmdProcess.Process.Kill(); err != nil {
			panic(err)
		}
		p.Status = "stopped"
		p.Uptime = time.Time{}
	}
}

func (p *Process) RestartProcess(shell, flag string) {
	p.StopProcess()
	time.Sleep(1 * time.Second)
	p.StartProcess(shell, flag)
}

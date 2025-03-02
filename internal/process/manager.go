package process

import (
	"time"

	"github.com/ABHINAV-JHA-27/pmx/pkg"
)

type Manager struct {
	Shell     string
	Flag      string
	Processes map[string]*Process
}

func NewManager() *Manager {
	shell, flag := pkg.GetCurrentShell()
	return &Manager{
		Shell:     shell,
		Flag:      flag,
		Processes: make(map[string]*Process),
	}
}

func (m *Manager) AddProcess(p *Process) {
	m.Processes[p.ID] = p
}

func (m *Manager) RemoveProcess(id string) {
	delete(m.Processes, id)
}

func (m *Manager) GetProcess(id string) *Process {
	return m.Processes[id]
}

func (m *Manager) GetProcesses() []*Process {
	var processes []*Process
	for _, p := range m.Processes {
		processes = append(processes, p)
	}
	return processes
}

func (m *Manager) StartProcess(name, cmd string, restart bool) {
	p := NewProcess(name, cmd, restart)
	p.StartProcess(m.Shell, m.Flag)
	m.AddProcess(p)
}

func (m *Manager) StopProcess(id string) {
	p := m.GetProcess(id)
	if p != nil {
		p.StopProcess()
	}
}

func (m *Manager) RestartProcess(id string) {
	p := m.GetProcess(id)
	if p != nil {
		p.RestartProcess(m.Shell, m.Flag)
	}
}

func (m *Manager) GetProcessStatus(id string) string {
	p := m.GetProcess(id)
	if p != nil {
		return p.Status
	}
	return ""
}

func (m *Manager) GetProcessCPUUsage(id string) float64 {
	p := m.GetProcess(id)
	if p != nil {
		return p.CPUUsage
	}
	return 0
}

func (m *Manager) GetProcessMemory(id string) float64 {
	p := m.GetProcess(id)
	if p != nil {
		return p.MemoryUsage
	}
	return 0
}

func (m *Manager) GetProcessUptime(id string) string {
	currentTime := time.Now()
	p := m.GetProcess(id)
	if p != nil {
		return currentTime.Sub(p.Uptime).String()
	}
	return ""
}

package process

type Manager struct {
	Processes []*Process
}

func NewManager() *Manager {
	return &Manager{}
}

func (m *Manager) AddProcess(p *Process) {
	m.Processes = append(m.Processes, p)
}

func (m *Manager) RemoveProcess(pid int) {
	for i, p := range m.Processes {
		if p.Pid == pid {
			m.Processes = append(m.Processes[:i], m.Processes[i+1:]...)
			break
		}
	}
}

func (m *Manager) GetProcess(pid int) *Process {
	for _, p := range m.Processes {
		if p.Pid == pid {
			return p
		}
	}
	return nil
}

func (m *Manager) GetProcesses() []*Process {
	return m.Processes
}

func (m *Manager) StartProcess(name, cmd string, restart bool) {
	p := NewProcess(name, cmd, restart)
	p.StartProcess()
	m.AddProcess(p)
}

func (m *Manager) StopProcess(pid int) {
	p := m.GetProcess(pid)
	if p != nil {
		p.StopProcess()
	}
}

func (m *Manager) RestartProcess(pid int) {
	p := m.GetProcess(pid)
	if p != nil {
		p.RestartProcess()
	}
}

func (m *Manager) GetProcessStatus(pid int) string {
	p := m.GetProcess(pid)
	if p != nil {
		return p.Status
	}
	return ""
}

func (m *Manager) GetProcessCPUUsage(pid int) float64 {
	p := m.GetProcess(pid)
	if p != nil {
		return p.CPUUsage
	}
	return 0
}

func (m *Manager) GetProcessMemory(pid int) float64 {
	p := m.GetProcess(pid)
	if p != nil {
		return p.MemoryUsage
	}
	return 0
}

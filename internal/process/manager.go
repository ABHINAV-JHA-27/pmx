package process

type Manager struct {
	Processes []*Process
}

func NewManager() *Manager {
	return &Manager{}
}

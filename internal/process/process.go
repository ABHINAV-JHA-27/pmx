package process

type Process struct {
	ID       string
	Cmd      string
	Args     []string
	Pid      int
	Status   string
	Restart  bool
	CPUUsage float64
	Memory   float64
}

func NewProcess(id, cmd string, args []string) *Process {
	return &Process{
		ID:     id,
		Cmd:    cmd,
		Args:   args,
		Status: "stopped",
	}
}

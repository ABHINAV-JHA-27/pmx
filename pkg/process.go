package pkg

import (
	"os"
	"time"

	"github.com/jedib0t/go-pretty/v6/table"
)

type Process struct {
	ID          string
	Pid         int
	Name        string
	Status      string
	CPUUsage    float64
	MemoryUsage float64
	Uptime      time.Time
}

func PrintProcess(process []Process) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "ID", "PID", "Name", "Status", "CPU Usage", "Memory Usage", "Uptime"})
	for i, p := range process {
		t.AppendRow([]interface{}{i, p.ID, p.Pid, p.Name, p.Status, p.CPUUsage, p.MemoryUsage, p.Uptime})
		t.AppendSeparator()
	}
	t.Render()
}

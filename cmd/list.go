package cmd

import (
	"fmt"
	"strconv"

	"github.com/ABHINAV-JHA-27/pmx/pkg"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all processes",
	Long:  "List all processes currently running",
	Run: func(cmd *cobra.Command, args []string) {
		_name := cmd.Flag("name").Value.String()
		_id := cmd.Flag("id").Value.String()
		_status := cmd.Flag("status").Value.String()
		_pid := cmd.Flag("pid").Value.String()

		if _name != "" {
			processes := manager.GetProcesses()
			print_proc := make([]pkg.Process, 0)
			for _, p := range processes {
				if p.Name == _name {
					print_proc = append(print_proc, pkg.Process{
						ID:          p.ID,
						Pid:         p.Pid,
						Name:        p.Name,
						Status:      p.Status,
						CPUUsage:    p.CPUUsage,
						MemoryUsage: p.MemoryUsage,
						Uptime:      p.Uptime,
					})
				}
			}

			pkg.PrintProcess(print_proc)
			return
		}
		if _id != "" {
			processes := manager.GetProcess(_id)
			print_proc := pkg.Process{
				ID:          processes.ID,
				Pid:         processes.Pid,
				Name:        processes.Name,
				Status:      processes.Status,
				CPUUsage:    processes.CPUUsage,
				MemoryUsage: processes.MemoryUsage,
				Uptime:      processes.Uptime,
			}

			pkg.PrintProcess([]pkg.Process{print_proc})
			return
		}
		if _status != "" {
			processes := manager.GetProcesses()
			print_proc := make([]pkg.Process, 0)
			for _, p := range processes {
				if p.Status == _status {
					print_proc = append(print_proc, pkg.Process{
						ID:          p.ID,
						Pid:         p.Pid,
						Name:        p.Name,
						Status:      p.Status,
						CPUUsage:    p.CPUUsage,
						MemoryUsage: p.MemoryUsage,
						Uptime:      p.Uptime,
					})
				}
			}

			pkg.PrintProcess(print_proc)
			return
		}
		if _pid != "" {
			pidInt, err := strconv.Atoi(_pid)
			if err != nil {
				fmt.Println("Invalid PID")
				return
			}
			processes := manager.GetProcesses()
			print_proc := make([]pkg.Process, 0)
			for _, p := range processes {
				if p.Pid == pidInt {
					print_proc = append(print_proc, pkg.Process{
						ID:          p.ID,
						Pid:         p.Pid,
						Name:        p.Name,
						Status:      p.Status,
						CPUUsage:    p.CPUUsage,
						MemoryUsage: p.MemoryUsage,
						Uptime:      p.Uptime,
					})
				}
			}

			pkg.PrintProcess(print_proc)
			return
		}

		processes := manager.GetProcesses()
		print_proc := make([]pkg.Process, 0)
		for _, p := range processes {
			print_proc = append(print_proc, pkg.Process{
				ID:          p.ID,
				Pid:         p.Pid,
				Name:        p.Name,
				Status:      p.Status,
				CPUUsage:    p.CPUUsage,
				MemoryUsage: p.MemoryUsage,
				Uptime:      p.Uptime,
			})
		}

		pkg.PrintProcess(print_proc)

	},
}

func init() {
	listCmd.Flags().StringP("name", "n", "", "Name of the process")
	listCmd.Flags().StringP("id", "i", "", "ID of the process")
	listCmd.Flags().StringP("status", "s", "", "Status of the process")
	listCmd.Flags().StringP("pid", "p", "", "PID of the process")
	rootCmd.AddCommand(listCmd)
}

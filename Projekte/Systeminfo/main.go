package main

import (
	"fmt"
	host2 "github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	"runtime"
)

func main() {
	mem, _ := mem.VirtualMemory()
	host, _ := host2.Info()
	fmt.Println("Arch: ", runtime.GOARCH)
	fmt.Println("OS: ", runtime.GOOS)
	fmt.Println("Kernel: ", host.KernelVersion)
	fmt.Println("CPUs:i ", runtime.NumCPU())
	fmt.Printf("RAM (Used): %f%%\n", mem.UsedPercent)
	fmt.Println("Hostname: ", host.Hostname)
	fmt.Println("Uptime: ", host.Uptime/60, "Minuten (", (host.Uptime/60)/60, "h)")
	fmt.Println("Host-ID: ", host.HostID)
}

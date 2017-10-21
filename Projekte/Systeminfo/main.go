package main

import (
	"runtime"
	"fmt"
	"github.com/shirou/gopsutil/mem"
	host2 "github.com/shirou/gopsutil/host"
)

func main(){
	mem, _ := mem.VirtualMemory()
	host, _ := host2.Info()
	fmt.Printf("Arch:\t%s\n", runtime.GOARCH)
	fmt.Printf("OS:\t%s\n", runtime.GOOS)
	fmt.Println("Kernel:", host.KernelVersion)
	fmt.Printf("CPUs:\t%d\n", runtime.NumCPU())
	fmt.Printf("RAM (Used): %f%%\n", mem.UsedPercent)
	fmt.Println("Hostname: ", host.Hostname)
	fmt.Println("Uptime: ", host.Uptime/60, "Minuten (", (host.Uptime/60)/60, "h)")
	fmt.Println("Host-ID: ", host.HostID)
}

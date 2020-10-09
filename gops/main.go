package main

import (
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
	"time"
)

func main() {
	fmt.Println("--------------------------------------------------------------------------------------")
	vmStat, _ := mem.VirtualMemory()
	strPrint := fmt.Sprintf("VIRTUAL MEMEORY [%+v]", vmStat)
	fmt.Println(strPrint)
	fmt.Println("--------------------------------------------------------------------------------------")
	cpuStat, _ := cpu.Info()
	strPrint = fmt.Sprintf("CPU INFORMATION [%+v]", cpuStat)
	fmt.Println(strPrint)
	fmt.Println("--------------------------------------------------------------------------------------")
	cpuPercent, _ := cpu.Percent(1*time.Second, true)
	strPrint = fmt.Sprintf("CPU PERCENT [%+v]", cpuPercent)
	fmt.Println(strPrint)
	fmt.Println("--------------------------------------------------------------------------------------")
	ioCounter, _ := disk.IOCounters()
	strPrint = fmt.Sprintf("DISK IO COUNTER [%+v]", ioCounter)
	fmt.Println(strPrint)
	fmt.Println("--------------------------------------------------------------------------------------")
	usageStat, _ := disk.Usage("D:\\")
	strPrint = fmt.Sprintf("DISK USAGE [%+v]", usageStat)
	fmt.Println(strPrint)
	fmt.Println("--------------------------------------------------------------------------------------")
	netStat, _ := net.Connections("tcp")
	strPrint = fmt.Sprintf("NET CONNECTIONS [%+v]", netStat)
	fmt.Println(strPrint)

	fmt.Println("--------------------------------------------------------------------------------------")
	netIO, _ := net.IOCounters(true)
	strPrint = fmt.Sprintf("NET IO COUNTER [%+v]", netIO)
	fmt.Println(strPrint)
}

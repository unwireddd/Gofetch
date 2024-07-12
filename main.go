package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"reflect"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/jaypipes/ghw"
	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/host"
	"github.com/shirou/gopsutil/v4/mem"
)

// Converting bytes to megabytes
func conv(unit uint64) float64 {
	result := unit / 1000000
	return float64(result)
}

func main() {

	bgDark := color.New(color.BgHiBlack)
	bgRed := color.New(color.BgHiRed)
	bgGreen := color.New(color.BgHiGreen)
	bgYellow := color.New(color.BgHiYellow)
	bgBlue := color.New(color.BgBlue)
	bgPurple := color.New(color.BgHiMagenta)
	bgAqua := color.New(color.BgHiBlue)
	bgWhite := color.New(color.BgHiWhite)

	// Obtaining the CPU info
	cpu, _ := cpu.Info()
	cpuModel := cpu[1].ModelName
	//fmt.Println(cpuModel)

	// Obtaining the hostname info
	host, _ := host.Info()
	hostname := host.Hostname
	uptime := host.Uptime
	uptime = uptime / 3600

	// Obtaining the OS and kernel version info
	osversion := host.Platform
	operating := host.OS
	kernel := host.KernelVersion
	kernelArch := host.KernelArch

	// Obtaining the GPU info
	graphs, _ := ghw.GPU(ghw.WithDisableWarnings())
	graphInfo := graphs.GraphicsCards
	firstElem := graphInfo[0]
	gpuout := graphInfo[0]
	elemType := reflect.TypeOf(firstElem)
	if elemType.Kind() == reflect.Struct {
		gpuout = graphInfo[0]
	} else {
		gpuout = graphInfo[1]
	}

	graphFinal := gpuout.DeviceInfo.Product.Name
	//fmt.Println(graphInfo[1])

	// Obtaining the BIOS info
	bios, _ := ghw.BIOS()
	biosVen := bios.Vendor

	// Obtaining the hard drive info
	drive, _ := ghw.Block()
	test2 := drive.Disks[0]
	driveFinal := test2.Model

	if driveFinal == "unknown" {
		driveFinal = "unknown / encrypted"
	}

	// Obtaining the timezone info
	timezone, _ := time.Now().Local().Zone()

	// Obtaining the memory info
	virmem, _ := mem.VirtualMemory()
	swap, _ := mem.SwapMemory()
	totalmem := virmem.Total
	usedmem := virmem.Used
	percentused := virmem.UsedPercent
	totalswap := swap.Total

	// Obtaining the init system info
	getInit := exec.Command("ps", "-p", "1", "-o", "comm=")
	init, _ := getInit.Output()
	//fmt.Println(init)
	initUn := string(init)
	initOut := strings.TrimSpace(initUn)
	//fmt.Println(initOut)

	// Obtaining the install date info

	getDate := exec.Command("sh", "-c", "ls -alct /|tail -1|awk '{print $6, $7}'")
	iDate, _ := getDate.Output()
	dateUn := string(iDate)
	dateOut := strings.TrimSpace(dateUn)
	//fmt.Println(dateOut)

	// Printing out the basic info
	var distro string
	kolor := color.New(color.BgBlack)
	switch {
	case osversion == "arch":
		distro = "/etc/Gofetch/Ascii/ascii"
		blue := color.New(color.FgHiBlue)
		kolor = blue.Add(color.Bold)
	case osversion == "artix":
		distro = "/etc/Gofetch/Ascii/artix"
		blue := color.New(color.FgHiBlue)
		kolor = blue.Add(color.Bold)
	case osversion == "debian":
		distro = "/etc/Gofetch/Ascii/debian"
		red := color.New(color.FgHiRed)
		kolor = red.Add(color.Bold)
	case osversion == "fedora":
		distro = "/etc/Gofetch/Ascii/fedora"
		blue := color.New(color.FgHiBlue)
		kolor = blue.Add(color.Bold)
	case osversion == "gentoo":
		distro = "/etc/Gofetch/Ascii/gentoo"
		purple := color.New(color.FgHiMagenta)
		kolor = purple.Add(color.Bold)
	case osversion == "popos":
		distro = "/etc/Gofetch/Ascii/popos"
		grey := color.New(color.FgHiWhite)
		kolor = grey.Add(color.Bold)
	case osversion == "slackware":
		distro = "/etc/Gofetch/Ascii/slackware"
		blue := color.New(color.FgHiBlue)
		kolor = blue.Add(color.Bold)
	case osversion == "ubuntu":
		distro = "/etc/Gofetch/Ascii/ubuntu"
		yellow := color.New(color.FgHiYellow)
		kolor = yellow.Add(color.Bold)
	case osversion == "void":
		distro = "/etc/Gofetch/Ascii/void"
		green := color.New(color.FgHiGreen)
		kolor = green.Add(color.Bold)
	default:
		distro = "/etc/Gofetch/Ascii/linux"
		grey := color.New(color.FgHiWhite)
		kolor = grey.Add(color.Bold)

	}
	fmt.Println("")
	ascii, _ := os.Open(distro)
	scanner := bufio.NewScanner(ascii)
	for scanner.Scan() {
		kolor.Println(scanner.Text())
	}

	kolor.Println("")
	kolor.Print("Hostname: ")
	fmt.Println(hostname)

	kolor.Print("Operating System: ")
	fmt.Println(osversion, operating)

	kolor.Print("Installation date: ")
	fmt.Println(dateOut)

	kolor.Print("Kernel: ")
	fmt.Println(kernel, kernelArch)

	kolor.Print("Init: ")
	fmt.Println(initOut)

	kolor.Print("RAM: ")
	fmt.Print(conv(usedmem), "M", " / ", conv(totalmem), "M", " (", int(percentused), "%", ")")
	fmt.Println()

	kolor.Print("Swap: ")
	fmt.Print(conv(totalswap), "M")
	fmt.Println()

	kolor.Print("Processor: ")
	fmt.Println(cpuModel)

	kolor.Print("Graphics: ")
	fmt.Println(graphFinal)

	kolor.Print("Hard Drive: ")
	fmt.Println(driveFinal)

	kolor.Print("BIOS: ")
	fmt.Println(biosVen)

	kolor.Print("Uptime: ")
	fmt.Println(uptime, "hours")
	kolor.Print("Timezone: ")
	fmt.Println(timezone)
	bgDark.Print("  ")
	bgRed.Print("  ")
	bgGreen.Print("  ")
	bgYellow.Print("  ")
	bgBlue.Print("  ")
	bgPurple.Print("  ")
	bgAqua.Print("  ")
	bgWhite.Print("  ")
	fmt.Println("")
	kolor.Println("")

}

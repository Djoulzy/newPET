package main

import (
	"fmt"
	"log"
	"newPET/config"
	"newPET/crtc"
	"os"
	"strconv"
	"time"

	"github.com/Djoulzy/Tools/clog"
	"github.com/Djoulzy/Tools/confload"
	"github.com/Djoulzy/emutools/mem"
	"github.com/Djoulzy/emutools/mos6510"
	"github.com/Djoulzy/emutools/render"
	"github.com/mattn/go-tty"
)

const (
	ramSize     = 65536
	kernalSize  = 4096
	basicSize   = 12288
	ioSize      = 2048
	chargenSize = 2048
	editorSize  = 2048
	blanckSize  = 12288

	nbMemLayout = 1

	Stopped = 0
	Paused  = 1
	Running = 2
)

var (
	conf = &config.ConfigData{}

	cpu mos6510.CPU

	RAM      []byte
	IO       []byte
	KERNAL   []byte
	BASIC    []byte
	EDITOR   []byte
	CHARGEN  []byte
	BLANK    []byte
	MEM      mem.BANK
	IOAccess mem.MEMAccess

	outputDriver render.SDL2Driver
	CRTC         crtc.CRTC
	run          bool
	trace        bool
)

// func init() {
// 	// This is needed to arrange that main() runs on main thread.
// 	// See documentation for functions that are only allowed to be called from the main thread.
// 	runtime.LockOSThread()
// }

func setup() {
	// ROMs & RAM Setup
	RAM = make([]byte, ramSize)
	IO = make([]byte, ioSize)
	BLANK = make([]byte, blanckSize)
	KERNAL = mem.LoadROM(kernalSize, "assets/roms/kernal-4.901465-22.bin")
	BASIC = mem.LoadROM(basicSize, "assets/roms/basic-4.901465-23-20-21.bin")
	// BASIC2 = mem.LoadROM(basicSize, "assets/roms/basic-4-c000.901465-20.bin")
	// BASIC3 = mem.LoadROM(basicSize, "assets/roms/basic-4-d000.901465-21.bin")
	CHARGEN = mem.LoadROM(chargenSize, "assets/roms/characters-2.901447-10.bin")
	EDITOR = mem.LoadROM(editorSize, "assets/roms/edit-4-40-n-50Hz.901498-01.bin")

	mem.Clear(RAM, 0, 0x00)
	mem.Clear(IO, 0, 0x00)
	// mem.DisplayCharRom(CHARGEN, 1, 8, 16)

	// RAM[0x0001] = 0x00
	// MEM = mem.InitBanks(nbMemLayout, &RAM[0x0001])
	var test byte = 0
	MEM = mem.InitBanks(nbMemLayout, &test)
	IOAccess = &accessor{}

	// MEM Setup
	memLayouts()

	outputDriver = render.SDL2Driver{}
	CRTC.Init(RAM, IO, CHARGEN, &outputDriver, conf)

	// CPU Setup
	cpu.Init(conf.CPUModel, conf.Mhz, &MEM, conf.Debug || conf.Disassamble)
}

func input() {
	dumpAddr := ""
	var keyb *tty.TTY
	keyb, _ = tty.Open()

	for {
		r, _ := keyb.ReadRune()
		switch r {
		case 's':
			MEM.DumpStack(cpu.SP)
		case 'z':
			MEM.Dump(0)
		case 'x':
			// DumpMem(&pla, "memDump.bin")
		case 'r':
			run = true
			trace = false
		case 'l':
			// LoadPRG(&pla, "./prg/GARDEN.prg")
			LoadPRG(&MEM, conf.LoadPRG)
			// addr, _ := LoadPRG(mem.Val, conf.LoadPRG)
			// cpu.GoTo(addr)
		case ' ':
			fmt.Printf("%s\n", cpu.FullDebug)
			trace = true
			run = true
		case 'w':
			fmt.Printf("\nFill Color RAM")
			for i := 0xD800; i < 0xDC00; i++ {
				MEM.Write(uint16(i), 0)
			}
			// for i := 0x0800; i < 0x0C00; i++ {
			// 	IO[uint16(i)] = 0
			// }
		case 'q':
			fmt.Printf("%s\n", cpu.FullDebug)
			os.Exit(0)
		default:
			dumpAddr += string(r)
			fmt.Printf("%c", r)
			if len(dumpAddr) == 4 {
				hx, _ := strconv.ParseInt(dumpAddr, 16, 64)
				MEM.Dump(uint16(hx))
				dumpAddr = ""
			}
		}

	}
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Now().Sub(start)
	log.Printf("%s took %s", name, elapsed)
}

func RunEmulation() {
	var speed float64
	// var key byte
	// defer timeTrack(time.Now(), "RunEmulation")
	for {
		CRTC.Run(!run)

		// if MEM.Read(0xC000) == 0 {
		// 	key = keyMap[InputLine.KeyCode]
		// 	if InputLine.Mode == 1073742048 {
		// 		key -= 0x40
		// 	}
		// 	MEM.Write(0xC000, key)
		// 	InputLine.KeyCode = 0
		// 	InputLine.Mode = 0
		// }

		speed = cpu.NextCycle()

		if cpu.CycleCount == 1 {
			if trace {
				run = false
			}
			outputDriver.DumpCode(cpu.FullInst)
			outputDriver.SetSpeed(speed)
			if conf.Breakpoint == cpu.InstStart {
				fmt.Printf("%s\n", cpu.FullDebug)
				trace = true
			}
		}
	}
}

func main() {
	// var exit chan bool
	// exit = make(chan bool)

	confload.Load("config.ini", conf)

	clog.LogLevel = conf.LogLevel
	clog.StartLogging = conf.StartLogging
	if conf.FileLog != "" {
		clog.EnableFileLog(conf.FileLog)
	}

	// f, err := os.Create("newC64.prof")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// pprof.StartCPUProfile(f)
	// defer pprof.StopCPUProfile()

	setup()
	go input()

	run = true
	trace = false
	outputDriver.ShowCode = true
	outputDriver.ShowFps = true

	go RunEmulation()
	outputDriver.Run(true)

	// cpu.DumpStats()
	// <-exit
}

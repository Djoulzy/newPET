package main

import (
	"io/ioutil"
	"newPET/config"
	"newPET/crtc"
	"newPET/graphic"
	"newPET/mem"
	"runtime"
)

const (
	ramSize     = 65536
	chargenSize = 4096
	ioSize      = 4096
)

var (
	conf             config.ConfigData
	RAM, IO, CHARGEN []byte
	CRTC             crtc.CRTC
	outputDriver     graphic.Driver
)

func init() {
	// This is needed to arrange that main() runs on main thread.
	// See documentation for functions that are only allowed to be called from the main thread.
	runtime.LockOSThread()
}

func LoadData(mem []byte, file string, memStart uint16) error {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}

	for i, val := range content {
		mem[memStart+uint16(i)] = val
	}
	return nil
}

func start() {
	conf.Disassamble = false

	RAM = make([]byte, ramSize)
	mem.Clear(RAM)
	IO = make([]byte, ioSize)
	mem.Clear(IO)
	CHARGEN = mem.LoadROM(chargenSize, "assets/roms/char.bin")

	LoadData(RAM, "assets/roms/bruce2.bin", 0xE000)

	outputDriver = &graphic.SDLDriver{}
	CRTC.Init(RAM, IO, CHARGEN, outputDriver, &conf)
	CRTC.BankSel = 0
	CRTC.Write(crtc.REG_MEM_LOC, 0x78)
	CRTC.Write(crtc.REG_CTRL1, 0x3B)
	CRTC.Write(crtc.REG_CTRL2, 0x18)
	CRTC.Write(crtc.REG_BORDER_COL, 0x0E)
	CRTC.Write(crtc.REG_BGCOLOR_0, 0x00)
	CRTC.Write(crtc.REG_BGCOLOR_1, 0x01)
	CRTC.Write(crtc.REG_BGCOLOR_2, 0x02)
	CRTC.Write(crtc.REG_BGCOLOR_3, 0x03)

}

func main() {
	start()
	for {
		CRTC.Run(false)
		// vic.Stats()
	}
}

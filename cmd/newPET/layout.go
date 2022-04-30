package main

import (
	"github.com/Djoulzy/emutools/mem"
)

func memLayouts() {
	MEM.Layouts[0] = mem.InitConfig(ramSize)
	MEM.Layouts[0].Attach("RAM", 0x0000, RAM, mem.READWRITE)
	MEM.Layouts[0].Attach("BASIC", 0xB000, BASIC, mem.READONLY)
	MEM.Layouts[0].Attach("EDITOR", 0xE000, EDITOR, mem.READONLY)
	MEM.Layouts[0].Attach("IO", 0xE800, IO, mem.READWRITE)
	MEM.Layouts[0].Attach("KERNAL", 0xF000, KERNAL, mem.READONLY)
	MEM.Layouts[0].Accessor("IO", IOAccess)
	MEM.Layouts[0].Show()
}

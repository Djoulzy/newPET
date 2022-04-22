package main

import (
	"newPET/mem"
)

func memLayouts() {
	MEM.Layouts[0] = mem.InitConfig(8, ramSize)
	MEM.Layouts[0].Attach("RAM", 0, 0, RAM, mem.READWRITE)
	MEM.Layouts[0].Attach("BLANK", 1, 0x9000, BLANK, mem.READONLY)
	MEM.Layouts[0].Attach("BASIC-1", 2, 0xB000, BASIC1, mem.READONLY)
	MEM.Layouts[0].Attach("BASIC-2", 3, 0xC000, BASIC2, mem.READONLY)
	MEM.Layouts[0].Attach("BASIC-3", 4, 0xD000, BASIC3, mem.READONLY)
	MEM.Layouts[0].Attach("EDITOR", 5, 0xE000, EDITOR, mem.READONLY)
	MEM.Layouts[0].Attach("IO", 6, 0xE800, IO, mem.READWRITE)
	MEM.Layouts[0].Attach("KERNAL", 7, 0xF000, KERNAL, mem.READONLY)
	MEM.Layouts[0].Accessor(6, IOAccess)
	// MEM.Show()
}

package main

import (
	"newPET/mem"
)

func memLayouts() {
	MEM.Layouts[0] = mem.InitConfig(5, ramSize)
	MEM.Layouts[0].Attach("RAM", 0, 0, RAM, mem.READWRITE)
	MEM.Layouts[0].Attach("BASIC-1", 1, 0xC000, BASIC1, mem.READONLY)
	MEM.Layouts[0].Attach("BASIC-2", 1, 0xC000, BASIC2, mem.READONLY)
	MEM.Layouts[0].Attach("BASIC-3", 2, 0xD000, BASIC3, mem.READONLY)
	MEM.Layouts[0].Attach("EDITOR", 3, 0xE000, EDITOR, mem.READONLY)
	MEM.Layouts[0].Attach("KERNAL", 4, 0xF000, KERNAL, mem.READONLY)
	MEM.Show()
}

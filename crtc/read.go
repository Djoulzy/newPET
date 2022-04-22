package crtc

func (C *CRTC) Read(addr uint16) byte {
	reg := addr - ((addr >> 6) << 6)
	// clog.Trace("VIC", "Read", "addr: %04X - Reg: %02X (%d)", addr, reg, reg)
	switch reg {
	case REG_MEM_LOC:
		// log.Printf("Read base: %04X", C.Reg[reg])
		return C.Reg[reg]
	case REG_RASTER:
		// log.Printf("RasterIRQ: %04X", C.Reg[reg])
		return C.Reg[reg]
	default:
		return C.Reg[reg]
	}
}

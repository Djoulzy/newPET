package crtc

func (C *CRTC) Write(addr uint16, val byte) {

	reg := addr - ((addr >> 6) << 6)
	// clog.Trace("VIC", "Write", "addr: %04X - Reg: %02X (%d)", addr, reg, reg)
	switch reg {
	case REG_X_SPRT_0:
		fallthrough
	case REG_Y_SPRT_0:
		fallthrough
	case REG_X_SPRT_1:
		fallthrough
	case REG_Y_SPRT_1:
		fallthrough
	case REG_X_SPRT_2:
		fallthrough
	case REG_Y_SPRT_2:
		fallthrough
	case REG_X_SPRT_3:
		fallthrough
	case REG_Y_SPRT_3:
		fallthrough
	case REG_X_SPRT_4:
		fallthrough
	case REG_Y_SPRT_4:
		fallthrough
	case REG_X_SPRT_5:
		fallthrough
	case REG_Y_SPRT_5:
		fallthrough
	case REG_X_SPRT_6:
		fallthrough
	case REG_Y_SPRT_6:
		fallthrough
	case REG_X_SPRT_7:
		fallthrough
	case REG_Y_SPRT_7:
		fallthrough
	case REG_MSBS_X_COOR:
		C.Reg[reg] = val
	case REG_CTRL1:
		// log.Printf("REG_CTRL1: %08b", val)
		newMode := (C.MODE & 0b00010000) | val&0b01100000
		if newMode != C.MODE {
			C.MODE = newMode
			// log.Printf("Graphic mode: %08b", C.MODE)
		}
		C.RasterIRQ &= 0x7FFF
		C.RasterIRQ |= uint16(val&RST8) << 8
		C.Reg[REG_CTRL1] = (C.Reg[REG_CTRL1] & 0b1000000) | (val & 0b01111111)
	case REG_RASTER:
		C.RasterIRQ = C.RasterIRQ&0x8000 + uint16(val)
		// log.Printf("RasterIRQ: %04X", C.RasterIRQ)
	case REG_LP_X:
		fallthrough
	case REG_LP_Y:
		fallthrough
	case REG_SPRT_ENABLED:
		fallthrough
	case REG_CTRL2:
		newMode := (C.MODE & 0b01100000) | val&0b00010000
		if newMode != C.MODE {
			C.MODE = newMode
			// log.Printf("Graphic mode: %08b", C.MODE)
		}
		C.Reg[reg] = val
	case REG_SPRT_Y_EXP:
		C.Reg[reg] = val
	case REG_MEM_LOC:
		C.ScreenBase = uint16(val&0b11110000) << 6
		C.CharBase = uint16(val&0b00001110) << 10
		// log.Printf("VIC Screenbase: %04X - Charbase: %04X", C.ScreenBase, C.CharBase)
		C.Reg[reg] = val
	case REG_IRQ:
		C.Reg[REG_IRQ] &= ^val
		if C.Reg[REG_IRQ]&0b10000000 == 0 {
			*C.IRQ_Pin = 0
		}
	case REG_IRQ_ENABLED:
		fallthrough
	case REG_SPRT_DATA_PRIORITY:
		fallthrough
	case REG_SPRT_MLTCOLOR:
		fallthrough
	case REG_SPRT_X_EXP:
		fallthrough
	case REG_SPRT_SPRT_COLL:
		fallthrough
	case REG_SPRT_DATA_COLL:
		fallthrough
	case REG_BORDER_COL:
		fallthrough
	case REG_BGCOLOR_0:
		fallthrough
	case REG_BGCOLOR_1:
		fallthrough
	case REG_BGCOLOR_2:
		fallthrough
	case REG_BGCOLOR_3:
		fallthrough
	case REG_SPRT_MLTCOLOR_0:
		fallthrough
	case REG_SPRT_MLTCOLOR_1:
		fallthrough
	case REG_COLOR_SPRT_0:
		fallthrough
	case REG_COLOR_SPRT_1:
		fallthrough
	case REG_COLOR_SPRT_2:
		fallthrough
	case REG_COLOR_SPRT_3:
		fallthrough
	case REG_COLOR_SPRT_4:
		fallthrough
	case REG_COLOR_SPRT_5:
		fallthrough
	case REG_COLOR_SPRT_6:
		fallthrough
	case REG_COLOR_SPRT_7:
		C.Reg[reg] = val
	}
}

func (C *CRTC) testBit(reg uint16, mask byte) bool {
	if C.Reg[reg]&mask == mask {
		return true
	}
	return false
}

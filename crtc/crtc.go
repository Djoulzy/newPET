package crtc

import (
	"fmt"
	"newPET/config"
	"newPET/graphic"
)

const (
	// Durée d'un cycle PAL : 1.015uS
	screenWidthPAL  = 504
	screenHeightPAL = 312
	rasterWidthPAL  = 403
	rasterHeightPAL = 284
	cyclesPerLine   = 63

	rasterTime = 1                  // Nb of cycle to put 1 byte on a line
	rasterLine = rasterWidthPAL / 8 // Nb of cycle to draw a full line
	fullRaster = rasterLine * rasterHeightPAL

	winWidth      = screenWidthPAL
	winHeight     = screenHeightPAL
	visibleWidth  = 320
	visibleHeight = 200

	firstVBlankLine  = 300
	lastVBlankLine   = 15
	firstDisplayLine = 51
	lastDisplayLine  = 250
)

func (C *CRTC) Init(ram []byte, io []byte, chargen []byte, video interface{}, conf *config.ConfigData) {
	R0 = 63
	R1 = 40
	R2 = 
	R3 
	R4 
	R5 
	R6 
	R7 
	R8 
	R9 

	C.graph = video.(graphic.Driver)
	C.graph.Init(winWidth, winHeight)
	C.conf = conf

	C.videoRam = ram[screenStart:screenStart+screenSize]
	C.charRom = chargen
}

func (C *CRTC) Disassemble() string {
	var buf string
	buf = fmt.Sprintf("RstX: %03d - RstY: %03d - RC: %02d - VC: %03X - VCBase: %03X - VMLI: %02d", C.BeamX, C.BeamY, C.RC, C.VC, C.VCBASE, C.VMLI)
	return buf
}

// func (C *CRTC) saveRasterPos(val int) {
// 	C.Reg[REG_RASTER] = byte(val)
// 	mask := byte(val>>1) & RST8
// 	res := C.Reg[REG_CTRL1] & 0b01111111
// 	C.Reg[REG_CTRL1] = res | mask
// }

func (C *CRTC) readVideoMatrix() {
	C.ColorBuffer[C.VMLI] = C.color[C.VC] & 0b00001111
	C.CharBuffer[C.VMLI] = C.bankMem.Read(C.ScreenBase + C.VC)
	// fmt.Printf("VMLI: %02X - VC: %02X - Screen Code: %d - Color: %04X\n", C.VMLI, C.VC, C.CharBuffer[C.VMLI], C.ColorBuffer[C.VMLI])
}

func (C *CRTC) drawChar(X int, Y int) {
	if C.drawArea && (C.Reg[REG_CTRL1]&DEN > 0) {
		C.StandardTextMode(X, Y)
		C.VMLI++
		C.VC++
	} else if C.visibleArea {
		for column := 0; column < 8; column++ {
			C.graph.DrawPixel(X+column, Y, Colors[C.Reg[REG_BORDER_COL]&0b00001111])
		}
	}
}

func (C *CRTC) Run(debug bool) bool {
	C.SystemClock++
	C.saveRasterPos(C.BeamY)

	C.visibleArea = (C.BeamY > lastVBlankLine) && (C.BeamY < firstVBlankLine)
	// C.displayArea = (C.BeamY >= firstDisplayLine) && (C.BeamY <= lastDisplayLine) && C.visibleArea
	C.displayArea = (C.BeamY >= firstDisplayLine) && (C.BeamY <= lastDisplayLine)
	C.BeamX = (C.cycle - 1) * 8
	C.drawArea = ((C.cycle > 15) && (C.cycle < 56)) && C.displayArea

	C.BA = !(((C.BeamY-firstDisplayLine)%8 == 0) && C.displayArea && (C.cycle > 11) && (C.cycle < 55))

	// if C.drawArea {
	// 	fmt.Printf("Raster: %d - Cycle: %d - BA: %t - VMLI: %d - VCBASE/VC: %d/%d - RC: %d - Char: %02X\n", C.BeamY, C.cycle, C.BA, C.VMLI, C.VCBASE, C.VC, C.RC, C.CharBuffer[C.VMLI])
	// }

	switch C.cycle {
	case 1:
		if C.testBit(REG_IRQ_ENABLED, IRQ_RST) {
			if C.RasterIRQ == uint16(C.BeamY) {
				// fmt.Printf("\nIRQ: %04X - %04X", C.RasterIRQ, uint16(C.BeamY))
				// fmt.Println("Rastrer Interrupt")
				C.Reg[REG_IRQ] = C.Reg[REG_IRQ] | 0b10000001
				*C.IRQ_Pin = 1
			}
		}
	case 2:
	case 3:
	case 4:
	case 5:
	case 6:
	case 7:
	case 8:
	case 9:
	case 10:
	case 11: // Debut de la zone visible
		C.drawChar(C.BeamX, C.BeamY)
	case 12:
		C.drawChar(C.BeamX, C.BeamY)
	case 13:
		C.drawChar(C.BeamX, C.BeamY)
	case 14:
		C.VC = C.VCBASE
		C.VMLI = 0
		if !C.BA {
			C.RC = 0
		}
		C.drawChar(C.BeamX, C.BeamY)
	case 15: // Debut de la lecture de la memoire video en mode BadLine
		C.drawChar(C.BeamX, C.BeamY)
		if !C.BA {
			C.readVideoMatrix()
		}
	case 16: // Debut de la zone d'affichage
		C.drawChar(C.BeamX, C.BeamY)
		if !C.BA {
			C.readVideoMatrix()
		}
	case 17:
		C.drawChar(C.BeamX, C.BeamY)
		if !C.BA {
			C.readVideoMatrix()
		}
	case 18:
		C.drawChar(C.BeamX, C.BeamY)
		if !C.BA {
			C.readVideoMatrix()
		}
	case 19:
		C.drawChar(C.BeamX, C.BeamY)
		if !C.BA {
			C.readVideoMatrix()
		}
	case 20:
		C.drawChar(C.BeamX, C.BeamY)
		if !C.BA {
			C.readVideoMatrix()
		}
	case 21:
		C.drawChar(C.BeamX, C.BeamY)
		if !C.BA {
			C.readVideoMatrix()
		}
	case 22:
		C.drawChar(C.BeamX, C.BeamY)
		if !C.BA {
			C.readVideoMatrix()
		}
	case 23:
		C.drawChar(C.BeamX, C.BeamY)
		if !C.BA {
			C.readVideoMatrix()
		}
	case 24:
		C.drawChar(C.BeamX, C.BeamY)
		if !C.BA {
			C.readVideoMatrix()
		}
	case 25:
		C.drawChar(C.BeamX, C.BeamY)
		if !C.BA {
			C.readVideoMatrix()
		}
	case 26:
		C.drawChar(C.BeamX, C.BeamY)
		if !C.BA {
			C.readVideoMatrix()
		}
	case 27:
		C.drawChar(C.BeamX, C.BeamY)
		if !C.BA {
			C.readVideoMatrix()
		}
	case 28:
		C.drawChar(C.BeamX, C.BeamY)
		if !C.BA {
			C.readVideoMatrix()
		}
	case 29:
		C.drawChar(C.BeamX, C.BeamY)
		if !C.BA {
			C.readVideoMatrix()
		}
	case 30:
		C.drawChar(C.BeamX, C.BeamY)
		if !C.BA {
			C.readVideoMatrix()
		}
	case 31:
		C.drawChar(C.BeamX, C.BeamY)
		if !C.BA {
			C.readVideoMatrix()
		}
	case 32:
		C.drawChar(C.BeamX, C.BeamY)
		if !C.BA {
			C.readVideoMatrix()
		}
	case 33:
		C.drawChar(C.BeamX, C.BeamY)
		if !C.BA {
			C.readVideoMatrix()
		}
	case 34:
		C.drawChar(C.BeamX, C.BeamY)
		if !C.BA {
			C.readVideoMatrix()
		}
	case 35:
		C.drawChar(C.BeamX, C.BeamY)
		if !C.BA {
			C.readVideoMatrix()
		}
	case 36:
		C.drawChar(C.BeamX, C.BeamY)
		if !C.BA {
			C.readVideoMatrix()
		}
	case 37:
		C.drawChar(C.BeamX, C.BeamY)
		if !C.BA {
			C.readVideoMatrix()
		}
	case 38:
		C.drawChar(C.BeamX, C.BeamY)
		if !C.BA {
			C.readVideoMatrix()
		}
	case 39:
		C.drawChar(C.BeamX, C.BeamY)
		if !C.BA {
			C.readVideoMatrix()
		}
	case 40:
		C.drawChar(C.BeamX, C.BeamY)
		if !C.BA {
			C.readVideoMatrix()
		}
	case 41:
		C.drawChar(C.BeamX, C.BeamY)
		if !C.BA {
			C.readVideoMatrix()
		}
	case 42:
		C.drawChar(C.BeamX, C.BeamY)
		if !C.BA {
			C.readVideoMatrix()
		}
	case 43:
		C.drawChar(C.BeamX, C.BeamY)
		if !C.BA {
			C.readVideoMatrix()
		}
	case 44:
		C.drawChar(C.BeamX, C.BeamY)
		if !C.BA {
			C.readVideoMatrix()
		}
	case 45:
		C.drawChar(C.BeamX, C.BeamY)
		if !C.BA {
			C.readVideoMatrix()
		}
	case 46:
		C.drawChar(C.BeamX, C.BeamY)
		if !C.BA {
			C.readVideoMatrix()
		}
	case 47:
		C.drawChar(C.BeamX, C.BeamY)
		if !C.BA {
			C.readVideoMatrix()
		}
	case 48:
		C.drawChar(C.BeamX, C.BeamY)
		if !C.BA {
			C.readVideoMatrix()
		}
	case 49:
		C.drawChar(C.BeamX, C.BeamY)
		if !C.BA {
			C.readVideoMatrix()
		}
	case 50:
		C.drawChar(C.BeamX, C.BeamY)
		if !C.BA {
			C.readVideoMatrix()
		}
	case 51:
		C.drawChar(C.BeamX, C.BeamY)
		if !C.BA {
			C.readVideoMatrix()
		}
	case 52:
		C.drawChar(C.BeamX, C.BeamY)
		if !C.BA {
			C.readVideoMatrix()
		}
	case 53:
		C.drawChar(C.BeamX, C.BeamY)
		if !C.BA {
			C.readVideoMatrix()
		}
	case 54: // Dernier lecture de la matrice video ram
		C.drawChar(C.BeamX, C.BeamY)
		if !C.BA {
			C.readVideoMatrix()
		}
	case 55: // Fin de la zone de display
		C.drawChar(C.BeamX, C.BeamY)
	case 56: // Debut de la zone visible
		C.drawChar(C.BeamX, C.BeamY)
	case 57:
		C.drawChar(C.BeamX, C.BeamY)
	case 58:
		if C.RC == 7 {
			C.VCBASE = C.VC
		}
		if C.displayArea {
			C.RC++
		}
		C.drawChar(C.BeamX, C.BeamY)
	case 59:
		C.drawChar(C.BeamX, C.BeamY)
	case 60:
	case 61:
	case 62:
	case 63:
	}
	// C.BeamX += 8
	C.cycle++
	if C.cycle > cyclesPerLine {
		C.cycle = 1
		C.BeamY++
		if C.BeamY >= screenHeightPAL {
			C.BeamY = 0
			C.VCBASE = 0
			C.graph.UpdateFrame()
		}
	}
	if debug {
		C.graph.UpdateFrame()
	}
	return C.BA
}

func (C *CRTC) Dump(addr uint16) {
	fmt.Printf("Bank: %d - VideoBase: %04X - CharBase: %04X", C.BankSel, C.ScreenBase, C.CharBase)
	C.bankMem.Show()
	C.bankMem.Dump(addr)
}

func (C *CRTC) Stats() {
	banks := [4]uint16{BankStart0, BankStart1, BankStart2, BankStart3}

	fmt.Printf("VIC:\n")
	fmt.Printf("Bank: %d - VideoBase: %04X (%04X) - CharBase: %04X (%04X)\n", C.BankSel, C.ScreenBase, banks[C.BankSel]+C.ScreenBase, C.CharBase, banks[C.BankSel]+C.CharBase)
	fmt.Printf("RstX: %04X - RstY: %04X - RC: %02d - VC: %03X - VCBase: %03X - VMLI: %02d\n", C.BeamX, C.BeamY, C.RC, C.VC, C.VCBASE, C.VMLI)
	fmt.Printf("IRQ Line: ")
	if C.Reg[REG_IRQ]&0b10000000 > 0 {
		fmt.Printf("On")
	} else {
		fmt.Printf("Off")
	}
	fmt.Printf(" - IRQ Enabled: ")
	if C.Reg[REG_IRQ_ENABLED]&0b00001111 > 0 {
		fmt.Printf("Yes")
	} else {
		fmt.Printf("None")
	}
	fmt.Printf(" - Raster IRQ: %04X\n", C.RasterIRQ)
}

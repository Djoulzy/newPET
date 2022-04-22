package crtc

func (C *CRTC) StandardTextMode(X int, Y int) {
	var pixelData byte
	var palette [4]byte

	charAddr := (uint16(C.CharBuffer[C.VMLI]) << 3) + uint16(C.RC)
	pixelData = C.bankMem.Read(C.CharBase + charAddr)
	palette[1] = C.ColorBuffer[C.VMLI] // C.color.Val[C.VC] & 0b00001111
	palette[0] = C.Reg[REG_BGCOLOR_0] & 0b00001111

	for column := 0; column < 8; column++ {
		bit := byte(0b10000000 >> column)
		if pixelData&bit > 0 {
			C.graph.DrawPixel(X+column, Y, Colors[palette[1]])
		} else {
			C.graph.DrawPixel(X+column, Y, Colors[palette[0]])
		}
	}
}
package crtc

func (C *CRTC) StandardTextMode(X int, Y int) {
	screenChar := C.videoRam[C.RasterLine*C.Reg[R6]+C.CCLK]
	pixelData := C.charRom[screenChar<<3+C.RasterCount]

	for column := 0; column < 8; column++ {
		bit := byte(0b10000000 >> column)
		if pixelData&bit > 0 {
			C.graph.DrawPixel(X+column, Y, Colors[Green])
		} else {
			C.graph.DrawPixel(X+column, Y, Colors[Black])
		}
	}
}

package main

var IORead_Mapper [](func(uint16) byte)
var IOWrite_Mapper []func(uint16, byte)

type accessor struct {
}

func (C *accessor) MRead(mem []byte, translatedAddr uint16) byte {
	// clog.Test("Accessor", "MRead", "Addr: %04X", addr)
	return IORead_Mapper[translatedAddr](translatedAddr)
}

func (C *accessor) MWrite(meme []byte, translatedAddr uint16, val byte) {
	// clog.Test("Accessor", "MWrite", "Addr: %04X", translatedAddr)
	IOWrite_Mapper[translatedAddr](translatedAddr, val)
}

func ioRead(translatedAddr uint16) byte {
	// clog.Trace("NullRead", "Mapper", "addr: %04X - Mapper: %d", translatedAddr, IORead_Mapper[translatedAddr])
	return IO[translatedAddr]
}

func ioWrite(translatedAddr uint16, val byte) {
	// clog.Trace("NullWrite", "Mapper", "addr: %04X - Mapper: %d", translatedAddr, IORead_Mapper[translatedAddr])
	IO[translatedAddr] = val
}

func fillIOMapper() {
	IORead_Mapper = make([]func(uint16) byte, 4096)
	IOWrite_Mapper = make([]func(uint16, byte), 4096)

	for i := 0x0400; i < 0x0800; i++ { // SID
		IORead_Mapper[i] = ioRead
		IOWrite_Mapper[i] = ioWrite
	}
	for i := 0x0800; i < 0x0C00; i++ { // Color
		IORead_Mapper[i] = ioRead
		IOWrite_Mapper[i] = ioWrite
	}
	for i := 0x0E00; i < 0x0F00; i++ { // IO1
		IORead_Mapper[i] = ioRead
		IOWrite_Mapper[i] = ioWrite
	}
	for i := 0x0F00; i < 0x1000; i++ { // IO2
		IORead_Mapper[i] = ioRead
		IOWrite_Mapper[i] = ioWrite
	}
}

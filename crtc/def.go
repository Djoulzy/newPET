package crtc

import (
	"newPET/config"
	"newPET/graphic"
	"newPET/mem"
)

var (
	Black      byte = 0
	White      byte = 1
	Red        byte = 2
	Cyan       byte = 3
	Violet     byte = 4
	Green      byte = 5
	Blue       byte = 6
	Yellow     byte = 7
	Orange     byte = 8
	Brown      byte = 9
	Lightred   byte = 10
	Darkgrey   byte = 11
	Grey       byte = 12
	Lightgreen byte = 13
	Lightblue  byte = 14
	Lightgrey  byte = 15
)

var Colors [16]graphic.RGB = [16]graphic.RGB{
	{R: 0, G: 0, B: 0},       // Black
	{R: 255, G: 255, B: 255}, // White
	{R: 137, G: 78, B: 67},   // Red
	{R: 146, G: 195, B: 203}, // Cyan
	{R: 138, G: 87, B: 176},  // Violet
	{R: 128, G: 174, B: 89},  // Green
	{R: 68, G: 63, B: 164},   // Blue
	{R: 215, G: 221, B: 137}, // Yellow
	{R: 146, G: 106, B: 56},  // Orange
	{R: 100, G: 82, B: 23},   // Brown
	{R: 184, G: 132, B: 122}, // Lightred
	{R: 96, G: 96, B: 96},    // Darkgrey
	{R: 138, G: 138, B: 138}, // Grey
	{R: 191, G: 233, B: 155}, // Lightgreen
	{R: 131, G: 125, B: 216}, // Lightblue
	{R: 179, G: 179, B: 179}, // Lightgrey
}

// VIC :
type CRTC struct {
	VML         [40]uint16 // Video Matrix Line
	VMLI        byte       // Video Matrix Line Indexer
	VC          uint16     // Vide Counter
	VCBASE      uint16     // Video Counter Base
	RC          byte       // Row counter
	BA          bool       // High: normal / Low: BadLine
	SystemClock uint16
	Reg         [64]byte

	conf  *config.ConfigData
	BeamX int
	BeamY int
	cycle int

	visibleArea bool
	displayArea bool
	drawArea    bool

	ColorBuffer [40]byte
	CharBuffer  [40]byte

	IRQ_Pin   *int
	RasterIRQ uint16
	graph     graphic.Driver

	color      []byte
	BankSel    byte
	ScreenBase uint16
	CharBase   uint16
	bankMem    mem.BANK

	ECM  bool
	BMM  bool
	MCM  bool
	MODE byte
}

var (
	R0 byte // Longueur d'une ligne (displayed + sync)
	R1 byte // Nb de characteres par ligne
	R2 byte // Pos du sync start par apport au debut de la ligne
	R3 byte // Sync control (0-3: Horizontal, 4-7: Vertical)
	R4 byte // Nb total de lignes
	R5 byte // Nb de scanlines à ajouter pour compléter l'ecran
	R6 byte // Nb de lignes visibles affichées
	R7 byte // Pos du vertical sync

	R8 byte
	R9 byte
)

const (
	colorStart  = 0x0800 // 0xD800 translated
	screenStart = 0x8000
	screenSize  = 4096

	PALNTSC uint16 = 0x02A6

	YSCROLL byte = 0b00000111 // From REG_CTRL1
	RSEL    byte = 0b00001000 // rom REG_CTRL1 : 0 = 24 rows; 1 = 25 rows.
	DEN     byte = 0b00010000 // rom REG_CTRL1 : 0 = Screen off, 1 = Screen on.
	// BMM     byte = 0b00100000 // rom REG_CTRL1 : 0 = Text mode; 1 = Bitmap mode.
	// ECM     byte = 0b01000000 // rom REG_CTRL1 : 1 = Extended background mode on.
	// MCM     byte = 0b00010000 // rom REG_CTRL2
	RST8 byte = 0b10000000 // rom REG_CTRL1 : Read: Current raster line (bit #8). Write: Raster line to generate interrupt at (bit #8).

	IRQ_RST byte = 0b00000001 // Raster line interrupt
	IRQ_MBC byte = 0b00000010 // Sprite collision with background
	IRQ_MMC byte = 0b00000100 // Sprite vs sprite collision
	IRQ_LP  byte = 0b00001000 // Light pen negative edge
)

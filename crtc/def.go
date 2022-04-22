package crtc

import (
	"newPET/config"
	"newPET/graphic"
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
	Reg [18]byte

	conf        *config.ConfigData
	BeamX       int
	BeamY       int
	RasterLine  byte
	RasterCount byte
	CCLK        byte

	visibleArea bool

	graph graphic.Driver
	MODE  byte

	videoRam []byte
	charRom  []byte
}

const (
	R0 byte = iota // Longueur d'une ligne (displayed + sync)
	R1             // Nb de characteres par ligne
	R2             // Pos du sync start par apport au debut de la ligne
	R3             // Sync control (0-3: Horizontal, 4-7: Vertical)
	R4             // Nb total de lignes
	R5             // Nb de scanlines à ajouter pour compléter l'ecran
	R6             // Nb de lignes visibles affichées
	R7             // Pos du vertical sync
	R8
	R9
	R10
	R11
	R12
	R13
	R14
	R15
	R16
	R17
)

const (
	colorStart  = 0x0800 // 0xD800 translated
	screenStart = 0x8000
	screenSize  = 4096
)

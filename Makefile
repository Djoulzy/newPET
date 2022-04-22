
all:
	make main

main:
	go build -o newPET cmd/newPET/*

crtc: cmd/crtc/main.go crtc/crtc.go
	go build -o TestCrtc cmd/crtc/*
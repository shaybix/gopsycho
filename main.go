package main

import (
	"flag"
	"fmt"
	"log"
)

var (
	drive    = flag.String("drive", "", "Select the drive the USB is connected to")
	action   = flag.String("action", "DumpFirmware", "Default: DumpFirmware (no other is supported at the moment)")
	burner   = flag.String("burner", "", "Path to the burner file")
	firmware = flag.String("firmware", "", "Path to the firmware file")
)

type PhisonDevice struct {
	DriveLetter string
}

func main() {

	flag.Parse()
	ok := flag.Parsed()

	if !ok {
		log.Println("The flags have not been parsed")

	}

	// At this point the flags are parsed and checked that they are parsed
	//
	err := OpenDrive(*drive)
	if err != nil {
		log.Println("Could not read from drive", err)
	}

	err = DumpFirmware(*firmware)
	if err != nil {
		log.Println("Could not dump the firmware")
	}
}

func OpenDrive(drive string) error {
	defer CloseDrive()

	device := new(PhisonDevice)
	device.DriveLetter = drive

	// Open a connection to the drive and handle the connection
	// check for errors and return accordingly

	return nil
}

func CloseDrive() {
	// function responsible for closing the drive

}

func DumpFirmware(firmware string) error {
	var address int = 0
	var data = make([]byte, 0x32400)
	var header = []byte{0x42, 0x74, 0x50, 0x72, 0x61, 0x6D, 0x43, 0x64, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x14, 0x10, 0x0B, 0x18}

	fmt.Println(address)
	fmt.Println(header)
	fmt.Println(data)
	return nil
}

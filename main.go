package main

import (
	"flag"
	"fmt"
	"log"
)

// Commandline flags/arguments specified
var (
	drive    = flag.String("drive", "", "Select the drive the USB is connected to")
	action   = flag.String("action", "DumpFirmware", "Default: DumpFirmware (no other is supported at the moment)")
	burner   = flag.String("burner", "", "Path to the burner file")
	firmware = flag.String("firmware", "", "Path to the firmware file")
)

type SafeFileHandle []byte

type PhisonDevice struct {
	DriveLetter string
}

func main() {

	flag.Parse()
	// Check if the flags have been parsed
	ok := flag.Parsed()

	// check if flag.Parsed() returned ok
	if !ok {
		log.Println("The flags have not been parsed")

	}

	// Try to read from the USB drive
	err := OpenDrive(*drive)
	if err != nil {
		log.Println("Could not read from drive", err)
	}

	// Try to dump (write) firmware to the drive
	err = DumpFirmware(*firmware)
	if err != nil {
		log.Println("Could not dump the firmware")
	}
}

// OpenDrive opens a connection with the device
func OpenDrive(drive string) error {
	defer CloseDrive()

	device := new(PhisonDevice)
	device.DriveLetter = drive

	// TODO: Open a connection to the drive and handle the connection

	// TODO: Check for errors and return accordingly

	return nil
}

// CloseDrive handles the graceful closing of the connection with the device
func CloseDrive() {
	// TODO: handle the closing of the connection gracefully

}

// DumpFirmware writes the custom firmware to the device
func DumpFirmware(firmware string) error {
	var address int = 0
	var data = make([]byte, 0x32400)
	var header = []byte{0x42, 0x74, 0x50, 0x72, 0x61, 0x6D, 0x43, 0x64, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x14, 0x10, 0x0B, 0x18}

	copy(data, header)

	// TODO: loop through the data array

	// TODO: While looping through the data send the data through

	// FIXME: remove these fmt.Println once you have implemented this function
	fmt.Println(address)
	fmt.Println(header)
	fmt.Println(data)

	return nil
}

func (d *PhisonDevice) SendCommand(handle SafeFileHandle, cmd []byte, data []byte, bytesExpected int) []byte {

	var ret []byte

	// TODO: Implement the SendCommand method
	return ret
}

package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

// Commandline flags/arguments specified
var (
	drive = flag.String("drive", "", "Select the drive the USB is connected to")
	// Change of focus on the setting the USB into bootmode first before implementing
	// other actions like DumpFirmware
	//action   = flag.String("action", "DumpFirmware", "Default: DumpFirmware (no other is supported at the moment)")
	action   = flag.String("action", "SetBootMode", "Default: SetBootMode (no other is supported at the moment)")
	burner   = flag.String("burner", "", "Path to the burner file")
	firmware = flag.String("firmware", "", "Path to the firmware file")
)

// SafeFileHandle of type []byte
// might not be neccessarily as it was a type in Windows for safe handling of files
// however keep it for now for ease of understanding code
type SafeFileHandle []byte

// PhisonDevice of type struct
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
	handle, err := OpenDrive(*drive)
	if err != nil {
		log.Println("Could not read from drive", err)
	}

	// Try to dump (write) firmware to the drive
	//err = DumpFirmware(*firmware)
	//if err != nil {
	//	log.Println("Could not dump the firmware")
	//}
	fmt.Println(handle)

}

// OpenDrive opens a connection with the device
// Look into whether there is a better way of doing this
// is it necessarilly to call another Open() method within this function?
// However in order to stick to the way done in the original Psychson code keep it as it is.
func OpenDrive(drive string) (*os.File, error) {
	defer CloseDrive()

	device := new(PhisonDevice)
	device.DriveLetter = drive

	// TODO: Open a connection to the drive and handle the connection
	handle, err := device.Open()

	// TODO: Check for errors and return accordingly
	if err != nil {
		log.Println("Could not open the device please make sure you selected correct drive")
		return nil, err
	}

	return handle, nil
}

// CloseDrive handles the graceful closing of the connection with the device
func CloseDrive() {
	// TODO: handle the closing of the connection gracefully

}

// DumpFirmware writes the custom firmware to the device
//func DumpFirmware(firmware string) error {
//	var address int = 0
//	var data = make([]byte, 0x32400)
//	var header = []byte{0x42, 0x74, 0x50, 0x72, 0x61, 0x6D, 0x43, 0x64, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x14, 0x10, 0x0B, 0x18}
//
//	copy(data, header)
//
// TODO: loop through the data array

// TODO: While looping through the data send the data through

// FIXME: remove these fmt.Println once you have implemented this function
//	fmt.Println(address)
//	fmt.Println(header)
//	fmt.Println(data)
//
//	return nil
//}

// Open is responsible for opening a connection
func (d *PhisonDevice) Open() (*os.File, error) {

	// TODO: open a connection
	// However getting a read-only filesystem error which I need to look into

	handle, err := os.OpenFile(d.DriveLetter, os.O_RDONLY, os.ModeDevice)
	if err != nil {
		log.Println("Could not open file to the device: ")
		return nil, err

	}

	return handle, nil

}

// JumpToBootMode sets the device to boot mode
// for when user sets --action=SetBootMode
func (d *PhisonDevice) JumpToBootMode() error {

	// TODO: SendCommand data to device to set it to boot mode
	return nil
}

func (d *PhisonDevice) ExecuteImage() error {

	// TODO: Execute the burner image
	return nil
}

func (d *PhisonDevice) JumpToPRAM() error {

	// TODO:
	return nil
}

// SendCommand sends data across to the device
func (d *PhisonDevice) SendCommand(handle SafeFileHandle, cmd []byte, data []byte, bytesExpected int) []byte {

	var ret []byte

	// TODO: Implement the SendCommand method
	return ret
}

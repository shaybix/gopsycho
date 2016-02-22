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
		log.Println("OpenDrive()", err)
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
		log.Println("Could not open the device: ", err)
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

	handle, err := os.OpenFile(d.DriveLetter, os.O_RDWR, 0777)
	if err != nil {
		log.Println("Could not open file to the device: ", err)
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

// ExecuteImage executes the burner image unto the device
func (d *PhisonDevice) ExecuteImage() error {

	// TODO: Execute the burner image
	return nil
}

// JumpToPRAM ....
func (d *PhisonDevice) JumpToPRAM() error {

	// TODO: Don't know exactly yet what this method does :p
	return nil

}

// GetInfo will get relevant information from the device and print it
// to the console
func (d *PhisonDevice) GetInfo() error {

	// TODO: Read and print the information of the device
	fmt.Println("Gathering information....")
	fmt.Println("Reported chip type: ", d.GetChipType())
	fmt.Println("Reported chip ID: ", d.GetChipID())
	fmt.Println("Reported firmware version: ", d.GetFirmwareVersion())

	return nil

}

// GetChipType ...
func (d *PhisonDevice) GetChipType() uint {

	// TODO: Fetch the Chip type
	var ret uint
	var info = d.RequestVendorInfo()
	if info != nil {
		//TODO: some more logic
	}
	return ret
}

// GetChipID ...
func (d *PhisonDevice) GetChipID() string {

	// TODO: Fetch the Chip type

	return ""

}

// GetFirmwareVersion ...
func (d *PhisonDevice) GetFirmwareVersion() string {

	// TODO: Fetch the Chip type

	return ""

}

// SendCommand sends data across to the device
func (d *PhisonDevice) SendCommand(handle SafeFileHandle, bytesExpected int) []byte {

	var ret []byte

	// TODO: Implement the SendCommand method
	return ret
}

func (d *PhisonDevice) RequestVendorInfo() []byte {

	var data = d.SendCommand([]byte{0x06, 0x05, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01}, 512+16)

	var ret []byte
	if data != nil {

		// TODO: convert content in data into []byte assign to ret to be returned.
	}
	return ret
}

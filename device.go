package main

import (
	"fmt"
	"log"
	"syscall"
)

// PhisonDevice of type struct
type PhisonDevice struct {
	DriveLetter string
}

// Open is responsible for opening a connection
func (d *PhisonDevice) Open() (*int, error) {

	// TODO : open a connection
	// However getting a read-only filesystem error which I need to look into

	//handle, err := os.OpenFile(d.DriveLetter, os.O_RDWR, 0777)
	//if err != nil {
	//log.Println("Could not open file to the device: ", err)
	//return nil, err

	//}

	handle, err := syscall.Open("/dev/sdb", syscall.O_RDONLY, 0777)

	if err != nil {
		log.Println("syscall.Open() tried : ")
		return nil, err
	}
	return &handle, nil

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

// SendCommand sends data across to the device however the shown below is private in the c# code
// There are overloaded methods in this case, and must consider how to perhaps factor out
// the other two methods. There are:
// 	SendCommand(cmd []byte, bytesExpected int) []byte
// 	SendCommand(cmd []byte)
// 	SendCommand(cmd, data []byte)
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

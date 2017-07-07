# gopsycho = [ABDANDONED PROJECT]
Go implementation of the Psychson code with support for Linux.


#### Summary of project

This is a reimplementation of Psychson's Drivecom in Go, with support for linux
and perhaps in the future with support for OSx and Windows. The tool initially will be 
entirely from the command line but eventually will work on a GUI version. 

The main purpose of this project is educational, to understand how BadUSB truly works
and secondarily how psychson is exploiting this vulnerable, and tertiary the goal is to 
how far we could push our exploitation.

Intentionally a lot of the function names, and variable names were kept in order for it 
to be as close as possible to the c# code. This will mean that the code may divert somewhat 
from Go's idiomatic way of doing things. :)


Sample usage from commandline: 

```
gopsycho --drive=/dev/sdb --action=DumpFirmware --burner=burnerfile --firmware=customfirmware

```
--------------------------------------------------------------------------------------

###### GOALS

- [ ]   Communicate with the USB device and obtain device information
- [ ]   Read and write from and to the device
....
- [ ]   GUI implementation 
- [ ]   GUI Firmware editor
- [ ]   Implementation of all actions


---------------------------------------------------------------------------------------

###### TODO 

- [ ]   Ability to read and write to USB drive (for linux /dev/)
- [ ]   GetInfo()
- [ ]   SetBootMode()
- [ ]   SendExecutable()
- [ ]   DumpFirmware()
- [ ]   SetPassword()
- [ ]   GetNumLBAs()
- [ ]   SendFirmware()

----------------------------------------------------------------------------------------

#### Authors

* Aboo "shaybix" Shayba


#### Contributors


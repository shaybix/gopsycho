# Progress Documentation
-----------------------------------------------------------

###### Friday 26-02-2016

Have refactored some of the code, putting all methods and code relevant to the PhisonDevice 
struct in its own file calle device.go and have been reading some more file descriptors. 
I do not if I have to shorten the pins on the device again to ensure it is working again, 
as at the moment i can not get into the read and write mode. The file descriptor 
I get returned is '3'. 

The aim for the moment is to get the method GetInfo() to work and before that open the device for 
read and write operations. 



###### Saturday 27-02-2016

Objectives for the day are:

- [ ]   Open device file with O_RDWR mode. 


Have so far only been able open device in 0_RDONLY mode and using the syscall package I 
only get a file descriptor in return, and in this case it was a '3'. Might have to consider
going back to the os package.



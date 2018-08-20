# go-mimikatz
A Go wrapper around a pre-compiled version of the Mimikatz executable for the purpose of anti-virus evasion.

### Requirements:
	MemoryModule => https://github.com/fancycode/MemoryModule

This application utilizes encryption to encrypt the main mimikatz binary. 

### Build Process:

The build process is pretty much completely automated in the Makefile. If you want to know more about how the build, 
take a look at the Makefile for more details.

Otherwise, to build an encrypted executable, just run `make` and you should see an `mk.exe` file appear. 

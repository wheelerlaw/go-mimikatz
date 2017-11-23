# go-mimikatz
A Go wrapper around a pre-compiled version of the Mimikatz executable for the purpose of anti-virus evasion.

### Requirements:
	MemoryModule => https://github.com/fancycode/MemoryModule

This application utilizes 3 segmented components to provide a Go wrapper for the Mimikatz application that is not considered malicious by most anti-virus software without additional packing, and can be dynamically built utilizing a repeatable build recipie. This is done by deviding the mimikatz executible into 2 randomly generated pads that are then stored as strings within the compiled Go binary and combined, and subsiquently loaded from within the existing process memory space at run time.

### Build Process:

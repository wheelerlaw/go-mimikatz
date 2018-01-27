# go-mimikatz
A Go wrapper around a pre-compiled version of the Mimikatz executable for the purpose of anti-virus evasion.

### Requirements:
	MemoryModule => https://github.com/fancycode/MemoryModule

This application utilizes 3 segmented components to provide a Go wrapper for the Mimikatz application that is not considered malicious by most anti-virus software without additional packing, and can be dynamically built utilizing a repeatable build recipie. This is done by dividing the mimikatz executible into 2 randomly generated pads that are then stored as strings within the compiled Go binary and combined, and subsiquently loaded from within the existing process memory space at run time.

### Build Process:

1. change the encryption key on both go files

2. pinpoint the SIZE constant:

```
make
$ du -bs mimikatz.exe
2637907	mimikatz.exe
```

3. encrypt the real mimikatz binary

```
./crypt mimi32.exe
```

4. build with the correct constants

```
make
```

5. put the encrypted mimikatz at the end of the unpacker

```
cat mimi32.exe >> mimikatz.exe
```

6. deploy the mimikatz.exe 


### Donations:
Bitcoin: 3GrtoFKp7UAf2eqTeUnN8eM3V7RS3n25Ae
Ether: 0x66DB9aCAEB85A08e34c04B4F290dE840E93dd08A



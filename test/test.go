package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

const SIZE int = 1984517

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
	fmt.Printf("hidden section size: %d\n", len(bin)-SIZE)
	a := bin[SIZE:]
	fmt.Println(a)
}

func main() {
	unpack(os.Args[0])
}

package main

import (
	"encoding/binary"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

// Automatically generated from daily build
const KEY string = "YmCNqH5wZikyzUwEVGLkiKscHCegdy8fjTPdz1uEHvynpMJjRgYwD6l71VADsAKgGiWTnbcMQV7mjOd3qQYL2LxMjYie5zJb"

func crypt(stage2 []byte) []byte {
	key := []byte(KEY)
	kl := len(key)

	for i := 0; i < len(stage2); i++ {
		stage2[i] = stage2[i] ^ key[i%kl]
	}

	return stage2
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	flag.Parse()
	var data []byte
	var err error
	switch flag.NArg() {
	case 0:
		data, err = ioutil.ReadAll(os.Stdin)
		check(err)
		binary.Write(os.Stdout, binary.LittleEndian, crypt(data))
		break
	case 1:
		data, err = ioutil.ReadFile(flag.Arg(0))
		check(err)
		ioutil.WriteFile(flag.Arg(0)+".encr", crypt(data), 0644)
		break
	default:
		fmt.Printf("input must be from stdin or file\n")
		os.Exit(1)
	}
}

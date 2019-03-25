package main

import (
	"encoding/binary"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

// Automatically generated from daily build
const ENCRYPT_KEY string = "gYHs9bygdKlf47I5hn97oYoxTjJ7Ug7BaGREsY9huQDENRerKeaM93sw0MRiSknjDlr0RKTqS1VB5IZQc2lXHOC8KmxYIdL4Xs"

func crypt(stage2 []byte) []byte {
	key := []byte(ENCRYPT_KEY)
	kl := len(key)

	for i := 0; i < len(stage2); i++ {
		stage2[i] = stage2[i] ^ key[i%kl]
	}

	return stage2
}

func cryptCheck(e error) {
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
		cryptCheck(err)
		binary.Write(os.Stdout, binary.LittleEndian, crypt(data))
		break
	case 1:
		data, err = ioutil.ReadFile(flag.Arg(0))
		cryptCheck(err)
		ioutil.WriteFile(flag.Arg(0)+".encr", crypt(data), 0644)
		break
	default:
		fmt.Printf("input must be from stdin or file\n")
		os.Exit(1)
	}
}

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

const KEY string = "changethiskey"

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
			fmt.Printf("%v", string(data))
			break
		case 1:
			data, err = ioutil.ReadFile(flag.Arg(0))
			check(err)
			ioutil.WriteFile(flag.Arg(0) + "-encrypted", crypt(data), 0644)
			break
		default:
			fmt.Printf("input must be from stdin or file\n")
			os.Exit(1)
	}
}

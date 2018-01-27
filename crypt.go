package main

import (
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

func main() {
	flag.Parse()
	var data []byte
	var err error
	switch flag.NArg() {
		case 0:
			data, err = ioutil.ReadAll(os.Stdin)
			check(err)
			fmt.Printf("stdin data: %v\n", string(data))
			break
		case 1:
			data, err = ioutil.ReadFile(flag.Arg(0))
			check(err)
			fmt.Printf("file data: %v\n", string(data))
			break
		default:
			fmt.Printf("input must be from stdin or file\n")
			os.Exit(1)
	}
	// ioutil.WriteFile(fname, crypt(data), 0644)
}

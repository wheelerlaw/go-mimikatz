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
	fname := os.Args[1]
	bin, _ := ioutil.ReadFile(fname)
	ioutil.WriteFile(fname, crypt(bin), 0644)
}

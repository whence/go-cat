package main

import (
	"io"
	"log"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		io.Copy(os.Stdout, os.Stdin)
	} else {
		for _, filename := range args {
			f, err := os.Open(filename)
			if err != nil {
				log.Fatal(err)
			}
			io.Copy(os.Stdout, f)
			f.Close()
		}
	}
}

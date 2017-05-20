package main

import (
	"fmt"
	"log"
	"os"

	unqfy "github.com/0x75960/unqfy/lib"
)

var srcdir string
var dstdir string

func init() {

	if len(os.Args) != 3 {
		fmt.Printf("\nusage:\n\t%s src_dir dest_dir\n\n", os.Args[0])
		os.Exit(-1)
	}

	srcdir = os.Args[1]
	dstdir = os.Args[2]

}

func main() {

	err := unqfy.Copy(dstdir, srcdir)
	if err != nil {
		log.Fatalln(err)
	}

}

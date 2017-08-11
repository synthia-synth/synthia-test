package main

import (
	"flag"
	"fmt"

	"github.com/synthia-synth/synthia"
)

const DefaultSampleRate = 44100

var glsampleRate float64 = DefaultSampleRate

func usage() {
	fmt.Println("synthia: The synth which goes...")
	fmt.Println("Usage: synthia FILE")
}

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) < 1 {
		usage()
		return
	}
	path := args[0]
	tune, err := synthia.FileToTune(path, glsampleRate)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Tune Generated. Playing...")
	playTune(tune, glsampleRate)
}

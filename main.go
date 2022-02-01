package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

var (
	config Config
)

type Config struct {
	File      string
	OutputDir string
}

func init() {
	flag.Usage = usage
	flag.StringVar(&config.File, "f", "example.txt", "source file")
	flag.StringVar(&config.OutputDir, "o", "dist", "output directory")
	flag.Parse()
}

func main() {
	fmt.Printf("File: %s\n", config.File)
	fmt.Printf("OutputDir: %s\n", config.OutputDir)
}

func usage() {
	if _, err := fmt.Fprintln(os.Stderr, "Usage: example [flags]"); err != nil {
		log.Fatal(err)
	}
	flag.PrintDefaults()
}

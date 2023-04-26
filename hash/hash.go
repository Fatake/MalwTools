package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"flag"
	"log"
	"os"

	"github.com/fatih/color"
)

type Options struct {
	inputFile string
}

func callOptions() *Options {
	inputFile := flag.String("f", "", "[i] File to calculate HASH")
	flag.Parse()

	return &Options{inputFile: *inputFile}
}

func main() {

	menu := callOptions()

	if menu.inputFile == "" {
		log.Fatal("[!] Error: No input file specified")
	}

	EntradaData, err := os.ReadFile(menu.inputFile)
	if err != nil {
		log.Fatal("[!] Error reading input file")
	}
	md5 := md5.Sum(EntradaData)
	sha1 := sha1.Sum(EntradaData)
	sha256 := sha256.Sum256(EntradaData)

	color.Yellow("[+] md5: %x\n", md5)
	color.Yellow("[+] sha1: %x\n", sha1)
	color.Yellow("[+] sha256: %x\n", sha256)
}

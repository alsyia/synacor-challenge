package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"os"
)

func loadBinary(state *VMState, path string) {
	fileReader, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer fileReader.Close()

	memoryAddress := 0
	var buffer = make([]byte, 2)

	for {
		// Read file two bytes by two bytes
		// TODO: Add buffering
		_, err := io.ReadFull(fileReader, buffer)
		if err != nil {
			if err == io.EOF {
				break
			}
			if err == io.ErrUnexpectedEOF {
				log.Fatal("Couldn't read binary in two bytes chunks")
			}
			log.Fatal(err)
		}

		entry := binary.LittleEndian.Uint16(buffer)
		// Sanity check
		if entry > 32775 {
			log.Fatalf("Found number higher than 32775 at %v address", memoryAddress+1)
		}
		// Store instruction in memory
		state.memory[memoryAddress] = entry
		memoryAddress++
	}
}

func main() {
	state := &VMState{}
	loadBinary(state, "./challenge.bin")

	fmt.Printf("%+v", state)
}

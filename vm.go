package main

import "log"

// MemoryAddressSpaceSize is 2**(adress space width)
const MemoryAddressSpaceSize int = 32768

// VMState contains the three different storage regions of our VM
// and a pointer to the next instruction
type VMState struct {
	mem     [MemoryAddressSpaceSize]uint16
	reg0    uint16
	reg1    uint16
	reg2    uint16
	reg3    uint16
	reg4    uint16
	reg5    uint16
	reg6    uint16
	reg7    uint16
	stack   []uint16
	nextPtr uint16
}

func run(state *VMState) {
	for {
		switch nextInstruction := state.mem[state.nextPtr]; nextInstruction {
		case opHalt:
			halt(state)
		case opOut:
			a := state.nextPtr + 1
			state.nextPtr += 2
			out(state, state.mem[a])
		case opNoop:
			state.nextPtr++
			noop(state)
		default:
			log.Fatalf("Instruction not implementend yet: %v", nextInstruction)
		}
	}
}

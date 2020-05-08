package main

import (
	"log"
)

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

func (state *VMState) resolveIfRegister(a uint16) uint16 {
	if a >= 32768 && a <= 32775 {
		switch a {
		case 32768:
			return state.reg0
		case 32769:
			return state.reg1
		case 32770:
			return state.reg2
		case 32771:
			return state.reg3
		case 32772:
			return state.reg4
		case 32773:
			return state.reg5
		case 32774:
			return state.reg6
		case 32775:
			return state.reg7
		}
	}
	return a
}

func (state *VMState) run() {
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

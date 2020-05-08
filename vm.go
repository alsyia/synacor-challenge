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
	reg0    uint16 //TODO: Maybe use an array?
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

//TODO: Rename to value
func (state *VMState) resolveIfRegister(val uint16) uint16 {
	if val >= 32768 && val <= 32775 {
		switch val {
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
	return val
}

func (state *VMState) writeToRegister(regAddress, value uint16) {
	switch regAddress {
	case 32768:
		state.reg0 = value
	case 32769:
		state.reg1 = value
	case 32770:
		state.reg2 = value
	case 32771:
		state.reg3 = value
	case 32772:
		state.reg4 = value
	case 32773:
		state.reg5 = value
	case 32774:
		state.reg6 = value
	case 32775:
		state.reg7 = value
	default:
		log.Fatalf("Not a valid register address: %v", regAddress)
	}
}

func (state *VMState) pushStack(value uint16) {
	state.stack = append(state.stack, value)
}

func (state *VMState) popStack() uint16 {
	if len(state.stack) == 0 {
		log.Fatal("Can't pop from an empty stack")
	}
	val := state.stack[len(state.stack)-1]
	state.stack = state.stack[:len(state.stack)-1]
	return val
}

func (state *VMState) run() {
	for {
		switch nextInstruction := state.mem[state.nextPtr]; nextInstruction {
		case opHalt:
			halt(state)
		case opSet:
			a := state.mem[state.nextPtr+1]
			b := state.mem[state.nextPtr+2]
			state.nextPtr += 3
			set(state, a, b)
		case opPush:
			a := state.mem[state.nextPtr+1]
			state.nextPtr += 2
			push(state, a)
		case opPop:
			a := state.mem[state.nextPtr+1]
			state.nextPtr += 2
			pop(state, a)
		case opEq:
			a := state.mem[state.nextPtr+1]
			b := state.mem[state.nextPtr+2]
			c := state.mem[state.nextPtr+3]
			state.nextPtr += 4
			eq(state, a, b, c)
		case opGt:
			a := state.mem[state.nextPtr+1]
			b := state.mem[state.nextPtr+2]
			c := state.mem[state.nextPtr+3]
			state.nextPtr += 4
			gt(state, a, b, c)
		case opJmp:
			a := state.mem[state.nextPtr+1]
			state.nextPtr += 2 // For consistency only
			jmp(state, a)
		case opJt:
			a := state.mem[state.nextPtr+1]
			b := state.mem[state.nextPtr+2]
			state.nextPtr += 3
			jt(state, a, b)
		case opJf:
			a := state.mem[state.nextPtr+1]
			b := state.mem[state.nextPtr+2]
			state.nextPtr += 3
			jf(state, a, b)
		case opAdd:
			a := state.mem[state.nextPtr+1]
			b := state.mem[state.nextPtr+2]
			c := state.mem[state.nextPtr+3]
			state.nextPtr += 4
			add(state, a, b, c)
		case opOut:
			a := state.mem[state.nextPtr+1]
			state.nextPtr += 2
			out(state, a)
		case opNoop:
			state.nextPtr++
			noop(state)
		default:
			log.Fatalf("Instruction not implemented yet: %v", nextInstruction)
		}
	}
}

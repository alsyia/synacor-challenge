package main

import (
	"fmt"
	"os"
)

const (
	opHalt = iota
	opSet  = iota
	opPush = iota
	opPop  = iota
	opEq   = iota
	opGt   = iota
	opJmp  = iota
	opJt   = iota
	opJf   = iota
	opAdd  = iota
	opMult = iota
	opMod  = iota
	opAnd  = iota
	opOr   = iota
	opNot  = iota
	opRmem = iota
	opWmem = iota
	opCall = iota
	opRet  = iota
	opOut  = iota
	opIn   = iota
	opNoop = iota
)

func halt(state *VMState) {
	os.Exit(0)
}

func set(state *VMState, a, b uint16) {
	b = state.resolveIfRegister(b)
	state.writeToRegister(a, b)
}

func push(state *VMState, a uint16) {
	a = state.resolveIfRegister(a)
	state.pushStack(a)
}

func pop(state *VMState, a uint16) {
	val := state.popStack()
	state.writeToRegister(a, val)
}

func eq(state *VMState, a, b, c uint16) {
	b = state.resolveIfRegister(b)
	c = state.resolveIfRegister(c)
	if b == c {
		state.writeToRegister(a, 1)
	} else {
		state.writeToRegister(a, 0)
	}
}

func gt(state *VMState, a, b, c uint16) {
	b = state.resolveIfRegister(b)
	c = state.resolveIfRegister(c)
	if b > c {
		state.writeToRegister(a, 1)
	} else {
		state.writeToRegister(a, 0)
	}
}

func jmp(state *VMState, a uint16) {
	a = state.resolveIfRegister(a)
	state.nextPtr = a
}

func jt(state *VMState, a, b uint16) {
	a = state.resolveIfRegister(a)
	if a != 0 {
		b = state.resolveIfRegister(b)
		state.nextPtr = b
	}
}

func jf(state *VMState, a, b uint16) {
	a = state.resolveIfRegister(a)
	if a == 0 {
		b = state.resolveIfRegister(b)
		state.nextPtr = b
	}
}

func add(state *VMState, a, b, c uint16) {
	b = state.resolveIfRegister(b)
	c = state.resolveIfRegister(c)
	state.writeToRegister(a, (b+c)%32768)
}

func mult(state *VMState, a, b, c uint16) {
	b = state.resolveIfRegister(b)
	c = state.resolveIfRegister(c)
	state.writeToRegister(a, (b*c)%32768)
}

func mod(state *VMState, a, b, c uint16) {
	b = state.resolveIfRegister(b)
	c = state.resolveIfRegister(c)
	state.writeToRegister(a, (b%c))
}

func and(state *VMState, a, b, c uint16) {
	b = state.resolveIfRegister(b)
	c = state.resolveIfRegister(c)
	state.writeToRegister(a, b&c)
}

func or(state *VMState, a, b, c uint16) {
	b = state.resolveIfRegister(b)
	c = state.resolveIfRegister(c)
	state.writeToRegister(a, b|c)
}

func not(state *VMState, a, b uint16) {
	b = state.resolveIfRegister(b)
	// Can't do ^b because we want a 15-bits NOT so we clear the first bit of a 16 bits XOR mask
	state.writeToRegister(a, 0b0111111111111111^b)
}

func rmem(state *VMState, a, b uint16) {
	b = state.resolveIfRegister(b)
	state.writeToRegister(a, state.mem[b])
}

func wmem(state *VMState, a, b uint16) {
	a = state.resolveIfRegister(a)
	b = state.resolveIfRegister(b)
	state.mem[a] = b
}

func out(state *VMState, a uint16) {
	a = state.resolveIfRegister(a)
	fmt.Printf("%c", a)
}

func call(state *VMState, a uint16) {
	a = state.resolveIfRegister(a)
	state.pushStack(state.nextPtr)
	state.nextPtr = a
}

func ret(state *VMState) {
	if len(state.stack) == 0 {
		os.Exit(0)
	}
	val := state.popStack()
	state.nextPtr = val
}

func noop(state *VMState) {
	// No op
}

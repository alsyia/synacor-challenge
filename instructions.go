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

func jmp(state *VMState, a uint16) {
	a = state.resolveIfRegister(a)
	state.nextPtr = a
}

func out(state *VMState, a uint16) {
	a = state.resolveIfRegister(a)
	fmt.Printf("%c", a)
}

func noop(state *VMState) {
	// No op
}

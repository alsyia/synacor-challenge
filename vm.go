package main

// MemoryAddressSpaceSize is 2**(adress space width)
const MemoryAddressSpaceSize int = 32768

// VMState contains the three different storage regions of our VM
type VMState struct {
	memory [MemoryAddressSpaceSize]uint16
	reg0   uint16
	reg1   uint16
	reg2   uint16
	reg3   uint16
	reg4   uint16
	reg5   uint16
	reg6   uint16
	reg7   uint16
	stack  []uint16
}

func run(state *VMState) {

}

# Analyzing challenge.bin

A dump of `challenge.bin` in opcode format can be obtained with `od -t u2 --endian little challenge.bin`.

```
0000000    21    21    19    87    19   101    19   108
0000020    19    99    19   111    19   109    19   101
[...]
```

If we rewrite the two first lines of the program using the opcode list, we get this:
```
noop
noop
out 87
out 101 
out 108
out 99
out 111
out 109
out 101
```

Using an ASCII table, we can translate those numbers and see that the first things the program will do is print "Welcome". 

Now that we get a grasp of how our VM should read the binary file and what the first instructions are, we can begin the implementation.

# Implementing the memory

Before implementing all the opcodes, we need to setup our architecture properly. This meands developping the memory, the registers, and the stack. 

## Memory

Since we are never going to loop over the memory and always going to access a specific integer index, I decided to implement the memory as a pre-allocated array. An hashmap could have been used but there would have been some overhead. 

A 15-bit address space means 32 768 entries in our memory. Each entry is a 16-bit unsigned integer, so the type of our memory array is `[32768]uint16``.

## Registers

Each register is an `uint16` variable.

## Stack

The stack is unbounded, so we can't use an array here. It will be a slice of 16-bits values `[]uint16` with some helper methods.

# Execution

The execution loop (`(*VMState) run()`) is straightforward: there is a next instruction pointer, initialized at zero, and each instruction is called sequentially (except if it modifies the pointer value).

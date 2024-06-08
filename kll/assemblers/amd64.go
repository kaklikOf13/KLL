package kll

import (
	"encoding/hex"
)

type AMD64_REG = byte

const (
	REG_AL AMD64_REG = iota // RAX
	REG_BL                  // RBX
	REG_CL                  // RCX
	REG_DL                  // RDX

	REG_DIL //RDI
	REG_SIL //RSI

	REG_8B  // R8
	REG_9B  // R9
	REG_10B // R10
	REG_11B // R11
	REG_12B // R12
	REG_13B // R13
	REG_14B // R14
	REG_15B // R15

	REG_AX // RAX
	REG_BX // RBX
	REG_CX // RCX
	REG_DX // RDX

	REG_DI //RDI
	REG_SI //RSI

	REG_8W  // R8
	REG_9W  // R9
	REG_10W // R10
	REG_11W // R11
	REG_12W // R12
	REG_13W // R13
	REG_14W // R14
	REG_15W // R15

	REG_EAX //RAX
	REG_EBX //RBX
	REG_ECX //RCX
	REG_EDX //RDX

	REG_EDI //RDI
	REG_ESI //RSI

	REG_8D  // R8
	REG_9D  // R9
	REG_10D // R10
	REG_11D // R11
	REG_12D // R12
	REG_13D // R13
	REG_14D // R14
	REG_15D // R15

	REG_RAX
	REG_RBX
	REG_RCX
	REG_RDX

	REG_RDI
	REG_RSI
	REG_RIP

	REG_8
	REG_9
	REG_10
	REG_11
	REG_12
	REG_13
	REG_14
	REG_15

	LAST_INT8  = REG_15B
	LAST_INT16 = REG_15W
	LAST_INT32 = REG_15D
	LAST_INT64 = REG_15
)

type X86_64_ASSEMBLER struct {
	Code []byte
}

func simpleREG(reg AMD64_REG) AMD64_REG {
	switch reg {
	case REG_AL:
		return REG_AL
	case REG_BL:
		return REG_BL
	case REG_CL:
		return REG_CL
	case REG_DL:
		return REG_DL

	case REG_8B:
		return REG_8B
	case REG_9B:
		return REG_9B
	case REG_10B:
		return REG_10B
	case REG_11B:
		return REG_11B
	case REG_12B:
		return REG_12B
	case REG_13B:
		return REG_13B
	case REG_14B:
		return REG_14B
	case REG_15B:
		return REG_15B

	case REG_EAX:
		return REG_AL
	case REG_EBX:
		return REG_BL
	case REG_ECX:
		return REG_CL
	case REG_EDX:
		return REG_DL

	case REG_RAX:
		return REG_AL
	case REG_RBX:
		return REG_BL
	case REG_RCX:
		return REG_CL
	case REG_RDX:
		return REG_DL
	case REG_RDI:
		return REG_RDI
	case REG_RSI:
		return REG_RSI

	case REG_8:
		return REG_8
	case REG_9:
		return REG_9
	case REG_10:
		return REG_10
	case REG_11:
		return REG_11
	case REG_12:
		return REG_12
	case REG_13:
		return REG_13
	case REG_14:
		return REG_14
	case REG_15:
		return REG_15
	}
	return 0
}

func (asm *X86_64_ASSEMBLER) MOV_VALUE(reg AMD64_REG, val []byte) {
	switch reg {
	case REG_AL:
		asm.Code = append(asm.Code, 0xB0)
	case REG_BL:
		asm.Code = append(asm.Code, 0xB3)
	case REG_CL:
		asm.Code = append(asm.Code, 0xB1)
	case REG_DL:
		asm.Code = append(asm.Code, 0xB2)

	case REG_8B:
		asm.Code = append(asm.Code, 0x41, 0xb0)
	case REG_9B:
		asm.Code = append(asm.Code, 0x41, 0xb1)
	case REG_10B:
		asm.Code = append(asm.Code, 0x41, 0xb2)
	case REG_11B:
		asm.Code = append(asm.Code, 0x41, 0xb3)
	case REG_12B:
		asm.Code = append(asm.Code, 0x41, 0xb4)
	case REG_13B:
		asm.Code = append(asm.Code, 0x41, 0xb5)
	case REG_14B:
		asm.Code = append(asm.Code, 0x41, 0xb6)
	case REG_15B:
		asm.Code = append(asm.Code, 0x41, 0xb7)

	case REG_EAX:
		asm.Code = append(asm.Code, 0xb8)
	case REG_EBX:
		asm.Code = append(asm.Code, 0xbb)
	case REG_ECX:
		asm.Code = append(asm.Code, 0xb9)
	case REG_EDX:
		asm.Code = append(asm.Code, 0xba)

	case REG_RAX:
		asm.Code = append(asm.Code, 0x48, 0xb8)
	case REG_RBX:
		asm.Code = append(asm.Code, 0x48, 0xbb)
	case REG_RCX:
		asm.Code = append(asm.Code, 0x48, 0xb9)
	case REG_RDX:
		asm.Code = append(asm.Code, 0x48, 0xba)

	case REG_RDI:
		asm.Code = append(asm.Code, 0xbf)
	case REG_RSI:
		asm.Code = append(asm.Code, 0xbe)
	}
	asm.Code = append(asm.Code, val...)
}

func (asm *X86_64_ASSEMBLER) SYSCALL() {
	asm.Code = append(asm.Code, 0x0f, 0x05)
}
func (asm *X86_64_ASSEMBLER) INSERT(val []byte) {
	asm.Code = append(asm.Code, val...)
}

func getRegCCode(to AMD64_REG, reg AMD64_REG) byte {
	switch to {
	case REG_AL:
		switch reg {
		default:
			return 0xd8
		case REG_CL:
			return 0xc8
		case REG_DL:
			return 0xd0
		case REG_RDI:
			return 0xf8
		}
	case REG_BL:
		switch reg {
		default:
			return 0xc3
		}
	}
	return 0xd8
}

func (asm *X86_64_ASSEMBLER) MOV_REG(to AMD64_REG, reg AMD64_REG) {
	if reg <= LAST_INT8 && to <= LAST_INT8 {
		// If both source and destination registers are 8-bit, use a different opcode
		asm.Code = append(asm.Code, 0x88)
	} else {
		// Otherwise, use the default 64-bit mov opcode
		asm.Code = append(asm.Code, 0x48, 0x89)
	}
	reg = simpleREG(reg)
	to = simpleREG(to)

	asm.Code = append(asm.Code, getRegCCode(to, reg))
}
func (asm *X86_64_ASSEMBLER) RET() {
	asm.Code = append(asm.Code, 0xC3)

}
func (asm *X86_64_ASSEMBLER) BREAK() {
	asm.Code = append(asm.Code, 0xCD, 0x03)
}
func (asm *X86_64_ASSEMBLER) String() string {
	ret := ""
	for i, v := range asm.Code {
		if i > 0 {
			ret += " "
		}
		ret += hex.EncodeToString([]byte{v})
	}
	return ret
}

func NewX86_64() *X86_64_ASSEMBLER {
	return &X86_64_ASSEMBLER{}
}

package kll_test

import (
	"fmt"
	"testing"

	"github.com/kaklikOf13/KLL/kll"
	kll_arch "github.com/kaklikOf13/KLL/kll/assemblers"
)

func TestX86_64Move(t *testing.T) {
	asm := kll_arch.NewX86_64()
	asm.MOV_VALUE(kll_arch.REG_AL, []byte{15}, 1)
	asm.RET()
	f, err := kll.GetJIT[func() uint8](asm.Stream.Value)
	if err != nil {
		panic(err)
	}
	ret := f()
	fmt.Println(ret)
	if ret != 15 {
		t.Fail()
	}
}

func TestX86_64MoveREG(t *testing.T) {
	asm := kll_arch.NewX86_64()
	asm.MOV_VALUE(kll_arch.REG_RBX, []byte{10, 0, 0, 0, 0, 0, 0, 0}, 8)
	asm.MOV_REG(kll_arch.REG_RAX, kll_arch.REG_RBX)
	asm.RET()
	fmt.Println(asm)
	f, err := kll.GetJIT[func() int64](asm.Stream.Value)
	if err != nil {
		panic(err)
	}
	ret := f()
	fmt.Println(ret)
	if ret != 10 {
		t.Fail()
	}
}

package kll_test

import (
	"fmt"
	"testing"

	"github.com/kaklikOf13/KLL/kll"
	kll_arch "github.com/kaklikOf13/KLL/kll/assemblers"
	kll_compiler "github.com/kaklikOf13/KLL/kll/compilers"
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

func TestX86_64Compiler1(t *testing.T) {
	comp := kll_compiler.NewX86_64_Compiler(kll.NewInterpreter(kll.NewScope()))
	nodes, errk := kll.Parse("10")
	if errk.Is(kll.Success) {
		comp.Inter.Panic(errk)
	}
	comp.Compile_Node(nodes[0], kll_arch.REG_AL, kll.Int8, []kll.Callstack{}, comp.Inter.Scope, false)
	comp.Assembler.RET()
	fmt.Println("code:", comp.Assembler)

	f, err := kll.GetJIT[func() int8](comp.Assembler.Stream.Value)
	if err != nil {
		panic(err)
	}
	ret := f()
	fmt.Println(ret)
	if ret != 10 {
		t.Fail()
	}
}
func TestX86_64Compiler2(t *testing.T) {
	comp := kll_compiler.NewX86_64_Compiler(kll.NewInterpreter(kll.NewScope()))
	nodes, errk := kll.Parse("10+5")
	if errk.Is(kll.Success) {
		comp.Inter.Panic(errk)
	}
	comp.Compile_Node(nodes[0], kll_arch.REG_AL, kll.Int8, []kll.Callstack{}, comp.Inter.Scope, false)
	comp.Assembler.RET()
	fmt.Println("code:", comp.Assembler)

	f, err := kll.GetJIT[func() int8](comp.Assembler.Stream.Value)
	if err != nil {
		panic(err)
	}
	ret := f()
	fmt.Println(ret)
	if ret != 15 {
		t.Fail()
	}
}
func TestX86_64Compiler3(t *testing.T) {
	comp := kll_compiler.NewX86_64_Compiler(kll.NewInterpreter(kll.NewScope()))
	nodes, errk := kll.Parse("10+5+5+3")
	if errk.Is(kll.Success) {
		comp.Inter.Panic(errk)
	}
	comp.Compile_Node(nodes[0], kll_arch.REG_AL, kll.Int8, []kll.Callstack{}, comp.Inter.Scope, false)
	comp.Assembler.RET()
	fmt.Println("code:", comp.Assembler)

	f, err := kll.GetJIT[func() int8](comp.Assembler.Stream.Value)
	if err != nil {
		panic(err)
	}
	ret := f()
	fmt.Println(ret)
	if ret != 23 {
		t.Fail()
	}
}

func TestX86_64Compiler4(t *testing.T) {
	comp := kll_compiler.NewX86_64_Compiler(kll.NewInterpreter(kll.NewScope()))
	nodes, errk := kll.Parse("10+5+5+3")
	if errk.Is(kll.Success) {
		comp.Inter.Panic(errk)
	}
	comp.Compile_Node(nodes[0], kll_arch.REG_AL, kll.Int8, []kll.Callstack{}, comp.Inter.Scope, false)
	comp.Assembler.RET()
	fmt.Println("code:", comp.Assembler)

	f, err := kll.GetJIT[func() int8](comp.Assembler.Stream.Value)
	if err != nil {
		panic(err)
	}
	ret := f()
	fmt.Println(ret)
	if ret != 23 {
		t.Fail()
	}
}
func TestX86_64Compiler5(t *testing.T) {
	comp := kll_compiler.NewX86_64_Compiler(kll.NewInterpreter(kll.NewScope()))
	nodes, errk := kll.Parse("10+5+5-3")
	if errk.Is(kll.Success) {
		comp.Inter.Panic(errk)
	}
	comp.Compile_Node(nodes[0], kll_arch.REG_AL, kll.Int8, []kll.Callstack{}, comp.Inter.Scope, false)
	comp.Assembler.RET()
	fmt.Println("code:", comp.Assembler)

	f, err := kll.GetJIT[func() int8](comp.Assembler.Stream.Value)
	if err != nil {
		panic(err)
	}
	ret := f()
	fmt.Println(ret)
	if ret != 17 {
		t.Fail()
	}
}

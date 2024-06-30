package kll_compiler

import (
	"github.com/kaklikOf13/KLL/kll"
	kll_arch "github.com/kaklikOf13/KLL/kll/assemblers"
)

type X86_64_Compiler struct {
	Assembler kll_arch.X86_64_ASSEMBLER
	Inter     kll.Interpreter
	Regs      map[kll_arch.AMD64_REG]kll.Value
}

func (c *X86_64_Compiler) MovIntToReg(reg kll_arch.AMD64_REG, val kll.Value) {
	c.Regs[reg] = val
	st := kll.NewStream(0)
	s := val.Bytes(st)
	c.Assembler.MOV_VALUE(reg, st.Read(s), 0)
}

func (c *X86_64_Compiler) Compile_Node(node kll.Node, reg kll_arch.AMD64_REG, expected_type kll.Type, callstack []kll.Callstack, scope *kll.Scope) kll.Error {
	_, err := c.Inter.Exec_Node(node, nil, callstack, scope)
	if err.Is(kll.Success) {
		return err
	}
	switch node.Tp {
	case kll.NT_INT:
		v, err := c.Inter.Exec_Node(node, expected_type, callstack, scope)
		if !err.Is(kll.Success) {
			return err
		}
		c.MovIntToReg(reg, v)
	case kll.NT_SUM:
		err := c.Compile_Node(node.NodeValue[0].(kll.Node), kll_arch.REG_AL, expected_type, callstack, scope)
		if !err.Is(kll.Success) {
			return err
		}
		err = c.Compile_Node(node.NodeValue[1].(kll.Node), kll_arch.REG_BL, expected_type, callstack, scope)
		if !err.Is(kll.Success) {
			return err
		}
		c.Assembler.ADD([]byte{}, 1)
		if !err.Is(kll.Success) {
			return err
		}
	}
	return kll.Success
}

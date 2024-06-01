package kll

import (
	"fmt"
	"os"
)

type Interpreter struct {
	Scope *Scope
}

func (inter *Interpreter) Exec_Node(node Node, expected_type Type, callstack []Callstack, scope *Scope) (Value, Error) {
	switch node.Tp {
	case NT_INT:
		if expected_type == nil {
			return Int.Instantiate(node.NodeValue[0].String())
		}
		return expected_type.Instantiate(node.NodeValue[0].String())
	case NT_FLOAT:
		if expected_type == nil {
			return Float.Instantiate(node.NodeValue[0].String())
		}
		return expected_type.Instantiate(node.NodeValue[0].String())
	case NT_STRING:
		if expected_type == nil {
			return String.Instantiate(node.NodeValue[0].String())
		}
		return expected_type.Instantiate(node.NodeValue[0].String())
	case NT_NAME:
		return scope.Get(node.NodeValue[0].String())
	case NT_SUBOP:
		var val Value
		var err Error
		val, err = inter.Exec_Node(node.NodeValue[0].(Node), expected_type, callstack, scope)
		if !err.Is(Success) {
			return nil, err
		}
		return val, Success
	case NT_SUM:
		var val1, val2 Value
		var err Error
		val1, err = inter.Exec_Node(node.NodeValue[0].(Node), expected_type, callstack, scope)
		if !err.Is(Success) {
			return nil, err
		}
		val2, err = inter.Exec_Node(node.NodeValue[1].(Node), expected_type, callstack, scope)
		if !err.Is(Success) {
			return nil, err
		}
		return val1.Sum(val2)
	case NT_SUB:
		var val1, val2 Value
		var err Error
		val1, err = inter.Exec_Node(node.NodeValue[0].(Node), expected_type, callstack, scope)
		if !err.Is(Success) {
			return nil, err
		}
		val2, err = inter.Exec_Node(node.NodeValue[1].(Node), expected_type, callstack, scope)
		if !err.Is(Success) {
			return nil, err
		}
		return val1.Sub(val2)
	case NT_MUL:
		var val1, val2 Value
		var err Error
		val1, err = inter.Exec_Node(node.NodeValue[0].(Node), expected_type, callstack, scope)
		if !err.Is(Success) {
			return nil, err
		}
		val2, err = inter.Exec_Node(node.NodeValue[1].(Node), expected_type, callstack, scope)
		if !err.Is(Success) {
			return nil, err
		}
		return val1.Mult(val2)
	case NT_DIV:
		var val1, val2 Value
		var err Error
		val1, err = inter.Exec_Node(node.NodeValue[0].(Node), expected_type, callstack, scope)
		if !err.Is(Success) {
			return nil, err
		}
		val2, err = inter.Exec_Node(node.NodeValue[1].(Node), expected_type, callstack, scope)
		if !err.Is(Success) {
			return nil, err
		}
		return val1.Div(val2)
	}
	return nil, Success
}
func (inter *Interpreter) Log(txt ...any) {
	fmt.Println(txt...)
}
func (inter *Interpreter) Panic(err Error) {
	if !err.Is(Success) {
		inter.Log(err.String())
		os.Exit(1)
	}
}
func (inter *Interpreter) Eval(code string, all_lines bool) ([]Value, Error) {
	nodes, err := Parse(code)
	if !err.Is(Success) {
		fmt.Println(err)
	}
	ret := []Value{}
	for i := range nodes {
		if !all_lines && i == len(nodes)-1 {
			scop := inter.Scope.AddChild("eval")
			val, err := inter.Exec_Node(nodes[i], nil, []Callstack{}, scop)
			inter.Scope.RemoveChild(scop.Name)
			if !err.Is(Success) {
				return nil, err
			}
			return []Value{val}, Success
		} else {
			var v Value
			v, err = inter.Exec_Node(nodes[i], nil, []Callstack{}, inter.Scope)
			if !err.Is(Success) {
				return nil, err
			}
			if all_lines {
				ret = append(ret, v)
			}
		}
	}
	return ret, Success
}
func (inter *Interpreter) Exec(code string) Error {
	_, err := inter.Eval(code, false)
	return err
}

func NewInterpreter(scope *Scope) *Interpreter {
	return &Interpreter{Scope: scope}
}

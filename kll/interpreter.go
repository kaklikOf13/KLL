package kll

import "fmt"

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
func (inter *Interpreter) Eval(code string) Value {
	nodes, err := Parse(code)
	if !err.Is(Success) {
		fmt.Println(err)
	}
	for i := range nodes {
		if i == len(nodes)-1 {
			scop := inter.Scope.AddChild("eval")
			val, err := inter.Exec_Node(nodes[i], nil, []Callstack{}, scop)
			inter.Scope.RemoveChild(scop.Name)
			if !err.Is(Success) {
				fmt.Println(err)
				return nil
			}
			return val
		} else {
			inter.Exec_Node(nodes[i], nil, []Callstack{}, inter.Scope)
		}
	}
	return Null
}

func NewInterpreter(scope *Scope) *Interpreter {
	return &Interpreter{Scope: scope}
}

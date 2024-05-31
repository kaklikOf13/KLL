package kll

import (
	"fmt"
	"math/rand"
)

type Scope struct {
	Variables map[string]Value
	Types     map[string]Type
	Constants map[string]Value
	Parent    *Scope
	Childs    map[string]*Scope
	Name      string
}

func (scope *Scope) Get(name string) (Value, Error) {
	if _, ok := scope.Variables[name]; ok {
		return scope.Parent.Variables[name], Success
	}
	if _, ok := scope.Constants[name]; ok {
		return scope.Parent.Constants[name], Success
	}
	if _, ok := scope.Types[name]; ok {
		return scope.Parent.Types[name], Success
	}
	if scope.Parent != nil {
		return scope.Parent.Get(name)
	} else {
		return nil, ScopeError.Create(ErrorArgs{Callstack: []Callstack{}, Args: []any{name}}).(Error)
	}
}
func (scope *Scope) AddChild(name string) *Scope {
	name = name + fmt.Sprint(rand.Int31())
	_, ok := scope.Childs[name]
	for ok {
		name = name + fmt.Sprint(rand.Int31())
		_, ok = scope.Childs[name]
	}
	scope.Childs[name] = NewScope()
	scope.Name = name
	scope.Childs[name].Parent = scope
	return scope.Childs[name]
}
func (scope *Scope) RemoveChild(name string) {
	delete(scope.Childs, name)
}
func (scope *Scope) String() string {
	if scope.Parent == nil {
		return scope.Name
	} else {
		return scope.Parent.String() + "." + scope.Name
	}
}

func NewScope() *Scope {
	return &Scope{Variables: make(map[string]Value), Types: make(map[string]Type), Constants: make(map[string]Value), Childs: make(map[string]*Scope), Name: "kll"}
}

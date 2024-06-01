package kll_test

import (
	"fmt"
	"testing"

	"github.com/kaklikOf13/KLL/kll"
)

func TestInterpreterSum(t *testing.T) {
	inter := kll.NewInterpreter(kll.NewScope())
	val, err := inter.Eval("11+13", false)
	inter.Panic(err)
	fmt.Println(val)
	if !val[0].Is(kll.Int.Create(24)) {
		t.Fail()
	}
}

func TestInterpreterSub(t *testing.T) {
	inter := kll.NewInterpreter(kll.NewScope())
	val, err := inter.Eval("15-13", false)
	inter.Panic(err)
	fmt.Println(val)
	if !val[0].Is(kll.Int.Create(2)) {
		t.Fail()
	}
}
func TestInterpreterMult(t *testing.T) {
	inter := kll.NewInterpreter(kll.NewScope())
	val, err := inter.Eval("15*2", false)
	inter.Panic(err)
	fmt.Println(val)
	if !val[0].Is(kll.Int.Create(30)) {
		t.Fail()
	}
}
func TestInterpreterDiv(t *testing.T) {
	inter := kll.NewInterpreter(kll.NewScope())
	val, err := inter.Eval("50/2", false)
	inter.Panic(err)
	fmt.Println(val)
	if !val[0].Is(kll.Int.Create(25)) {
		t.Fail()
	}
}

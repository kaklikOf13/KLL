package kll_test

import (
	"fmt"
	"testing"

	"github.com/kaklikOf13/KLL/kll"
)

func TestInterpreterSum(t *testing.T) {
	inter := kll.NewInterpreter(kll.NewScope())
	fmt.Println(inter.Eval("11+13"))
}

func TestInterpreterSub(t *testing.T) {
	inter := kll.NewInterpreter(kll.NewScope())
	fmt.Println(inter.Eval("15-13"))
}
func TestInterpreterMult(t *testing.T) {
	inter := kll.NewInterpreter(kll.NewScope())
	fmt.Println(inter.Eval("15*2"))
}
func TestInterpreterDiv(t *testing.T) {
	inter := kll.NewInterpreter(kll.NewScope())
	fmt.Println(inter.Eval("50/2"))
}

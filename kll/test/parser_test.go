package kll_test

import (
	"fmt"
	"testing"

	"github.com/kaklikOf13/KLL/kll"
)

func TestParserMath(t *testing.T) {
	code, err := kll.Parse("1+13\n20-3\n2*2\n3/7\n3+2+7")
	if !err.Is(kll.Success) {
		fmt.Println(err)
		t.Fail()
	}
	fmt.Println(code)
	if fmt.Sprint(code) != "[<SUM:[<INT:[1]> <INT:[13]>]> <SUB:[<INT:[20]> <INT:[3]>]> <MUL:[<INT:[2]> <INT:[2]>]> <DIV:[<INT:[3]> <INT:[7]>]> <SUM:[<SUM:[<INT:[3]> <INT:[2]>]> <INT:[7]>]>]" {
		t.Fail()
	}
}

func TestParserGet(t *testing.T) {
	code, err := kll.Parse("console.log")
	if !err.Is(kll.Success) {
		fmt.Println(err)
		t.Fail()
	}
	fmt.Println(code)
	if fmt.Sprint(code) != "[<GET:[<NAME:[console]> <NAME:[log]>]>]" {
		t.Fail()
	}
}

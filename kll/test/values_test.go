package kll_test

import (
	"fmt"
	"testing"

	"github.com/kaklikOf13/KLL/kll"
)

func TestValueSum(t *testing.T) {
	val1 := kll.Int8.Create("4")
	val2 := kll.Int8.Create(7)
	fmt.Println(val1, val2)
	result, _ := val1.Sum(val2)
	fmt.Println(result)
	if !result.Is(kll.Int8.Create(11)) {
		t.Fail()
	}
}
func TestValueSub(t *testing.T) {
	val1 := kll.Int8.Create("4")
	val2 := kll.Int8.Create(7)
	fmt.Println(val1, val2)
	result, _ := val1.Sub(val2)
	fmt.Println(result)
	if !result.Is(kll.Int8.Create(-3)) {
		t.Fail()
	}
}

func TestValueMult(t *testing.T) {
	val1 := kll.Int8.Create("5")
	val2 := kll.Int8.Create(8)
	fmt.Println(val1, val2)
	result, _ := val1.Mult(val2)
	fmt.Println(result)
	if !result.Is(kll.Int8.Create(40)) {
		t.Fail()
	}
}

func TestValueDiv(t *testing.T) {
	val1 := kll.Int8.Create("50")
	val2 := kll.Int8.Create(2)
	fmt.Println(val1, val2)
	result, _ := val1.Div(val2)
	fmt.Println(result)
	if !result.Is(kll.Int8.Create(25)) {
		t.Fail()
	}
}

func TestValueErrorSum(t *testing.T) {
	val1 := kll.Int8.Create("4")
	fmt.Println(val1, kll.Null)
	result, err := val1.Sum(kll.Null)
	fmt.Println(result, err)
	if result != nil && err.String() != "ErrorImpossibleOperation:Can t Sum int8 with null" {
		t.Fail()
	}
}

func TestValueErrorCreate(t *testing.T) {
	err := kll.ImpossibleOperation.Create(kll.ErrorArgs{Callstack: []kll.Callstack{{Line: 5, Col: 5, Show: "123456"}}, Args: []any{"Sum", "int8", "null"}})
	fmt.Println(err)
	const sn = "Error:Impossible Operation\nLine: 5, Col: 5\n123456\n    ^\nCan t Sum int8 with null"
	if err.String() != sn {
		t.Fail()
	}
}

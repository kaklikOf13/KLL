package kll_test

import (
	"fmt"
	"testing"

	"github.com/kaklikOf13/KLL/kll"
)

func TestLexerSum(t *testing.T) {
	toks, err := kll.Tokenizer("1+ 10")
	if !err.Is(kll.Success) {
		t.Fail()
	}
	fmt.Println(toks[0].Value.Type())
	fmt.Println(toks)
	if fmt.Sprint(toks) != "[{FLOAT:1} {SUM} {FLOAT:10}]" {
		t.Fail()
	}
}
func TestLexerString(t *testing.T) {
	toks, err := kll.Tokenizer("'hua'-\"zaaa\"")
	if !err.Is(kll.Success) {
		t.Fail()
	}
	fmt.Println(toks[0].Value.Type())
	fmt.Println(toks)
	if fmt.Sprint(toks) != "[{STRING:hua} {SUB} {STRING:zaaa}]" {
		t.Fail()
	}
}

func TestLexerName(t *testing.T) {
	toks, err := kll.Tokenizer("hello*world")
	if !err.Is(kll.Success) {
		t.Fail()
	}
	fmt.Println(toks[0].Value.Type())
	fmt.Println(toks)
	if fmt.Sprint(toks) != "[{NAME:hello} {MUL} {NAME:world}]" {
		t.Fail()
	}
}

func TestLexerPoint(t *testing.T) {
	toks, err := kll.Tokenizer("hello.world")
	if !err.Is(kll.Success) {
		t.Fail()
	}
	fmt.Println(toks[0].Value.Type())
	fmt.Println(toks)
	if fmt.Sprint(toks) != "[{NAME:hello} {POINT} {NAME:world}]" {
		t.Fail()
	}
}

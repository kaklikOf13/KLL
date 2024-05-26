package kll

import (
	"strings"
)

const (
	TT_INT       = "INT"
	TT_FLOAT     = "FLOAT"
	TT_SUM       = "SUM"
	TT_SUB       = "SUB"
	TT_MUL       = "MUL"
	TT_DIV       = "DIV"
	TT_INTERATOR = "INTERATOR"
	TT_NEWLINE   = "NEWLINE"
	TT_STRING    = "STRING"
	TT_NAME      = "NAME"
	TT_POINT     = "POINT"
	TT_LPAREN    = "LPAREN"
	TT_RPAREN    = "RPAREN"
)

const NUMBERS = "0123456789"
const VARS_NAME = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_"

type Token struct {
	Type  string
	Value Value
}

func (t Token) String() string {
	if t.Value == nil {
		return "{" + t.Type + "}"
	}
	return "{" + t.Type + ":" + t.Value.String() + "}"
}
func NewToken(tp string, value Value) Token {
	return Token{Type: tp, Value: value}
}

type Lexer struct {
	Pos    uint64
	Ch     rune
	Input  string
	Tokens []Token
	Line   uint32
	Col    uint32
	Lines  []string
}

func (l *Lexer) NextToken() {
	if l.Pos >= uint64(len(l.Input)) {
		l.Ch = 0
	} else {
		l.Ch = rune(l.Input[l.Pos])
		l.Pos++
		l.Col++
	}
}

func (l *Lexer) GetString(closer rune) Error {
	col := l.Col
	line := l.Line
	l.NextToken()
	value := ""
	running := true
	for running {
		switch l.Ch {
		case 0:
			return Unclosed.Instantiate(ErrorArgs{Args: []any{string(closer)}, Callstack: []Callstack{{Line: line, Col: col, Show: l.Lines[line-1]}}}).(Error)
		case '\n':
			return InvalidNewLine.Instantiate(ErrorArgs{Args: []any{}, Callstack: []Callstack{{Line: line, Col: col, Show: l.Lines[line-1]}}}).(Error)
		case closer:
			running = false
		default:
			value += string(l.Ch)
			l.NextToken()
		}
	}
	l.Tokens = append(l.Tokens, NewToken(TT_STRING, String.Instantiate(value)))
	return Success
}

func (l *Lexer) Main() Error {
	l.NextToken()
	l.Col = 1
	l.Line = 1
	l.Lines = strings.Split(l.Input, "\n")
	for l.Ch != 0 {
		ok := false
		switch l.Ch {
		case '\n':
			l.Tokens = append(l.Tokens, NewToken(TT_NEWLINE, nil))
			l.Line++
			l.Col = 0
		case '+':
			l.Tokens = append(l.Tokens, NewToken(TT_SUM, nil))
		case '-':
			l.Tokens = append(l.Tokens, NewToken(TT_SUB, nil))
		case '*':
			l.Tokens = append(l.Tokens, NewToken(TT_MUL, nil))
		case '/':
			l.Tokens = append(l.Tokens, NewToken(TT_DIV, nil))
		case '(':
			l.Tokens = append(l.Tokens, NewToken(TT_LPAREN, nil))
		case ')':
			l.Tokens = append(l.Tokens, NewToken(TT_RPAREN, nil))
		case '"':
			err := l.GetString('"')
			if !err.Is(Success) {
				return err
			}
		case '\'':
			err := l.GetString('\'')
			if !err.Is(Success) {
				return err
			}
		case ' ', '\t', rune(13):
			ok = false
		default:
			ok = true
		}
		if strings.ContainsRune(VARS_NAME, l.Ch) {
			content := ""
			for l.Ch != 0 && strings.ContainsRune(VARS_NAME+NUMBERS, l.Ch) {
				content += string(l.Ch)
				l.NextToken()
			}
			l.Tokens = append(l.Tokens, NewToken(TT_NAME, String.Instantiate(content)))
			continue
		} else if strings.ContainsRune(NUMBERS+".", l.Ch) {
			float := false
			content := ""
			for l.Ch != 0 && strings.ContainsRune(NUMBERS+".", l.Ch) {
				if l.Ch == '.' {
					if float {
						l.Tokens = append(l.Tokens, NewToken(TT_INTERATOR, nil))
						break
					}
				} else {
					float = true
				}
				content += string(l.Ch)
				l.NextToken()
			}
			if content == "." {
				l.Tokens = append(l.Tokens, NewToken(TT_POINT, nil))
			} else if float {
				l.Tokens = append(l.Tokens, NewToken(TT_FLOAT, String.Instantiate(content)))
			} else {
				l.Tokens = append(l.Tokens, NewToken(TT_INT, String.Instantiate(content)))
			}
			continue
		} else if ok {
			return InvalidChar.Instantiate(ErrorArgs{Args: []any{l.Ch}, Callstack: []Callstack{}}).(Error)
		}
		l.NextToken()
	}
	return Success
}

func Tokenizer(txt string) ([]Token, Error) {
	l := Lexer{Input: txt}
	err := l.Main()
	if !err.Is(Success) && len(err.Callstack) == 0 {
		err.Callstack = append(err.Callstack, Callstack{Line: l.Line, Col: l.Col, Show: l.Lines[l.Line-1]})
	}
	return l.Tokens, err
}

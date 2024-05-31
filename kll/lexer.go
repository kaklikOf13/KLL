package kll

import (
	"strings"
)

const (
	TT_INT         = "INT"
	TT_FLOAT       = "FLOAT"
	TT_SUM         = "SUM"
	TT_SUB         = "SUB"
	TT_MUL         = "MUL"
	TT_DIV         = "DIV"
	TT_INTERATOR   = "INTERATOR"
	TT_NEWLINE     = "NEWLINE"
	TT_STRING      = "STRING"
	TT_NAME        = "NAME"
	TT_POINT       = "POINT"
	TT_LPAREN      = "LPAREN"
	TT_RPAREN      = "RPAREN"
	TT_NULL        = "null"
	TT_SPLIT       = "SPLIT"
	TT_DOUBLEPOINT = "DOUBLEPOINT"
)

const NUMBERS = "0123456789"
const VARS_NAME = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_"

type Token struct {
	Type      string
	Value     Value
	Callstack Callstack
}

func (t Token) String() string {
	if t.Value == nil {
		return "{" + t.Type + "}"
	}
	return "{" + t.Type + ":" + t.Value.String() + "}"
}
func NewToken(tp string, value Value, callstack Callstack) Token {
	return Token{Type: tp, Value: value, Callstack: callstack}
}

type Lexer struct {
	Pos       uint64
	Ch        rune
	Input     string
	Tokens    []Token
	Callstack Callstack
	Lines     []string
}

func (l *Lexer) LoadToken() {
	if l.Pos >= uint64(len(l.Input)) {
		l.Ch = 0
	} else {
		l.Ch = rune(l.Input[l.Pos])
	}
}
func (l *Lexer) NextToken() {
	l.Pos++
	l.Callstack.Col++
	l.LoadToken()
}

func (l *Lexer) GetString(closer rune) Error {
	callstack := l.Callstack
	l.NextToken()
	value := ""
	running := true
	for running {
		switch l.Ch {
		case 0:
			return Unclosed.Create(ErrorArgs{Args: []any{string(closer)}, Callstack: []Callstack{callstack}}).(Error)
		case '\n':
			return InvalidNewLine.Create(ErrorArgs{Args: []any{}, Callstack: []Callstack{callstack}}).(Error)
		case closer:
			running = false
		default:
			value += string(l.Ch)
			l.NextToken()
		}
	}
	l.Tokens = append(l.Tokens, NewToken(TT_STRING, String.Create(value), callstack))
	return Success
}

func (l *Lexer) Main() Error {
	l.LoadToken()
	l.Lines = strings.Split(l.Input, "\n")
	l.Callstack = Callstack{Line: 1, Col: 1, Show: l.Lines[0]}
	for l.Ch != 0 {
		ok := false
		switch l.Ch {
		case '\n':
			l.Tokens = append(l.Tokens, NewToken(TT_NEWLINE, nil, l.Callstack))
			l.Callstack.Line++
			l.Callstack.Col = 0
			l.Callstack.Show = l.Lines[l.Callstack.Line-1]
		case '+':
			l.Tokens = append(l.Tokens, NewToken(TT_SUM, nil, l.Callstack))
		case '-':
			l.Tokens = append(l.Tokens, NewToken(TT_SUB, nil, l.Callstack))
		case '*':
			l.Tokens = append(l.Tokens, NewToken(TT_MUL, nil, l.Callstack))
		case '/':
			l.Tokens = append(l.Tokens, NewToken(TT_DIV, nil, l.Callstack))
		case '(':
			l.Tokens = append(l.Tokens, NewToken(TT_LPAREN, nil, l.Callstack))
		case ')':
			l.Tokens = append(l.Tokens, NewToken(TT_RPAREN, nil, l.Callstack))
		case ';':
			l.Tokens = append(l.Tokens, NewToken(TT_SPLIT, nil, l.Callstack))
		case ':':
			l.Tokens = append(l.Tokens, NewToken(TT_DOUBLEPOINT, nil, l.Callstack))
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
			callstack := l.Callstack
			for l.Ch != 0 && strings.ContainsRune(VARS_NAME+NUMBERS, l.Ch) {
				content += string(l.Ch)
				l.NextToken()
			}
			l.Tokens = append(l.Tokens, NewToken(TT_NAME, String.Create(content), callstack))
			continue
		} else if strings.ContainsRune(NUMBERS+".", l.Ch) {
			float := false
			content := ""
			callstack := l.Callstack
			for l.Ch != 0 && strings.ContainsRune(NUMBERS+".", l.Ch) {
				if l.Ch == '.' {
					if float {
						l.Tokens = append(l.Tokens, NewToken(TT_INTERATOR, nil, callstack))
						break
					}
					float = true
				}
				content += string(l.Ch)
				l.NextToken()
			}
			if content == "." {
				l.Tokens = append(l.Tokens, NewToken(TT_POINT, nil, callstack))
			} else if float {
				l.Tokens = append(l.Tokens, NewToken(TT_FLOAT, String.Create(content), callstack))
			} else {
				l.Tokens = append(l.Tokens, NewToken(TT_INT, String.Create(content), callstack))
			}
			continue
		} else if ok {
			return InvalidChar.Create(ErrorArgs{Args: []any{l.Ch}, Callstack: []Callstack{}}).(Error)
		}
		l.NextToken()
	}
	return Success
}

func Tokenizer(txt string) ([]Token, Error) {
	l := Lexer{Input: txt}
	err := l.Main()
	if !err.Is(Success) && len(err.Callstack) == 0 {
		err.Callstack = append(err.Callstack, l.Callstack)
	}
	return l.Tokens, err
}

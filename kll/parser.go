package kll

import "fmt"

const (
	NT_NULL   = "NULL"
	NT_INT    = "INT"
	NT_FLOAT  = "FLOAT"
	NT_STRING = "STRING"
	NT_SUM    = "SUM"
	NT_SUB    = "SUB"
	NT_MUL    = "MUL"
	NT_DIV    = "DIV"
	NT_NAME   = "NAME"
	NT_GET    = "GET"
	NT_TYPE   = "TYPE"
	NT_SUBOP  = "SUBOP"
)

type Node struct {
	ValueBase
	Tp        string
	NodeValue []Value
	Callstack Callstack
}

func (n Node) String() string {
	return fmt.Sprintf("<%v:%v>", n.Tp, n.NodeValue)
}
func NewNode(Type string, Value []Value, callstack Callstack) Node {
	return Node{Tp: Type, NodeValue: Value, Callstack: callstack}
}

type Block []Token
type Code []Node

type Parser struct {
	Blocks   []Block
	Block    uint
	Tok      uint
	CurTok   Token
	CurBlock Block
	Main     func() (Node, Error)
}

func (p *Parser) NextToken() {
	if uint(len(p.CurBlock)) <= p.Tok {
		p.CurTok = NewToken(TT_NULL, nil, Callstack{0, 0, ""})
	} else {
		p.CurTok = p.CurBlock[p.Tok]
	}
	p.Tok++
}
func (p *Parser) LoadBlock() {
	if uint(len(p.Blocks)) <= p.Block {
		p.CurBlock = Block{}
	} else {
		p.CurBlock = p.Blocks[p.Block]
		p.Tok = 0
		p.NextToken()
	}
}
func (p *Parser) NextBlock() {
	p.Block++
	p.LoadBlock()
}

func (p *Parser) Value() (Node, Error) {
	tok := p.CurTok
	p.NextToken()
	switch tok.Type {
	case TT_FLOAT:
		return NewNode(NT_FLOAT, []Value{tok.Value}, tok.Callstack), Success
	case TT_INT:
		return NewNode(NT_INT, []Value{tok.Value}, tok.Callstack), Success
	case TT_STRING:
		return NewNode(NT_STRING, []Value{tok.Value}, tok.Callstack), Success
	case TT_NAME:
		return NewNode(NT_NAME, []Value{tok.Value}, tok.Callstack), Success
	case TT_LPAREN:
		node, err := p.Main()
		if !err.Is(Success) {
			return node, err
		}
		node = NewNode(NT_SUBOP, []Value{node}, tok.Callstack)
		if p.CurTok.Type != TT_RPAREN {
			return node, Unclosed.Create(ErrorArgs{Callstack: []Callstack{tok.Callstack}}).(Error)
		}
		return node, Success
	default:
		return NewNode(NT_NULL, nil, tok.Callstack), InvalidPosition.Create(ErrorArgs{Callstack: []Callstack{tok.Callstack}, Args: []any{}}).(Error)
	}
}

func (p *Parser) Name() (Node, Error) {
	node, err := p.Value()
	if !err.Is(Success) {
		return node, err
	}
	for containsString([]string{TT_POINT}, p.CurTok.Type) {
		tok := p.CurTok
		p.NextToken()
		switch tok.Type {
		case TT_POINT:
			node2, err2 := p.Value()
			if !err2.Is(Success) {
				return node2, err2
			}
			node = NewNode(NT_GET, []Value{node, node2}, tok.Callstack)
		}
	}
	return node, err
}
func (p *Parser) Type() (Node, Error) {
	node, err := p.Name()
	if !err.Is(Success) {
		return node, err
	}
	for containsString([]string{TT_DOUBLEPOINT}, p.CurTok.Type) {
		tok := p.CurTok
		p.NextToken()
		switch tok.Type {
		case TT_DOUBLEPOINT:
			node2, err2 := p.Name()
			if !err2.Is(Success) {
				return node2, err2
			}
			node = NewNode(NT_TYPE, []Value{node, node2}, tok.Callstack)
		}
	}
	return node, err
}

func (p *Parser) OperationsNLevel1() (Node, Error) {
	node, err := p.Type()
	if !err.Is(Success) {
		return node, err
	}
	for containsString([]string{TT_MUL, TT_DIV}, p.CurTok.Type) {
		tok := p.CurTok
		p.NextToken()
		switch tok.Type {
		case TT_MUL:
			node2, err2 := p.Type()
			if !err2.Is(Success) {
				return node2, err2
			}
			node = NewNode(NT_MUL, []Value{node, node2}, tok.Callstack)
		case TT_DIV:
			node2, err2 := p.Type()
			if !err2.Is(Success) {
				return node2, err2
			}
			node = NewNode(NT_DIV, []Value{node, node2}, tok.Callstack)
		}
	}
	return node, err
}
func (p *Parser) OperationsNLevel2() (Node, Error) {
	node, err := p.OperationsNLevel1()
	if !err.Is(Success) {
		return node, err
	}
	for containsString([]string{TT_SUM, TT_SUB}, p.CurTok.Type) {
		tok := p.CurTok
		p.NextToken()
		switch tok.Type {
		case TT_SUM:
			node2, err2 := p.OperationsNLevel1()
			if !err2.Is(Success) {
				return node2, err2
			}
			node = NewNode(NT_SUM, []Value{node, node2}, tok.Callstack)
		case TT_SUB:
			node2, err2 := p.OperationsNLevel1()
			if !err2.Is(Success) {
				return node2, err2
			}
			node = NewNode(NT_SUB, []Value{node, node2}, tok.Callstack)
		}
	}
	return node, err
}

func Parse(txt string) (Code, Error) {
	toks, err := Tokenizer(txt)
	if !err.Is(Success) {
		return Code{}, err
	}
	return ParseToks(toks)
}

func ParseToks(tokens []Token) (Code, Error) {
	p := &Parser{}
	p.Blocks = splitWithSeparators(tokens, []string{TT_NEWLINE, TT_SPLIT})
	p.Main = p.OperationsNLevel2
	p.Block = 0
	p.LoadBlock()
	ret := Code{}
	for p.Block < uint(len(p.Blocks)) {
		n, err := p.Main()
		if !err.Is(Success) {
			return Code{}, err
		}
		ret = append(ret, n)
		p.NextBlock()
	}
	return ret, Success
}

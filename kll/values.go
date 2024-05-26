package kll

import (
	"fmt"
	"strconv"
	"strings"
)

// #region Value Base

type Value interface {
	String() string
	Type() Type
	Is(other Value) bool

	Sum(Value) (Value, Error)
	Sub(Value) (Value, Error)
	Mult(Value) (Value, Error)
	Div(Value) (Value, Error)

	Convert(Type) (Value, Error)
}

type ValueBase struct {
	Value Value
}

func (v ValueBase) String() string {
	return "null"
}
func (v ValueBase) Type() Type {
	return NullType
}
func (v ValueBase) Is(other Value) bool {
	switch other.(type) {
	case TypeBase, *TypeBase:
		return true
	}
	return false
}
func (v ValueBase) Sum(other Value) (Value, Error) {
	return nil, ImpossibleOperation.Instantiate(ErrorArgs{[]Callstack{}, []any{"Sum", v.Value.Type().String(), other.Type().String()}}).(Error)
}
func (v ValueBase) Sub(other Value) (Value, Error) {
	return nil, ImpossibleOperation.Instantiate(ErrorArgs{[]Callstack{}, []any{"Subtract", v.Value.Type().String(), other.Type().String()}}).(Error)
}
func (v ValueBase) Mult(other Value) (Value, Error) {
	return nil, ImpossibleOperation.Instantiate(ErrorArgs{[]Callstack{}, []any{"Multiply", v.Value.Type().String(), other.Type().String()}}).(Error)
}
func (v ValueBase) Div(other Value) (Value, Error) {
	return nil, ImpossibleOperation.Instantiate(ErrorArgs{[]Callstack{}, []any{"Divade", v.Value.Type().String(), other.Type().String()}}).(Error)
}
func (v ValueBase) Convert(other Type) (Value, Error) {
	if other.Is(NullType) {
		return Null, Success
	}
	return nil, CantConvert.Instantiate(ErrorArgs{[]Callstack{}, []any{v.Value.Type().String(), other.Type().String()}}).(Error)
}

// #endregion

// #region Type Base
type Type interface {
	Value
	Instantiate(arg any) Value
	ConvertibleWith(Type) bool
}

type TypeBase struct {
	ValueBase
}

func (t TypeBase) Instantiate(arg any) Value {
	return Null
}
func (t TypeBase) Is(other Value) bool {
	switch other.(type) {
	case *TypeBase, TypeBase:
		return true
	}
	return false
}
func (t TypeBase) ConvertibleWith(tp Type) bool {
	return t.Is(tp)
}

//#endregion

// #region Error

type Callstack struct {
	Line uint32
	Col  uint32
	Show string
}
type ErrorArgs struct {
	Callstack []Callstack
	Args      []any
}

type ErrorType struct {
	TypeBase
	Tp              string
	MessageCallback func(...any) string
}

func (et ErrorType) ConvertibleWith(other Type) bool {
	return et.Is(other)
}

func (et ErrorType) Instantiate(arg any) Value {
	args := arg.(ErrorArgs)
	var ret = Error{tp: et, ErrorArgs: args, Message: et.MessageCallback(args.Args...)}
	ret.Value = ret
	return ret
}
func (et ErrorType) Is(other Value) bool {
	switch other := other.(type) {
	case *ErrorType:
		return et.Tp == other.Tp
	case ErrorType:
		return et.Tp == other.Tp
	}
	return false
}
func (et ErrorType) String() string {
	return "Error:" + et.Tp
}
func (et ErrorType) Type() Type {
	return et
}
func NewErrorType(tp string, message_callback func(...any) string) ErrorType {
	return ErrorType{Tp: tp, MessageCallback: message_callback}
}

type Error struct {
	ValueBase
	ErrorArgs
	Message string
	tp      ErrorType
}

func (e Error) Type() Type {
	return e.tp
}
func (e Error) Is(other Value) bool {
	switch other := other.(type) {
	case *Error:
		return e.tp.Is(other.tp)
	case Error:
		return e.tp.Is(other.tp)
	}
	return false
}
func (e Error) String() string {
	ret := e.tp.String()
	for i, c := range e.Callstack {
		//+strings.Repeat(" ", len(c.Show)-int(c.Col))
		ret += "\n" + fmt.Sprintf("Line: %v, Col: %v\n%v\n%v\n", c.Line, c.Col, c.Show, strings.Repeat(" ", int(c.Col-1))+"^")
		if i == len(e.Callstack)-1 {
			ret += e.Message
		}
	}
	return ret
}

// #endregion

// #region Int
type int8Type struct {
	TypeBase
}

func (t int8Type) ConvertibleWith(other Type) bool {
	switch other.(type) {
	case int8Type, *int8Type /*, uint8Type, *uint8Type, int16Type, *int16Type, uint16Type, *uint16Type, int32Type, *int32Type, uint32Type, *uint32Type, int64Type, *int64Type, uint64Type, *uint64Type*/ :
		return true
	}
	return false
}
func (t int8Type) String() string {
	return "int8"
}
func (t int8Type) Is(other Value) bool {
	switch other.(type) {
	case *int8Type, int8Type:
		return true
	}
	return false
}
func (t int8Type) Type() Type {
	return t
}
func (t int8Type) Instantiate(arg any) Value {
	var val int8 = 0
	switch arg := arg.(type) {
	case int:
		val = int8(arg)
	case int8:
		val = int8(arg)
	case int16:
		val = int8(arg)
	case int32:
		val = int8(arg)
	case int64:
		val = int8(arg)
	case uint:
		val = int8(arg)
	case uint8:
		val = int8(arg)
	case uint16:
		val = int8(arg)
	case uint32:
		val = int8(arg)
	case uint64:
		val = int8(arg)
	case string:
		valb, _ := strconv.ParseInt(arg, 0, 8)
		val = int8(valb)
	}
	var ret = Int8Value{Number: val, ValueBase: ValueBase{Value: nil}}
	ret.Value = &ret
	return ret
}

type Int8Value struct {
	ValueBase
	Number int8
}

func (iv Int8Value) String() string {
	return fmt.Sprint(iv.Number)
}

func (iv Int8Value) Type() Type {
	return Int8
}

func (iv Int8Value) Is(other Value) bool {
	switch other := other.(type) {
	case Int8Value:
		return iv.Number == other.Number
	case *Int8Value:
		return iv.Number == other.Number
	}
	return false
}
func (iv Int8Value) Sum(other Value) (Value, Error) {
	switch other := other.(type) {
	case Int8Value:
		return Int8.Instantiate(iv.Number + other.Number), Success
	case *Int8Value:
		return Int8.Instantiate(iv.Number + other.Number), Success
	}
	return nil, ImpossibleOperation.Instantiate(ErrorArgs{[]Callstack{}, []any{"Sum", iv.Value.Type().String(), other.Type().String()}}).(Error)
}
func (iv Int8Value) Sub(other Value) (Value, Error) {
	switch other := other.(type) {
	case Int8Value:
		return Int8.Instantiate(iv.Number - other.Number), Success
	case *Int8Value:
		return Int8.Instantiate(iv.Number - other.Number), Success
	}
	return nil, ImpossibleOperation.Instantiate(ErrorArgs{[]Callstack{}, []any{"Sub", iv.Value.Type().String(), other.Type().String()}}).(Error)
}
func (iv Int8Value) Mult(other Value) (Value, Error) {
	switch other := other.(type) {
	case Int8Value:
		return Int8.Instantiate(iv.Number * other.Number), Success
	case *Int8Value:
		return Int8.Instantiate(iv.Number * other.Number), Success
	}
	return nil, ImpossibleOperation.Instantiate(ErrorArgs{[]Callstack{}, []any{"Mult", iv.Value.Type().String(), other.Type().String()}}).(Error)
}
func (iv Int8Value) Div(other Value) (Value, Error) {
	switch other := other.(type) {
	case Int8Value:
		return Int8.Instantiate(iv.Number / other.Number), Success
	case *Int8Value:
		return Int8.Instantiate(iv.Number / other.Number), Success
	}
	return nil, ImpossibleOperation.Instantiate(ErrorArgs{[]Callstack{}, []any{"Div", iv.Value.Type().String(), other.Type().String()}}).(Error)
}
func (iv Int8Value) Convert(other Type) (Value, Error) {
	if other.Type().ConvertibleWith(iv.Type()) {
		return other.Instantiate(iv.Number), Success
	}
	return nil, CantConvert.Instantiate(ErrorArgs{[]Callstack{}, []any{iv.Value.Type().String(), other.Type().String()}}).(Error)
}

// #endregion

// #region Uint

// #endregion

// #region String

type StringType struct {
	TypeBase
	Limit uint32
}

func (st StringType) String() string {
	if st.Limit == 0 {
		return "string"
	} else {
		return fmt.Sprintf("string[%v]", st.Limit)
	}
}
func (st StringType) Type() Type {
	return st
}
func (st StringType) Instantiate(arg any) Value {
	ret := StringValue{tp: st, value: fmt.Sprint(arg), ValueBase: ValueBase{Value: nil}}
	ret.Value = &ret
	return ret
}
func (st StringType) ConvertibleWith(other Type) bool {
	switch other := other.(type) {
	case StringType:
		return other.Limit >= st.Limit
	case *StringType:
		return other.Limit >= st.Limit
	}
	return false
}
func NewStringType(limit uint32) Type {
	ret := StringType{Limit: limit, TypeBase: TypeBase{ValueBase: ValueBase{Value: nil}}}
	ret.Value = &ret
	return ret
}

type StringValue struct {
	ValueBase
	tp    StringType
	value string
}

func (s StringValue) Type() Type {
	return s.tp
}
func (s StringValue) String() string {
	return s.value
}
func (s StringValue) Is(other Value) bool {
	switch other := other.(type) {
	case StringValue:
		return s.value == other.value
	case *StringValue:
		return s.value == other.value
	}
	return false
}

func (s StringValue) Sum(other Value) (Value, Error) {
	switch other := other.(type) {
	case StringValue:
		ret := s.value + other.value
		if s.tp.Limit != 0 && s.tp.Limit < uint32(len(ret)) {
			break
		}
		return nil, StringOutOfLimits.Instantiate(ErrorArgs{[]Callstack{}, []any{s.tp.String(), s.tp.Limit}}).(Error)
	}
	return nil, ImpossibleOperation.Instantiate(ErrorArgs{[]Callstack{}, []any{"Sum", s.tp.String(), other.Type().String()}}).(Error)
}

// #endregion

// #region Values

var (
	NullType Type = TypeBase{ValueBase{}}
	Int8     Type = int8Type{}
	String   Type = NewStringType(0)

	SuccessType Type = NewErrorType("Success", func(a ...any) string {
		return "Success To Realise This Operation"
	})
	Unclosed Type = NewErrorType("Unclosed", func(a ...any) string {
		return fmt.Sprintf("\"%v\" was not closed", a[0])
	})
	InvalidNewLine Type = NewErrorType("Invalid New Line", func(a ...any) string {
		return "You Can T Use NewLine On This Moment"
	})
	CantConvert Type = NewErrorType("Can'T Convert", func(a ...any) string {
		return fmt.Sprintf("Can t Convert %s to %s", a[0].(string), a[1].(string))
	})
	InvalidChar Type = NewErrorType("Invalid Char", func(a ...any) string {
		return fmt.Sprintf("The Char \"%v\" is Invalid", a[0])
	})
	ImpossibleOperation Type = NewErrorType("Impossible Operation", func(a ...any) string {
		return fmt.Sprintf("Can t %s %s with %s", a[0].(string), a[1].(string), a[2].(string))
	})
	StringOutOfLimits Type = NewErrorType("String Out Of Limits", func(a ...any) string {
		return fmt.Sprintf("The String %s can t ultrapass %v", a[0].(string), a[1].(uint32))
	})
)
var (
	Null Value = ValueBase{Value: &ValueBase{}}
)
var (
	Success Error = SuccessType.Instantiate(ErrorArgs{}).(Error)
)

// #endregion

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
	return "Error" + et.Tp
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
	ret := e.tp.String() + ":" + e.Message + "\n"
	for i, c := range e.Callstack {
		if i != 0 {
			ret += "\n"
		}
		ret += "\n" + fmt.Sprintf("Line: %v, Col: %v\n%v\n%v", c.Line, c.Col, c.Show, strings.Repeat("_", int(c.Col-1))+"^"+strings.Repeat("_", len(c.Show)-int(c.Col)))
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

// #region Values

var (
	NullType Type = TypeBase{ValueBase{}}
	Int8     Type = int8Type{}

	SuccessType Type = NewErrorType("Success", func(a ...any) string {
		return "Success To Realise This Operation"
	})
	CantConvert Type = NewErrorType("CanTConvert", func(a ...any) string {
		return fmt.Sprintf("Can t Convert %s to %s", a[0].(string), a[1].(string))
	})
	ImpossibleOperation Type = NewErrorType("ImpossibleOperation", func(a ...any) string {
		return fmt.Sprintf("Can t %s %s with %s", a[0].(string), a[1].(string), a[2].(string))
	})
)
var (
	Null Value = ValueBase{Value: &ValueBase{}}
)
var (
	Success Error = SuccessType.Instantiate(ErrorArgs{}).(Error)
)

// #endregion

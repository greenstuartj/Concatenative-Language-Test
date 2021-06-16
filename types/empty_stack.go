package types

import "errors"

type EmptyStack struct {
}

func MakeEmptyStack() *EmptyStack{
	return &EmptyStack{}
}

func (es *EmptyStack) String() string {
	return "."
}

func (es *EmptyStack) ToNumber() (*Number, error) {
	return nil, errors.New("Not a number")
}

func (es *EmptyStack) ToString() (*String, error) {
	return MakeString("."), nil
}

func (es *EmptyStack) Apply(i *Interpreter) (*Interpreter, error) {
	return nil, errors.New("Cannot be applied")
}

func (es *EmptyStack) Type() uint8 {
	return EmptyStackT
}

func (es *EmptyStack) GetName() uint32 {
	return 0
}

func (es *EmptyStack) Unify(i *Interpreter) (*Interpreter, bool) {
	if i.Stack.Nullp() {
		return i, true
	} else {
		return i, false
	}
}

func (es *EmptyStack) StrictCompare(t2 Type) int {
	if t2.Type() != EmptyStackT {
		if EmptyStackT < t2.Type() {
			return -1
		} else {
			return 1
		}
	} else {
		return 0
	}
}

func (es *EmptyStack) Compare(t2 Type) (int, error) {
	return 0, errors.New("Cannot compare")
}

func (es *EmptyStack) ToList() (*List, error) {
	return nil, errors.New("Cannot be enlisted")
}

func (es *EmptyStack) ToChar() (*Char, error) {
	return nil, errors.New("Cannot make a char")
}

func (es *EmptyStack) ToBool() (*Bool, error) {
	return MakeBool(true), nil
}

func (es *EmptyStack) ToTuple() (*Tuple, error) {
	return nil, errors.New("Cannot convert to tuple")
}

func (es *EmptyStack) ToMap() (*Tree, error) {
	return nil, errors.New("Cannot convert to a dict")
}

func (es *EmptyStack) ToSymbol() (*Symbol, error) {
	return nil, errors.New("Cannot convert to symbol")
}

func (es *EmptyStack) ToSet() (*Tree, error) {
	return nil, errors.New("Cannot convert to set")
}

func (es *EmptyStack) ToCore() (*Core, error) {
	return nil, errors.New("Cannot make core")
}

func (es *EmptyStack) Get(t2 Type) (Type, error) {
	return nil, errors.New("Cannot get")
}

func (es *EmptyStack) Set(t2 Type, t3 Type) (Type, error) {
	return nil, errors.New("Cannot set")
}

func (es *EmptyStack) ToFstream() (*Fstream, error) {
	return nil, errors.New("Not a file stream")
}

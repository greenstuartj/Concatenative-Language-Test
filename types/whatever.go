package types

import "errors"

type Whatever struct {
}

func MakeWhatever() *Whatever {
	return &Whatever{}
}

func (w *Whatever) String() string {
	return "_"
}

func (w *Whatever) ToNumber() (*Number, error) {
	return nil, errors.New("Cannot convert to a number")
}

func (w *Whatever) ToString() (*String, error) {
	return MakeString("_"), nil
}

func (w *Whatever) Apply(i *Interpreter) (*Interpreter, error) {
	return nil, errors.New("Cannot be applied")
}

func (w *Whatever) Type() uint8 {
	return WhateverT
}

func (w *Whatever) GetName() uint32 {
	return 0
}

func (w *Whatever) Unify(i *Interpreter) (*Interpreter, bool) {
	i.Stack = i.Stack.Cdr()
	return i, true
}

func (w *Whatever) StrictCompare(t2 Type) int {
	if t2.Type() != WhateverT {
		if WhateverT < t2.Type() {
			return -1
		} else {
			return 1
		}
	} else {
		return 0
	}
}

func (w *Whatever) Compare(t2 Type) (int, error) {
	return 0, errors.New("Cannot compare")
}

func (w *Whatever) ToList() (*List, error) {
	return nil, errors.New("Cannot be enlisted")
}

func (w *Whatever) ToChar() (*Char, error) {
	return nil, errors.New("Cannot make a char")
}

func (w *Whatever) ToBool() (*Bool, error) {
	return MakeBool(true), nil
}

func (w *Whatever) ToTuple() (*Tuple, error) {
	return nil, errors.New("Cannot convert to tuple")
}

func (w *Whatever) ToMap() (*Tree, error) {
	return nil, errors.New("Cannot convert to a dict")
}

func (w *Whatever) ToSymbol() (*Symbol, error) {
	return nil, errors.New("Cannot convert to symbol")
}

func (w *Whatever) ToSet() (*Tree, error) {
	return nil, errors.New("Cannot convert to set")
}

func (w *Whatever) ToCore() (*Core, error) {
	return nil, errors.New("Cannot make core")
}

func (w *Whatever) Get(t2 Type) (Type, error) {
	return nil, errors.New("Cannot get")
}

func (w *Whatever) Set(t2 Type, t3 Type) (Type, error) {
	return nil, errors.New("Cannot set")
}

func (w *Whatever) ToFstream() (*Fstream, error) {
	return nil, errors.New("Not a file stream")
}

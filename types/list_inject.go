package types

import "errors"

type ListInject struct {
	inject Type
}

func MakeListInject(inject Type) *ListInject {
	return &ListInject{inject}
}

func (li *ListInject) String() string {
	return ":" + li.inject.String()
}

func (li *ListInject) ToNumber() (*Number, error) {
	return nil, errors.New("Cannot convert to number")
}

func (li *ListInject) ToString() (*String, error) {
	return nil, errors.New("Cannot make a string")
}

func (li *ListInject) Apply(i *Interpreter) (*Interpreter, error) {
	new_prog := MakeList()
	new_prog = new_prog.Cons(li.inject)
	new_i := MakeInterpreter(new_prog, i.Env)
	lst, err := i.Stack.Car().ToList()
	i.Stack = i.Stack.Cdr()
	if err != nil {
		return i, err
	}
	new_i.Stack = lst
	new_i, err = new_i.Run()
	if err != nil {
		return i, err
	}
	i.Stack = i.Stack.Cons(new_i.Stack)
	return i, nil
}

func (li *ListInject) Type() uint8 {
	return ListInjectT
}

func (li *ListInject) GetName() uint32 {
	return 0
}

func (li *ListInject) Unify(i *Interpreter) (*Interpreter, bool) {
	return i, false
}

func (li *ListInject) StrictCompare(t2 Type) int {
	if t2.Type() != ListInjectT {
		if ListInjectT < t2.Type() {
			return -1
		} else {
			return 1
		}
	} else {
		return 0
	}
}

func (li *ListInject) Compare(t2 Type) (int, error) {
	return 0, errors.New("Cannot compare")
}

func (li *ListInject) ToList() (*List, error) {
	return nil, errors.New("Cannot enlist")
}

func (li *ListInject) ToChar() (*Char, error) {
	return nil, errors.New("Cannot make a char")
}

func (li *ListInject) ToBool() (*Bool, error) {
	return MakeBool(true), nil
}

func (li *ListInject) ToTuple() (*Tuple, error) {
	return nil, errors.New("Cannot convert to tuple")
}

func (li *ListInject) ToMap() (*Tree, error) {
	return nil, errors.New("Cannot convert to a dict")
}

func (li *ListInject) ToSymbol() (*Symbol, error) {
	return nil, errors.New("Cannot convert to symbol")
}

func (li *ListInject) ToSet() (*Tree, error) {
	return nil, errors.New("Cannot convert to set")
}

func (li *ListInject) ToCore() (*Core, error) {
	return nil, errors.New("Cannot make core")
}

func (li *ListInject) Get(t2 Type) (Type, error) {
	return nil, errors.New("Cannot get")
}

func (li *ListInject) Set(t2 Type, t3 Type) (Type, error) {
	return nil, errors.New("Cannot set")
}

func (li *ListInject) ToFstream() (*Fstream, error) {
	return nil, errors.New("Not a file stream")
}

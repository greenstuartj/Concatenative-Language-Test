package types

import "errors"

type At struct {
	g Type
}

func MakeAt(t Type) *At {
	return &At{t}
}

func (a *At) String() string {
	return "@" + a.g.String()
}

func (a *At) Apply(i *Interpreter) (*Interpreter, error) {
	var get Type
	if a.g.Type() == FunctionT || a.g.Type() == CoreT {
		get = a.g
	} else {
		prog := MakeList()
		prog = prog.Cons(a.g)
		new_i := MakeInterpreter(prog, i.Env)
		new_i, err := new_i.Run()
		if err != nil {
			return i, err
		}
		val := new_i.Stack.Car()
		get = val
	}
	top := i.Stack.Car()
	got, err := top.Get(get)
	if err != nil {
		return i, err
	}
	i.Stack = i.Stack.Cdr().Cons(got)
	return i, nil
}

func (a *At) Type() uint8 {
	return AtT
}

func (a *At) GetName() uint32 {
	return 0
}

func (a *At) Unify(i *Interpreter) (*Interpreter, bool) {
	return i, false
}

func (a *At) StrictCompare(t2 Type) int {
	if t2.Type() != AtT {
		if AtT < t2.Type() {
			return -1
		} else {
			return 1
		}
	}
	return 0
}

func (a *At) Compare(t2 Type) (int, error) {
	return -1, errors.New("Cannot compare")
}

func (a *At) ToNumber() (*Number, error) {
	return nil, errors.New("Cannot make number")
}

func (a *At) ToString() (*String, error) {
	return nil, errors.New("Cannot make string")
}

func (a *At) ToList() (*List, error) {
	return nil, errors.New("Cannot make list")
}

func (a *At) ToChar() (*Char, error) {
	return nil, errors.New("Cannot make char")
}

func (a *At) ToBool() (*Bool, error) {
	return MakeBool(true), nil
}

func (a *At) ToTuple() (*Tuple, error) {
	return nil, errors.New("Cannot make tuple")
}

func (a *At) ToMap() (*Tree, error) {
	return nil, errors.New("Cannot make map")
}

func (a *At) ToSymbol() (*Symbol, error) {
	return nil, errors.New("Cannot make symbol")
}

func (a *At) ToSet() (*Tree, error) {
	return nil, errors.New("Cannot make set")
}

func (a *At) ToCore() (*Core, error) {
	return nil, errors.New("Cannot make core")
}

func (a *At) Get(t2 Type) (Type, error) {
	return nil, errors.New("Cannot get")
}

func (a *At) Set(t2 Type, t3 Type) (Type, error) {
	return nil, errors.New("Cannot set")
}

func (a *At) ToFstream() (*Fstream, error) {
	return nil, errors.New("Not a file stream")
}

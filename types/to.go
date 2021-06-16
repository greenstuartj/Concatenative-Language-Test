package types

import "errors"

type To struct {
	s Type
}

func MakeTo(t Type) *To {
	return &To{t}
}

func (t *To) String() string {
	return "!" + t.s.String()
}

func (t *To) Apply(i *Interpreter) (*Interpreter, error) {
	prog := MakeList()
	prog = prog.Cons(t.s)
	new_i := MakeInterpreter(prog, i.Env)
	new_i, err := new_i.Run()
	if err != nil {
		return i, err
	}
	val := new_i.Stack.Car()
	if val.Type() == SetT {
		top := i.Stack.Car()
		s,err := val.Set(top, nil)
		if err != nil {
			return i, err
		}
		i.Stack = i.Stack.Cdr().Cons(s)
		return i, nil
	}
	top := i.Stack.Car()
	top2 := i.Stack.Cdr().Car()
	s, err := val.Set(top, top2)
	if err != nil {
		return i, err
	}
	i.Stack = i.Stack.Cdr().Cdr().Cons(s)
	return i, nil
}

func (t *To) Type() uint8 {
	return ToT
}

func (t *To) GetName() uint32 {
	return 0
}

func (t *To) Unify(i *Interpreter) (*Interpreter, bool) {
	return i, false
}

func (t *To) StrictCompare(t2 Type) int {
	if t2.Type() != ToT {
		if ToT < t2.Type() {
			return -1
		} else {
			return 1
		}
	}
	return 0
}

func (t *To) Compare(t2 Type) (int, error) {
	return -1, errors.New("Cannot compare")
}


func (t *To) ToNumber() (*Number, error) {
	return nil, errors.New("Cannot make number")
}

func (t *To) ToString() (*String, error) {
	return nil, errors.New("Cannot make string")
}

func (t *To) ToList() (*List, error) {
	return nil, errors.New("Cannot make list")
}

func (t *To) ToChar() (*Char, error) {
	return nil, errors.New("Cannot make char")
}

func (t *To) ToBool() (*Bool, error) {
	return MakeBool(true), nil
}

func (t *To) ToTuple() (*Tuple, error) {
	return nil, errors.New("Cannot make tuple")
}

func (t *To) ToMap() (*Tree, error) {
	return nil, errors.New("Cannot make map")
}

func (t *To) ToSymbol() (*Symbol, error) {
	return nil, errors.New("Cannot make symbol")
}

func (t *To) ToSet() (*Tree, error) {
	return nil, errors.New("Cannot make set")
}

func (t *To) ToCore() (*Core, error) {
	return nil, errors.New("Cannot make core")
}

func (t *To) Get(t2 Type) (Type, error) {
	return nil, errors.New("Cannot get")
}

func (t *To) Set(t2 Type, t3 Type) (Type, error) {
	return nil, errors.New("Cannot set")
}

func (t *To) ToFstream() (*Fstream, error) {
	return nil, errors.New("Not a file stream")
}

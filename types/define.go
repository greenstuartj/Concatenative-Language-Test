package types

import "errors"

type Define struct {
	name uint32
	value Type
}

func MakeDefine(name uint32, value Type) *Define {
	return &Define{ name, value }
}

func (d *Define) String() string {
	return ""
}

func (d *Define) ToNumber() (*Number, error) {
	return nil, errors.New("Not a number")
}

func (d *Define) ToString() (*String, error) {
	return nil, errors.New("Cannot make a string")
}

func (d *Define) Apply(i *Interpreter) (*Interpreter, error) {
	if d.value.Type() == FunctionT || d.value.Type() == CoreT {
		i.Env.Add(d.name, d.value)
	} else {
		prog := MakeList()
		prog = prog.Cons(d.value)
		new_i := MakeInterpreter(prog, i.Env)
		new_i, err := new_i.Run()
		if err != nil {
			return i, err
		}
		val := new_i.Stack.Car()
		i.Env.Add(d.name, val)
	}
	return i, nil
}

func (d *Define) Type() uint8 {
	return DefineT
}

func (d *Define) GetName() uint32 {
	return 0
}

func (d *Define) Unify(i *Interpreter) (*Interpreter, bool) {
	return i, false
}

func (d *Define) StrictCompare(t2 Type) int {
	if t2.Type() != DefineT {
		if DefineT < t2.Type() {
			return -1
		} else {
			return 1
		}
	} else {
		return 0
	}
}

func (d *Define) Compare(t2 Type) (int, error) {
	return 0, errors.New("Cannot compare")
}

func (d *Define) ToList() (*List, error) {
	return nil, errors.New("Cannot convert definition to list")
}

func (d *Define) ToChar() (*Char, error) {
	return nil, errors.New("Cannot make a char")
}

func (d *Define) ToBool() (*Bool, error) {
	return MakeBool(true), nil
}

func (d *Define) ToTuple() (*Tuple, error) {
	return nil, errors.New("Cannot convert to tuple")
}

func (d *Define) ToMap() (*Tree, error) {
	return nil, errors.New("Cannot convert to a dict")
}

func (d *Define) ToSymbol() (*Symbol, error) {
	return nil, errors.New("Cannot convert to symbol")
}

func (d *Define) ToSet() (*Tree, error) {
	return nil, errors.New("Cannot convert to set")
}

func (d *Define) ToCore() (*Core, error) {
	return nil, errors.New("Cannot make core")
}

func (d *Define) Get(t2 Type) (Type, error) {
	return nil, errors.New("Cannot get")
}

func (d *Define) Set(t2 Type, t3 Type) (Type, error) {
	return nil, errors.New("Cannot set")
}

func (d *Define) ToFstream() (*Fstream, error) {
	return nil, errors.New("Not a file stream")
}

package types

import (
	"errors"
)

func MakeRestVariable(s string) *RestVariable {
	n := Variable_Table.Assign(s)
	return &RestVariable{ n }
}

type RestVariable struct {
	name uint32
}

func (rv *RestVariable)	String() string {
	return "<rest_var>"
}

func (rv *RestVariable)	ToNumber() (*Number, error) {
	return nil, errors.New("Cannot convert to number")
}

func (rv *RestVariable) ToString() (*String, error) {
	return nil, errors.New("Cannot make a string")
}

func (rv *RestVariable)	Apply(i *Interpreter) (*Interpreter, error) {
	t, ok := i.Env.Lookup(rv.name)
	if ok {
		if t.Type() == ListT {
			lst, _ := t.ToList()
			new_i := MakeInterpreter(lst, i.Env)
			new_i.Stack = i.Stack
			new_i, err := new_i.Run()
			if err != nil {
				return i, err
			}
			new_i.Program = i.Program
			return new_i, nil
		}
		return t.Apply(i)
	} else {
		return i, nil
	}
}

func (rv *RestVariable)	Type() uint8 {
	return RestVariableT
}

func (rv *RestVariable)	GetName() uint32 {
	return rv.name
}

func (rv *RestVariable)	Unify(i *Interpreter) (*Interpreter, bool) {
	success := i.Env.UnifyEnv(rv.name, i.Stack)
	if success {
		i.Stack = MakeList()
		return i, true
	}
	return i, false
}

func (rv *RestVariable)	StrictCompare(t2 Type) int {
	if t2.Type() != RestVariableT {
		if RestVariableT < t2.Type() {
			return -1
		} else {
			return 1
		}
	} else {
		return 0
	}
}

func (rv *RestVariable) Compare(t2 Type) (int, error) {
	return 0, errors.New("Cannot compare")
}

func (rv *RestVariable)	ToList() (*List, error) {
	return nil, errors.New("???")
}

func (rv *RestVariable) ToChar() (*Char, error) {
	return nil, errors.New("Cannot make a char")
}

func (rv *RestVariable) ToBool() (*Bool, error) {
	return MakeBool(true), nil
}

func (rv *RestVariable) ToTuple() (*Tuple, error) {
	return nil, errors.New("Cannot convert to tuple")
}

func (rv *RestVariable) ToMap() (*Tree, error) {
	return nil, errors.New("Cannot convert to a dict")
}

func (rv *RestVariable) ToSymbol() (*Symbol, error) {
	return nil, errors.New("Cannot convert to symbol")
}

func (rv *RestVariable) ToSet() (*Tree, error) {
	return nil, errors.New("Cannot convert to set")
}

func (rv *RestVariable) ToCore() (*Core, error) {
	return nil, errors.New("Cannot make core")
}

func (rv *RestVariable) Get(t2 Type) (Type, error) {
	return nil, errors.New("Cannot get")
}

func (rv *RestVariable) Set(t2 Type, t3 Type) (Type, error) {
	return nil, errors.New("Cannot set")
}

func (rv *RestVariable) ToFstream() (*Fstream, error) {
	return nil, errors.New("Not a file stream")
}

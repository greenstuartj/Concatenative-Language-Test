package types

import (
	"math/big"
	"errors"
)

type Bool struct {
	b bool
}

func MakeBool(b bool) *Bool {
	return &Bool{b}
}

func (b *Bool) String() string {
	if b.b {
		return "True"
	} else {
		return "False"
	}
}

func (b *Bool) ToNumber() (*Number, error) {
	r := new(big.Rat)
	if b.b {
		r.SetInt64(1)
	} else {
		r.SetInt64(0)
	}
	return &Number{r}, nil
}

func (b *Bool) ToString() (*String, error) {
	return MakeString(b.String()), nil
}

func (b *Bool) Apply(i *Interpreter) (*Interpreter, error) {
	i.Stack = i.Stack.Cons(b)
	return i, nil
}

func (b *Bool) Type() uint8 {
	return BoolT
}

func (b *Bool) GetName() uint32 {
	return 0
}

func (b *Bool) Unify(i *Interpreter) (*Interpreter, bool) {
	b2 := i.Stack.Car()
	if b2.Type() != BoolT {
		return i, false
	}
	b3,_ := b2.ToBool()
	if b.b == b3.b {
		i.Stack = i.Stack.Cdr()
		return i, true
	}
	return i, false
}

func (b *Bool) StrictCompare(t2 Type) int {
	if t2.Type() != BoolT {
		if BoolT < t2.Type() {
			return -1
		} else {
			return 1
		}
	} else {
		b2,_ := t2.ToBool()
		if b.b == b2.b {
			return 0
		} else if b.b {
			return 1
		} else {
			return -1
		}
	}
}

func (b *Bool) Compare(t2 Type) (int, error) {
	b2, err := t2.ToBool()
	if err != nil {
		return 0, err
	}
	return b.StrictCompare(b2), nil
}

func (b *Bool) ToList() (*List, error) {
	lst := MakeList()
	lst.Cons(b)
	return lst, nil
}

func (b *Bool) ToChar() (*Char, error) {
	return nil, errors.New("Cannot convert to char")
}

func (b *Bool) ToBool() (*Bool, error) {
	return b, nil
}

func (b *Bool) ToTuple() (*Tuple, error) {
	return nil, errors.New("Cannot convert to tuple")
}

func (b *Bool) ToMap() (*Tree, error) {
	return nil, errors.New("Cannot convert to a dict")
}

func (b *Bool) ToSymbol() (*Symbol, error) {
	return nil, errors.New("Cannot convert to a symbol")
}

func (b *Bool) ToSet() (*Tree, error) {
	return nil, errors.New("Cannot convert to a set")
}

func (b *Bool) ToCore() (*Core, error) {
	return nil, errors.New("Cannot make core")
}

func (b *Bool) Get(t2 Type) (Type, error) {
	return nil, errors.New("Cannot get")
}

func (b *Bool) Set(t2 Type, t3 Type) (Type, error) {
	return nil, errors.New("Cannot set")
}

func (b *Bool) ToFstream() (*Fstream, error) {
	return nil, errors.New("Not a file stream")
}

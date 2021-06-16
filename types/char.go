package types

import (
	"math/big"
	"errors"
)

type Char struct {
	c byte
}

func MakeChar(c byte) *Char {
	return &Char{c}
}

func (c *Char) String() string {
	return "#\\" + string(c.c)
}

func (c *Char) ToNumber() (*Number, error) {
	r := new(big.Rat)
	r.SetInt64(int64(c.c))
	return &Number{r}, nil
}

func (c *Char) ToString() (*String, error) {
	return MakeString(string(c.c)), nil
}

func (c *Char) Apply(i *Interpreter) (*Interpreter, error) {
	i.Stack = i.Stack.Cons(c)
	return i, nil
}

func (c *Char) Type() uint8 {
	return CharT
}

func (c *Char) GetName() uint32 {
	return 0
}

func (c *Char) Unify(i *Interpreter) (*Interpreter, bool) {
	c2 := i.Stack.Car()
	if c2.Type() != CharT {
		return i, false
	}
	if c.StrictCompare(c2) == 0 {
		i.Stack = i.Stack.Cdr()
		return i, true
	}
	return i, false
}

func (c *Char) StrictCompare(t2 Type) int {
	if t2.Type() != CharT {
		if CharT < t2.Type() {
			return -1
		} else {
			return 1
		}
	} else {
		c2,_ := t2.ToChar()
		if c.c == c2.c {
			return 0
		} else if c.c < c2.c {
			return -1
		} else {
			return 1
		}
	}
}

func (c *Char) Compare(t2 Type) (int, error) {
	c2, err := t2.ToChar()
	if err != nil {
		return 0, err
	}
	return c.StrictCompare(c2), nil
}

func (c *Char) ToList() (*List, error) {
	lst := MakeList()
	lst.Cons(c)
	return lst, nil
}

func (c *Char) ToChar() (*Char, error) {
	return c, nil
}

func (c *Char) ToBool() (*Bool, error) {
	return MakeBool(true), nil
}

func (c *Char) ToTuple() (*Tuple, error) {
	return nil, errors.New("Cannot convert to tuple")
}

func (c *Char) ToMap() (*Tree, error) {
	return nil, errors.New("Cannot convert to a dict")
}

func (c *Char) ToSymbol() (*Symbol, error) {
	return nil, errors.New("Cannot convert to symbol")
}

func (c *Char) ToSet() (*Tree, error) {
	return nil, errors.New("Cannot convert to set")
}

func (c *Char) ToCore() (*Core, error) {
	return nil, errors.New("Cannot make core")
}

func (c *Char) Get(t2 Type) (Type, error) {
	return nil, errors.New("Cannot get")
}

func (c *Char) Set(t2 Type, t3 Type) (Type, error) {
	return nil, errors.New("Cannot set")
}

func (c *Char) ToFstream() (*Fstream, error) {
	return nil, errors.New("Not a file stream")
}

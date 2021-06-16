package types

import (
	"math/big"
	"errors"
)

type String struct {
	s string
}

func MakeString(s string) *String {
	return &String{s}
}

func (s *String) String() string {
	return s.s
}

func (s *String) StringWithQuotes() string {
	return "\"" + s.s + "\""
}

func (s *String) ToNumber() (*Number, error) {
	r := new(big.Rat)
	_, success := r.SetString(s.s)
	if !success {
		return nil, errors.New("Cannot convert to number")
	}
	return &Number{r}, nil
}

func (s *String) ToString() (*String, error) {
	return s, nil
}

func (s *String) Apply(i *Interpreter) (*Interpreter, error) {
	i.Stack = i.Stack.Cons(s)
	return i, nil
}

func (s *String) Type() uint8 {
	return StringT
}

func (s *String) GetName() uint32 {
	return 0
}

func (s *String) Unify(i *Interpreter) (*Interpreter, bool) {
	s2 := i.Stack.Car()
	if s.Type() != StringT {
		return i, false
	}
	s3 := s2.String()
	if s.s == s3 {
		i.Stack = i.Stack.Cdr()
		return i, true
	}
	return i, false
}

func (s *String) StrictCompare(t2 Type) int {
	if t2.Type() != StringT {
		if StringT < t2.Type() {
			return -1
		} else {
			return 1
		}
	} else {
		if s.s == t2.String() {
			return 0
		} else if s.s < t2.String() {
			return -1
		} else {
			return 1
		}
	}
}

func (s *String) Compare(t2 Type) (int, error) {
	s2, err := t2.ToString()
	if err != nil {
		return 0, err
	}
	return s.StrictCompare(s2), nil
}

func (s *String) ToList() (*List, error) {
	lst := MakeList()
	lst.Cons(s)
	return lst, nil
}

func (s *String) ToChar() (*Char, error) {
	if len(s.s) != 1 {
		return nil, errors.New("Cannot make a char")
	} else {
		return MakeChar(s.s[0]), nil
	}	
}

func (s *String) ToBool() (*Bool, error) {
	return MakeBool(true), nil
}

func (s *String) ToTuple() (*Tuple, error) {
	return nil, errors.New("Cannot convert to tuple")
}

func (s *String) ToMap() (*Tree, error) {
	return nil, errors.New("Cannot convert to a dict")
}

func (s *String) ToSymbol() (*Symbol, error) {
	return nil, errors.New("Cannot convert to symbol")
}

func (s *String) ToSet() (*Tree, error) {
	return nil, errors.New("Cannot convert to set")
}

func (s *String) ToCore() (*Core, error) {
	return nil, errors.New("Cannot make core")
}

func (s *String) Get(t2 Type) (Type, error) {
	return nil, errors.New("Cannot get")
}

func (s *String) Set(t2 Type, t3 Type) (Type, error) {
	return nil, errors.New("Cannot set")
}

func (s *String) ToFstream() (*Fstream, error) {
	return nil, errors.New("Not a file stream")
}

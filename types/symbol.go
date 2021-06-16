package types

import "errors"

type SymbolTable struct {
	SymbolToInt map[string]uint32
	IntToSymbol []string
}

func (st *SymbolTable) Assign(name string) uint32 {
	i, ok := st.SymbolToInt[name]
	if ok {
		return i
	} else {
		current := uint32(len(st.IntToSymbol))
		st.SymbolToInt[name] = current
		st.IntToSymbol = append(st.IntToSymbol, name)
		return current
	}
}

var Symbol_Table = &SymbolTable{
	make(map[string]uint32),
	[]string{},
}


type Symbol struct {
	s uint32
}

func MakeSymbol(s string) *Symbol {
	n := Symbol_Table.Assign(s)
	return &Symbol{n}
}

func (s *Symbol) String() string {
	return Symbol_Table.IntToSymbol[s.s]
}

func (s *Symbol) ToNumber() (*Number, error) {
	return nil, errors.New("Cannot convert to number")
}

func (s *Symbol) ToString() (*String, error) {
	// bar symbols |Symbol Name| to be handled
	return MakeString(Symbol_Table.IntToSymbol[s.s]), nil
}

func (s *Symbol) Apply(i *Interpreter) (*Interpreter, error) {
	i.Stack = i.Stack.Cons(s)
	return i, nil
}

func (s *Symbol) Type() uint8 {
	return SymbolT
}

func (s *Symbol) GetName() uint32 {
	return 0
}

func (s *Symbol) Unify(i *Interpreter) (*Interpreter, bool) {
	sym2 := i.Stack.Car()
	if sym2.Type() != SymbolT {
		return i, false
	}
	if s.GetName() == sym2.GetName() {
		i.Stack = i.Stack.Cdr()
		return i, true
	}
	return i, false
}

func (s *Symbol) StrictCompare(t2 Type) int {
	if t2.Type() != SymbolT {
		if SymbolT < t2.Type() {
			return -1
		} else {
			return 1
		}
	} else {
		s2,_ := t2.ToSymbol()
		if s.s == s2.s {
			return 0
		} else if s.s < s2.s {
			return -1
		} else {
			return 1
		}
	}
}

func (s *Symbol) Compare(t2 Type) (int, error) {
	return 0, errors.New("Cannot compare")
}

func (s *Symbol) ToList() (*List, error) {
	lst := MakeList()
	lst.Cons(s)
	return lst, nil
}

func (s *Symbol) ToChar() (*Char, error) {
	return nil, errors.New("Cannot make a char")
}

func (s *Symbol) ToBool() (*Bool, error) {
	return MakeBool(true), nil
}

func (s *Symbol) ToTuple() (*Tuple, error) {
	return nil, errors.New("Cannot convert to tuple")
}

func (s *Symbol) ToMap() (*Tree, error) {
	return nil, errors.New("Cannot convert to a dict")
}

func (s *Symbol) ToSymbol() (*Symbol, error) {
	return s, nil
}

func (s *Symbol) ToSet() (*Tree, error) {
	return nil, errors.New("Cannot convert to set")
}

func (s *Symbol) ToCore() (*Core, error) {
	return nil, errors.New("Cannot make core")
}

func (s *Symbol) Get(t2 Type) (Type, error) {
	return nil, errors.New("Cannot get")
}

func (s *Symbol) Set(t2 Type, t3 Type) (Type, error) {
	return nil, errors.New("Cannot set")
}

func (s *Symbol) ToFstream() (*Fstream, error) {
	return nil, errors.New("Not a file stream")
}

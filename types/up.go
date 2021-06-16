package types

import "errors"

type Up struct {
	u Type
}

func MakeUp(t Type) *Up {
	return &Up{t}
}

func (u *Up) String() string {
	return "^" + u.u.String()
}

// Needs work -- Default should have function applied to it after insert?
func (u *Up) Apply(i *Interpreter) (*Interpreter, error) {
	k := i.Stack.Car()
	m := i.Stack.Cdr().Car()
	if m.Type() != MapT {
		return i, errors.New("Not a map")
	}
	m2,_ := m.ToMap()
	v,s := m2.Lookup(k)
	if !s && u.u.Type() != TupleT {
		return i, errors.New("Key not found")
	} else if !s {
		tup,err := u.u.ToTuple()
		if err != nil {
			return i, err
		}
		def := tup.t.Car()
		i.Stack = i.Stack.Cdr().Cdr().Cons(m2.Insert(k, def))
		return i, nil
	}
	temp_i := MakeInterpreter(MakeList(), i.Env)
	temp_i.Stack = temp_i.Stack.Cons(v)
	if u.u.Type() == FunctionT {
		temp_i,err := u.u.Apply(temp_i)
		if err != nil {
			return i, err
		}
		i.Stack = i.Stack.Cdr().Cdr().Cons(m2.Insert(k, temp_i.Stack.Car()))
		return i, nil
	} else if u.u.Type() == TupleT {
		tup,_ := u.u.ToTuple()
		f := tup.t.Cdr().Car()
		temp_i,err := f.Apply(temp_i)
		if err != nil {
			return i, err
		}
		a,_ := CoreTable.IsCore("apply")
		temp_i,err = CoreTable.Lookup(a)(temp_i)
		if err != nil {
			return i, err
		}
		i.Stack = i.Stack.Cdr().Cdr().Cons(m2.Insert(k, temp_i.Stack.Car()))
		return i, nil
	} else {
		return i, errors.New("Not function or tuple")
	}
}

func (u *Up) Type() uint8 {
	return UpT
}

func (u *Up) GetName() uint32 {
	return 0
}

func (u *Up) Unify(i *Interpreter) (*Interpreter, bool) {
	return i, false
}

func (u *Up) StrictCompare(t2 Type) int {
	if t2.Type() != UpT {
		if UpT < t2.Type() {
			return -1
		} else {
			return 1
		}
	}
	return 0
}

func (u *Up) Compare(t2 Type) (int, error) {
	return -1, errors.New("Cannot compare")
}

func (u *Up) ToNumber() (*Number, error) {
	return nil, errors.New("Cannot make number")
}

func (u *Up) ToString() (*String, error) {
	return nil, errors.New("Cannot make string")
}

func (u *Up) ToList() (*List, error) {
	return nil, errors.New("Cannot make list")
}

func (u *Up) ToChar() (*Char, error) {
	return nil, errors.New("Cannot make char")
}

func (u *Up) ToBool() (*Bool, error) {
	return MakeBool(true), nil
}

func (u *Up) ToTuple() (*Tuple, error) {
	return nil, errors.New("Cannot make tuple")
}

func (u *Up) ToMap() (*Tree, error) {
	return nil, errors.New("Cannot make map")
}

func (u *Up) ToSymbol() (*Symbol, error) {
	return nil, errors.New("Cannot make symbol")
}

func (u *Up) ToSet() (*Tree, error) {
	return nil, errors.New("Cannot make set")
}

func (u *Up) ToCore() (*Core, error) {
	return nil, errors.New("Cannot make core")
}

func (u *Up) Get(t2 Type) (Type, error) {
	return nil, errors.New("Cannot get")
}

func (u *Up) Set(t2 Type, t3 Type) (Type, error) {
	return nil, errors.New("Cannot set")
}

func (u *Up) ToFstream() (*Fstream, error) {
	return nil, errors.New("Not a file stream")
}

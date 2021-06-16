package types

import (
	"errors"
)

type Tuple struct {
	t *List
}

func MakeTuple(t *List) *Tuple {
	return &Tuple{t}
}

func (t *Tuple)	String() string {
	var s string = "("
	temp := t.t
	first := true
	for !temp.Nullp() {
		if !first {
			s += " "
		}
		first = false
		s += temp.Element.String()
		temp = temp.Next
	}
	s += ")"
	return s
}

func (t *Tuple)	ToNumber() (*Number, error) {
	return nil, errors.New("Cannot convert to a number")
}

func (t *Tuple)	ToString() (*String, error) {
	return MakeString(t.String()), nil
}

func (t *Tuple)	Apply(i *Interpreter) (*Interpreter, error) {
	new_i := MakeInterpreter(t.t, i.Env)
	new_i, err := new_i.Run()
	if err != nil {
		return nil, err
	}
	temp_lst := new_i.Stack
	tuple_lst := MakeList()
	for !temp_lst.Nullp() {
		tuple_lst = tuple_lst.Cons(temp_lst.Element)
		temp_lst = temp_lst.Next
	}
	i.Stack = i.Stack.Cons(MakeTuple(tuple_lst))
	return i, nil
}

func (t *Tuple)	Type() uint8 {
	return TupleT
}

func (t *Tuple)	GetName() uint32 {
	return 0
}

// may need double checking
func (t *Tuple)	Unify(i *Interpreter) (*Interpreter, bool) {
	if t.t.Nullp() {
		return i, true
	}
	if i.Stack.Car().Type() != TupleT {
		return i, false
	}
	car_stack, err := i.Stack.Car().ToList()
	if err != nil {
		return i, false
	}
	temp_i := MakeInterpreter(MakeList(), i.Env)
	temp_i.Stack = car_stack
	temp_args := t.t
	temp_i, unified := temp_args.Element.Unify(temp_i)
	temp_args = temp_args.Next
	for unified && !temp_args.Nullp() {
		temp_i, unified = temp_args.Element.Unify(temp_i)
		temp_args = temp_args.Next
	}
	i.Stack = i.Stack.Cdr()
	i.Env = temp_i.Env
	return i, unified
}

func (t *Tuple)	StrictCompare(t2 Type) int {
	if t2.Type() != TupleT {
		if TupleT < t2.Type() {
			return -1
		} else {
			return 1
		}
	} else {
		lst1 := t.t
		lst2,_ := t2.ToTuple()
		lst3 := lst2.t
		for !lst1.Nullp() && !lst3.Nullp() && lst1.Element.StrictCompare(lst3.Element) == 0 {
			lst1 = lst1.Next
			lst3 = lst3.Next
		}
		if lst1.Nullp() && !lst3.Nullp() {
			return -1
		} else if !lst1.Nullp() && lst3.Nullp() {
			return 1
		} else {
			return lst1.Element.StrictCompare(lst3.Element)
		}
	}
}

func (t *Tuple) Compare(t2 Type) (int, error) {
	tu2, err := t2.ToTuple()
	if err != nil {
		return 0, err
	}
	return t.StrictCompare(tu2), nil
}

func (t *Tuple)	ToList() (*List, error) {
	return t.t, nil
}

func (t *Tuple)	ToChar() (*Char, error) {
	return nil, errors.New("Cannot convert to a char")
}

func (t *Tuple)	ToBool() (*Bool, error) {
	return MakeBool(true), nil
}

func (t *Tuple) ToTuple() (*Tuple, error) {
	return t, nil
}

func (t *Tuple) ToMap() (*Tree, error) {
	return nil, errors.New("Cannot convert to a dict")
}

func (t *Tuple) ToSymbol() (*Symbol, error) {
	return nil, errors.New("Cannot convert to symbol")
}

func (t *Tuple) ToSet() (*Tree, error) {
	return nil, errors.New("Cannot convert to a set")
}

func (t *Tuple) ToCore() (*Core, error) {
	return nil, errors.New("Cannot make core")
}

func (t *Tuple) Get(t2 Type) (Type, error) {
	return nil, errors.New("Cannot get")
}

func (t *Tuple) Set(t2 Type, t3 Type) (Type, error) {
	return nil, errors.New("Cannot set")
}

func (t *Tuple) ToFstream() (*Fstream, error) {
	return nil, errors.New("Not a file stream")
}

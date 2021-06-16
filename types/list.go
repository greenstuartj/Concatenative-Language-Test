package types

// need to handle eval of tuples, dicts and sets in list

import (
	"errors"
)

type List struct {
	Element Type
	Next *List
}

func MakeList() *List {
	return &List{ nil, nil }
}

func (lst *List) Nullp() bool {
	return lst.Element == nil && lst.Next == nil
}

func (lst *List) Cons(e Type) *List {
	return &List{ e, lst }
}

func (lst *List) Car() Type {
	return lst.Element
}

func (lst *List) Cdr() *List {
	return lst.Next
}

func (lst *List) String() string {
	var s string = "["
	temp := lst
	for !temp.Nullp() {
		if temp.Element.Type() == StringT {
			temp_s,_ := temp.Element.ToString()
			s += " " + temp_s.StringWithQuotes()
		} else {
			s += " " + temp.Element.String()
		}
		temp = temp.Next
	}
	s += " ]"
	return s
}

func (lst *List) ToNumber() (*Number, error) {
	return nil, errors.New("Not a number")
}

func (lst *List) ToString() (*String, error) {
	return nil, errors.New("Cannot make a string")
}

func (lst *List) Apply(i *Interpreter) (*Interpreter, error) {
	i.Stack = i.Stack.Cons(lst)
	return i, nil
	// dead code
	lst2 := MakeList()
	temp_lst := lst
	for !temp_lst.Nullp() {
		if temp_lst.Element.Type() == VariableT {
			value, found := i.Env.Lookup(temp_lst.Element.GetName())
			if found {
				lst2 = lst2.Cons(value)
			} else {
				return i, errors.New("Variable not defined")
			}
		} else if temp_lst.Element.Type() == RestVariableT {
			value, found := i.Env.Lookup(temp_lst.Element.GetName())
			if found {
				lst3,_ := value.ToList()
				for !lst3.Nullp() {
					lst2 = lst2.Cons(lst3.Element)
					lst3 = lst3.Next
				}
			} else {
				return i, errors.New("Variable not defined")
			}
		} else {
			lst2 = lst2.Cons(temp_lst.Element)
		}
		temp_lst = temp_lst.Next
	}
	out_list := MakeList()
	for !lst2.Nullp() {
		out_list = out_list.Cons(lst2.Element)
		lst2 = lst2.Next
	}
	i.Stack = i.Stack.Cons(out_list)
	return i, nil
}

func (lst *List) Type() uint8 {
	return ListT
}

func (lst *List) GetName() uint32 {
	return 0
}

func (lst *List) Unify(i *Interpreter) (*Interpreter, bool) {
	return i, false
	// dead code
	if lst.Nullp() {
		return i, true
	}
	car_stack, err := i.Stack.Car().ToList()
	if err != nil {
		return i, false
	}
	temp_i := MakeInterpreter(MakeList(), i.Env)
	temp_i.Stack = car_stack
	temp_args := lst
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

func (lst *List) StrictCompare(t2 Type) int {
	if t2.Type() == ListT {
		lst1 := lst
		lst2,_ := t2.ToList()
		for !lst1.Nullp() && !lst2.Nullp() && lst1.Element.StrictCompare(lst2.Element) == 0 {
			lst1 = lst1.Next
			lst2 = lst2.Next
		}
		if lst1.Nullp() && !lst2.Nullp() {
			return -1
		} else if !lst1.Nullp() && lst2.Nullp() {
			return 1
		} else {
			return lst1.Element.StrictCompare(lst2.Element)
		}
	} else {
		return -1
	}
}

func (lst *List) Compare(t2 Type) (int, error) {
	lst2, err := t2.ToList()
	if err != nil {
		return 0, err
	}
	return lst.StrictCompare(lst2), nil
}

func (lst *List) ToList() (*List, error) {
	return lst, nil
}

func (lst *List) ToChar() (*Char, error) {
	return nil, errors.New("Cannot make a char")
}

func (lst *List) ToBool() (*Bool, error) {
	return MakeBool(true), nil
}

func (lst *List) ToTuple() (*Tuple, error) {
	return MakeTuple(lst), nil
}

func (lst *List) ToMap() (*Tree, error) {
	return nil, errors.New("Cannot convert to a dict")
}


func (lst *List) ToSymbol() (*Symbol, error) {
	return nil, errors.New("Cannot convert to symbol")
}

func (lst *List) ToSet() (*Tree, error) {
	return nil, errors.New("Cannot convert to set")
}

func (lst *List) ToCore() (*Core, error) {
	return nil, errors.New("Cannot make core")
}

func (lst *List) Get(t2 Type) (Type, error) {
	return nil, errors.New("Cannot get")
}

func (lst *List) Set(t2 Type, t3 Type) (Type, error) {
	return nil, errors.New("Cannot set")
}

func (lst *List) ToFstream() (*Fstream, error) {
	return nil, errors.New("Not a file stream")
}

type ListLiteral struct {
	lst *List
}

func MakeListLiteral(lst *List) *ListLiteral {
	return &ListLiteral{ lst }
}

func (lst *ListLiteral) String() string {
	var s string = "["
	temp := lst.lst
	for !temp.Nullp() {
		if temp.Element.Type() == StringT {
			temp_s,_ := temp.Element.ToString()
			s += " " + temp_s.StringWithQuotes()
		} else {
			s += " " + temp.Element.String()
		}
		temp = temp.Next
	}
	s += " ]"
	return s
}

func (lst *ListLiteral) ToNumber() (*Number, error) {
	return nil, errors.New("Not a number")
}

func (lst *ListLiteral) ToString() (*String, error) {
	return nil, errors.New("Cannot make a string")
}

func (lst *ListLiteral) Apply(i *Interpreter) (*Interpreter, error) {
	lst2 := MakeList()
	temp_lst := lst.lst
	for !temp_lst.Nullp() {
		if temp_lst.Element.Type() == VariableT {
			value, found := i.Env.Lookup(temp_lst.Element.GetName())
			if found {
				lst2 = lst2.Cons(value)
			} else {
				return i, errors.New("Variable not defined")
			}
		} else if temp_lst.Element.Type() == RestVariableT {
			value, found := i.Env.Lookup(temp_lst.Element.GetName())
			if found {
				lst3,_ := value.ToList()
				for !lst3.Nullp() {
					lst2 = lst2.Cons(lst3.Element)
					lst3 = lst3.Next
				}
			} else {
				return i, errors.New("Variable not defined")
			}
		} else {
			lst2 = lst2.Cons(temp_lst.Element)
		}
		temp_lst = temp_lst.Next
	}
	out_list := MakeList()
	for !lst2.Nullp() {
		out_list = out_list.Cons(lst2.Element)
		lst2 = lst2.Next
	}
	i.Stack = i.Stack.Cons(out_list)
	return i, nil
}

func (lst *ListLiteral) Type() uint8 {
	return ListLiteralT
}

func (lst *ListLiteral) GetName() uint32 {
	return 0
}

func (lst *ListLiteral) Unify(i *Interpreter) (*Interpreter, bool) {
	if lst.lst.Nullp() {
		return i, true
	}
	car_stack, err := i.Stack.Car().ToList()
	if err != nil {
		return i, false
	}
	temp_i := MakeInterpreter(MakeList(), i.Env)
	temp_i.Stack = car_stack
	temp_args := lst.lst
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

func (lst *ListLiteral) StrictCompare(t2 Type) int {
	return -1
}

func (lst *ListLiteral) Compare(t2 Type) (int, error) {
	return -1, errors.New("Cannot compare list literal")
}

func (lst *ListLiteral) ToList() (*List, error) {
	return lst.lst, nil
}

func (lst *ListLiteral) ToChar() (*Char, error) {
	return nil, errors.New("Cannot make a char")
}

func (lst *ListLiteral) ToBool() (*Bool, error) {
	return MakeBool(true), nil
}

func (lst *ListLiteral) ToTuple() (*Tuple, error) {
	return MakeTuple(lst.lst), nil
}

func (lst *ListLiteral) ToMap() (*Tree, error) {
	return nil, errors.New("Cannot convert to a dict")
}

func (lst *ListLiteral) ToSymbol() (*Symbol, error) {
	return nil, errors.New("Cannot convert to symbol")
}

func (lst *ListLiteral) ToSet() (*Tree, error) {
	return nil, errors.New("Cannot convert to set")
}

func (lst *ListLiteral) ToCore() (*Core, error) {
	return nil, errors.New("Cannot make core")
}

func (lst *ListLiteral) Get(t2 Type) (Type, error) {
	return nil, errors.New("Cannot get")
}

func (lst *ListLiteral) Set(t2 Type, t3 Type) (Type, error) {
	return nil, errors.New("Cannot set")
}

func (lst *ListLiteral) ToFstream() (*Fstream, error) {
	return nil, errors.New("Not a file stream")
}

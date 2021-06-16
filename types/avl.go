package types

import (
	"errors"
)

// delete method

type Tree struct {
	Key Type
	Value Type
	Left *Tree
	Right *Tree
	Height int
	CompareF *List
	TypeT uint8
}

func (t *Tree) get_height() int {
	if t == nil {
		return -1
	} else {
		return t.Height
	}
}

func (t *Tree) new_height() int {
	if t.Left.get_height() > t.Right.get_height() {
		return t.Left.get_height() + 1
	} else {
		return t.Right.get_height() + 1
	}
}

func (t *Tree) rotate_right() *Tree {
	var new_right *Tree
	if t.Left == nil {
		new_right = &Tree{ t.Key, t.Value, t.Left, t.Right, 0, t.CompareF, t.TypeT }
	} else {
		new_right = &Tree{ t.Key, t.Value, t.Left.Right, t.Right, 0, t.CompareF, t.TypeT }
	}
	new_right.Height = new_right.new_height()

	var new_root *Tree
	new_root = &Tree{
		t.Left.Key,
		t.Left.Value,
		t.Left.Left,
		new_right,
		0, t.Left.CompareF, t.Left.TypeT }
	new_root.Height = new_root.new_height()
	return new_root
}

func (t *Tree) rotate_left() *Tree {
	var new_left *Tree
	if t.Right == nil {
		new_left = &Tree{ t.Key, t.Value, t.Left, t.Right, 0, t.CompareF, t.TypeT }
	} else {
		new_left = &Tree{ t.Key, t.Value, t.Left, t.Right.Left, 0, t.CompareF, t.TypeT }
	}
	new_left.Height = new_left.new_height()

	var new_root *Tree
	new_root = &Tree{
		t.Right.Key,
		t.Right.Value,
		new_left,
		t.Right.Right,
		0, t.Right.CompareF, t.Right.TypeT }
	new_root.Height = new_root.new_height()
	return new_root
}

func (t *Tree) get_balance() int {
	if t == nil {
		return 0
	}
	if t.Right == nil && t.Left == nil {
		return 0
	}
	if t.Right == nil {
		return -1 -t.Left.get_height()
	}
	if t.Left == nil {
		return t.Right.get_height() - -1
		
	}
	return t.Right.get_height() - t.Left.get_height()
}

func (t *Tree) balance() *Tree {
	if t.get_balance() < 2 && t.get_balance() > -2 {
		return t
	} else if t.get_balance() <= -2 {
		if t.Left.get_balance() <= 0 {
			return t.rotate_right()
		} else {
			new_sub := t.Left.rotate_left()
			new_root := &Tree{t.Key,t.Value,new_sub,t.Right,0,t.CompareF,t.TypeT}
			new_root.Height = new_root.new_height()
			return new_root.rotate_right()
		}
	} else {
		if t.Right.get_balance() >= 0 {
			return t.rotate_left()
		} else {
			new_sub := t.Right.rotate_right()
			new_root := &Tree{t.Key,t.Value,t.Left,new_sub,0,t.CompareF,t.TypeT}
			new_root.Height = new_root.new_height()
			return new_root.rotate_left()
		}
	}
}

func (t *Tree) Insert(key Type, value Type) *Tree {
	if t == nil || t.Key == nil {
		return &Tree{ key, value, nil, nil, 0, t.CompareF, t.TypeT }
	}
	var comp int
	if t.CompareF == nil {
		comp = key.StrictCompare(t.Key)
	} else {
		comp = 0
	}
	switch (comp) {
	case 0:
		return &Tree{ key, value, t.Left, t.Right, t.Height, t.CompareF, t.TypeT }
	case -1:
		var new_left *Tree
		if t.Left == nil {
			new_left = &Tree{ key, value, nil, nil, 0, t.CompareF, t.TypeT }
		} else {
			new_left = t.Left.Insert(key, value).balance()
		}
		new_t := &Tree{t.Key, t.Value, new_left, t.Right, t.Height, t.CompareF, t.TypeT}
		new_t.Height = new_t.new_height()
		return new_t.balance()
	case 1:
		var new_right *Tree
		if t.Right == nil {
			new_right = &Tree{ key, value, nil, nil, 0, t.CompareF, t.TypeT }
		} else {
			new_right = t.Right.Insert(key, value).balance()
		}
		new_t := &Tree{t.Key, t.Value, t.Left, new_right, t.Height, t.CompareF, t.TypeT}
		new_t.Height = new_t.new_height()
		return new_t.balance()
	default:
		return t
	}
}

func (t *Tree) Lookup(key Type) (Type, bool) {
	if t == nil || t.Key == nil {
		return nil, false
	}
	if t.CompareF == nil {
		switch (key.StrictCompare(t.Key)) {
		case 0:
			return t.Value, true
		case -1:
			return t.Left.Lookup(key)
		case 1:
			return t.Right.Lookup(key)
		default:
			return nil, false
		}
	} else {
		return nil, false
	}
}

func (t *Tree) String() string {
	if t.Key == nil {
		if t.TypeT == MapT {
			return "#[ ]"
		} else if t.TypeT == SetT {
			return "~[ ]"
		}
	}
	var s string
	if t.TypeT == MapT {
		s = "#[ "
	} else if t.TypeT == SetT {
		s = "~[ "
	}
	s += t.StringAux()
	s += " ]"
	return s
}

func (t *Tree) StringAux() string {
	if t.Key == nil {
		return ""
	}
	var s string
	if t.Value == nil {
		s = t.Key.String()
	} else {
		s = "(" + t.Key.String() + " " + t.Value.String() + ")"
	}
	if t.Left != nil {
		s += " " + t.Left.StringAux()
	}
	if t.Right != nil {
		s += " " + t.Right.StringAux()
	}
	return s
}

func (m *Tree) Apply(i *Interpreter) (*Interpreter, error) {
	i.Stack = i.Stack.Cons(m)
	return i, nil
}

func (m *Tree) Type() uint8 {
	return m.TypeT
}

func (m *Tree) GetName() uint32 {
	return 0
}

func (m *Tree) Unify(i *Interpreter) (*Interpreter, bool) {
	return i, false
}

func (m *Tree) StrictCompare(t2 Type) int {
	return -1
}

func (m *Tree) Compare(t2 Type) (int, error) {
	return -1, errors.New("Not done")
}

func (m *Tree) ToNumber() (*Number, error) {
	return nil, errors.New("Cannot make a number")
}

func (m *Tree) ToString() (*String, error) {
	return nil, errors.New("Cannot make a string")
}

func (m *Tree) ToList() (*List, error) {
	return nil, errors.New("Cannot make a list")
}

func (m *Tree) ToChar() (*Char, error) {
	return nil, errors.New("Cannot make a char")
}

func (m *Tree) ToBool() (*Bool, error) {
	return MakeBool(true), nil
}

func (m *Tree) ToTuple() (*Tuple, error) {
	return nil, errors.New("Cannot make a tuple")
}

func (m *Tree) ToMap() (*Tree, error) {
	if m.TypeT == MapT {
		return m, nil
	}
	return nil, errors.New("Cannot convert to map")
}

func (m *Tree) ToSymbol() (*Symbol, error) {
	return nil, errors.New("Cannot convert to symbol")
}

func (m *Tree) ToSet() (*Tree, error) {
	if m.TypeT == SetT {
		return m, nil
	}
	return nil, errors.New("Cannot convert to set")
}

func (m *Tree) ToCore() (*Core, error) {
	return nil, errors.New("Cannot make core")
}

func (m *Tree) Get(t2 Type) (Type, error) {
	if m.TypeT == MapT {
		r,_ := m.Lookup(t2)
		return r, nil
	} else {
		_,s := m.Lookup(t2)
		return MakeBool(s), nil
	}
}

func (m *Tree) Set(t2 Type, t3 Type) (Type, error) {
	return m.Insert(t2, t3), nil
}

func (m *Tree) ToFstream() (*Fstream, error) {
	return nil, errors.New("Not a file stream")
}

type MapLiteral struct {
	m *List
}

func MakeMapLiteral(m *List) *MapLiteral {
	return &MapLiteral{ m }
}

func (m *MapLiteral) String() string {
	return "#" + m.m.String()
}

func (m *MapLiteral) Apply(i *Interpreter) (*Interpreter, error) {
	temp := MakeInterpreter(m.m, i.Env)
	temp, err := temp.Run()
	if err != nil {
		return i, err
	}
	lst := temp.Stack
	t := &Tree{ nil, nil, nil, nil, 0, nil, MapT }
	for !lst.Nullp() {
		if lst.Element.Type() != TupleT {
			return i, errors.New("Map must be made using tuple")
		}
		tup,_ := lst.Element.ToTuple()
		t = t.Insert(tup.t.Car(), tup.t.Cdr().Car())
		lst = lst.Next
	}
	i.Stack = i.Stack.Cons(t)
	return i, nil
}

func (m *MapLiteral) Type() uint8 {
	return MapLiteralT
}

func (m *MapLiteral) GetName() uint32 {
	return 0
}

func (m *MapLiteral) Unify(i *Interpreter) (*Interpreter, bool) {
	mt := i.Stack.Car()
	if mt.Type() != MapT {
		return i, false
	}
	mt2,_ := mt.ToMap()
	lst := m.m
	// can be implemented in terms of other Unify methods? Can get functions for free
	for !lst.Nullp() {
		tup,_ := lst.Element.ToList()
		key := tup.Car()
		val := tup.Cdr().Car()
		result, success := mt2.Lookup(key)
		c := val.StrictCompare(result)
		if !success {
			return i, false
		} else if val.Type() == VariableT {
			s := i.Env.UnifyEnv(val.GetName(), result)
			if !s {
				return i, false
			}
		} else if val.Type() != VariableT && c != 0 {
			return i, false
		} else if c != 0 {
			return i, false
		}
		lst = lst.Next
	}
	i.Stack = i.Stack.Cdr()
	return i, true
}

func (m *MapLiteral) StrictCompare(t2 Type) int {
	return -1
}

func (m *MapLiteral) Compare(t2 Type) (int, error) {
	return -1, errors.New("Cannot comapre map literal")
}

func (m *MapLiteral) ToNumber() (*Number, error) {
	return nil, errors.New("Cannot make a number")
}

func (m *MapLiteral) ToString() (*String, error) {
	return nil, errors.New("Cannot make a string")
}

func (m *MapLiteral) ToList() (*List, error) {
	return nil, errors.New("Cannot make a list")
}

func (m *MapLiteral) ToChar() (*Char, error) {
	return nil, errors.New("Cannot make a char")
}

func (m *MapLiteral) ToBool() (*Bool, error) {
	return MakeBool(true), nil
}

func (m *MapLiteral) ToTuple() (*Tuple, error) {
	return nil, errors.New("Cannot make a tuple")
}

func (m *MapLiteral) ToMap() (*Tree, error) {
	return nil, errors.New("Cannot make map")
}

func (m *MapLiteral) ToSymbol() (*Symbol, error) {
	return nil, errors.New("Cannot convert to symbol")
}

func (m *MapLiteral) ToSet() (*Tree, error) {
	return nil, errors.New("Cannot convert to set")
}

func (m *MapLiteral) ToCore() (*Core, error) {
	return nil, errors.New("Cannot make core")
}

func (m *MapLiteral) Get(t2 Type) (Type, error) {
	return nil, errors.New("Cannot get")
}

func (m *MapLiteral) Set(t2 Type, t3 Type) (Type, error) {
	return nil, errors.New("Cannot set")
}

func (m *MapLiteral) ToFstream() (*Fstream, error) {
	return nil, errors.New("Not a file stream")
}

type SetLiteral struct {
	m *List
}

func MakeSetLiteral(m *List) *SetLiteral {
	return &SetLiteral{ m }
}

func (m *SetLiteral) String() string {
	return "~" + m.m.String()
}

func (m *SetLiteral) Apply(i *Interpreter) (*Interpreter, error) {
	temp := MakeInterpreter(m.m, i.Env)
	temp, err := temp.Run()
	if err != nil {
		return i, err
	}
	lst := temp.Stack
	t := &Tree{ nil, nil, nil, nil, 0, nil, SetT }
	for !lst.Nullp() {
		t = t.Insert(lst.Element, nil)
		lst = lst.Next
	}
	i.Stack = i.Stack.Cons(t)
	return i, nil
}

func (m *SetLiteral) Type() uint8 {
	return SetLiteralT
}

func (m *SetLiteral) GetName() uint32 {
	return 0
}

func (m *SetLiteral) Unify(i *Interpreter) (*Interpreter, bool) {
	mt := i.Stack.Car()
	if mt.Type() != SetT {
		return i, false
	}
	mt2,_ := mt.ToSet()
	lst := m.m
	// can be implemented in terms of other Unify methods? Can get functions for free
	for !lst.Nullp() {
		_, success := mt2.Lookup(lst.Element)
		if !success {
			return i, false
		}
		lst = lst.Next
	}
	i.Stack = i.Stack.Cdr()
	return i, true
}

func (m *SetLiteral) StrictCompare(t2 Type) int {
	return -1
}

func (m *SetLiteral) Compare(t2 Type) (int, error) {
	return -1, errors.New("Cannot comapre map literal")
}

func (m *SetLiteral) ToNumber() (*Number, error) {
	return nil, errors.New("Cannot make a number")
}

func (m *SetLiteral) ToString() (*String, error) {
	return nil, errors.New("Cannot make a string")
}

func (m *SetLiteral) ToList() (*List, error) {
	return nil, errors.New("Cannot make a list")
}

func (m *SetLiteral) ToChar() (*Char, error) {
	return nil, errors.New("Cannot make a char")
}

func (m *SetLiteral) ToBool() (*Bool, error) {
	return MakeBool(true), nil
}

func (m *SetLiteral) ToTuple() (*Tuple, error) {
	return nil, errors.New("Cannot make a tuple")
}

func (m *SetLiteral) ToMap() (*Tree, error) {
	return nil, errors.New("Cannot make map")
}

func (m *SetLiteral) ToSymbol() (*Symbol, error) {
	return nil, errors.New("Cannot convert to symbol")
}

func (m *SetLiteral) ToSet() (*Tree, error) {
	return nil, errors.New("Cannot convert to set")
}

func (m *SetLiteral) ToCore() (*Core, error) {
	return nil, errors.New("Cannot make core")
}

func (m *SetLiteral) Get(t2 Type) (Type, error) {
	return nil, errors.New("Cannot get")
}

func (m *SetLiteral) Set(t2 Type, t3 Type) (Type, error) {
	return nil, errors.New("Cannot set")
}

func (m *SetLiteral) ToFstream() (*Fstream, error) {
	return nil, errors.New("Not a file stream")
}

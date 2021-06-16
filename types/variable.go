package types

import "errors"

type Variable struct {
	v uint32
}

func MakeVariable(s string) *Variable {
	n := Variable_Table.Assign(s)
	return &Variable{n}
}

func (v *Variable) Apply(i *Interpreter) (*Interpreter, error) {
	t, ok := i.Env.Lookup(v.v)
	if ok {
		return t.Apply(i)
	} else {
		return i, errors.New("undefined")
	}
}

func (v *Variable) String() string {
	name,_ := Variable_Table.ints[v.v]
	return name
}

func (v *Variable) ToNumber() (*Number, error) {
	return nil, errors.New("Not a number")
}

func (v *Variable) ToString() (*String, error) {
	return nil, errors.New("Cannot make a string")
}

func (v *Variable) Type() uint8 {
	return VariableT
}

func (v *Variable) GetName() uint32 {
	return v.v
}

func (v *Variable) Unify(i *Interpreter) (*Interpreter, bool) {
	success := i.Env.UnifyEnv(v.v, i.Stack.Car())
	if success {
		i.Stack = i.Stack.Cdr()
		return i, true
	}
	return i, false
}

func (v *Variable) StrictCompare(t2 Type) int {
	if t2.Type() != VariableT {
		if VariableT < t2.Type() {
			return -1
		} else {
			return 1
		}
	} else {
		return 0
	}
}

func (v *Variable) Compare(t2 Type) (int, error) {
	return 0, errors.New("Cannot compare")
}

func (v *Variable) ToList() (*List, error) {
	lst := MakeList()
	lst.Cons(v)
	return lst, nil
}

func (v *Variable) ToChar() (*Char, error) {
	return nil, errors.New("Cannot make a char")
}

func (v *Variable) ToBool() (*Bool, error) {
	return MakeBool(true), nil
}

func (v *Variable) ToTuple() (*Tuple, error) {
	return nil, errors.New("Cannot convert to tuple")
}

func (v *Variable) ToMap() (*Tree, error) {
	return nil, errors.New("Cannot convert to a dict")
}

func (v *Variable) ToSymbol() (*Symbol, error) {
	return nil, errors.New("Cannot convert to symbol")
}

func (v *Variable) ToSet() (*Tree, error) {
	return nil, errors.New("Cannot convert to set")
}

func (v *Variable) ToCore() (*Core, error) {
	return nil, errors.New("Cannot make core")
}

func (v *Variable) Get(t2 Type) (Type, error) {
	return nil, errors.New("Cannot get")
}

func (v *Variable) Set(t2 Type, t3 Type) (Type, error) {
	return nil, errors.New("Cannot set")
}

func (v *Variable) ToFstream() (*Fstream, error) {
	return nil, errors.New("Not a file stream")
}

type Recur struct {}

func MakeRecur() Type {
	return &Recur{}
}

func (r *Recur)	String() string {
	return "recur"
}

func (r *Recur)	ToNumber() (*Number, error) {
	return nil, errors.New("Not a number")
}

func (r *Recur) ToString() (*String, error) {
	return MakeString("recur"), nil
}

func (r *Recur)	Apply(i *Interpreter) (*Interpreter, error) {
	return i, nil
}

func (r *Recur)	Type() uint8 {
	return RecurT
}

func (r *Recur)	GetName() uint32 {
	return 0
}

func (r *Recur)	Unify(i *Interpreter) (*Interpreter, bool) {
	return i, false
}

func (r *Recur)	StrictCompare(t2 Type) int {
	if t2.Type() != RecurT {
		if RecurT < t2.Type() {
			return -1
		} else {
			return 1
		}
	} else {
		return 0
	}
}

func (r *Recur) Compare(t2 Type) (int, error) {
	return 0, errors.New("Cannot compare")
}

func (r *Recur) ToList() (*List, error) {
	lst := MakeList()
	lst.Cons(r)
	return lst, nil
}

func (r *Recur) ToChar() (*Char, error) {
	return nil, errors.New("Cannot make a char")
}

func (r *Recur) ToBool() (*Bool, error) {
	return MakeBool(true), nil
}

func (r *Recur) ToTuple() (*Tuple, error) {
	return nil, errors.New("Cannot convert to tuple")
}

func (r *Recur) ToMap() (*Tree, error) {
	return nil, errors.New("Cannot convert to a dict")
}

func (r *Recur) ToSymbol() (*Symbol, error) {
	return nil, errors.New("Cannot convert to symbol")
}

func (r *Recur) ToSet() (*Tree, error) {
	return nil, errors.New("Cannot convert to set")
}

func (r *Recur) ToCore() (*Core, error) {
	return nil, errors.New("Cannot make core")
}

func (r *Recur) Get(t2 Type) (Type, error) {
	return nil, errors.New("Cannot get")
}

func (r *Recur) Set(t2 Type, t3 Type) (Type, error) {
	return nil, errors.New("Cannot set")
}

func (r *Recur) ToFstream() (*Fstream, error) {
	return nil, errors.New("Not a file stream")
}

type VariableTable struct {
	current_n uint32
	names map[string]uint32
	ints map[uint32]string
}

func (vt *VariableTable) Assign(name string) uint32 {
	i, ok := vt.names[name]
	if ok {
		return i
	} else {
		vt.current_n++
		vt.names[name] = vt.current_n
		vt.ints[vt.current_n] = name
		return vt.current_n
	}
}

var Variable_Table = &VariableTable{
	0,
	make(map[string]uint32),
	make(map[uint32]string)}

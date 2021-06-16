package types

const (
	VariableT uint8 = iota
	NumberT
	StringT
	SymbolT
	BoolT
	RestVariableT
	ListT
	ListLiteralT
	TupleT
	FunctionT
	MapT
	MapLiteralT
	SetT
	SetLiteralT
	ListInjectT
	AtT
	ToT
	UpT
	WhateverT
	EmptyStackT
	DefineT
	RecurT
	CoreT
	CharT
	FstreamT
)

// figure out floor for big.Rat
//  floor is num // denom over 1
type Type interface {
	String() string
	Apply(i *Interpreter) (*Interpreter, error)
	Type() uint8
	GetName() uint32
	Unify(i *Interpreter) (*Interpreter, bool)
	StrictCompare(t2 Type) int
	Compare(t2 Type) (int, error)
	ToNumber() (*Number, error)
	ToString() (*String, error)
	ToList() (*List, error)
	ToChar() (*Char, error)
	ToBool() (*Bool, error)
	ToTuple() (*Tuple, error)
	ToMap() (*Tree, error)
	ToSymbol() (*Symbol, error)
	ToSet() (*Tree, error)
	ToCore() (*Core, error)
	Get(t2 Type) (Type, error)
	Set(t2 Type, t3 Type) (Type, error)
	ToFstream() (*Fstream, error)
}


type Environment struct {
	Env map[uint32]Type
	Next *Environment
}

type Interpreter struct {
	Program *List
	Stack   *List
	Env     *Environment
}

func NewEnv() *Environment {
	return &Environment{ make(map[uint32]Type), nil}
}

func FreshEnv(e *Environment) *Environment {
	return &Environment{ make(map[uint32]Type), e }
}

func (e *Environment) Add(name uint32, t Type) {
	e.Env[name] = t
}

func (e *Environment) UnifyEnv(name uint32, t Type) bool {
	t2, ok := e.Env[name]
	if ok {
		if t.StrictCompare(t2) == 0 {
			return true
		} else {
			return false
		}
	}
	e.Add(name, t)
	return true
}

func (e *Environment) Lookup(name uint32) (Type, bool) {
	temp_env := e
	for temp_env != nil {
		t, ok := temp_env.Env[name]
		if ok {
			return t, true
		}
		temp_env = temp_env.Next
	}
	return nil, false
}

func MakeInterpreter(program *List, current_environment *Environment) *Interpreter {
	return &Interpreter{program, MakeList(), current_environment}
}

func (interpreter *Interpreter) Step() (*Interpreter, error) {
	word := interpreter.Program.Car()
	interpreter.Program = interpreter.Program.Cdr()
	i, err := word.Apply(interpreter)
	return i, err
}

func (interpreter *Interpreter) Run() (*Interpreter, error) {
	var err error
	for err == nil && !interpreter.Program.Nullp() {
		interpreter, err = interpreter.Step()
	}
	return interpreter, err
}


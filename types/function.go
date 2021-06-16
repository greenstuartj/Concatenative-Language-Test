package types

import (
	"errors"
)

type Function struct {
	Name uint32
	bodies []FunctionBody
}

type FunctionBody struct {
	args *List
	body *List
}

func MakeFunctionBody(args *List, body *List) FunctionBody {
	return FunctionBody{ args, body }
}

func ToTailCall(p *List) bool {
	return !p.Nullp() && p.Next.Nullp() && p.Element.Type() == RecurT
}

func (fb *FunctionBody) Apply(i *Interpreter, f *Function) (*Interpreter, bool, error) {
	var err error
	prog := fb.body
	new_i := i
	for err == nil && !prog.Nullp() && !ToTailCall(prog) {
		e := prog.Element
		prog = prog.Cdr()
		if e.Type() == RecurT {
			new_i, err = f.Apply(new_i)
		} else if e.Type() == ListT {
			lst,_ := e.ToList()
			if !lst.Nullp() && lst.Cdr().Nullp() && lst.Car().Type() == RestVariableT {
				name := lst.Car().GetName()
				rv,_ := i.Env.Lookup(name)
				new_i, err = rv.Apply(new_i)
			} else {
				new_i, err = e.Apply(new_i)
			}
		} else {
			new_i, err = e.Apply(new_i)
		}
	}
	return new_i, ToTailCall(prog), err
}

func (fb *FunctionBody) UnifyArgs(i *Interpreter) (*Interpreter, bool) {
	if fb.args.Nullp() {
		return i, true
	}
	temp_args := fb.args
	i, unified := temp_args.Element.Unify(i)
	temp_args = temp_args.Next
	for unified && !temp_args.Nullp() {
		i, unified = temp_args.Element.Unify(i)
		temp_args = temp_args.Next
	}
	return i, unified
}

func (fb *FunctionBody) String() string {
	s := ""
	a := fb.args
	for !a.Nullp() {
		s += a.Element.String()
		s += " "
		a = a.Next
	}
	s += "-- "
	b := fb.body
	for !b.Nullp() {
		s += b.Element.String()
		s += " "
		b = b.Next
	}
	return s
}

// unification of function as a guard, 1 arg only?

func MakeFunction(bodies []FunctionBody) *Function {
	return &Function{ 0, bodies }
}

func (f *Function) Unify(i *Interpreter) (*Interpreter, bool) {
	temp_i := MakeInterpreter(MakeList(), i.Env)
	temp_i.Stack = temp_i.Stack.Cons(i.Stack.Car())
	temp_i, err := f.Apply(temp_i)
	if err != nil {
		return i, false
	}
	b,err := temp_i.Stack.Car().ToBool()
	if err != nil {
		return i, false
	}
	if b.b {
		if f.bodies[0].args.Nullp() {
			i.Stack = i.Stack.Cdr()
			return i, true
		}
		v := f.bodies[0].args.Car()
		i, unified := v.Unify(i)
		return i, unified
	}
	return i, false
}

func (f *Function) Apply(i *Interpreter) (*Interpreter, error) {
	var new_i *Interpreter
	var recur bool
	var err error
	new_i = i
tco:
	for j := 0; j < len(f.bodies); j++ {
		temp_i := MakeInterpreter(new_i.Program, FreshEnv(new_i.Env))
		temp_i.Stack = new_i.Stack
		temp_i, unified := f.bodies[j].UnifyArgs(temp_i)
		if unified {
			new_i, recur, err = f.bodies[j].Apply(temp_i, f)
			new_i.Env = new_i.Env.Next
			if recur {
				goto tco
			} else {
				return new_i, err
			}
		}
	}
	return nil, errors.New("no match")
}

func (f *Function) String() string {
	s := "{ "
	for i,b := range f.bodies {
		s += b.String()
		if i != len(f.bodies)-1 {
			s += "; "
		}
	}
	s += "}"
	return s
}

func (f *Function) ToNumber() (*Number, error) {
	return nil, errors.New("Cannot make number")
}

func (f *Function) ToString() (*String, error) {
	return nil, errors.New("Cannot make string")
}

func (f *Function) Type() uint8 {
	return FunctionT
}

func (f *Function) GetName() uint32 {
	return f.Name
}

func (f *Function) StrictCompare(t2 Type) int {
	if t2.Type() != FunctionT {
		if FunctionT < t2.Type() {
			return -1
		} else {
			return 1
		}
	} else {
		return 0
	}
}

func (f *Function) Compare(t2 Type) (int, error) {
	return 0, errors.New("Cannot compare")
}

func (f *Function) ToList() (*List, error) {
	lst := MakeList()
	lst = lst.Cons(f)
	return lst, nil
}

func (f *Function) ToChar() (*Char, error) {
	return nil, errors.New("Cannot make a char")
}

func (f *Function) ToBool() (*Bool, error) {
	return MakeBool(true), nil
}

func (f *Function) ToTuple() (*Tuple, error) {
	return nil, errors.New("Cannot convert to tuple")
}

func (f *Function) ToMap() (*Tree, error) {
	return nil, errors.New("Cannot convert to a dict")
}

func (f *Function) ToSymbol() (*Symbol, error) {
	return nil, errors.New("Cannot convert to symbol")
}

func (f *Function) ToSet() (*Tree, error) {
	return nil, errors.New("Cannot convert to set")
}

func (f *Function) ToCore() (*Core, error) {
	return nil, errors.New("Cannot make core")
}

func (f *Function) Get(t2 Type) (Type, error) {
	return nil, errors.New("Cannot get")
}

func (f *Function) Set(t2 Type, t3 Type) (Type, error) {
	return nil, errors.New("Cannot set")
}

func (f *Function) ToFstream() (*Fstream, error) {
	return nil, errors.New("Not a file stream")
}

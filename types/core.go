package types

import (
	"fmt"
	"time"
	"errors"
	"os"
	"bufio"
	"strings"
	"io/ioutil"
)

type Core struct {
	v uint32
}

func MakeCore(i uint32) *Core {
	return &Core{ i }
}

func (c *Core) String() string {
	return CoreTable.GetName(c.v)
}

func (c *Core) ToNumber() (*Number, error) {
	return nil, errors.New("Not a number")
}

func (c *Core) ToString() (*String, error) {
	return MakeString(c.String()), nil // no error?
}

func (c *Core) Apply(i *Interpreter) (*Interpreter, error) {
	return CoreTable.Lookup(c.v)(i)
}

func (c *Core) Type() uint8 {
	return CoreT
}

func (c *Core) GetName() uint32 {
	return c.v
}

// TODO -- needs ToCore?
func (c *Core) Unify(i *Interpreter) (*Interpreter, bool) {
	return i, false
}

// TODO -- needs ToCore?
func (c *Core) StrictCompare(t2 Type) int {
	if t2.Type() != CoreT {
		if CoreT < t2.Type() {
			return -1
		} else {
			return 1
		}
	} else {
		return 0
	}
}

func (c *Core) Compare(t2 Type) (int, error) {
	return 0, errors.New("Cannot compare")
}

func (c *Core) ToList() (*List, error) {
	lst := MakeList()
	lst.Cons(c)
	return lst, nil
}

func (c *Core) ToChar() (*Char, error) {
	return nil, errors.New("Cannot make a char")
}

func (c *Core) ToBool() (*Bool, error) {
	return MakeBool(true), nil
}

func (c *Core) ToTuple() (*Tuple, error) {
	return nil, errors.New("Cannot convert to tuple")
}

func (c *Core) ToMap() (*Tree, error) {
	return nil, errors.New("Cannot convert to a dict")
}

func (c *Core) ToSymbol() (*Symbol, error) {
	return nil, errors.New("Cannot convert to symbol")
}

func (c *Core) ToSet() (*Tree, error) {
	return nil, errors.New("Cannot convert to set")
}

func (c *Core) ToCore() (*Core, error) {
	return c, nil
}

func (c *Core) Get(t2 Type) (Type, error) {
	return nil, errors.New("Cannot get")
}

func (c *Core) Set(t2 Type, t3 Type) (Type, error) {
	return nil, errors.New("Cannot set")
}

func (c *Core) ToFstream() (*Fstream, error) {
	return nil, errors.New("Not a file stream")
}


type stack_function = func(*Interpreter) (*Interpreter, error)
type core_env = map[uint32]stack_function

var Core_names map[string]uint32
var core_old core_env

var CoreTable *Core_Table

type Core_Table struct {
	current_n uint32
	names map[string]uint32
	ints map[uint32]string
	functions map[uint32]stack_function
}

func MakeCoreTable() *Core_Table {
	return &Core_Table{
		0,
		make(map[string]uint32),
		make(map[uint32]string),
		make(map[uint32]stack_function)}
}

func (ct *Core_Table) Add(name string, function stack_function) {
	ct.current_n++
	ct.names[name] = ct.current_n
	ct.ints[ct.current_n] = name
	ct.functions[ct.current_n] = function
}

func (ct *Core_Table) Lookup(name uint32) stack_function {
	f,_ := ct.functions[name]
	return f
}

func (ct *Core_Table) GetName(name uint32) string {
	sname,_ := ct.ints[name]
	return sname
}

func (ct *Core_Table) IsCore(name string) (uint32, bool) {
	i, ok := ct.names[name]
	return i, ok
}

func Init_core() {
	CoreTable = MakeCoreTable()
	CoreTable.Add("show",
		func(i *Interpreter) (*Interpreter, error) {
			var stack_slice []Type
			temp_stack := i.Stack
			for !temp_stack.Nullp() {
				stack_slice = append(stack_slice, temp_stack.Car())
				temp_stack = temp_stack.Cdr()
			}
			var shown_stack string
			shown_stack += " =>"
			for p := len(stack_slice) ; p > 0 ; p-- {
				shown_stack += " "
				if stack_slice[p-1].Type() == StringT {
					str,_ := stack_slice[p-1].ToString()
					shown_stack += str.StringWithQuotes()
				} else {
					shown_stack += stack_slice[p-1].String()
				}
			}
			fmt.Println(shown_stack)
			return i, nil
		})
	CoreTable.Add("println",
		func(i *Interpreter) (*Interpreter, error) {
			top := i.Stack.Car()
			fmt.Println(top)
			return i, nil
		})
	CoreTable.Add("dup",
		func(i *Interpreter) (*Interpreter, error) {
			top := i.Stack.Car()
			i.Stack = i.Stack.Cons(top)
			return i, nil
		})
	CoreTable.Add("drop",
		func(i *Interpreter) (*Interpreter, error) {
			i.Stack = i.Stack.Cdr()
			return i, nil
		})
	CoreTable.Add("swap",
		func(i *Interpreter) (*Interpreter, error) {
			top := i.Stack.Car()
			i.Stack = i.Stack.Cdr()
			top2 := i.Stack.Car()
			i.Stack = i.Stack.Cdr()
			i.Stack = i.Stack.Cons(top)
			i.Stack = i.Stack.Cons(top2)
			return i, nil
		})
	CoreTable.Add("append",
		func(i *Interpreter) (*Interpreter, error) {
			top2 := i.Stack.Car()
			i.Stack = i.Stack.Cdr()
			if top2.Type() != ListT {
				return nil, errors.New("Not a list")
			}
			lst2,_ := top2.ToList()
			top := i.Stack.Car()
			if top.Type() != ListT {
				return nil, errors.New("Not a list")
			}
			lst,_ := top.ToList()
			i.Stack = i.Stack.Cdr()
			temp := MakeList()
			for !lst.Nullp() {
				temp = temp.Cons(lst.Element)
				lst = lst.Next
			}
			for !temp.Nullp() {
				lst2 = lst2.Cons(temp.Element)
				temp = temp.Next
			}
			i.Stack = i.Stack.Cons(lst2)
			return i, nil
		})
	CoreTable.Add("+",
		func(i *Interpreter) (*Interpreter, error) {
			b := i.Stack.Car()
			a := i.Stack.Cdr().Car()
			an, err := a.ToNumber()
			if err != nil {
				return nil, err
			}
			bn, err := b.ToNumber()
			if err != nil {
				return nil, err
			}
			c := an.Add(bn)
			i.Stack = i.Stack.Cdr().Cdr().Cons(c)
			return i, nil
		})
	CoreTable.Add("-",
		func(i *Interpreter) (*Interpreter, error) {
			b := i.Stack.Car()
			a := i.Stack.Cdr().Car()
			an, err := a.ToNumber()
			if err != nil {
				return nil, err
			}
			bn, err := b.ToNumber()
			if err != nil {
				return nil, err
			}
			c := an.Minus(bn)
			i.Stack = i.Stack.Cdr().Cdr().Cons(c)
			return i, nil
		})
	CoreTable.Add("*",
		func(i *Interpreter) (*Interpreter, error) {
			b := i.Stack.Car()
			a := i.Stack.Cdr().Car()
			an, err := a.ToNumber()
			if err != nil {
				return nil, err
			}
			bn, err := b.ToNumber()
			if err != nil {
				return nil, err
			}
			c := an.Multiply(bn)
			i.Stack = i.Stack.Cdr().Cdr().Cons(c)
			return i, nil
		})
	CoreTable.Add("/",
		func(i *Interpreter) (*Interpreter, error) {
			b := i.Stack.Car()
			a := i.Stack.Cdr().Car()
			an, err := a.ToNumber()
			if err != nil {
				return nil, err
			}
			bn, err := b.ToNumber()
			if err != nil {
				return nil, err
			}
			c := an.Divide(bn)
			i.Stack = i.Stack.Cdr().Cdr().Cons(c)
			return i, nil
		})
	CoreTable.Add("%",
		func(i *Interpreter) (*Interpreter, error) {
			b := i.Stack.Car()
			a := i.Stack.Cdr().Car()
			an, err := a.ToNumber()
			if err != nil {
				return nil, err
			}
			bn, err := b.ToNumber()
			if err != nil {
				return nil, err
			}
			if !an.n.IsInt() || !bn.n.IsInt() {
				return nil, errors.New("Not integer")
			}
			c := an.Modulo(bn)
			i.Stack = i.Stack.Cdr().Cdr().Cons(c)
			return i, nil
		})
	CoreTable.Add("sleep",
		func(i *Interpreter) (*Interpreter, error) {
			n := i.Stack.Car()
			x, err := n.ToNumber()
			if err != nil {
				return i, err
			}
			if !x.n.IsInt() {
				return i, err
			}
			time.Sleep(time.Duration(x.n.Num().Int64()) * time.Millisecond)
			i.Stack = i.Stack
			return i, nil
		})
	CoreTable.Add("cons",
		func(i *Interpreter) (*Interpreter, error) {
			lstT := i.Stack.Car()
			e := i.Stack.Cdr().Car()
			lst, err := lstT.ToList()
			if err != nil {
				return i, err
			}
			lst = lst.Cons(e)
			i.Stack = i.Stack.Cdr().Cdr().Cons(lst)
			return i, nil
		})
	CoreTable.Add("uncons",
		func(i *Interpreter) (*Interpreter, error) {
			lstT := i.Stack.Car()
			lst, err := lstT.ToList()
			if err != nil {
				return i, err
			}
			e := lst.Car()
			rest := lst.Cdr()
			i.Stack = i.Stack.Cdr().Cons(e).Cons(rest)
			return i, nil
		})
	CoreTable.Add("apply",
		func(i *Interpreter) (*Interpreter, error) {
			lstT := i.Stack.Car()
			if lstT.Type() != ListT {
				i.Stack = i.Stack.Cdr()
				return lstT.Apply(i)
			}
			lst, err := lstT.ToList()
			if err != nil {
				return i, err
			}
			new_i := MakeInterpreter(lst, i.Env)
			new_i.Stack = i.Stack.Cdr()
			new_i, err = new_i.Run()
			if err != nil {
				return new_i, err
			}
			i.Stack = new_i.Stack
			i.Env = new_i.Env
			return i, nil
		})
	CoreTable.Add("clear",
		func(i *Interpreter) (*Interpreter, error) {
			i.Stack = MakeList()
			return i, nil
		})
	CoreTable.Add("print",
		func(i *Interpreter) (*Interpreter, error) {
			top := i.Stack.Car()
			fmt.Print(top)
			return i, nil
		})
	CoreTable.Add("string->list",
		func(i *Interpreter) (*Interpreter, error) {
			top := i.Stack.Car()
			i.Stack = i.Stack.Cdr()
			str, err := top.ToString()
			if err != nil {
				return nil, err
			}
			s := str.String()
			lst := MakeList()
			for i:= len(s); i > 0; i-- {
				lst = lst.Cons(MakeChar(s[i-1]))
			}
			i.Stack = i.Stack.Cons(lst)
			return i, nil
		})
	CoreTable.Add("list->string",
		func (i *Interpreter) (*Interpreter, error) {
			top := i.Stack.Car()
			i.Stack = i.Stack.Cdr()
			if top.Type() != ListT {
				return i, errors.New("Not a list")
			}
			s := ""
			lst,_ := top.ToList()
			for !lst.Nullp() {
				c,err := lst.Element.ToChar()
				if err != nil {
					return nil, err
				}
				s += string(c.c)
				lst = lst.Next
			}
			i.Stack = i.Stack.Cons(MakeString(s))
			return i, nil
		})
	CoreTable.Add("eq?",
		func (i *Interpreter) (*Interpreter, error) {
			top := i.Stack.Car()
			i.Stack = i.Stack.Cdr()
			top2 := i.Stack.Car()
			i.Stack = i.Stack.Cdr()
			if top.StrictCompare(top2) == 0 {
				i.Stack = i.Stack.Cons(MakeBool(true))
			} else {
				i.Stack = i.Stack.Cons(MakeBool(false))
			}
			return i, nil
		})
	CoreTable.Add("=",
		func (i *Interpreter) (*Interpreter, error) {
			top := i.Stack.Car()
			i.Stack = i.Stack.Cdr()
			top2 := i.Stack.Car()
			n, err := top2.ToNumber()
			if err != nil {
				return nil, err
			}
			i.Stack = i.Stack.Cdr()
			c, err := n.Compare(top)
			if err != nil {
				return nil, err
			}
			if c == 0 {
				i.Stack = i.Stack.Cons(MakeBool(true))
			} else {
				i.Stack = i.Stack.Cons(MakeBool(false))
			}
			return i, nil
		})
	CoreTable.Add("<",
		func (i *Interpreter) (*Interpreter, error) {
			top := i.Stack.Car()
			i.Stack = i.Stack.Cdr()
			top2 := i.Stack.Car()
			n, err := top2.ToNumber()
			if err != nil {
				return nil, err
			}
			i.Stack = i.Stack.Cdr()
			c, err := n.Compare(top)
			if err != nil {
				return nil, err
			}
			if c == -1 {
				i.Stack = i.Stack.Cons(MakeBool(true))
			} else {
				i.Stack = i.Stack.Cons(MakeBool(false))
			}
			return i, nil
		})
	CoreTable.Add("<=",
		func (i *Interpreter) (*Interpreter, error) {
			top := i.Stack.Car()
			i.Stack = i.Stack.Cdr()
			top2 := i.Stack.Car()
			n, err := top2.ToNumber()
			if err != nil {
				return nil, err
			}
			i.Stack = i.Stack.Cdr()
			c, err := n.Compare(top)
			if err != nil {
				return nil, err
			}
			if c == -1 || c == 0 {
				i.Stack = i.Stack.Cons(MakeBool(true))
			} else {
				i.Stack = i.Stack.Cons(MakeBool(false))
			}
			return i, nil
		})
	CoreTable.Add(">",
		func (i *Interpreter) (*Interpreter, error) {
			top := i.Stack.Car()
			i.Stack = i.Stack.Cdr()
			top2 := i.Stack.Car()
			n, err := top2.ToNumber()
			if err != nil {
				return nil, err
			}
			i.Stack = i.Stack.Cdr()
			c, err := n.Compare(top)
			if err != nil {
				return nil, err
			}
			if c == 1 {
				i.Stack = i.Stack.Cons(MakeBool(true))
			} else {
				i.Stack = i.Stack.Cons(MakeBool(false))
			}
			return i, nil
		})
	CoreTable.Add(">=",
		func (i *Interpreter) (*Interpreter, error) {
			top := i.Stack.Car()
			i.Stack = i.Stack.Cdr()
			top2 := i.Stack.Car()
			n, err := top2.ToNumber()
			if err != nil {
				return nil, err
			}
			i.Stack = i.Stack.Cdr()
			c, err := n.Compare(top)
			if err != nil {
				return nil, err
			}
			if c == 0 || c == 1 {
				i.Stack = i.Stack.Cons(MakeBool(true))
			} else {
				i.Stack = i.Stack.Cons(MakeBool(false))
			}
			return i, nil
		})
	CoreTable.Add("=s",
		func (i *Interpreter) (*Interpreter, error) {
			top := i.Stack.Car()
			i.Stack = i.Stack.Cdr()
			top2 := i.Stack.Car()
			s, err := top2.ToString()
			if err != nil {
				return nil, err
			}
			i.Stack = i.Stack.Cdr()
			c, err := s.Compare(top)
			if err != nil {
				return nil, err
			}
			if c == 0 {
				i.Stack = i.Stack.Cons(MakeBool(true))
			} else {
				i.Stack = i.Stack.Cons(MakeBool(false))
			}
			return i, nil
		})
	CoreTable.Add("<s",
		func (i *Interpreter) (*Interpreter, error) {
			top := i.Stack.Car()
			i.Stack = i.Stack.Cdr()
			top2 := i.Stack.Car()
			s, err := top2.ToString()
			if err != nil {
				return nil, err
			}
			i.Stack = i.Stack.Cdr()
			c, err := s.Compare(top)
			if err != nil {
				return nil, err
			}
			if c == -1 {
				i.Stack = i.Stack.Cons(MakeBool(true))
			} else {
				i.Stack = i.Stack.Cons(MakeBool(false))
			}
			return i, nil
		})
	CoreTable.Add("<=s",
		func (i *Interpreter) (*Interpreter, error) {
			top := i.Stack.Car()
			i.Stack = i.Stack.Cdr()
			top2 := i.Stack.Car()
			s, err := top2.ToString()
			if err != nil {
				return nil, err
			}
			i.Stack = i.Stack.Cdr()
			c, err := s.Compare(top)
			if err != nil {
				return nil, err
			}
			if c == -1 || c == 0 {
				i.Stack = i.Stack.Cons(MakeBool(true))
			} else {
				i.Stack = i.Stack.Cons(MakeBool(false))
			}
			return i, nil
		})
	CoreTable.Add(">s",
		func (i *Interpreter) (*Interpreter, error) {
			top := i.Stack.Car()
			i.Stack = i.Stack.Cdr()
			top2 := i.Stack.Car()
			s, err := top2.ToString()
			if err != nil {
				return nil, err
			}
			i.Stack = i.Stack.Cdr()
			c, err := s.Compare(top)
			if err != nil {
				return nil, err
			}
			if c == 1 {
				i.Stack = i.Stack.Cons(MakeBool(true))
			} else {
				i.Stack = i.Stack.Cons(MakeBool(false))
			}
			return i, nil
		})
	CoreTable.Add(">=s",
		func (i *Interpreter) (*Interpreter, error) {
			top := i.Stack.Car()
			i.Stack = i.Stack.Cdr()
			top2 := i.Stack.Car()
			s, err := top2.ToString()
			if err != nil {
				return nil, err
			}
			i.Stack = i.Stack.Cdr()
			c, err := s.Compare(top)
			if err != nil {
				return nil, err
			}
			if c == 0 || c == 1 {
				i.Stack = i.Stack.Cons(MakeBool(true))
			} else {
				i.Stack = i.Stack.Cons(MakeBool(false))
			}
			return i, nil
		})
	CoreTable.Add("read-line",
		func (i *Interpreter) (*Interpreter, error) {
			reader := bufio.NewReader(os.Stdin)
			text, _ := reader.ReadString('\n')
			i.Stack = i.Stack.Cons(MakeString(text[:len(text)-1]))
			return i, nil
		})
	CoreTable.Add("exit",
		func (i *Interpreter) (*Interpreter, error) {
			fmt.Println(" Leaving FOLDEX interpreter...")
			os.Exit(0)
			return nil, nil
		})
	CoreTable.Add("stack->list",
		func (i *Interpreter) (*Interpreter, error) {
			stack := i.Stack
			i.Stack = MakeList()
			out := MakeList()
			for !stack.Nullp() {
				out = out.Cons(stack.Element)
				stack = stack.Next
			}
			i.Stack = i.Stack.Cons(out)
			return i, nil
		})
	CoreTable.Add("+s",
		func (i *Interpreter) (*Interpreter, error) {
			top2 := i.Stack.Car()
			i.Stack = i.Stack.Cdr()
			top := i.Stack.Car()
			i.Stack = i.Stack.Cdr()
			s, err := top.ToString()
			if err != nil {
				return i, err
			}
			s2, err := top2.ToString()
			if err != nil {
				return i, err
			}
			i.Stack = i.Stack.Cons(MakeString(s.String() + s2.String()))
			return i, nil
		})
	CoreTable.Add("split",
		func (i *Interpreter) (*Interpreter, error) {
			top2 := i.Stack.Car()
			i.Stack = i.Stack.Cdr()
			top := i.Stack.Car()
			i.Stack = i.Stack.Cdr()
			s, err := top.ToString()
			if err != nil {
				return i, err
			}
			s2, err := top2.ToString()
			if err != nil {
				return i, err
			}
			arr := strings.Split(s.String(), s2.String())
			sList := MakeList()
			for j := len(arr); j > 0; j-- {
				sList = sList.Cons(MakeString(arr[j-1]))
			}
			i.Stack = i.Stack.Cons(sList)
			return i, nil
		})
	CoreTable.Add("replace",
		func (i *Interpreter) (*Interpreter, error) {
			top3 := i.Stack.Car()
			i.Stack = i.Stack.Cdr()
			top2 := i.Stack.Car()
			i.Stack = i.Stack.Cdr()
			top := i.Stack.Car()
			i.Stack = i.Stack.Cdr()
			s, err := top.ToString()
			if err != nil {
				return i, err
			}
			s2, err := top2.ToString()
			if err != nil {
				return i, err
			}
			s3, err := top3.ToString()
			if err != nil {
				return i, err
			}
			s4 := strings.ReplaceAll(s.String(), s2.String(), s3.String())
			i.Stack = i.Stack.Cons(MakeString(s4))
			return i, nil
		})
	CoreTable.Add("emit",
		func (i *Interpreter) (*Interpreter, error) {
			top := i.Stack.Car()
			c, err := top.ToChar()
			if err != nil {
				return i, err
			}
			s, err := c.ToString()
			if err != nil {
				return i, err
			}
			fmt.Print(s)
			return i, nil
		})
	CoreTable.Add("read-file",
		func (i *Interpreter) (*Interpreter, error) {
			top := i.Stack.Car()
			if top.Type() != StringT {
				return i, errors.New("Not a string")
			}
			s, err := ioutil.ReadFile(top.String())
			if err != nil {
				return i, err
			}
			i.Stack = i.Stack.Cdr()
			i.Stack = i.Stack.Cons(MakeString(string(s)))
			return i, nil
		})
	CoreTable.Add("not",
		func (i *Interpreter) (*Interpreter, error) {
			top := i.Stack.Car()
			b,_ := top.ToBool()
			i.Stack = i.Stack.Cdr()
			i.Stack = i.Stack.Cons(MakeBool(!b.b))
			return i, nil
		})
	CoreTable.Add("insert",
		func (i *Interpreter) (*Interpreter, error) {
			key := i.Stack.Car()
			i.Stack = i.Stack.Cdr()
			value := i.Stack.Car()
			i.Stack = i.Stack.Cdr()
			dict := i.Stack.Car()
			i.Stack = i.Stack.Cdr()
			if dict.Type() != MapT {
				return i, errors.New("Not a dict")
			}
			m,_ := dict.ToMap()
			i.Stack = i.Stack.Cons(m.Insert(key, value))
			return i, nil
		})
	CoreTable.Add("lookup",
		func (i *Interpreter) (*Interpreter, error) {
			key := i.Stack.Car()
			i.Stack = i.Stack.Cdr()
			dict := i.Stack.Car()
			i.Stack = i.Stack.Cdr()
			if dict.Type() != MapT {
				return i, errors.New("Not a dict")
			}
			m,_ := dict.ToMap()
			t,ok := m.Lookup(key)
			// should return (OK t)
			if ok {
				i.Stack = i.Stack.Cons(t)
				return i, nil
			} else {
				return i, nil
			}
		})
	CoreTable.Add("add",
		func (i *Interpreter) (*Interpreter, error) {
			key := i.Stack.Car()
			i.Stack = i.Stack.Cdr()
			set := i.Stack.Car()
			i.Stack = i.Stack.Cdr()
			if set.Type() != SetT {
				return i, errors.New("Not a set")
			}
			s,_ := set.ToSet()
			s = s.Insert(key, nil)
			i.Stack = i.Stack.Cons(s)
			return i, nil
		})
	CoreTable.Add("member?",
		func (i *Interpreter) (*Interpreter, error) {
			key := i.Stack.Car()
			i.Stack = i.Stack.Cdr()
			set := i.Stack.Car()
			i.Stack = i.Stack.Cdr()
			if set.Type() != SetT {
				return i, errors.New("Not a set")
			}
			s,_ := set.ToSet()
			_,ok := s.Lookup(key)
			if ok {
				i.Stack = i.Stack.Cons(MakeBool(true))
			} else {
				i.Stack = i.Stack.Cons(MakeBool(false))
			}
			return i, nil
		})
	CoreTable.Add("newline",
		func (i *Interpreter) (*Interpreter, error) {
			fmt.Println("")
			return i, nil
		})
	CoreTable.Add("open",
		func (i *Interpreter) (*Interpreter, error) {
			top := i.Stack.Car()
			fs,err := MakeFstream(top)
			if err != nil {
				return i, err
			}
			i.Stack = i.Stack.Cdr().Cons(fs)
			return i, nil
		})
	CoreTable.Add("close",
		func (i *Interpreter) (*Interpreter, error) {
			top := i.Stack.Car()
			fs, err := top.ToFstream()
			if err != nil {
				return i, err
			}
			fs.Close()
			i.Stack = i.Stack.Cdr()
			return i, nil
		})
	CoreTable.Add("read-char",
		func (i *Interpreter) (*Interpreter, error) {
			top := i.Stack.Car()
			fs, err := top.ToFstream()
			if err != nil {
				return i, err
			}
			c,err := fs.GetChar()
			if err != nil {
				i.Stack = i.Stack.Cdr().Cons(MakeSymbol("EOF"))
				return i, nil
			}
			i.Stack = i.Stack.Cons(c)
			return i, nil
		})
}

// functions for arity to abstract car and cdr behaviours

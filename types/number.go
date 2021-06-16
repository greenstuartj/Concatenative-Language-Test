package types

import (
	"math/big"
	"errors"
)

type Number struct {
	n *big.Rat
}

func MakeNumber(s string) *Number {
	r := new(big.Rat)
	r.SetString(s)
	return &Number{r}
}

func (num *Number) Apply(i *Interpreter) (*Interpreter, error) {
	i.Stack = i.Stack.Cons(num)
	return i, nil
}

func (num *Number) String() string {
	if num.n.IsInt() {
		return num.n.Num().String()
	} else {
		return num.n.FloatString(5)
	}
}

func (num *Number) ToNumber() (*Number, error) {
	return num, nil
}

func (num *Number) ToString() (*String, error) {
	return MakeString(num.String()), nil
}

func (num *Number) Add(t *Number) *Number {
	new_num := new(big.Rat)
	new_num.Add(num.n, t.n)
	return &Number{new_num}
}

func (num *Number) Minus(t *Number) *Number {
	new_num := new(big.Rat)
	new_num.Sub(num.n, t.n)
	return &Number{new_num}
}

func (num *Number) Multiply(t *Number) *Number {
	new_num := new(big.Rat)
	new_num.Mul(num.n, t.n)
	return &Number{new_num}
}

func (num *Number) Divide(t *Number) *Number {
	new_num := new(big.Rat)
	new_num.Quo(num.n, t.n)
	return &Number{new_num}
}

func (num *Number) Modulo(t *Number) *Number {
	new_num := new(big.Rat)
	temp_int := new(big.Int)
	temp_int.Mod(num.n.Num(), t.n.Num())
	new_num.SetInt(temp_int)
	return &Number{new_num}
}

func (num *Number) Type() uint8 {
	return NumberT
}

func (num *Number) GetName() uint32 {
	return 0
}

func (num *Number) Unify(i *Interpreter) (*Interpreter, bool) {
	num2 := i.Stack.Car()
	if num2.Type() != NumberT {
		return i, false
	}
	num3,_ := num2.ToNumber()
	if num.n.Cmp(num3.n) == 0 {
		i.Stack = i.Stack.Cdr()
		return i, true
	}
	return i, false
}

func (num *Number) StrictCompare(t2 Type) int {
	if t2.Type() != NumberT {
		if NumberT < t2.Type() {
			return -1
		} else {
			return 1
		}
	} else {
		n,_ := t2.ToNumber()
		return num.n.Cmp(n.n)
	}
}

func (num *Number) Compare(t2 Type) (int, error) {
	num2, err := t2.ToNumber()
	if err != nil {
		return 0, err
	}
	return num.StrictCompare(num2), nil
}

func (num *Number) ToList() (*List, error) {
	lst := MakeList()
	lst.Cons(num)
	return lst, nil
}

func (num *Number) ToChar() (*Char, error) {
	if num.n.IsInt() {
		n := num.n.Num()
		return MakeChar(byte(n.Int64())), nil
	} else {
		return nil, errors.New("Cannot make a char")
	}
}

func (num *Number) ToBool() (*Bool, error) {
	return MakeBool(true), nil
}

func (num *Number) ToTuple() (*Tuple, error) {
	return nil, errors.New("Cannot convert to tuple")
}

func (num *Number) ToMap() (*Tree, error) {
	return nil, errors.New("Cannot convert to a dict")
}

func (num *Number) ToSymbol() (*Symbol, error) {
	return nil, errors.New("Cannot convert to symbol")
}

func (num *Number) ToSet() (*Tree, error) {
	return nil, errors.New("Cannot convert to set")
}

func (num *Number) ToCore() (*Core, error) {
	return nil, errors.New("Cannot make core")
}

func (num *Number) Get(t2 Type) (Type, error) {
	return nil, errors.New("Cannot get")
}

func (num *Number) Set(t2 Type, t3 Type) (Type, error) {
	return nil, errors.New("Cannot set")
}

func (num *Number) ToFstream() (*Fstream, error) {
	return nil, errors.New("Not a file stream")
}

package types

import (
	"os"
	"errors"
)

type Fstream struct {
	name string
	fs *os.File
}

func MakeFstream(s Type) (*Fstream, error) {
	if s.Type() != StringT {
		return nil, errors.New("Not a string")
	}
	str,_ := s.ToString()
	file := str.s
	fs,_ := os.Open(file)
	return &Fstream{file, fs}, nil
}

func (fs *Fstream) Close() {
	fs.fs.Close()
}

// should it close?
func (fs *Fstream) GetChar() (*Char, error) {
	var c []byte = make([]byte, 1)
	n, err := fs.fs.Read(c)
	if n == 0 || err != nil {
		fs.Close()
		return nil, err
	}
	return MakeChar(c[0]), nil
}

func (fs *Fstream) String() string {
	return "<" + fs.name + ">"
}

func (fs *Fstream) Apply(i *Interpreter) (*Interpreter, error) {
	i.Stack = i.Stack.Cons(fs)
	return i,nil
}

func (fs *Fstream) Type() uint8 {
	return FstreamT
}

func (fs *Fstream) GetName() uint32 {
	return 0
}

func (fs *Fstream) Unify(i *Interpreter) (*Interpreter, bool) {
	return i, false
}

func (fs *Fstream) StrictCompare(t2 Type) int {
	return -1
}

func (fs *Fstream) Compare(t2 Type) (int, error) {
	return -1, errors.New("Cannot compare")
}

func (fs *Fstream) ToNumber() (*Number, error) {
	return nil, errors.New("Cannot make number")
}

func (fs *Fstream) ToString() (*String, error) {
	return nil, errors.New("Cannot make string")
}

func (fs *Fstream) ToList() (*List, error) {
	return nil, errors.New("Cannot make list")
}

func (fs *Fstream) ToChar() (*Char, error) {
	return nil, errors.New("Cannot make char")
}

func (fs *Fstream) ToBool() (*Bool, error) {
	return MakeBool(true), nil
}

func (fs *Fstream) ToTuple() (*Tuple, error) {
	return nil, errors.New("Cannot make tuple")
}

func (fs *Fstream) ToMap() (*Tree, error) {
	return nil, errors.New("Cannot make map")
}

func (fs *Fstream) ToSymbol() (*Symbol, error) {
	return nil, errors.New("Cannot make symbol")
}

func (fs *Fstream) ToSet() (*Tree, error) {
	return nil, errors.New("Cannot make set")
}

func (fs *Fstream) ToCore() (*Core, error) {
	return nil, errors.New("Cannot make core")
}

func (fs *Fstream) Get(t2 Type) (Type, error) {
	return nil, errors.New("Cannot get")
}

func (fs *Fstream) Set(t2 Type, t3 Type) (Type, error) {
	return nil, errors.New("Cannot set")
}

func (fs *Fstream) ToFstream() (*Fstream, error) {
	return fs, nil
}

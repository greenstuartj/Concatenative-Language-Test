package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"lexer/lexer"
	"parser/parser"
	"types/types"
	"bufio"
	"errors"
)

func main() {
	types.Init_core()
	types.CoreTable.Add("load", load)
	if len(os.Args) > 1 {
		prog_string, _ := ioutil.ReadFile(os.Args[1])
		lex := lexer.New(string(prog_string))
		parser, err := parser.New(lex)
		prog, err := parser.Evaluatable()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		stack := types.MakeList()
		for _,arg := range os.Args[2:] {
			stack = stack.Cons(types.MakeString(arg))
		}
		interpreter := types.MakeInterpreter(prog, types.NewEnv())
		interpreter.Stack = stack
		_, err = interpreter.Run()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		os.Exit(0)
	}
	fmt.Println("")
	fmt.Println(" FOLDEX Concatenative Language Interpreter ")
	fmt.Println("")
	interpreter := types.MakeInterpreter(types.MakeList(), types.NewEnv())
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">>> ")
		text, _ := reader.ReadString('\n')
		lex := lexer.New(text)
		parser, err := parser.New(lex)
		if err != nil {
			fmt.Println(err)
			continue
		}
		prog, err := parser.Evaluatable()
		if err != nil || prog == nil {
			fmt.Println(err)
			continue
		}
		interpreter.Program = prog
		interpreter, err = interpreter.Run()
		if err != nil {
			fmt.Println(err)
			continue
		}
	}
}

func load(i *types.Interpreter) (*types.Interpreter, error) {
	filename := i.Stack.Car()
	if filename.Type() != types.StringT {
		return i, errors.New("Not a string")
	}
	f,err := filename.ToString()
	if err != nil {
		return i, err
	}
	prog_string,err := ioutil.ReadFile(f.String())
	if err != nil {
		return i, errors.New("File does not exist")
	}
	lex := lexer.New(string(prog_string))
	parser, err := parser.New(lex)
	if err != nil {
		return i, err
	}
	prog, err := parser.Evaluatable()
	if err != nil {
		return i, err
	}
	interpreter := types.MakeInterpreter(prog, i.Env)
	interpreter, err = interpreter.Run()
	if err != nil {
		return i, err
	}
	interpreter.Program = i.Program
	interpreter.Stack = i.Stack.Cdr()
	return interpreter, nil
}

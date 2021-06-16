package parser

import (
	"errors"
	"lexer/lexer"
	"types/types"
)

const (
	VariableT uint8 = iota
	NumberT
	StringT
	CharT
	SymbolT
	BoolT
	RestVariableT
	ListT
	TupleT
	FunctionT
	MapT
	SetT
	ListInjectT
	AtT
	ToT
	UpT
	WhateverT
	EmptyStackT
	DefineT
	EndT
)

type AST interface {
	GetType() uint8
	Show(indent int) string
	ToType() (types.Type, error)
}

type End struct {
}

func (ast End) GetType() uint8 {
	return EndT
}

func (ast End) Show(indent int) string {
	return ""
}

func IsEOF(ast AST) bool {
	return ast == nil || ast.GetType() == EndT
}

func (ast End) ToType() (types.Type, error) {
	return nil, errors.New("Not a valid type")
}

type Variable struct {
	vrbl string
}

func (ast Variable) GetType() uint8 {
	return VariableT
}

func (ast Variable) Show(indent int) string {
	var s string
	i := indent
	for i > 0 {
		s += " "
		i--
	}
	s += "{ Variable: "
	s += ast.vrbl
	s += " }"
	return s
}

func (ast Variable) ToType() (types.Type, error) {
	if ast.vrbl == "recur" {
		return types.MakeRecur(), nil
	}
	i, ok := types.CoreTable.IsCore(ast.vrbl)
	if ok {
		return types.MakeCore(i), nil
	}
	return types.MakeVariable(ast.vrbl), nil
}

type Number struct {
	num string
}

func (ast Number) GetType() uint8 {
	return NumberT
}

func (ast Number) Show(indent int) string {
	var s string
	i := indent
	for i > 0 {
		s += " "
		i--
	}
	s += "{ Number: "
	s += ast.num
	s += " }"
	return s
}

func (ast Number) ToType() (types.Type, error) {
	return types.MakeNumber(ast.num), nil
}

type String struct {
	str string
}

func (ast String) GetType() uint8 {
	return StringT
}

func (ast String) Show(indent int) string {
	var s string
	i := indent
	for i > 0 {
		s += " "
		i--
	}
	s += "{ String: "
	s += "\"" + ast.str + "\""
	s += " }"
	return s
}

func (ast String) ToType() (types.Type, error) {
	return types.MakeString(ast.str), nil
}

type Char struct {
	c string
}

func (ast Char) GetType() uint8 {
	return CharT
}

func (ast Char) Show(indent int) string {
	var s string
	i := indent
	for i > 0 {
		s += " "
		i--
	}
	s += "{ Char: "
	s += ast.c
	s += " }"
	return s
}

func (ast Char) ToType() (types.Type, error) {
	if len(ast.c) == 1 {
		return types.MakeChar(ast.c[0]), nil
	} else {
		switch ast.c {
		case "space":
			return types.MakeChar(' '), nil
		case "newline":
			return types.MakeChar('\n'), nil
		case "return":
			return types.MakeChar('\r'), nil
		case "tab":
			return types.MakeChar('\t'), nil
		default:
			return nil, errors.New("Invalid char")
		}
	}
}

type Symbol struct {
	sym string
}

func (ast Symbol) GetType() uint8 {
	return SymbolT
}

func (ast Symbol) Show(indent int) string {
	var s string
	i := indent
	for i > 0 {
		s += " "
		i--
	}
	s += "{ Symbol: "
	s += ast.sym
	s += " }"
	return s
}

func (ast Symbol) ToType() (types.Type, error) {
	return types.MakeSymbol(ast.sym), nil
}

type Bool struct {
	b bool
}

func (ast Bool) GetType() uint8 {
	return BoolT
}

func (ast Bool) Show(indent int) string {
	var s string
	i := indent
	for i > 0 {
		s += " "
		i--
	}
	s += "{ Bool: "
	if ast.b {
		s += "True"
	} else {
		s += "False"
	}
	s += " }"
	return s
}

func (ast Bool) ToType() (types.Type, error) {
	return types.MakeBool(ast.b), nil
}

type RestVariable struct {
	rv string
}

func (ast RestVariable) GetType() uint8 {
	return RestVariableT
}

func (ast RestVariable) Show(indent int) string {
	var s string
	i := indent
	for i > 0 {
		s += " "
		i--
	}
	s += "{ RestVariable: "
	s += ast.rv
	s += " }"
	return s
}

func (ast RestVariable) ToType() (types.Type, error) {
	return types.MakeRestVariable(ast.rv), nil
}

type List struct {
	lst []AST
}

func (ast List) GetType() uint8 {
	return ListT
}

func (ast List) Show(indent int) string {
	var s string
	i := indent
	for i > 0 {
		s += " "
		i--
	}
	s += "{ List: "
	if len(ast.lst) != 0 {
		s += "\n"
		for j, a := range ast.lst {
			s += a.Show(indent + 2)
			if j != len(ast.lst)-1 {
				s += "\n"
			}
		}
	}
	s += " }"
	return s
}

func (ast List) ToType() (types.Type, error) {
	lst := types.MakeList()
	if len(ast.lst) == 0 {
		return lst, nil
	}
	for i := len(ast.lst); i > 0; i-- {
		t, err := ast.lst[i-1].ToType()
		if err != nil {
			return nil, err
		}
		lst = lst.Cons(t)
	}
	return types.MakeListLiteral(lst), nil
}

type Tuple struct {
	tup []AST
}

func (ast Tuple) GetType() uint8 {
	return TupleT
}

func (ast Tuple) Show(indent int) string {
	var s string
	i := indent
	for i > 0 {
		s += " "
		i--
	}
	s += "{ Tuple: "
	if len(ast.tup) != 0 {
		s += "\n"
		for j, a := range ast.tup {
			s += a.Show(indent + 2)
			if j != len(ast.tup)-1 {
				s += "\n"
			}
		}
	}
	s += " }"
	return s
}

func (ast Tuple) ToType() (types.Type, error) {
	lst := types.MakeList()
	if len(ast.tup) == 0 {
		return types.MakeTuple(lst), nil
	}
	for i := len(ast.tup); i > 0; i-- {
		t, err := ast.tup[i-1].ToType()
		if err != nil {
			return nil, err
		}
		lst = lst.Cons(t)
	}
	return types.MakeTuple(lst), nil
}

type FunctionBody struct {
	args []AST
	body []AST
}

func (fb FunctionBody) Show(indent int) string {
	var s string
	i := indent
	for i > 0 {
		s += " "
		i--
	}
	s += "{ FunctionBody: "
	s += "\n"
	i = indent + 2
	for i > 0 {
		s += " "
		i--
	}
	s += "{ Args: "
	if len(fb.args) != 0 {
		s += "\n"
		for j, arg := range fb.args {
			s += arg.Show(indent + 4)
			if j != len(fb.args)-1 {
				s += "\n"
			}
		}
	}
	s += " }"
	s += "\n"
	i = indent + 2
	for i > 0 {
		s += " "
		i--
	}
	s += "{ Body: "
	if len(fb.body) != 0 {
		s += "\n"
		for j, bdy := range fb.body {
			s += bdy.Show(indent + 4)
			if j != len(fb.body)-1 {
				s += "\n"
			}
		}
	}
	s += " }"
	return s
}

func (fb FunctionBody) ToFunctionBody() types.FunctionBody {
	var t types.Type
	//var err error
	args := types.MakeList()
	body := types.MakeList()
	for i := 0; i < len(fb.args); i++ {
		t, _ = fb.args[i].ToType()
		args = args.Cons(t)
	}
	for i := len(fb.body); i > 0; i-- {
		t, _ = fb.body[i-1].ToType()
		body = body.Cons(t)
	}
	return types.MakeFunctionBody(args, body)
}

type Function struct {
	func_bodies []FunctionBody
}

func (ast Function) GetType() uint8 {
	return FunctionT
}

func (ast Function) Show(indent int) string {
	var s string
	i := indent
	for i > 0 {
		s += " "
		i--
	}
	s += "{ Function: "
	s += "\n"
	if len(ast.func_bodies) != 0 {
		for j, fb := range ast.func_bodies {
			s += fb.Show(indent + 2)
			if j != len(ast.func_bodies)-1 {
				s += "\n"
			}
		}
	}
	s += " }"
	return s
}

func (ast Function) ToType() (types.Type, error) {
	fbs := make([]types.FunctionBody, len(ast.func_bodies))
	for i, body := range ast.func_bodies {
		fbs[i] = body.ToFunctionBody()
	}
	return types.MakeFunction(fbs), nil // errors.New("Not a valid type")
}

type Map struct {
	m []AST
}

func (ast Map) GetType() uint8 {
	return MapT
}

func (ast Map) Show(indent int) string {
	return ""
}

func (ast Map) ToType() (types.Type, error) {
	lst := types.MakeList()
	for _,t := range ast.m {
		t2,_ := t.ToType()
		lst = lst.Cons(t2)
	}
	return types.MakeMapLiteral(lst), nil
}

type Set struct {
	m []AST
}

func (ast Set) GetType() uint8 {
	return SetT
}

func (ast Set) Show(indent int) string {
	return ""
}

func (ast Set) ToType() (types.Type, error) {
	lst := types.MakeList()
	for _,t := range ast.m {
		t2,_ := t.ToType()
		lst = lst.Cons(t2)
	}
	return types.MakeSetLiteral(lst), nil
}

type ListInject struct {
	inject AST
}

func (ast ListInject) GetType() uint8 {
	return ListInjectT
}

func (ast ListInject) Show(indent int) string {
	var s string
	i := indent
	for i > 0 {
		s += " "
		i--
	}
	s += "{ ListInject: "
	s += "\n"
	s += ast.inject.Show(indent + 2)
	s += " }"
	return s
}

func (ast ListInject) ToType() (types.Type, error) {
	t, err := ast.inject.ToType()
	if err != nil {
		return nil, err
	}
	return types.MakeListInject(t), nil
}

type At struct {
	g AST
}

func (ast At) GetType() uint8 {
	return AtT
}

func (ast At) Show(indent int) string {
	return ""
}

func (ast At) ToType() (types.Type, error) {
	t, err := ast.g.ToType()
	if err != nil {
		return nil, err
	}
	return types.MakeAt(t), nil
}

type To struct {
	s AST
}

func (ast To) GetType() uint8 {
	return ToT
}

func (ast To) Show(indent int) string {
	return ""
}

func (ast To) ToType() (types.Type, error) {
	t, err := ast.s.ToType()
	if err != nil {
		return nil, err
	}
	return types.MakeTo(t), nil
}

type Up struct {
	u AST
}

func (ast Up) GetType() uint8 {
	return UpT
}

func (ast Up) Show(indent int) string {
	return ""
}

func (ast Up) ToType() (types.Type, error) {
	t, err := ast.u.ToType()
	if err != nil {
		return nil, err
	}
	return types.MakeUp(t), nil
}

type Whatever struct {
}

func (ast Whatever) GetType() uint8 {
	return WhateverT
}

func (ast Whatever) Show(indent int) string {
	var s string
	i := indent
	for i > 0 {
		s += " "
		i--
	}
	s += "{ Whatever }"
	return s
}

func (ast Whatever) ToType() (types.Type, error) {
	return types.MakeWhatever(), nil
}

type EmptyStack struct {
}

func (ast EmptyStack) GetType() uint8 {
	return EmptyStackT
}

func (ast EmptyStack) Show(indent int) string {
	var s string
	i := indent
	for i > 0 {
		s += " "
		i--
	}
	s += "{ EmptyStack }"
	return s
}

func (ast EmptyStack) ToType() (types.Type, error) {
	return types.MakeEmptyStack(), nil
}

type Define struct {
	vrbl AST
	def  AST
}

func (ast Define) GetType() uint8 {
	return DefineT
}

func (ast Define) Show(indent int) string {
	var s string
	i := indent
	for i > 0 {
		s += " "
		i--
	}
	s += "{ Define: "
	s += "\n"
	s += ast.vrbl.Show(indent + 2)
	s += "\n"
	s += ast.def.Show(indent + 2)
	s += " }"
	return s
}

func (ast Define) ToType() (types.Type, error) {
	v,err := ast.vrbl.ToType()
	if err != nil {
		return nil, err
	}
	if v.Type() != VariableT {
		return nil, errors.New("Cannot define non-variable")
	}
	d, err := ast.def.ToType()
	if err != nil {
		return nil, err
	}
	return types.MakeDefine(v.GetName(), d), nil
}

type Parser struct {
	lexer   lexer.Lexer
	current lexer.Token
}

func (p *Parser) advance() error {
	t, err := p.lexer.Next()
	if err != nil {
		return err
	}
	p.current = t
	return nil
}

func (p *Parser) parseList() (AST, error) {
	var lst []AST
	var err error
	var ast AST = Variable{"TEMP"} // Initialised to start loop
	for !IsEOF(ast) && err == nil && p.current.Type != lexer.Terminator {
		ast, err = p.Next()
		lst = append(lst, ast)
	}
	if IsEOF(ast) {
		error_message := "ERROR - unexpected EOF"
		return List{lst}, errors.New(error_message)
	}
	if err != nil {
		return List{lst}, err
	}
	if p.current.Type == lexer.Terminator && p.current.Data != "]" {
		error_message := "ERROR - unmatched terminator: " + p.current.Data
		return List{lst}, errors.New(error_message)
	}
	p.advance()
	return List{lst}, nil
}

func (p *Parser) parseTuple() (AST, error) {
	var tup []AST
	var err error
	var ast AST = Variable{"TEMP"} // Initialised to start loop
	for !IsEOF(ast) && err == nil && p.current.Type != lexer.Terminator {
		ast, err = p.Next()
		tup = append(tup, ast)
	}
	if IsEOF(ast) {
		error_message := "ERROR - unexpected EOF"
		return Tuple{tup}, errors.New(error_message)
	}
	if err != nil {
		return Tuple{tup}, err
	}
	if p.current.Type == lexer.Terminator && p.current.Data != ")" {
		error_message := "ERROR - unmatched terminator: " + p.current.Data
		return Tuple{tup}, errors.New(error_message)
	}
	p.advance()
	return Tuple{tup}, nil
}

func (p *Parser) parseMap() (AST, error) {
	var m []AST
	var err error
	var ast AST = Variable{"TEMP"} // Initialised to start loop
	for !IsEOF(ast) && err == nil && p.current.Type != lexer.Terminator {
		ast, err = p.Next()
		m = append(m, ast)
	}
	if IsEOF(ast) {
		error_message := "ERROR - unexpected EOF"
		return Map{m}, errors.New(error_message)
	}
	if err != nil {
		return Map{m}, err
	}
	if p.current.Type == lexer.Terminator && p.current.Data != "]" {
		error_message := "ERROR - unmatched terminator: " + p.current.Data
		return Map{m}, errors.New(error_message)
	}
	p.advance()
	return Map{m}, nil
}

func (p *Parser) parseSet() (AST, error) {
	var m []AST
	var err error
	var ast AST = Variable{"TEMP"} // Initialised to start loop
	for !IsEOF(ast) && err == nil && p.current.Type != lexer.Terminator {
		ast, err = p.Next()
		m = append(m, ast)
	}
	if IsEOF(ast) {
		error_message := "ERROR - unexpected EOF"
		return Set{m}, errors.New(error_message)
	}
	if err != nil {
		return Set{m}, err
	}
	if p.current.Type == lexer.Terminator && p.current.Data != "]" {
		error_message := "ERROR - unmatched terminator: " + p.current.Data
		return Set{m}, errors.New(error_message)
	}
	p.advance()
	return Set{m}, nil
}

func (p *Parser) parseFunctionBody() (FunctionBody, bool, error) {
	var args []AST
	var body []AST
	var err error
	var ast AST = Variable{"TEMP"} // Initialised to start loop
	swaped_args := false
	for !IsEOF(ast) && err == nil && (p.current.Type != lexer.BodySeparator && p.current.Type != lexer.Terminator) {
		if p.current.Type == lexer.ArgSeparator && !swaped_args {
			swaped_args = true
			p.advance()
			args = body
			body = make([]AST, 0)
			continue
		} else if p.current.Type == lexer.ArgSeparator {
			err = errors.New("ERROR - multiple argument separators used")
			return FunctionBody{args, body}, false, err
		}
		ast, err = p.Next()
		body = append(body, ast)
	}
	if IsEOF(ast) {
		error_message := "ERROR - unexpected EOF"
		return FunctionBody{args, body}, false, errors.New(error_message)
	}
	final_blank := false
	if p.current.Type == lexer.BodySeparator {
		p.advance()
		if p.current.Type == lexer.Terminator && p.current.Data == "}" {
			final_blank = true
		}
	}
	return FunctionBody{args, body}, final_blank, err
}

func (p *Parser) parseFunction() (AST, error) {
	var bodies []FunctionBody
	var err error
	var body FunctionBody
	var final_blank bool
	for err == nil && p.current.Type != lexer.Terminator {
		body, final_blank, err = p.parseFunctionBody()
		if err != nil {
			return Function{bodies}, err
		}
		bodies = append(bodies, body)
	}
	if p.current.Type == lexer.Terminator && p.current.Data == "}" {
		p.advance()
	} else if p.current.Type == lexer.Terminator {
		err = errors.New("ERROR - unmatched terminator: " + p.current.Data)
	}
	if final_blank {
		var args []AST
		var body []AST
		fb := FunctionBody{args, body}
		bodies = append(bodies, fb)
	}
	return Function{bodies}, err
}

func (p *Parser) parseListInject() (AST, error) {
	ast, err := p.Next()
	if IsEOF(ast) {
		error_message := "ERROR - unexpeced EOF"
		err = errors.New(error_message)
	}
	return ListInject{ast}, err
}

func (p *Parser) parseAt() (AST, error) {
	ast, err := p.Next()
	if IsEOF(ast) {
		error_message := "ERROR - unexpeced EOF"
		err = errors.New(error_message)
	}
	return At{ast}, err
}

func (p *Parser) parseTo() (AST, error) {
	ast, err := p.Next()
	if IsEOF(ast) {
		error_message := "ERROR - unexpeced EOF"
		err = errors.New(error_message)
	}
	return To{ast}, err
}

func (p *Parser) parseUp() (AST, error) {
	ast, err := p.Next()
	if IsEOF(ast) {
		error_message := "ERROR - unexpeced EOF"
		err = errors.New(error_message)
	}
	return Up{ast}, err
}

func (p *Parser) parseDefine() (AST, error) {
	vrbl, err := p.Next()
	if err != nil {
		return nil, err
	}
	def, err := p.Next()
	if err != nil {
		return nil, err
	}
	err = nil
	if vrbl.GetType() != VariableT {
		error_message := "ERROR - trying to define non-variable: " + vrbl.Show(0)
		err = errors.New(error_message)
	}
	if IsEOF(vrbl) || IsEOF(def) {
		error_message := "ERROR - unexpeced EOF"
		err = errors.New(error_message)
	}
	return Define{vrbl, def}, err
}

func (p *Parser) Next() (AST, error) {
	current := p.current
	err := p.advance()
	if err != nil {
		return nil, err
	}
	switch current.Type {
	case lexer.End:
		return End{}, nil
	case lexer.Variable:
		return Variable{current.Data}, nil
	case lexer.Number:
		return Number{current.Data}, nil
	case lexer.String:
		return String{current.Data}, nil
	case lexer.Char:
		return  Char{current.Data}, nil
	case lexer.Symbol:
		return Symbol{current.Data}, nil
	case lexer.Bool:
		return Bool{current.Data == "True"}, nil
	case lexer.RestVariable:
		return RestVariable{current.Data}, nil
	case lexer.ListBegin:
		return p.parseList()
	case lexer.TupleBegin:
		return p.parseTuple()
	case lexer.FunctionBegin:
		return p.parseFunction()
	case lexer.MapBegin:
		return p.parseMap()
	case lexer.SetBegin:
		return p.parseSet()
	case lexer.ListInject:
		return p.parseListInject()
	case lexer.At:
		return p.parseAt()
	case lexer.To:
		return p.parseTo()
	case lexer.Up:
		return p.parseUp()
	case lexer.Whatever:
		return Whatever{}, nil
	case lexer.EmptyStack:
		return EmptyStack{}, nil
	case lexer.Define:
		return p.parseDefine()
	case lexer.Terminator:
		error_message := "ERROR - unmatched terminator: " + current.Data
		return nil, errors.New(error_message)
	default:
		return nil, errors.New("ERROR - unhandled token: " + current.Data)
	}
}

func (p *Parser) Parse() ([]AST, error) {
	var def_list []AST
	var ast_list []AST
	ast, err := p.Next()
	for err == nil && !IsEOF(ast) {
		if ast.GetType() == DefineT {
			def_list = append(def_list, ast)
		} else {
			ast_list = append(ast_list, ast)
		}
		ast, err = p.Next()
	}
	if err != nil {
		return nil, err
	}
	def_list = append(def_list, ast_list...)
	return def_list, nil
}

func (p *Parser) Evaluatable() (*types.List, error) {
	ast, err := p.Parse()
	if err != nil {
		return nil, err
	}
	lst := types.MakeList()
	if len(ast) == 0 {
		return lst, nil
	}
	i := len(ast)-1
	t, err := ast[i].ToType()
	for err == nil && i >= 0 {
		lst = lst.Cons(t)
		i--
		if i != -1 {
			t, err = ast[i].ToType()
		}
	}
	if err != nil {
		return nil, err
	}
	return lst, nil
}
	

func New(lexer lexer.Lexer) (Parser, error) {
	t, err := lexer.Next()
	if err != nil {
		return Parser{lexer, t}, err
	}
	return Parser{lexer, t}, nil
}

package lexer

import (
	"errors"
)

const (
	Variable uint8 = iota
	Number
	String
	Symbol
	Bool
	Char
	RestVariable
	ListBegin
	TupleBegin
	FunctionBegin
	MapBegin
	SetBegin
	Terminator
	ListInject
	At
	To
	Up
	Whatever
	EmptyStack
	ArgSeparator
	BodySeparator
	Define
	Unknown
	End
)

type Token struct {
	Type uint8
	Data string
	Line int
}

type Lexer struct {
	source              string
	position            int
	current             byte
	line                int
	valid_variable_char map[byte]bool
}

func (lex *Lexer) advance() error {
	lex.position++
	if lex.position >= len(lex.source) {
		return errors.New("EOF")
	}
	lex.current = lex.source[lex.position]
	if lex.current == '\n' {
		lex.line++
	}
	return nil
}

func (lex *Lexer) is_eof() bool {
	return lex.position >= len(lex.source)
}

func (lex *Lexer) skip_whitespace() {
	var err error
	for err == nil && is_whitespace(lex.current) {
		err = lex.advance()
	}
}

func (lex *Lexer) skip_comment() {
	for lex.current != '\n' {
		lex.advance()
	}
	lex.advance()
}

func (lex *Lexer) number() (Token, error) {
	num := Token{Number, "", lex.line}
	var err error
	for err == nil && ((lex.current >= '0' && lex.current <= '9') || lex.current == '.') {
		num.Data += string(lex.current)
		err = lex.advance()
	}
	return num, nil
}

func quote(c byte) byte {
	switch c {
	case 'n':
		return '\n'
	case 't':
		return '\t'
	case 'r':
		return '\r'
	case '\\':
		return '\\'
	default:
		return 0
	}
}

func (lex *Lexer) string() (Token, error) {
	str := Token{String, "", lex.line}
	var err error
	for err == nil && (lex.current != '"') {
		if lex.current == '\\' {
			err = lex.advance()
			quoted := quote(lex.current)
			str.Data += string(quoted)
		} else {
			str.Data += string(lex.current)
		}
		err = lex.advance()
	}
	if err != nil {
		return Token{Unknown, "", lex.line}, errors.New("Unexpected EOF")
	}
	lex.advance()
	return str, err
}

func (lex *Lexer) char() (Token, error) {
	chr := Token{Char, "", lex.line}
	var err error
	for err == nil && !is_whitespace(lex.current) {
		chr.Data += string(lex.current)
		err = lex.advance()
	}
	if err != nil {
		return Token{Unknown, "", lex.line}, errors.New("Unexpected EOF")
	}
	if chr.Data == "" {
		return Token{Unknown, "", lex.line}, errors.New("Empty Char")
	}
	return chr, err
}		

func (lex *Lexer) variable() (Token, error) {
	vrbl := Token{Variable, "", lex.line}
	if !in_set(lex.current, lex.valid_variable_char) {
		if is_whitespace(lex.current) || lex.current == ']' || lex.current == '}' || lex.current == ')' {
			return vrbl, nil
		}
		for !is_whitespace(lex.current) {
			vrbl.Data += string(lex.current)
			lex.advance()
		}
		vrbl.Type = Unknown
		return vrbl, errors.New("Invalid input")
	}
	var err error
	for err == nil && in_set(lex.current, lex.valid_variable_char) {
		vrbl.Data += string(lex.current)
		err = lex.advance()
	}
	if vrbl.Data == "def" {
		vrbl.Type = Define
		return vrbl, nil
	}
	return vrbl, nil
}

func (lex *Lexer) symbol() (Token, error) {
	t, err := lex.variable()
	t.Type = Symbol
	if t.Data == "True" || t.Data == "False" {
		t.Type = Bool
	}
	return t, err
}

func (lex *Lexer) quoted_symbol() (Token, error) {
	sym := Token{Symbol, "", lex.line}
	var err error
	for err == nil && (lex.current != '|') {
		sym.Data += string(lex.current)
		err = lex.advance()
	}
	if err != nil {
		return Token{Unknown, "", lex.line}, errors.New("Unexpected EOF")
	}
	lex.advance()
	sym.Data = "|" + sym.Data + "|"
	return sym, err
}

func (lex *Lexer) hyphen_option() (Token, error) {
	lex.advance()
	if is_whitespace(lex.current) {
		return Token{Variable, "-", lex.line}, nil
	}
	if lex.current == '-' {
		lex.advance()
		return Token{ArgSeparator, "--", lex.line}, nil
	}
	if lex.current >= '0' && lex.current <= '9' {
		t, err := lex.number()
		t.Data = "-" + t.Data
		return t, err
	} else {
		t, err := lex.variable()
		t.Data = "-" + t.Data
		return t, err
	}
}

func (lex *Lexer) underscore_option() (Token, error) {
	lex.advance()
	if is_whitespace(lex.current) {
		return Token{Whatever, "_", lex.line}, nil
	} else {
		t, err := lex.variable()
		t.Type = RestVariable
		t.Data = "_" + t.Data
		return t, err
	}
}

func (lex *Lexer) forward_slash_option() (Token, error) {
	lex.advance()
	if lex.current == '/' {
		lex.skip_comment()
		return lex.Next()
	} else {
		t, _ := lex.variable()
		t.Data = "/" + t.Data
		return t, nil
	}
}

func (lex *Lexer) hash_option() (Token, error) {
	lex.advance()
	if lex.current == '\\' {
		lex.advance()
		return lex.char()
	} else if lex.current == '[' {
		lex.advance()
		return Token{MapBegin, "", lex.line}, nil
	}
	return Token{Unknown, "", lex.line}, errors.New("Unhandled hash option")
}

func (lex *Lexer) tilde_option() (Token, error) {
	lex.advance()
	if lex.current == '[' {
		lex.advance()
		return Token{SetBegin, "", lex.line}, nil
	} else {
		t, _ := lex.variable()
		t.Data = "~" + t.Data
		return t, nil
	}
}

func (lex *Lexer) Next() (Token, error) {
	lex.skip_whitespace()
	if lex.is_eof() {
		return Token{End, "", lex.line}, nil
	}
	if lex.current >= '0' && lex.current <= '9' {
		return lex.number()
	}
	if lex.current == ']' || lex.current == ')' || lex.current == '}' {
		t := Token{Terminator, string(lex.current), lex.line}
		lex.advance()
		return t, nil
	}
	if lex.current >= 'A' && lex.current <= 'Z' {
		return lex.symbol()
	}
	switch lex.current {
	case '-':
		return lex.hyphen_option()
	case '_':
		return lex.underscore_option()
	case '/':
		return lex.forward_slash_option()
	case '#':
		return lex.hash_option()
	case '~':
		return lex.tilde_option()
	case '"':
		lex.advance()
		return lex.string()
	case '|':
		lex.advance()
		return lex.quoted_symbol()
	case '[':
		t := Token{ListBegin, "[", lex.line}
		lex.advance()
		return t, nil
	case '(':
		t := Token{TupleBegin, "(", lex.line}
		lex.advance()
		return t, nil
	case '{':
		t := Token{FunctionBegin, "{", lex.line}
		lex.advance()
		return t, nil
	case ':':
		t := Token{ListInject, ":", lex.line}
		lex.advance()
		return t, nil
	case '@':
		t := Token{At, "@", lex.line}
		lex.advance()
		return t, nil
	case '!':
		t := Token{To, "!", lex.line}
		lex.advance()
		return t, nil
	case '^':
		t := Token{Up, "^", lex.line}
		lex.advance()
		return t, nil
	case ';':
		t := Token{BodySeparator, ";", lex.line}
		lex.advance()
		return t, nil
	case '.':
		t := Token{EmptyStack, ".", lex.line}
		lex.advance()
		return t, nil
	default:
		return lex.variable()
	}
}

func (lex *Lexer) Tokenise() ([]Token, error) {
	var tokens []Token
	t, err := lex.Next()
	for err == nil && t.Type != End {
		tokens = append(tokens, t)
		t, err = lex.Next()
	}
	return tokens, err
}

func New(source string) Lexer {
	var valid_var_chars map[byte]bool
	valid_var_chars = make_char_set("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz1234567890<>/?~!£$€%^&*-_=+")
	if source == "" {
		source = " "
	}
	return Lexer{source, 0, source[0], 1, valid_var_chars}
}

func is_whitespace(char byte) bool {
	space := char == ' '
	tab := char == '\t'
	newline := char == '\n'
	ret := char == '\r'
	return space || tab || newline || ret
}

func make_char_set(chars string) map[byte]bool {
	var set = make(map[byte]bool)
	for i, _ := range chars {
		set[chars[i]] = true
	}
	return set
}

func in_set(char byte, set map[byte]bool) bool {
	_, ok := set[char]
	return ok
}

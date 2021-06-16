module main

go 1.15

replace lexer/lexer => ../lexer

replace parser/parser => ../parser

replace types/types => ../types

require (
	lexer/lexer v0.0.0-00010101000000-000000000000
	parser/parser v0.0.0-00010101000000-000000000000
	types/types v0.0.0-00010101000000-000000000000
)

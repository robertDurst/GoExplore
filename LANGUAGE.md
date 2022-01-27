## Data Language

```
<S-expression> ::= <atomic symbol> | <atomic list>
<atomic symbol> ::= <LETTER><atom part>
<atomic list>   ::= (<S-expression> ... <S-expression>)
<atom part> ::= <empty> | <LETTER><atom part> | <number><atom part>
<number> ::= 0 | 1 | 2 | ... | 9
<LETTER> ::= A | B | C | ... | Z
```

## Meta-Language

```
<form> ::= <S-expression> | 
           <variable> |
           <function>[[<form>]...[<form>]] |
           [[<form>~<form>]...[<form>~<form>]]
<variable> ::= <identifier>
<function> ::= <identifier> | 
               @[<var list> <form>] |
               label[<identifier><function>]
<var list> ::= [<variable> ... <variable>]
<argument> ::= <form>
<identifier> ::= <letter><id part>
<id part> ::= <empty> | <letter><id part> | <number><id part>
<letter> := a | b | c | ... | z
```

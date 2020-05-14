# shgo: A shell written in Go

## What' `shgo`?
`shgo` (酒豪 【ʃəgóʊ】; heavy drinker in Japanese) is a bash-like shell program written in Go for the pupose of learning the mechanisms
for the shell and the interpreter.

### shgo's BNF
```
<letter> ::= a|b|c|d|e|f|g|h|i|j|k|l|m|n|o|p|q|r|s|t|u|v|w|x|y|z|
             A|B|C|D|E|F|G|H|I|J|K|L|M|N|O|P|Q|R|S|T|U|V|W|X|Y|Z

<digit> ::= 0|1|2|3|4|5|6|7|8|9

<number> ::= <digit>
           | <number> <digit>

<word> ::= <letter>
         | <word> <letter>
         | <word> '_'
         | <word> '-'

<simple_cmd> ::= <simple_cmd_elem>
				| <simple_cmd_elem> <simple_cmd_elem>

<cmd_stmt> ::= <simple_cmd>
			| <simple_cmd> <simple_cmd>
```

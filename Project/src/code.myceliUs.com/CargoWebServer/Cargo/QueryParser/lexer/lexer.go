
// generated by gocc; DO NOT EDIT.

package lexer

import (
	// "fmt"
	"io/ioutil"
	"unicode/utf8"

	// "code.myceliUs.com/CargoWebServer/Cargo/QueryParser/util"
	"code.myceliUs.com/CargoWebServer/Cargo/QueryParser/token"
)

const(
	NoState = -1
	NumStates = 90
	NumSymbols = 136
)

type Lexer struct {
	src             []byte
	pos             int
	line            int
	column          int
}

func NewLexer(src []byte) *Lexer {
	lexer := &Lexer{
		src:    src,
		pos:    0,
		line:   1,
		column: 1,
	}
	return lexer
}

func NewLexerFile(fpath string) (*Lexer, error) {
	src, err := ioutil.ReadFile(fpath)
	if err != nil {
		return nil, err
	}
	return NewLexer(src), nil
}

func (this *Lexer) Scan() (tok *token.Token) {
	
	// fmt.Printf("Lexer.Scan() pos=%d\n", this.pos)
	
	tok = new(token.Token)
	if this.pos >= len(this.src) {
		tok.Type = token.EOF
		tok.Pos.Offset, tok.Pos.Line, tok.Pos.Column = this.pos, this.line, this.column
		return
	}
	start, startLine, startColumn, end := this.pos, this.line, this.column, 0
	tok.Type = token.INVALID
	state, rune1, size := 0, rune(-1), 0
	for state != -1 {
	
		// fmt.Printf("\tpos=%d, line=%d, col=%d, state=%d\n", this.pos, this.line, this.column, state)
	
		if this.pos >= len(this.src) {
			rune1 = -1
		} else {
			rune1, size = utf8.DecodeRune(this.src[this.pos:])
			this.pos += size
		}

	
		// Production start
		if rune1 != -1 {
			state = TransTab[state](rune1)
		} else {
			state = -1
		}
		// Production end

		// Debug start
		// nextState := -1
		// if rune1 != -1 {
		// 	nextState = TransTab[state](rune1)
		// }
		// fmt.Printf("\tS%d, : tok=%s, rune == %s(%x), next state == %d\n", state, token.TokMap.Id(tok.Type), util.RuneToString(rune1), rune1, nextState)
		// fmt.Printf("\t\tpos=%d, size=%d, start=%d, end=%d\n", this.pos, size, start, end)
		// if nextState != -1 {
		// 	fmt.Printf("\t\taction:%s\n", ActTab[nextState].String())
		// }
		// state = nextState
		// Debug end
	

		if state != -1 {

			switch rune1 {
			case '\n':
				this.line++
				this.column = 1
			case '\r':
				this.column = 1
			case '\t':
				this.column += 4
			default:
				this.column++
			}

			switch {
			case ActTab[state].Accept != -1:
				tok.Type = ActTab[state].Accept
				// fmt.Printf("\t Accept(%s), %s(%d)\n", string(act), token.TokMap.Id(tok), tok)
				end = this.pos
			case ActTab[state].Ignore != "":
				// fmt.Printf("\t Ignore(%s)\n", string(act))
				start, startLine, startColumn = this.pos, this.line, this.column
				state = 0
				if start >= len(this.src) {
					tok.Type = token.EOF
				}

			}
		} else {
			if tok.Type == token.INVALID {
				end = this.pos
			}
		}
	}
	if end > start {
		this.pos = end
		tok.Lit = this.src[start:end]
	} else {
		tok.Lit = []byte{}
	}
	tok.Pos.Offset, tok.Pos.Line ,tok.Pos.Column = start, startLine, startColumn

	
	// fmt.Printf("Token at %s: %s \"%s\"\n", tok.String(), token.TokMap.Id(tok.Type), tok.Lit)
	

	return
}

func (this *Lexer) Reset() {
	this.pos = 0
}

/*
Lexer symbols:
0: '-'
1: '-'
2: '.'
3: '"'
4: ' '
5: '"'
6: '/'
7: '\'
8: '|'
9: ' '
10: '/'
11: 'n'
12: 'u'
13: 'l'
14: 'l'
15: 't'
16: 'r'
17: 'u'
18: 'e'
19: 'f'
20: 'a'
21: 'l'
22: 's'
23: 'e'
24: '_'
25: '_'
26: '!'
27: '='
28: '='
29: '='
30: '~'
31: '='
32: '>'
33: '<'
34: '>'
35: '='
36: '<'
37: '='
38: '^'
39: '='
40: '$'
41: '='
42: '&'
43: '&'
44: '|'
45: '|'
46: '('
47: ')'
48: '.'
49: 'e'
50: 'E'
51: '+'
52: '-'
53: '\'
54: '\'
55: '\'
56: '\'
57: 'u'
58: '\'
59: 'b'
60: 't'
61: 'n'
62: 'f'
63: 'r'
64: '"'
65: '''
66: '\'
67: '\'
68: '{'
69: '}'
70: '('
71: ')'
72: '['
73: ']'
74: '/'
75: '\'
76: '|'
77: \u00c9
78: \u00e9
79: \u00e8
80: \u00c8
81: \u00ea
82: \u00ca
83: \u00e0
84: \u00c0
85: \u00c7
86: \u00eb
87: \u00cb
88: \u00d4
89: \u00f4
90: '?'
91: '!'
92: '.'
93: ','
94: ':'
95: ';'
96: '`'
97: '_'
98: '~'
99: '@'
100: '#'
101: '`'
102: '^'
103: '%'
104: '-'
105: '+'
106: '*'
107: '$'
108: '='
109: '<'
110: '>'
111: '!'
112: \u2019
113: '\t'
114: '\n'
115: '\r'
116: ' '
117: '1'-'9'
118: 'a'-'z'
119: 'A'-'Z'
120: 'a'-'z'
121: 'A'-'Z'
122: '0'-'9'
123: '0'-'9'
124: '0'-'9'
125: 'a'-'f'
126: 'A'-'F'
127: '0'-'3'
128: '0'-'7'
129: '0'-'7'
130: '0'-'7'
131: '0'-'7'
132: '0'-'7'
133: 'a'-'z'
134: 'A'-'Z'
135: .

*/

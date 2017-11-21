package main

//go:generate peg -switch -inline urn.peg

import (
	"fmt"
	"math"
	"sort"
	"strconv"
)

const endSymbol rune = 1114112

/* The rule types inferred from the grammar are below. */
type pegRule uint8

const (
	ruleUnknown pegRule = iota
	ruleURN
	ruleURN_PREFIX
	ruleNID
	ruleNSS
	ruleLET_NUM
	ruleLET_NUM_HYP
	ruleCHARS
	ruleTRANS
	ruleHEX
	ruleOTHER
	ruleRESERVED
	rulecolon
	ruleeot
	ruleupper
	rulelower
	rulenumber
	rulehyp
	ruleperc
)

var rul3s = [...]string{
	"Unknown",
	"URN",
	"URN_PREFIX",
	"NID",
	"NSS",
	"LET_NUM",
	"LET_NUM_HYP",
	"CHARS",
	"TRANS",
	"HEX",
	"OTHER",
	"RESERVED",
	"colon",
	"eot",
	"upper",
	"lower",
	"number",
	"hyp",
	"perc",
}

type token32 struct {
	pegRule
	begin, end uint32
}

func (t *token32) String() string {
	return fmt.Sprintf("\x1B[34m%v\x1B[m %v %v", rul3s[t.pegRule], t.begin, t.end)
}

type node32 struct {
	token32
	up, next *node32
}

func (node *node32) print(pretty bool, buffer string) {
	var print func(node *node32, depth int)
	print = func(node *node32, depth int) {
		for node != nil {
			for c := 0; c < depth; c++ {
				fmt.Printf(" ")
			}
			rule := rul3s[node.pegRule]
			quote := strconv.Quote(string(([]rune(buffer)[node.begin:node.end])))
			if !pretty {
				fmt.Printf("%v %v\n", rule, quote)
			} else {
				fmt.Printf("\x1B[34m%v\x1B[m %v\n", rule, quote)
			}
			if node.up != nil {
				print(node.up, depth+1)
			}
			node = node.next
		}
	}
	print(node, 0)
}

func (node *node32) Print(buffer string) {
	node.print(false, buffer)
}

func (node *node32) PrettyPrint(buffer string) {
	node.print(true, buffer)
}

type tokens32 struct {
	tree []token32
}

func (t *tokens32) Trim(length uint32) {
	t.tree = t.tree[:length]
}

func (t *tokens32) Print() {
	for _, token := range t.tree {
		fmt.Println(token.String())
	}
}

func (t *tokens32) AST() *node32 {
	type element struct {
		node *node32
		down *element
	}
	tokens := t.Tokens()
	var stack *element
	for _, token := range tokens {
		if token.begin == token.end {
			continue
		}
		node := &node32{token32: token}
		for stack != nil && stack.node.begin >= token.begin && stack.node.end <= token.end {
			stack.node.next = node.up
			node.up = stack.node
			stack = stack.down
		}
		stack = &element{node: node, down: stack}
	}
	if stack != nil {
		return stack.node
	}
	return nil
}

func (t *tokens32) PrintSyntaxTree(buffer string) {
	t.AST().Print(buffer)
}

func (t *tokens32) PrettyPrintSyntaxTree(buffer string) {
	t.AST().PrettyPrint(buffer)
}

func (t *tokens32) Add(rule pegRule, begin, end, index uint32) {
	if tree := t.tree; int(index) >= len(tree) {
		expanded := make([]token32, 2*len(tree))
		copy(expanded, tree)
		t.tree = expanded
	}
	t.tree[index] = token32{
		pegRule: rule,
		begin:   begin,
		end:     end,
	}
}

func (t *tokens32) Tokens() []token32 {
	return t.tree
}

type URN struct {
	Buffer string
	buffer []rune
	rules  [19]func() bool
	parse  func(rule ...int) error
	reset  func()
	Pretty bool
	tokens32
}

func (p *URN) Parse(rule ...int) error {
	return p.parse(rule...)
}

func (p *URN) Reset() {
	p.reset()
}

type textPosition struct {
	line, symbol int
}

type textPositionMap map[int]textPosition

func translatePositions(buffer []rune, positions []int) textPositionMap {
	length, translations, j, line, symbol := len(positions), make(textPositionMap, len(positions)), 0, 1, 0
	sort.Ints(positions)

search:
	for i, c := range buffer {
		if c == '\n' {
			line, symbol = line+1, 0
		} else {
			symbol++
		}
		if i == positions[j] {
			translations[positions[j]] = textPosition{line, symbol}
			for j++; j < length; j++ {
				if i != positions[j] {
					continue search
				}
			}
			break search
		}
	}

	return translations
}

type parseError struct {
	p   *URN
	max token32
}

func (e *parseError) Error() string {
	tokens, error := []token32{e.max}, "\n"
	positions, p := make([]int, 2*len(tokens)), 0
	for _, token := range tokens {
		positions[p], p = int(token.begin), p+1
		positions[p], p = int(token.end), p+1
	}
	translations := translatePositions(e.p.buffer, positions)
	format := "parse error near %v (line %v symbol %v - line %v symbol %v):\n%v\n"
	if e.p.Pretty {
		format = "parse error near \x1B[34m%v\x1B[m (line %v symbol %v - line %v symbol %v):\n%v\n"
	}
	for _, token := range tokens {
		begin, end := int(token.begin), int(token.end)
		error += fmt.Sprintf(format,
			rul3s[token.pegRule],
			translations[begin].line, translations[begin].symbol,
			translations[end].line, translations[end].symbol,
			strconv.Quote(string(e.p.buffer[begin:end])))
	}

	return error
}

func (p *URN) PrintSyntaxTree() {
	if p.Pretty {
		p.tokens32.PrettyPrintSyntaxTree(p.Buffer)
	} else {
		p.tokens32.PrintSyntaxTree(p.Buffer)
	}
}

func (p *URN) Init() {
	var (
		max                  token32
		position, tokenIndex uint32
		buffer               []rune
	)
	p.reset = func() {
		max = token32{}
		position, tokenIndex = 0, 0

		p.buffer = []rune(p.Buffer)
		if len(p.buffer) == 0 || p.buffer[len(p.buffer)-1] != endSymbol {
			p.buffer = append(p.buffer, endSymbol)
		}
		buffer = p.buffer
	}
	p.reset()

	_rules := p.rules
	tree := tokens32{tree: make([]token32, math.MaxInt16)}
	p.parse = func(rule ...int) error {
		r := 1
		if len(rule) > 0 {
			r = rule[0]
		}
		matches := p.rules[r]()
		p.tokens32 = tree
		if matches {
			p.Trim(tokenIndex)
			return nil
		}
		return &parseError{p, max}
	}

	add := func(rule pegRule, begin uint32) {
		tree.Add(rule, begin, position, tokenIndex)
		tokenIndex++
		if begin != position && position > max.end {
			max = token32{rule, begin, position}
		}
	}

	matchDot := func() bool {
		if buffer[position] != endSymbol {
			position++
			return true
		}
		return false
	}

	/*matchChar := func(c byte) bool {
		if buffer[position] == c {
			position++
			return true
		}
		return false
	}*/

	/*matchRange := func(lower byte, upper byte) bool {
		if c := buffer[position]; c >= lower && c <= upper {
			position++
			return true
		}
		return false
	}*/

	_rules = [...]func() bool{
		nil,
		/* 0 URN <- <(URN_PREFIX NID colon NSS eot)> */
		func() bool {
			position0, tokenIndex0 := position, tokenIndex
			{
				position1 := position
				{
					position2 := position
					{
						position3, tokenIndex3 := position, tokenIndex
						if buffer[position] != rune('u') {
							goto l4
						}
						position++
						goto l3
					l4:
						position, tokenIndex = position3, tokenIndex3
						if buffer[position] != rune('U') {
							goto l0
						}
						position++
					}
				l3:
					{
						position5, tokenIndex5 := position, tokenIndex
						if buffer[position] != rune('r') {
							goto l6
						}
						position++
						goto l5
					l6:
						position, tokenIndex = position5, tokenIndex5
						if buffer[position] != rune('R') {
							goto l0
						}
						position++
					}
				l5:
					{
						position7, tokenIndex7 := position, tokenIndex
						if buffer[position] != rune('n') {
							goto l8
						}
						position++
						goto l7
					l8:
						position, tokenIndex = position7, tokenIndex7
						if buffer[position] != rune('N') {
							goto l0
						}
						position++
					}
				l7:
					if !_rules[rulecolon]() {
						goto l0
					}
					add(ruleURN_PREFIX, position2)
				}
				{
					position9 := position
					if !_rules[ruleLET_NUM]() {
						goto l0
					}
					{
						position10, tokenIndex10 := position, tokenIndex
						{
							position12 := position
							{
								position13, tokenIndex13 := position, tokenIndex
								if !_rules[ruleLET_NUM]() {
									goto l14
								}
								goto l13
							l14:
								position, tokenIndex = position13, tokenIndex13
								if !_rules[rulehyp]() {
									goto l10
								}
							}
						l13:
							add(ruleLET_NUM_HYP, position12)
						}
						goto l11
					l10:
						position, tokenIndex = position10, tokenIndex10
					}
				l11:
					add(ruleNID, position9)
				}
				if !_rules[rulecolon]() {
					goto l0
				}
				{
					position15 := position
					{
						position18 := position
						{
							position19, tokenIndex19 := position, tokenIndex
							{
								position21 := position
								{
									switch buffer[position] {
									case '#', '%', '/', '?':
										{
											position23 := position
											{
												switch buffer[position] {
												case '#':
													if buffer[position] != rune('#') {
														goto l20
													}
													position++
													break
												case '?':
													if buffer[position] != rune('?') {
														goto l20
													}
													position++
													break
												case '/':
													if buffer[position] != rune('/') {
														goto l20
													}
													position++
													break
												default:
													if !_rules[ruleperc]() {
														goto l20
													}
													break
												}
											}

											add(ruleRESERVED, position23)
										}
										break
									case '!', '$', '\'', '(', ')', '*', '+', ',', '-', '.', ':', ';', '=', '@', '_':
										{
											position25 := position
											{
												switch buffer[position] {
												case '\'':
													if buffer[position] != rune('\'') {
														goto l20
													}
													position++
													break
												case '*':
													if buffer[position] != rune('*') {
														goto l20
													}
													position++
													break
												case '!':
													if buffer[position] != rune('!') {
														goto l20
													}
													position++
													break
												case '_':
													if buffer[position] != rune('_') {
														goto l20
													}
													position++
													break
												case '$':
													if buffer[position] != rune('$') {
														goto l20
													}
													position++
													break
												case ';':
													if buffer[position] != rune(';') {
														goto l20
													}
													position++
													break
												case '@':
													if buffer[position] != rune('@') {
														goto l20
													}
													position++
													break
												case '=':
													if buffer[position] != rune('=') {
														goto l20
													}
													position++
													break
												case ':':
													if !_rules[rulecolon]() {
														goto l20
													}
													break
												case '.':
													if buffer[position] != rune('.') {
														goto l20
													}
													position++
													break
												case '-':
													if !_rules[rulehyp]() {
														goto l20
													}
													break
												case ',':
													if buffer[position] != rune(',') {
														goto l20
													}
													position++
													break
												case '+':
													if buffer[position] != rune('+') {
														goto l20
													}
													position++
													break
												case ')':
													if buffer[position] != rune(')') {
														goto l20
													}
													position++
													break
												default:
													if buffer[position] != rune('(') {
														goto l20
													}
													position++
													break
												}
											}

											add(ruleOTHER, position25)
										}
										break
									default:
										if !_rules[ruleLET_NUM]() {
											goto l20
										}
										break
									}
								}

								add(ruleTRANS, position21)
							}
							goto l19
						l20:
							position, tokenIndex = position19, tokenIndex19
							if !_rules[ruleperc]() {
								goto l0
							}
							if !_rules[ruleHEX]() {
								goto l0
							}
							if !_rules[ruleHEX]() {
								goto l0
							}
						}
					l19:
						add(ruleCHARS, position18)
					}
				l16:
					{
						position17, tokenIndex17 := position, tokenIndex
						{
							position27 := position
							{
								position28, tokenIndex28 := position, tokenIndex
								{
									position30 := position
									{
										switch buffer[position] {
										case '#', '%', '/', '?':
											{
												position32 := position
												{
													switch buffer[position] {
													case '#':
														if buffer[position] != rune('#') {
															goto l29
														}
														position++
														break
													case '?':
														if buffer[position] != rune('?') {
															goto l29
														}
														position++
														break
													case '/':
														if buffer[position] != rune('/') {
															goto l29
														}
														position++
														break
													default:
														if !_rules[ruleperc]() {
															goto l29
														}
														break
													}
												}

												add(ruleRESERVED, position32)
											}
											break
										case '!', '$', '\'', '(', ')', '*', '+', ',', '-', '.', ':', ';', '=', '@', '_':
											{
												position34 := position
												{
													switch buffer[position] {
													case '\'':
														if buffer[position] != rune('\'') {
															goto l29
														}
														position++
														break
													case '*':
														if buffer[position] != rune('*') {
															goto l29
														}
														position++
														break
													case '!':
														if buffer[position] != rune('!') {
															goto l29
														}
														position++
														break
													case '_':
														if buffer[position] != rune('_') {
															goto l29
														}
														position++
														break
													case '$':
														if buffer[position] != rune('$') {
															goto l29
														}
														position++
														break
													case ';':
														if buffer[position] != rune(';') {
															goto l29
														}
														position++
														break
													case '@':
														if buffer[position] != rune('@') {
															goto l29
														}
														position++
														break
													case '=':
														if buffer[position] != rune('=') {
															goto l29
														}
														position++
														break
													case ':':
														if !_rules[rulecolon]() {
															goto l29
														}
														break
													case '.':
														if buffer[position] != rune('.') {
															goto l29
														}
														position++
														break
													case '-':
														if !_rules[rulehyp]() {
															goto l29
														}
														break
													case ',':
														if buffer[position] != rune(',') {
															goto l29
														}
														position++
														break
													case '+':
														if buffer[position] != rune('+') {
															goto l29
														}
														position++
														break
													case ')':
														if buffer[position] != rune(')') {
															goto l29
														}
														position++
														break
													default:
														if buffer[position] != rune('(') {
															goto l29
														}
														position++
														break
													}
												}

												add(ruleOTHER, position34)
											}
											break
										default:
											if !_rules[ruleLET_NUM]() {
												goto l29
											}
											break
										}
									}

									add(ruleTRANS, position30)
								}
								goto l28
							l29:
								position, tokenIndex = position28, tokenIndex28
								if !_rules[ruleperc]() {
									goto l17
								}
								if !_rules[ruleHEX]() {
									goto l17
								}
								if !_rules[ruleHEX]() {
									goto l17
								}
							}
						l28:
							add(ruleCHARS, position27)
						}
						goto l16
					l17:
						position, tokenIndex = position17, tokenIndex17
					}
					add(ruleNSS, position15)
				}
				{
					position36 := position
					{
						position37, tokenIndex37 := position, tokenIndex
						if !matchDot() {
							goto l37
						}
						goto l0
					l37:
						position, tokenIndex = position37, tokenIndex37
					}
					add(ruleeot, position36)
				}
				add(ruleURN, position1)
			}
			return true
		l0:
			position, tokenIndex = position0, tokenIndex0
			return false
		},
		/* 1 URN_PREFIX <- <(('u' / 'U') ('r' / 'R') ('n' / 'N') colon)> */
		nil,
		/* 2 NID <- <(LET_NUM LET_NUM_HYP?)> */
		nil,
		/* 3 NSS <- <CHARS+> */
		nil,
		/* 4 LET_NUM <- <((&('0' | '1' | '2' | '3' | '4' | '5' | '6' | '7' | '8' | '9') number) | (&('a' | 'b' | 'c' | 'd' | 'e' | 'f' | 'g' | 'h' | 'i' | 'j' | 'k' | 'l' | 'm' | 'n' | 'o' | 'p' | 'q' | 'r' | 's' | 't' | 'u' | 'v' | 'w' | 'x' | 'y' | 'z') lower) | (&('A' | 'B' | 'C' | 'D' | 'E' | 'F' | 'G' | 'H' | 'I' | 'J' | 'K' | 'L' | 'M' | 'N' | 'O' | 'P' | 'Q' | 'R' | 'S' | 'T' | 'U' | 'V' | 'W' | 'X' | 'Y' | 'Z') upper))> */
		func() bool {
			position41, tokenIndex41 := position, tokenIndex
			{
				position42 := position
				{
					switch buffer[position] {
					case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
						if !_rules[rulenumber]() {
							goto l41
						}
						break
					case 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z':
						{
							position44 := position
							if c := buffer[position]; c < rune('a') || c > rune('z') {
								goto l41
							}
							position++
						l45:
							{
								position46, tokenIndex46 := position, tokenIndex
								if c := buffer[position]; c < rune('a') || c > rune('z') {
									goto l46
								}
								position++
								goto l45
							l46:
								position, tokenIndex = position46, tokenIndex46
							}
							add(rulelower, position44)
						}
						break
					default:
						{
							position47 := position
							if c := buffer[position]; c < rune('A') || c > rune('Z') {
								goto l41
							}
							position++
						l48:
							{
								position49, tokenIndex49 := position, tokenIndex
								if c := buffer[position]; c < rune('A') || c > rune('Z') {
									goto l49
								}
								position++
								goto l48
							l49:
								position, tokenIndex = position49, tokenIndex49
							}
							add(ruleupper, position47)
						}
						break
					}
				}

				add(ruleLET_NUM, position42)
			}
			return true
		l41:
			position, tokenIndex = position41, tokenIndex41
			return false
		},
		/* 5 LET_NUM_HYP <- <(LET_NUM / hyp)> */
		nil,
		/* 6 CHARS <- <(TRANS / (perc HEX HEX))> */
		nil,
		/* 7 TRANS <- <((&('#' | '%' | '/' | '?') RESERVED) | (&('!' | '$' | '\'' | '(' | ')' | '*' | '+' | ',' | '-' | '.' | ':' | ';' | '=' | '@' | '_') OTHER) | (&('0' | '1' | '2' | '3' | '4' | '5' | '6' | '7' | '8' | '9' | 'A' | 'B' | 'C' | 'D' | 'E' | 'F' | 'G' | 'H' | 'I' | 'J' | 'K' | 'L' | 'M' | 'N' | 'O' | 'P' | 'Q' | 'R' | 'S' | 'T' | 'U' | 'V' | 'W' | 'X' | 'Y' | 'Z' | 'a' | 'b' | 'c' | 'd' | 'e' | 'f' | 'g' | 'h' | 'i' | 'j' | 'k' | 'l' | 'm' | 'n' | 'o' | 'p' | 'q' | 'r' | 's' | 't' | 'u' | 'v' | 'w' | 'x' | 'y' | 'z') LET_NUM))> */
		nil,
		/* 8 HEX <- <((&('a' | 'b' | 'c' | 'd' | 'e' | 'f') [a-f]+) | (&('A' | 'B' | 'C' | 'D' | 'E' | 'F') [A-F]+) | (&('0' | '1' | '2' | '3' | '4' | '5' | '6' | '7' | '8' | '9') number))> */
		func() bool {
			position53, tokenIndex53 := position, tokenIndex
			{
				position54 := position
				{
					switch buffer[position] {
					case 'a', 'b', 'c', 'd', 'e', 'f':
						if c := buffer[position]; c < rune('a') || c > rune('f') {
							goto l53
						}
						position++
					l56:
						{
							position57, tokenIndex57 := position, tokenIndex
							if c := buffer[position]; c < rune('a') || c > rune('f') {
								goto l57
							}
							position++
							goto l56
						l57:
							position, tokenIndex = position57, tokenIndex57
						}
						break
					case 'A', 'B', 'C', 'D', 'E', 'F':
						if c := buffer[position]; c < rune('A') || c > rune('F') {
							goto l53
						}
						position++
					l58:
						{
							position59, tokenIndex59 := position, tokenIndex
							if c := buffer[position]; c < rune('A') || c > rune('F') {
								goto l59
							}
							position++
							goto l58
						l59:
							position, tokenIndex = position59, tokenIndex59
						}
						break
					default:
						if !_rules[rulenumber]() {
							goto l53
						}
						break
					}
				}

				add(ruleHEX, position54)
			}
			return true
		l53:
			position, tokenIndex = position53, tokenIndex53
			return false
		},
		/* 9 OTHER <- <((&('\'') '\'') | (&('*') '*') | (&('!') '!') | (&('_') '_') | (&('$') '$') | (&(';') ';') | (&('@') '@') | (&('=') '=') | (&(':') colon) | (&('.') '.') | (&('-') hyp) | (&(',') ',') | (&('+') '+') | (&(')') ')') | (&('(') '('))> */
		nil,
		/* 10 RESERVED <- <((&('#') '#') | (&('?') '?') | (&('/') '/') | (&('%') perc))> */
		nil,
		/* 11 colon <- <':'> */
		func() bool {
			position62, tokenIndex62 := position, tokenIndex
			{
				position63 := position
				if buffer[position] != rune(':') {
					goto l62
				}
				position++
				add(rulecolon, position63)
			}
			return true
		l62:
			position, tokenIndex = position62, tokenIndex62
			return false
		},
		/* 12 eot <- <!.> */
		nil,
		/* 13 upper <- <[A-Z]+> */
		nil,
		/* 14 lower <- <[a-z]+> */
		nil,
		/* 15 number <- <[0-9]+> */
		func() bool {
			position67, tokenIndex67 := position, tokenIndex
			{
				position68 := position
				if c := buffer[position]; c < rune('0') || c > rune('9') {
					goto l67
				}
				position++
			l69:
				{
					position70, tokenIndex70 := position, tokenIndex
					if c := buffer[position]; c < rune('0') || c > rune('9') {
						goto l70
					}
					position++
					goto l69
				l70:
					position, tokenIndex = position70, tokenIndex70
				}
				add(rulenumber, position68)
			}
			return true
		l67:
			position, tokenIndex = position67, tokenIndex67
			return false
		},
		/* 16 hyp <- <'-'> */
		func() bool {
			position71, tokenIndex71 := position, tokenIndex
			{
				position72 := position
				if buffer[position] != rune('-') {
					goto l71
				}
				position++
				add(rulehyp, position72)
			}
			return true
		l71:
			position, tokenIndex = position71, tokenIndex71
			return false
		},
		/* 17 perc <- <'%'> */
		func() bool {
			position73, tokenIndex73 := position, tokenIndex
			{
				position74 := position
				if buffer[position] != rune('%') {
					goto l73
				}
				position++
				add(ruleperc, position74)
			}
			return true
		l73:
			position, tokenIndex = position73, tokenIndex73
			return false
		},
	}
	p.rules = _rules
}

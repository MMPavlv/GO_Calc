package Grammar

import (
	"fmt"
	"strconv"
)

type Triplet struct {
	l *Token
	m *Token
	r *Token
}

func NewTriplet() *Triplet {
	return &Triplet{l: new(Token), m: new(Token), r: new(Token)}
}

func (t *Triplet) SetTriplet(l *Token, m *Token, r *Token) {
	t.l = l
	t.m = m
	t.r = r
}

type Tokenizer struct {
	tokens []Token
}

func (t *Tokenizer) uParse(importString string) {
	t.tokens = append(t.tokens, &TokenOperator{BEGIN})

	tempString := ""
	minusOp := false
	for _, symbol := range importString {
		if symbol == '(' {
			if len(tempString) > 0 {
				parsedFloat, parsedError := strconv.ParseFloat(tempString, 64)
				if parsedError == nil {
					t.tokens = append(t.tokens, &TokenOperand{Number, parsedFloat})
				} else {
					panic(fmt.Sprintf("WRONG NUMBER: %s", tempString))
				}

				tempString = ""
			}
			t.tokens = append(t.tokens, &TokenOperator{LParentheses})
			minusOp = false
		} else if symbol == ')' {
			if len(tempString) > 0 {
				parsedFloat, parsedError := strconv.ParseFloat(tempString, 64)
				if parsedError == nil {
					t.tokens = append(t.tokens, &TokenOperand{Number, parsedFloat})
				} else {
					panic(fmt.Sprintf("WRONG NUMBER: %s", tempString))
				}

				tempString = ""
			}
			t.tokens = append(t.tokens, &TokenOperator{RParentheses})
			minusOp = true
		} else if symbol == '+' {
			if len(tempString) > 0 {
				parsedFloat, parsedError := strconv.ParseFloat(tempString, 64)
				if parsedError == nil {
					t.tokens = append(t.tokens, &TokenOperand{Number, parsedFloat})
				} else {
					panic(fmt.Sprintf("WRONG NUMBER: %s", tempString))
				}

				tempString = ""
			}
			t.tokens = append(t.tokens, &TokenOperator{Plus})
			minusOp = false
		} else if symbol == '-' {
			if len(tempString) > 0 {
				parsedFloat, parsedError := strconv.ParseFloat(tempString, 64)
				if parsedError == nil {
					t.tokens = append(t.tokens, &TokenOperand{Number, parsedFloat})
					t.tokens = append(t.tokens, &TokenOperator{Minus})
				} else {
					panic(fmt.Sprintf("WRONG NUMBER: %s", tempString))
				}

				tempString = ""
			} else {
				if minusOp {
					t.tokens = append(t.tokens, &TokenOperator{Minus})
				} else {
					tempString += string(symbol)
				}
			}
			minusOp = false
		} else if symbol == '*' {
			if len(tempString) > 0 {
				parsedFloat, parsedError := strconv.ParseFloat(tempString, 64)
				if parsedError == nil {
					t.tokens = append(t.tokens, &TokenOperand{Number, parsedFloat})
				} else {
					panic(fmt.Sprintf("WRONG NUMBER: %s", tempString))
				}

				tempString = ""
			}
			t.tokens = append(t.tokens, &TokenOperator{Multi})
			minusOp = false
		} else if symbol == '/' {
			if len(tempString) > 0 {
				parsedFloat, parsedError := strconv.ParseFloat(tempString, 64)
				if parsedError == nil {
					t.tokens = append(t.tokens, &TokenOperand{Number, parsedFloat})
				} else {
					panic(fmt.Sprintf("WRONG NUMBER: %s", tempString))
				}

				tempString = ""
			}
			t.tokens = append(t.tokens, &TokenOperator{Divide})
			minusOp = false
		} else if symbol == '.' {
			tempString += string(symbol)
			minusOp = true
		} else if symbol >= '0' && symbol <= '9' {
			tempString += string(symbol)
			minusOp = true
		} else {
			panic(string(symbol))
		}
	}
	if len(tempString) > 0 {
		parsedFloat, parsedError := strconv.ParseFloat(tempString, 64)
		if parsedError == nil {
			t.tokens = append(t.tokens, &TokenOperand{Number, parsedFloat})
		} else {
			panic(fmt.Sprintf("WRONG NUMBER: %s", tempString))
		}
	}

	t.tokens = append(t.tokens, &TokenOperator{END})
}

func (t *Triplet) Solve() Token {
	num := TokenOperand{
		Info: Number,
	}

	left, _ := strconv.ParseFloat((*t.l).GetString(), 64)
	right, _ := strconv.ParseFloat((*t.r).GetString(), 64)

	switch (*t.m).GetType() {
	case Plus:
		num.Number = left+right
	case Minus:
		num.Number = left-right
	case Multi:
		num.Number = left*right
	case Divide:
		num.Number = left/right
	}

	return &num
}

func (t *Tokenizer) Parse(s string) string {
	t.uParse(s)

	if len(t.tokens) == 2 {
		panic("EMPTY EXPRESSION")
	}
	localTriplet := Triplet{}
	for ; len(t.tokens) != 1; {
		for i, symbol := range t.tokens {
			if localTriplet.l == nil {
				localTriplet.l = new(Token)
				*localTriplet.l = symbol
				continue
			}

			if localTriplet.m == nil {
				localTriplet.m = new(Token)
				*localTriplet.m = symbol
				continue
			}

			if localTriplet.r == nil {
				localTriplet.r = new(Token)
				*localTriplet.r = symbol
				continue
			}

			if (*localTriplet.l).GetType() == BEGIN &&
				(*localTriplet.m).GetType() == Number &&
				(*localTriplet.r).GetType() == END {
				t.tokens = nil
				t.tokens = append(t.tokens, *localTriplet.m)
				break
			}

			if symbol.GetType() == END || (symbol.GetType() >= RParentheses && symbol.GetType() <= Divide) {
				if (*localTriplet.l).GetType() == Number &&
					(*localTriplet.r).GetType() == Number &&
					((*localTriplet.m).GetType() >= Plus && (*localTriplet.m).GetType() <= Divide) {
						if (symbol.GetType() == Multi || symbol.GetType() == Divide) &&
							((*localTriplet.m).GetType() == Plus || (*localTriplet.m).GetType() == Minus ) {
							*localTriplet.l = *localTriplet.m
							*localTriplet.m = *localTriplet.r
							*localTriplet.r = symbol
							continue
						}
						before := t.tokens[0:i-3]
						after := t.tokens[i:len(t.tokens)]
						t.tokens = append(before, localTriplet.Solve())
						t.tokens = append(t.tokens, after...)
						localTriplet = Triplet{}
						break
				} else if (*localTriplet.l).GetType() == LParentheses &&
					(*localTriplet.m).GetType() == Number &&
					(*localTriplet.r).GetType() == RParentheses {
					before := t.tokens[0:i-3]
					after := t.tokens[i:len(t.tokens)]
					t.tokens = append(before, *localTriplet.m)
					t.tokens = append(t.tokens, after...)
					localTriplet = Triplet{}
					break
				} else if (*localTriplet.r).GetType() == LParentheses ||
					(*localTriplet.r).GetType() >= Plus {
					panic("WRONG EXPRESSION")
				}
			}

			*localTriplet.l = *localTriplet.m
			*localTriplet.m = *localTriplet.r
			*localTriplet.r = symbol
		}
	}

	return t.tokens[0].GetString()
}

func (t *Tokenizer) GetTokens() []Token {
	return t.tokens
}
package Grammar

import (
	"fmt"
)


type Token interface {
	GetType() uint8
	GetString() string
}

type TokenOperator struct {
	Info uint8
}

func (t *TokenOperator) GetType() uint8 {
	return t.Info
}

func (t *TokenOperator) GetString() string {
	switch t.Info {
	case LParentheses: return "("
	case RParentheses: return ")"
	case Plus: return "+"
	case Minus: return "-"
	case Multi: return "*"
	case Divide: return "/"
	default: return "ERROR"
	}
}

type TokenOperand struct {
	Info uint8
	Number float64
}

func (t *TokenOperand) GetType() uint8 {
	return t.Info
}

func (t *TokenOperand) GetString() string {
	return fmt.Sprintf("%f", t.Number)
}
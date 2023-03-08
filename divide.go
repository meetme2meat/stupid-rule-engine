package main

import "fmt"

func divBuilder(rule interface{}) Valuer {
	unaryRules, ok := rule.([]interface{})
	if !ok {
		panic("not a valid divide rule")
	}

	if len(unaryRules) != 2 {
		panic("divide must be 2 operand")
	}

	plus := new(divideOperator)
	plus.leftValue = buildValue(unaryRules[0])
	plus.rightValue = buildValue(unaryRules[1])
	return plus
}

type divideOperator struct {
	leftValue  Valuer
	rightValue Valuer
}

func (p *divideOperator) GetValue(data map[string]interface{}) interface{} {
	leftValue := p.leftValue.GetValue(data)
	rightValue := p.rightValue.GetValue(data)
	typeOf := p.leftValue.GetType()

	switch typeOf {
	case "int64":
		a := leftValue.(int64)
		b := rightValue.(int64)
		return a / b
	case "float64":
		a := leftValue.(float64)
		b := rightValue.(float64)
		return a + b
	default:
		panic(fmt.Sprintf("divide does not support given type %s", typeOf))
	}

}

func (p *divideOperator) GetType() string {
	return p.leftValue.GetType()
}
